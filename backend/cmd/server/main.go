package main

import (
	"bufio"
	"database/sql"
	"log"
	"net/http"
	"os"
	"scp-breach-simulator/backend/db"
	"scp-breach-simulator/backend/handlers"
	"scp-breach-simulator/backend/seed"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	loadEnv(".env")

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is required in .env")
	}

	model := os.Getenv("MODEL")
	if model == "" {
		log.Fatal("MODEL is required in .env")
	}

	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = "https://models.github.ai/inference"
	aiClient := openai.NewClientWithConfig(cfg)

	database, err := sql.Open("sqlite3", "./scp.db")
	if err != nil {
		log.Fatal("failed to open database:", err)
	}
	defer database.Close()

	queries := db.New(database)

	if err := applySchema(database); err != nil {
		log.Fatal("failed to apply schema:", err)
	}
	if err := seed.Run(queries); err != nil {
		log.Fatal("failed to seed database:", err)
	}

	scpsHandler := handlers.NewSCPsHandler(queries)
	locationsHandler := handlers.NewLocationsHandler(queries)
	gameHandler := handlers.NewGameHandler(queries, aiClient, model)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/scps", scpsHandler.List)
	mux.HandleFunc("GET /api/locations", locationsHandler.List)
	mux.HandleFunc("POST /api/game/start", gameHandler.Start)
	mux.HandleFunc("POST /api/game/turn", gameHandler.Turn)

	log.Println("Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}

// loadEnv reads a .env file and sets any unset environment variables from it.
func loadEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		if os.Getenv(strings.TrimSpace(key)) == "" {
			os.Setenv(strings.TrimSpace(key), strings.TrimSpace(value))
		}
	}
}

func applySchema(database *sql.DB) error {
	_, err := database.Exec(`
	CREATE TABLE IF NOT EXISTS scps (
		id                TEXT PRIMARY KEY,
		name              TEXT NOT NULL,
		containment_class TEXT NOT NULL,
		description       TEXT NOT NULL,
		lore              TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS locations (
		id          TEXT PRIMARY KEY,
		name        TEXT NOT NULL,
		description TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS rooms (
		id          TEXT PRIMARY KEY,
		location_id TEXT NOT NULL REFERENCES locations(id),
		name        TEXT NOT NULL,
		image_path  TEXT NOT NULL
	);`)
	return err
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
