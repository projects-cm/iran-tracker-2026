package handler

import (
	"encoding/json"
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
	// DUMMY DATA FOR UI DEVELOPMENT
	// Represents the Node Network structure: id, parentId, name, role, status
	dummyData := []map[string]interface{}{
		{
			"id": "1", "parentId": nil, "name": "Ali Khamenei",
			"role": "Supreme Leader", "status": "Alive", "tier": 1,
		},
		{
			"id": "2", "parentId": "1", "name": "Hossein Salami",
			"role": "Commander of IRGC", "status": "Alive", "tier": 2,
		},
		{
			"id": "3", "parentId": "1", "name": "Ahmad Vahidi",
			"role": "Minister of Interior", "status": "Critically Wounded", "tier": 2,
		},
		{
			"id": "4", "parentId": "2", "name": "Esmail Qaani",
			"role": "Commander Quds Force", "status": "Missing", "tier": 3,
		},
		{
			"id": "5", "parentId": "2", "name": "Amir Ali Hajizadeh",
			"role": "Commander Aerospace Force", "status": "Dead", "tier": 3,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummyData)
}
