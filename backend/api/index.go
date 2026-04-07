package handler

import (
	"log"
	"net/http"

	"iranian-tracker/backend/pkg/infra"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Pulse API Request: %s %s", r.Method, r.URL.Path)

	// 1. Initialize Clients from Environment
	clients, err := infra.InitClients()
	if err != nil {
		log.Printf("Failed to init clients: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer clients.Close()

	// 2. Compose logic from our packages
	router, _ := infra.Compose(clients)

	// 3. Delegate to chi router
	router.ServeHTTP(w, r)
}
