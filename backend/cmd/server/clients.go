package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/generative-ai-go/genai"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"google.golang.org/api/option"
)

// Clients holds all external service configuration
type Clients struct {
	TursoURL       string
	TursoToken     string
	Gemini         *genai.Client
	Telegram       *telegram.Client
	SimulationMode bool
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

	// Gemini Initialization
	geminiKey := os.Getenv("GEMINI_API_KEY")
	if geminiKey == "" {
		log.Fatal("GEMINI_API_KEY is required")
	}
	genaiClient, err := genai.NewClient(context.Background(), option.WithAPIKey(geminiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to init gemini: %w", err)
	}

	// Telegram Initialization
	apiIDStr := os.Getenv("TELEGRAM_API_ID")
	apiHash := os.Getenv("TELEGRAM_API_HASH")
	if apiIDStr == "" || apiHash == "" {
		log.Fatal("TELEGRAM_API_ID and TELEGRAM_API_HASH are required")
	}
	apiID, err := strconv.Atoi(apiIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid TELEGRAM_API_ID: %w", err)
	}

	// Setup session storage
	sessionDir := ".session"
	if err := os.MkdirAll(sessionDir, 0700); err != nil {
		return nil, err
	}
	sessionFile := filepath.Join(sessionDir, "session.json")

	telegramClient := telegram.NewClient(apiID, apiHash, telegram.Options{
		SessionStorage: &session.FileStorage{Path: sessionFile},
	})

	// Optional: Simulation Mode for testing without real keys
	simulationMode := os.Getenv("SIMULATION_MODE") == "true"

	log.Println("Client configuration loaded successfully")

	return &Clients{
		TursoURL:       tursoURL,
		TursoToken:     tursoToken,
		Gemini:         genaiClient,
		Telegram:       telegramClient,
		SimulationMode: simulationMode,
	}, nil
}
