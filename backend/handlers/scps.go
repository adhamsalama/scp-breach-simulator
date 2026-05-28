package handlers

import (
	"net/http"
	"scp-breach-simulator/backend/db"
)

type SCPsHandler struct {
	queries *db.Queries
}

func NewSCPsHandler(q *db.Queries) *SCPsHandler {
	return &SCPsHandler{queries: q}
}

type SCPResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ContainmentClass string `json:"containment_class"`
	Description      string `json:"description"`
	Lore             string `json:"lore"`
}

func (h *SCPsHandler) List(w http.ResponseWriter, r *http.Request) {
	scps, err := h.queries.ListSCPs(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := make([]SCPResponse, len(scps))
	for i, s := range scps {
		result[i] = SCPResponse{
			ID:               s.ID,
			Name:             s.Name,
			ContainmentClass: s.ContainmentClass,
			Description:      s.Description,
			Lore:             s.Lore,
		}
	}
	writeJSON(w, result)
}
