import { useState } from 'react';
import type { Room } from '../types/game';

interface Props {
  room: Room;
}

export default function LocationImage({ room }: Props) {
  const [failed, setFailed] = useState(false);

  return (
    <div className="relative w-full h-64 md:h-full md:min-h-80 bg-gray-950 overflow-hidden">
      {failed ? (
        <div className="absolute inset-0 flex flex-col items-center justify-center border border-gray-800">
          <div className="text-gray-700 text-6xl mb-4">▪</div>
          <div className="text-gray-600 font-mono text-xs tracking-widest uppercase">{room.name}</div>
          <div className="text-gray-800 font-mono text-xs mt-1">[IMAGE PLACEHOLDER]</div>
        </div>
      ) : (
        <img
          key={room.id}
          src={room.image_path}
          alt={room.name}
          className="w-full h-full object-contain opacity-80"
          onError={() => setFailed(true)}
        />
      )}
      <div className="absolute inset-0 bg-gradient-to-t from-[#0a0a0b] to-transparent pointer-events-none" />
      <div className="absolute bottom-3 left-3">
        <span className="bg-black/70 text-gray-300 font-mono text-xs px-2 py-1 tracking-widest uppercase">
          {room.name}
        </span>
      </div>
    </div>
  );
}
