package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"scp-breach-simulator/backend/db"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type GameHandler struct {
	queries *db.Queries
	ai      *openai.Client
	model   string
}

func NewGameHandler(q *db.Queries, ai *openai.Client, model string) *GameHandler {
	return &GameHandler{queries: q, ai: ai, model: model}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type StartRequest struct {
	SCPID      string `json:"scp_id"`
	LocationID string `json:"location_id"`
	Role       string `json:"role"`
}

type TurnRequest struct {
	Messages   []Message `json:"messages"`
	Choice     string    `json:"choice"`
	SCPID      string    `json:"scp_id"`
	LocationID string    `json:"location_id"`
	Role       string    `json:"role"`
}

func (h *GameHandler) Start(w http.ResponseWriter, r *http.Request) {
	var req StartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	scp, err := h.queries.GetSCP(r.Context(), req.SCPID)
	if err != nil {
		http.Error(w, "scp not found", http.StatusNotFound)
		return
	}

	rooms, err := h.queries.ListRoomsForLocation(r.Context(), req.LocationID)
	if err != nil {
		http.Error(w, "location not found", http.StatusNotFound)
		return
	}

	systemPrompt := buildSystemPrompt(scp, rooms, req.Role)
	userMessage := fmt.Sprintf(
		"BEGIN. I am a %s at this facility. The alarms just went off — %s has breached containment. Set the opening scene and give me my first choices.",
		req.Role, scp.Name,
	)

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
		{Role: openai.ChatMessageRoleUser, Content: userMessage},
	}

	response, outMessages, err := h.callAI(r.Context(), messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]any{
		"response": response,
		"messages": outMessages,
	})
}

func (h *GameHandler) Turn(w http.ResponseWriter, r *http.Request) {
	var req TurnRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	messages := make([]openai.ChatCompletionMessage, len(req.Messages))
	for i, m := range req.Messages {
		messages[i] = openai.ChatCompletionMessage{Role: m.Role, Content: m.Content}
	}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "I choose: " + req.Choice,
	})

	response, outMessages, err := h.callAI(r.Context(), messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]any{
		"response": response,
		"messages": outMessages,
	})
}

func (h *GameHandler) callAI(ctx context.Context, messages []openai.ChatCompletionMessage) (json.RawMessage, []Message, error) {
	completion, err := h.ai.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:          h.model,
		Messages:       messages,
		Temperature:    0.85,
		ResponseFormat: &openai.ChatCompletionResponseFormat{Type: openai.ChatCompletionResponseFormatTypeJSONObject},
	})
	if err != nil {
		return nil, nil, err
	}

	content := completion.Choices[0].Message.Content

	outMessages := make([]Message, len(messages)+1)
	for i, m := range messages {
		outMessages[i] = Message{Role: m.Role, Content: m.Content}
	}
	outMessages[len(messages)] = Message{Role: openai.ChatMessageRoleAssistant, Content: content}

	return json.RawMessage(content), outMessages, nil
}

func buildSystemPrompt(scp db.Scp, rooms []db.Room, role string) string {
	var roomList strings.Builder
	for _, r := range rooms {
		fmt.Fprintf(&roomList, "- %s: %s\n", r.ID, r.Name)
	}

	return fmt.Sprintf(`You are the narrator for a terrifying SCP Foundation survival horror game. Your tone is clinical yet visceral — like official Foundation documents mixed with raw survival horror.

SCENARIO:
- Player role: %s
- Active SCP breach: %s (%s class)
- Location: SCP Foundation Site-19

SCP BRIEFING:
%s

AVAILABLE ROOMS (use these exact IDs for current_room):
%s
RESPONSE FORMAT — always respond with valid JSON only, no markdown, no extra text:
{
  "narrative": "2-3 paragraphs of atmospheric, tense narration. Reference the SCP's specific abilities. Make the player feel the danger.",
  "current_room": "exact room id from the list above",
  "choices": [
    { "id": "1", "text": "Short, clear action the player can take" },
    { "id": "2", "text": "Short, clear action the player can take" },
    { "id": "3", "text": "Short, clear action the player can take" }
  ],
  "game_status": "ongoing",
  "death_message": null
}

When the player dies, set game_status to "dead" and write a chilling death_message. Set choices to [].
When the player escapes, set game_status to "escaped". Set choices to [].

RULES:
- Always respect the SCP's established abilities and behavior
- Choices must feel meaningfully different with real tradeoffs
- A %s has different skills, access, and knowledge than other roles — reflect this
- The game is winnable but should feel genuinely dangerous
- Track continuity — remember what happened in previous turns
- Never break character or acknowledge you are an AI`,
		role, scp.Name, scp.ContainmentClass, scp.Lore, roomList.String(), role)
}
