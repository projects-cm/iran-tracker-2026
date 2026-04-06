package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Initial configuration
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found at ../.env, checking current directory...")
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found in current directory either.")
		}
	}

	// 2. Initialize Clients (Database, external APIs, etc.)
	clients, err := InitClients()
	if err != nil {
		log.Fatalf("Initialization error: %v", err)
	}

	// 3. Compose layers (Handlers -> Services -> DAL) and build router
	router, scraper := Compose(clients)

	// 4. Define Target Channels to Monitor
	targets := []string{"amitsegal", "abualiexpress"}

	// 5. Start the Scraper in the background
	scraperCtx, scraperCancel := context.WithCancel(context.Background())
	defer scraperCancel()
	go func() {
		if err := scraper.StartScraping(scraperCtx, targets); err != nil {
			log.Printf("⚠️ Scraper error: %v", err)
		}
	}()

	// 6. Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Iranian Casualty Tracker Backend starting on port %s...\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
