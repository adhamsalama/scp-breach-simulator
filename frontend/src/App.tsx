import { useState } from 'react';
import type { SCP, Location, PlayerRole, GameState } from './types/game';
import { startGame } from './services/api';
import SetupScreen from './components/SetupScreen';
import GameScreen from './components/GameScreen';

type AppScreen = 'setup' | 'loading' | 'game';

export default function App() {
  const [screen, setScreen] = useState<AppScreen>('setup');
  const [gameState, setGameState] = useState<GameState | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function handleStart(scp: SCP, location: Location, role: PlayerRole) {
    setError(null);
    setScreen('loading');

    try {
      const { response, messages } = await startGame({ scp, location, role });

      const startingRoom = location.rooms.find(r => r.id === response.current_room)
        ? response.current_room
        : location.rooms[0].id;

      setGameState({
        scp,
        location,
        role,
        currentRoom: startingRoom,
        narrative: response.narrative,
        choices: response.choices,
        gameStatus: response.game_status,
        deathMessage: response.death_message,
        messages,
        isLoading: false,
        turnCount: 1,
      });
      setScreen('game');
    } catch (err) {
      console.error(err);
      setError('Failed to connect to the AI. Check your API key and try again.');
      setScreen('setup');
    }
  }

  function handleRestart() {
    setGameState(null);
    setScreen('setup');
  }

  function handleStateUpdate(updates: Partial<GameState>) {
    setGameState(prev => prev ? { ...prev, ...updates } : prev);
  }

  return (
    <>
      {screen === 'setup' && (
        <>
          {error && (
            <div className="bg-red-950 border-b border-red-800 text-red-300 font-mono text-xs px-4 py-2 text-center">
              {error}
            </div>
          )}
          <SetupScreen onStart={handleStart} />
        </>
      )}

      {screen === 'loading' && (
        <div className="min-h-screen bg-[#0a0a0b] flex items-center justify-center">
          <div className="text-center font-mono">
            <div className="text-red-500 text-xs tracking-[0.3em] uppercase animate-pulse mb-4">
              ■ Initializing breach scenario
            </div>
            <div className="text-gray-700 text-xs">Connecting to Foundation AI...</div>
          </div>
        </div>
      )}

      {screen === 'game' && gameState && (
        <GameScreen
          gameState={gameState}
          onStateUpdate={handleStateUpdate}
          onRestart={handleRestart}
        />
      )}
    </>
  );
}
