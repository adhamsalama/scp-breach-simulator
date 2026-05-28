import { useEffect, useState } from 'react';
import { fetchSCPs, fetchLocations } from '../services/api';
import type { SCP, Location, PlayerRole } from '../types/game';

const ROLES: PlayerRole[] = ['D-Class Personnel', 'Security Guard', 'Researcher', 'MTF Agent'];

const ROLE_DESCRIPTIONS: Record<PlayerRole, string> = {
  'D-Class Personnel': 'Expendable test subject. No weapons, no access. Your wits are all you have.',
  'Security Guard': 'Armed and trained, but protocol only gets you so far against the anomalous.',
  'Researcher': 'You know the SCPs better than anyone. Knowledge is your only weapon.',
  'MTF Agent': 'Elite containment specialist. Best equipped, best trained. Still mortal.',
};

const CLASS_COLORS: Record<string, string> = {
  Keter: 'text-red-400',
  Euclid: 'text-yellow-400',
  Safe: 'text-green-400',
  Thaumiel: 'text-purple-400',
};

interface Props {
  onStart: (scp: SCP, location: Location, role: PlayerRole) => void;
}

export default function SetupScreen({ onStart }: Props) {
  const [scps, setSCPs] = useState<SCP[]>([]);
  const [locations, setLocations] = useState<Location[]>([]);
  const [selectedSCP, setSelectedSCP] = useState<SCP | null>(null);
  const [selectedLocation, setSelectedLocation] = useState<Location | null>(null);
  const [selectedRole, setSelectedRole] = useState<PlayerRole | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    Promise.all([fetchSCPs(), fetchLocations()])
      .then(([s, l]) => { setSCPs(s); setLocations(l); })
      .catch(() => setError('Failed to load game data. Is the backend running?'))
      .finally(() => setLoading(false));
  }, []);

  const canStart = selectedSCP && selectedLocation && selectedRole;

  if (loading) {
    return (
      <div className="min-h-screen bg-[#0a0a0b] flex items-center justify-center font-mono">
        <p className="text-gray-600 text-xs tracking-widest animate-pulse uppercase">Loading...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0a0a0b] text-gray-200 p-6 flex flex-col items-center">
      <div className="w-full max-w-4xl">
        <div className="text-center mb-10">
          <p className="text-red-500 text-xs tracking-[0.3em] uppercase mb-2">Secure. Contain. Protect.</p>
          <h1 className="text-5xl font-bold text-white tracking-tight font-mono">SCP Foundation</h1>
          <p className="text-gray-500 mt-3 text-sm tracking-widest uppercase">Containment Breach Simulator</p>
        </div>

        {error && (
          <div className="border border-red-800 bg-red-950/30 text-red-300 font-mono text-xs px-4 py-3 mb-6 text-center">
            {error}
          </div>
        )}

        <div className="grid grid-cols-1 gap-8">
          {/* SCP Selection */}
          <section>
            <h2 className="text-xs text-red-400 tracking-[0.25em] uppercase mb-3 font-mono">
              01 — Select Breached SCP
            </h2>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3">
              {scps.map(scp => (
                <button
                  key={scp.id}
                  onClick={() => setSelectedSCP(scp)}
                  className={`text-left p-4 border rounded transition-all ${
                    selectedSCP?.id === scp.id
                      ? 'border-red-500 bg-red-950/30 text-white'
                      : 'border-gray-800 bg-gray-900/50 text-gray-400 hover:border-gray-600 hover:text-gray-200'
                  }`}
                >
                  <div className="font-mono font-bold text-sm">{scp.name}</div>
                  <div className={`text-xs mt-1 ${CLASS_COLORS[scp.containment_class] ?? 'text-gray-400'}`}>
                    {scp.containment_class}
                  </div>
                  <div className="text-xs mt-2 text-gray-500 leading-relaxed">{scp.description}</div>
                </button>
              ))}
            </div>
          </section>

          {/* Location Selection */}
          <section>
            <h2 className="text-xs text-red-400 tracking-[0.25em] uppercase mb-3 font-mono">
              02 — Select Location
            </h2>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
              {locations.map(loc => (
                <button
                  key={loc.id}
                  onClick={() => setSelectedLocation(loc)}
                  className={`text-left p-4 border rounded transition-all ${
                    selectedLocation?.id === loc.id
                      ? 'border-red-500 bg-red-950/30 text-white'
                      : 'border-gray-800 bg-gray-900/50 text-gray-400 hover:border-gray-600 hover:text-gray-200'
                  }`}
                >
                  <div className="font-mono font-bold text-sm">{loc.name}</div>
                  <div className="text-xs mt-2 text-gray-500 leading-relaxed">{loc.description}</div>
                </button>
              ))}
            </div>
          </section>

          {/* Role Selection */}
          <section>
            <h2 className="text-xs text-red-400 tracking-[0.25em] uppercase mb-3 font-mono">
              03 — Select Role
            </h2>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
              {ROLES.map(role => (
                <button
                  key={role}
                  onClick={() => setSelectedRole(role)}
                  className={`text-left p-4 border rounded transition-all ${
                    selectedRole === role
                      ? 'border-red-500 bg-red-950/30 text-white'
                      : 'border-gray-800 bg-gray-900/50 text-gray-400 hover:border-gray-600 hover:text-gray-200'
                  }`}
                >
                  <div className="font-mono font-bold text-sm">{role}</div>
                  <div className="text-xs mt-2 text-gray-500 leading-relaxed">{ROLE_DESCRIPTIONS[role]}</div>
                </button>
              ))}
            </div>
          </section>

          <div className="flex justify-center mt-2 mb-8">
            <button
              disabled={!canStart}
              onClick={() => canStart && onStart(selectedSCP, selectedLocation, selectedRole)}
              className={`px-12 py-4 font-mono text-sm tracking-widest uppercase transition-all ${
                canStart
                  ? 'bg-red-700 hover:bg-red-600 text-white cursor-pointer'
                  : 'bg-gray-800 text-gray-600 cursor-not-allowed'
              }`}
            >
              Initiate Breach
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
