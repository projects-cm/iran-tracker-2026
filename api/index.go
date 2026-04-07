package handler

import (
	"log"
	"net/http"
	"os"

	"iranian-tracker/pkg/infra"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Pulse API Request: %s %s", r.Method, r.URL.Path)

	// 1. Initialize Clients from Environment
	clients := infra.InitClients(
		os.Getenv("TURSO_DATABASE_URL"),
		os.Getenv("TURSO_AUTH_TOKEN"),
		os.Getenv("GEMINI_API_KEY"),
		os.Getenv("SIMULATION_MODE") == "true",
	)
	defer clients.Close()

	// 2. Compose logic from the backend package
	router, _ := infra.Compose(clients)

	// 3. Delegate to chi router
	router.ServeHTTP(w, r)
}
