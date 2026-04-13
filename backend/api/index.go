package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"iranian-tracker/backend/pkg/infra"
)

var (
	globalRouter http.Handler
	initErr      error
	once         sync.Once
)

func initialize() {
	log.Printf("Initializing backend on cold start...")
	clients, err := infra.InitClients()
	if err != nil {
		initErr = err
		return
	}

	r, _, err := infra.Compose(clients)
	if err != nil {
		initErr = err
		return
	}
	globalRouter = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Pulse API Request: %s %s", r.Method, r.URL.Path)

	once.Do(initialize)

	if initErr != nil {
		log.Printf("Initialization error: %v", initErr)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if globalRouter == nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// 3. Delegate to chi router
	globalRouter.ServeHTTP(w, r)
}

func main() {
	once.Do(initialize)
	if initErr != nil {
		log.Fatalf("Fatal initialization error: %v", initErr)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local or specific testing
	}

	log.Printf("Starting Vercel Go Microservice on :%s", port)
	if err := http.ListenAndServe(":"+port, globalRouter); err != nil {
		log.Fatalf("Server exited ungracefully: %v", err)
	}
}
