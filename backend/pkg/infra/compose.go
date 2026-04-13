package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"iranian-tracker/backend/pkg/dal"
	"iranian-tracker/backend/pkg/handler"
	"iranian-tracker/backend/pkg/service"
)

// Compose wires all layers together and returns the router and scraper
func Compose(clients *Clients) (http.Handler, *service.ScraperService, error) {
	r := chi.NewRouter()

	// 1. Common Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// 2. CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// 2. Initialize DAL
	dalRepo, err := dal.NewDB(clients.TursoURL, clients.TursoToken)
	if err != nil {
		return nil, nil, fmt.Errorf("dal init error: %w", err)
	}

	// 3. Initialize Services
	casualtyService := service.NewCasualtyService(dalRepo)
	processorService := service.NewProcessorService(clients.Gemini)
	scraperService := service.NewScraperService(clients.Telegram, dalRepo, processorService, casualtyService, clients.SimulationMode)

	// 4. Initialize Handlers
	casualtyHandler := handler.NewCasualtyHandler(casualtyService)
	scraperHandler := handler.NewScraperHandler(scraperService, TargetChannels)

	// 5. Define Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "up"})
	})

	r.Route("/v1", func(r chi.Router) {
		r.Get("/figures", casualtyHandler.GetFigures)
		r.Get("/pulse", scraperHandler.TriggerPulse)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/figures", casualtyHandler.GetFigures)
		r.Get("/pulse", scraperHandler.TriggerPulse)
	})

	return r, scraperService, nil
}
