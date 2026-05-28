# SCP Breach Simulator

A text-based survival horror game set in the SCP Foundation universe. An AI narrator generates a branching story where you must survive a containment breach and escape with your life.

## How it works

1. Pick an SCP, a location, and a role (D-Class, Guard, Researcher, or MTF Agent)
2. The AI sets the scene and presents you with choices
3. Every choice has consequences — survive long enough to escape

## Stack

- **Frontend** — React + TypeScript + Tailwind CSS (Vite)
- **Backend** — Go + SQLite + sqlc
- **AI** — OpenAI-compatible API for narrative generation

## Setup

### Backend

Create `backend/.env`:

```
OPENAI_API_KEY=your-key-here
MODEL=your-model-here
```

### Frontend

No `.env` needed. The frontend calls `http://localhost:8080` by default. To override:

```
VITE_API_URL=http://your-backend-url
```

## Running

```bash
./run.sh
```

Frontend runs on http://localhost:5173 — backend on http://localhost:8080.

## Adding images

Drop images into `public/images/{location_id}/{room_id}.jpg`. See `IMAGE_PROMPTS.md` for generation prompts.
