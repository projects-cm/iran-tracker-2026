package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"iranian-tracker/backend/pkg/service"
)

// ScraperHandler handles explicit HTTP triggers for the background scraper
type ScraperHandler struct {
	scraper *service.ScraperService
	targets []string
}

// NewScraperHandler creates a new instance of ScraperHandler
func NewScraperHandler(scraper *service.ScraperService, targets []string) *ScraperHandler {
	return &ScraperHandler{
		scraper: scraper,
		targets: targets,
	}
}

// TriggerPulse executes a single immediate run of the scraper (used for cron-job pings)
func (h *ScraperHandler) TriggerPulse(w http.ResponseWriter, r *http.Request) {
	log.Println("Received HTTP trigger for Scraper Pulse...")
	
	// We run directly within the HTTP context. Vercel allows up to 10-60s 
	// for serverless functions, which is enough time for a single scrape pass.
	err := h.scraper.StartScrapingPulse(r.Context(), h.targets)
	if err != nil {
		log.Printf("⚠️ Pulse scrape failed: %v", err)
		http.Error(w, "Scrape failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"message": "Scraping pulse completed successfully",
	})
}
