package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"iranian-tracker/pkg/service"
)

// CasualtyHandler handles HTTP requests related to the tracker dashboard
type CasualtyHandler struct {
	service *service.CasualtyService
}

// NewCasualtyHandler creates a new instance of CasualtyHandler
func NewCasualtyHandler(svc *service.CasualtyService) *CasualtyHandler {
	return &CasualtyHandler{
		service: svc,
	}
}

// GetFigures handles the request to retrieve the dashboard figures
func (h *CasualtyHandler) GetFigures(w http.ResponseWriter, r *http.Request) {
	figures, err := h.service.GetTrackerDashboard(r.Context())
	if err != nil {
		log.Printf("Error fetching tracker dashboard: %v", err)
		http.Error(w, "Failed to fetch figures", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(figures)
}
