package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"iranian-tracker/pkg/dal"
	"iranian-tracker/pkg/handler"
	"iranian-tracker/pkg/service"
)

// Compose wires all layers together and returns the router and scraper
func Compose(clients *Clients) (http.Handler, *service.ScraperService) {
	r := chi.NewRouter()

	// 1. Common Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// 2. Initialize DAL
	dalRepo, err := dal.NewDB(clients.TursoURL, clients.TursoToken)
	if err != nil {
		panic(err)
	}

	// 3. Initialize Services
	casualtyService := service.NewCasualtyService(dalRepo)
	processorService := service.NewProcessorService(clients.Gemini)
	scraperService := service.NewScraperService(clients.Telegram, dalRepo, processorService, casualtyService, clients.SimulationMode)

	// 4. Initialize Handlers
	casualtyHandler := handler.NewCasualtyHandler(casualtyService)

	// 5. Define Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "up"})
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/figures", casualtyHandler.GetFigures)
	})

	return r, scraperService
}
