package infra

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	tursoURL := strings.TrimSpace(os.Getenv("TURSO_DATABASE_URL"))
	if tursoURL == "" {
		return nil, fmt.Errorf("TURSO_DATABASE_URL is required")
	}

	tursoToken := strings.TrimSpace(os.Getenv("TURSO_AUTH_TOKEN"))
	if tursoToken == "" {
		return nil, fmt.Errorf("TURSO_AUTH_TOKEN is required")
	}

	// Gemini Initialization
	geminiKey := strings.TrimSpace(os.Getenv("GEMINI_API_KEY"))
	if geminiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is required")
	}
	genaiClient, err := genai.NewClient(context.Background(), option.WithAPIKey(geminiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to init gemini: %w", err)
	}

	// Telegram Initialization
	apiIDStr := strings.TrimSpace(os.Getenv("TELEGRAM_API_ID"))
	apiHash := strings.TrimSpace(os.Getenv("TELEGRAM_API_HASH"))
	
	// Optional in serverless mode (if we don't need real-time scraping)
	var telegramClient *telegram.Client
	if apiIDStr != "" && apiHash != "" {
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

		telegramClient = telegram.NewClient(apiID, apiHash, telegram.Options{
			SessionStorage: &session.FileStorage{Path: sessionFile},
		})
	}

	// Optional: Simulation Mode for testing without real keys
	simulationMode := os.Getenv("SIMULATION_MODE") == "true"

	return &Clients{
		TursoURL:       tursoURL,
		TursoToken:     tursoToken,
		Gemini:         genaiClient,
		Telegram:       telegramClient,
		SimulationMode: simulationMode,
	}, nil
}

// Close releases all external service resources
func (c *Clients) Close() {
	if c.Gemini != nil {
		c.Gemini.Close()
	}
}
