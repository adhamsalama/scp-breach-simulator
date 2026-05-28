package handlers

import (
	"encoding/json"
	"net/http"
	"scp-breach-simulator/backend/db"
)

type LocationsHandler struct {
	queries *db.Queries
}

func NewLocationsHandler(q *db.Queries) *LocationsHandler {
	return &LocationsHandler{queries: q}
}

type RoomResponse struct {
	ID         string `json:"id"`
	LocationID string `json:"location_id"`
	Name       string `json:"name"`
	ImagePath  string `json:"image_path"`
}

type LocationResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Rooms       []RoomResponse `json:"rooms"`
}

func (h *LocationsHandler) List(w http.ResponseWriter, r *http.Request) {
	locations, err := h.queries.ListLocations(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := make([]LocationResponse, 0, len(locations))
	for _, loc := range locations {
		rooms, err := h.queries.ListRoomsForLocation(r.Context(), loc.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		roomResponses := make([]RoomResponse, len(rooms))
		for i, room := range rooms {
			roomResponses[i] = RoomResponse{
				ID:         room.ID,
				LocationID: room.LocationID,
				Name:       room.Name,
				ImagePath:  room.ImagePath,
			}
		}
		result = append(result, LocationResponse{
			ID:          loc.ID,
			Name:        loc.Name,
			Description: loc.Description,
			Rooms:       roomResponses,
		})
	}

	writeJSON(w, result)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
