package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"iranian-tracker/backend/pkg/infra"
	"iranian-tracker/backend/pkg/service"
)

var (
	globalRouter  http.Handler
	globalScraper *service.ScraperService
	initErr       error
	once          sync.Once
)

func initialize() {
	log.Printf("Initializing backend on cold start...")
	clients, err := infra.InitClients()
	if err != nil {
		initErr = err
		return
	}

	r, scraper, err := infra.Compose(clients)
	if err != nil {
		initErr = err
		return
	}
	globalRouter = r
	globalScraper = scraper
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

	if globalScraper != nil {
		scraperCtx := context.Background() // Safe as this is stay-alive service
		go func() {
			log.Printf("Starting background scraper on Vercel for targets: %v", infra.TargetChannels)
			if err := globalScraper.StartScraping(scraperCtx, infra.TargetChannels); err != nil {
				log.Printf("⚠️ Scraper error: %v", err)
			}
		}()
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
