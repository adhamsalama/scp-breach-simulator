export interface SCP {
  id: string;
  name: string;
  containment_class: 'Safe' | 'Euclid' | 'Keter' | 'Thaumiel';
  description: string;
  lore: string;
}

export interface Room {
  id: string;
  location_id: string;
  name: string;
  image_path: string;
}

export interface Location {
  id: string;
  name: string;
  description: string;
  rooms: Room[];
}

export type PlayerRole = 'D-Class Personnel' | 'Security Guard' | 'Researcher' | 'MTF Agent';

export interface Choice {
  id: string;
  text: string;
}

export type GameStatus = 'ongoing' | 'dead' | 'escaped';

export interface AIResponse {
  narrative: string;
  current_room: string;
  choices: Choice[];
  game_status: GameStatus;
  death_message?: string;
}

export interface Message {
  role: 'user' | 'assistant' | 'system';
  content: string;
}

export interface GameState {
  scp: SCP;
  location: Location;
  role: PlayerRole;
  currentRoom: string;
  narrative: string;
  choices: Choice[];
  gameStatus: GameStatus;
  deathMessage?: string;
  messages: Message[];
  isLoading: boolean;
  turnCount: number;
}
