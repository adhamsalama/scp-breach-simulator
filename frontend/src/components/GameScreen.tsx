import { useEffect, useState } from 'react';
import type { GameState, Choice } from '../types/game';
import { makeChoice } from '../services/api';
import LocationImage from './LocationImage';

interface Props {
  gameState: GameState;
  onStateUpdate: (updates: Partial<GameState>) => void;
  onRestart: () => void;
}

export default function GameScreen({ gameState, onStateUpdate, onRestart }: Props) {
  const [typedNarrative, setTypedNarrative] = useState('');
  const [isTyping, setIsTyping] = useState(false);

  const currentRoom =
    gameState.location.rooms.find(r => r.id === gameState.currentRoom) ??
    gameState.location.rooms[0];

  useEffect(() => {
    if (!gameState.narrative) return;
    setIsTyping(true);
    setTypedNarrative('');
    let i = 0;
    const interval = setInterval(() => {
      setTypedNarrative(gameState.narrative.slice(0, i + 1));
      i++;
      if (i >= gameState.narrative.length) {
        clearInterval(interval);
        setIsTyping(false);
      }
    }, 12);
    return () => clearInterval(interval);
  }, [gameState.narrative]);

  async function handleChoice(choice: Choice) {
    if (gameState.isLoading || isTyping) return;
    onStateUpdate({ isLoading: true });

    try {
      const { response, messages } = await makeChoice(
        choice.text,
        gameState.messages,
        gameState.scp.id,
        gameState.location.id,
        gameState.role,
      );

      const newRoom = gameState.location.rooms.find(r => r.id === response.current_room)
        ? response.current_room
        : gameState.currentRoom;

      onStateUpdate({
        narrative: response.narrative,
        currentRoom: newRoom,
        choices: response.choices,
        gameStatus: response.game_status,
        deathMessage: response.death_message,
        messages,
        isLoading: false,
        turnCount: gameState.turnCount + 1,
      });
    } catch (err) {
      console.error(err);
      onStateUpdate({ isLoading: false });
    }
  }

  const isDead = gameState.gameStatus === 'dead';
  const isEscaped = gameState.gameStatus === 'escaped';
  const isOver = isDead || isEscaped;

  return (
    <div className="min-h-screen bg-[#0a0a0b] flex flex-col">
      <div className="border-b border-gray-800 px-4 py-2 flex items-center justify-between">
        <div className="flex items-center gap-4">
          <span className="text-red-500 font-mono text-xs tracking-widest uppercase">
            {gameState.scp.name}
          </span>
          <span className="text-gray-700">|</span>
          <span className="text-gray-500 font-mono text-xs">{gameState.role}</span>
          <span className="text-gray-700">|</span>
          <span className="text-gray-600 font-mono text-xs">Turn {gameState.turnCount}</span>
        </div>
        <button
          onClick={onRestart}
          className="text-gray-600 hover:text-gray-400 font-mono text-xs tracking-widest uppercase transition-colors"
        >
          Abandon
        </button>
      </div>

      <div className="flex flex-col md:flex-row flex-1">
        <div className="md:w-2/5 lg:w-1/2">
          <LocationImage room={currentRoom} />
        </div>

        <div className="flex-1 flex flex-col p-6 gap-6 overflow-y-auto">
          {isOver && (
            <div
              className={`border px-4 py-3 font-mono text-sm ${
                isDead
                  ? 'border-red-800 bg-red-950/30 text-red-300'
                  : 'border-green-800 bg-green-950/30 text-green-300'
              }`}
            >
              {isDead ? '— PERSONNEL DECEASED —' : '— CONTAINMENT BREACH SURVIVED —'}
            </div>
          )}

          <div className="flex-1">
            <p className="text-gray-300 font-mono text-sm leading-relaxed whitespace-pre-wrap">
              {typedNarrative}
              {isTyping && <span className="animate-pulse text-red-500">▌</span>}
            </p>
            {isDead && gameState.deathMessage && (
              <p className="text-red-400 font-mono text-sm leading-relaxed mt-4 border-t border-red-900 pt-4">
                {gameState.deathMessage}
              </p>
            )}
          </div>

          {!isOver && (
            <div className="flex flex-col gap-2">
              {gameState.isLoading ? (
                <div className="text-gray-600 font-mono text-xs animate-pulse tracking-widest">
                  ■ ■ ■ PROCESSING
                </div>
              ) : (
                gameState.choices.map((choice, i) => (
                  <button
                    key={choice.id}
                    onClick={() => handleChoice(choice)}
                    disabled={isTyping}
                    className={`text-left px-4 py-3 border font-mono text-sm transition-all ${
                      isTyping
                        ? 'border-gray-800 text-gray-700 cursor-not-allowed'
                        : 'border-gray-700 text-gray-300 hover:border-red-600 hover:text-white hover:bg-red-950/20 cursor-pointer'
                    }`}
                  >
                    <span className="text-red-600 mr-3">[{i + 1}]</span>
                    {choice.text}
                  </button>
                ))
              )}
            </div>
          )}

          {isOver && (
            <button
              onClick={onRestart}
              className="border border-gray-700 hover:border-gray-500 text-gray-400 hover:text-gray-200 font-mono text-xs px-6 py-3 tracking-widest uppercase transition-all self-start"
            >
              New Session
            </button>
          )}
        </div>
      </div>
    </div>
  );
}
