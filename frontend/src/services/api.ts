import type { SCP, Location, Message, AIResponse, PlayerRole } from '../types/game';

const BASE_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8080';

async function get<T>(path: string): Promise<T> {
  const res = await fetch(`${BASE_URL}${path}`);
  if (!res.ok) throw new Error(`GET ${path} failed: ${res.statusText}`);
  return res.json();
}

async function post<T>(path: string, body: unknown): Promise<T> {
  const res = await fetch(`${BASE_URL}${path}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body),
  });
  if (!res.ok) throw new Error(`POST ${path} failed: ${res.statusText}`);
  return res.json();
}

export function fetchSCPs(): Promise<SCP[]> {
  return get<SCP[]>('/api/scps');
}

export function fetchLocations(): Promise<Location[]> {
  return get<Location[]>('/api/locations');
}

export async function startGame(params: {
  scp: SCP;
  location: Location;
  role: PlayerRole;
}): Promise<{ response: AIResponse; messages: Message[] }> {
  return post('/api/game/start', {
    scp_id: params.scp.id,
    location_id: params.location.id,
    role: params.role,
  });
}

export async function makeChoice(
  choice: string,
  messages: Message[],
  scpId: string,
  locationId: string,
  role: PlayerRole,
): Promise<{ response: AIResponse; messages: Message[] }> {
  return post('/api/game/turn', {
    choice,
    messages,
    scp_id: scpId,
    location_id: locationId,
    role,
  });
}
