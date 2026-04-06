package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Clients holds all external service clients
type Clients struct {
	DB *sql.DB
	// Add Gemini and Telegram clients here later
}

// InitClients initializes all external connections
func InitClients() (*Clients, error) {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./database/tracker.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database client initialized successfully")

	return &Clients{
		DB: db,
	}, nil
}

// Close closes all client connections
func (c *Clients) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
