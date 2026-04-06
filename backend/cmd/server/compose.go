package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"iranian-tracker/pkg/dal"
	"iranian-tracker/pkg/handler"
	"iranian-tracker/pkg/service"
)

// Compose wires all layers together and returns the router
func Compose(clients *Clients) http.Handler {
	r := chi.NewRouter()

	// 1. Common Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30))

	// 2. Initialize DAL
	dalRepo, err := dal.NewDB(clients.TursoURL, clients.TursoToken)
	if err != nil {
		panic(err)
	}

	// 3. Initialize Services
	casualtyService := service.NewCasualtyService(dalRepo)
	// scraperService := service.NewScraperService(clients.Telegram, clients.Gemini, dalRepo)

	// 4. Initialize Handlers
	casualtyHandler := handler.NewCasualtyHandler(casualtyService)

	// 5. Define Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "up"})
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/figures", casualtyHandler.GetFigures)
		// r.Get("/figures/{id}/reports", casualtyHandler.GetReports)
		// r.Get("/stats", casualtyHandler.GetStats)
	})

	return r
}
