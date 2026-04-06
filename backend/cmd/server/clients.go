package main

import (
	"log"
	"os"
)

// Clients holds all external service configuration
type Clients struct {
	TursoURL   string
	TursoToken string
	// Add Gemini and Telegram clients here later
}

// InitClients reads environment variables for all external connections
func InitClients() (*Clients, error) {
	tursoURL := os.Getenv("TURSO_DATABASE_URL")
	if tursoURL == "" {
		log.Fatal("TURSO_DATABASE_URL is required")
	}

	tursoToken := os.Getenv("TURSO_AUTH_TOKEN")
	if tursoToken == "" {
		log.Fatal("TURSO_AUTH_TOKEN is required")
	}

	log.Println("Client configuration loaded successfully")

	return &Clients{
		TursoURL:   tursoURL,
		TursoToken: tursoToken,
	}, nil
}
