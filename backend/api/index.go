package handler

import (
	"net/http"

	"iranian-tracker/pkg/infra"
)

// NOTE: For better serverless support, I am moving Compose to pkg/infra
// to avoid circular dependencies or main package issues.

func Handler(w http.ResponseWriter, r *http.Request) {
	clients, err := infra.InitClients()
	if err != nil {
		http.Error(w, "Infrastructure initialization failed", http.StatusInternalServerError)
		return
	}

	// We'll use a version of Compose that only returns the router
	router, _ := infra.Compose(clients)
	router.ServeHTTP(w, r)
}
