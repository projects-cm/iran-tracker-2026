package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
	"iranian-tracker/pkg/dal"
)

// ScraperService handles scraping messages from Telegram channels
type ScraperService struct {
	client    *telegram.Client
	api       *tg.Client
	db        *dal.DB
	processor *ProcessorService
	casualty  *CasualtyService
}

// NewScraperService creates a new ScraperService instance
func NewScraperService(client *telegram.Client, db *dal.DB, processor *ProcessorService, casualty *CasualtyService) *ScraperService {
	return &ScraperService{
		client:    client,
		api:       client.API(),
		db:        db,
		processor: processor,
		casualty:  casualty,
	}
}

// StartScraping begins monitoring the target channels
func (s *ScraperService) StartScraping(ctx context.Context, targetChannels []string) error {
	log.Println("Starting Telegram Scraper...")

	// Connect to Telegram
	return s.client.Run(ctx, func(ctx context.Context) error {
		log.Println("Successfully connected and authenticated with Telegram.")

		// For each target channel, we would resolve its ID and fetch history
		for _, username := range targetChannels {
			go s.scrapeChannelRoutine(ctx, username)
		}

		// Keep the connection alive
		<-ctx.Done()
		return ctx.Err()
	})
}

// scrapeChannelRoutine polls a specific channel with jittering to avoid bans
func (s *ScraperService) scrapeChannelRoutine(ctx context.Context, username string) {
	log.Printf("Initializing scraping routine for %s", username)
	
	// Resolve the username to an InputPeer
	peer, err := s.resolveUsername(ctx, username)
	if err != nil {
		log.Printf("Failed to resolve channel %s: %v", username, err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Fetch the latest messages (e.g., limit 10)
			err := s.fetchRecentMessages(ctx, peer, username)
			if err != nil {
				log.Printf("Error fetching messages from %s: %v", username, err)
			}

			// Implement human-like jitter (random delay between 2 and 10 seconds)
			jitterSecs := rand.Intn(9) + 2
			jitterDuration := time.Duration(jitterSecs) * time.Second
			log.Printf("Sleeping for %v before next fetch for %s...", jitterDuration, username)
			time.Sleep(jitterDuration)
		}
	}
}

// resolveUsername uses contacts.resolveUsername to get peer details
func (s *ScraperService) resolveUsername(ctx context.Context, username string) (tg.InputPeerClass, error) {
	res, err := s.api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{Username: username})
	if err != nil {
		return nil, err
	}
	for _, chat := range res.GetChats() {
		if channel, ok := chat.(*tg.Channel); ok {
			return &tg.InputPeerChannel{
				ChannelID:  channel.ID,
				AccessHash: channel.AccessHash,
			}, nil
		}
	}
	return nil, fmt.Errorf("channel not found")
}

// fetchRecentMessages retrieves and processes new messages from the peer
func (s *ScraperService) fetchRecentMessages(ctx context.Context, peer tg.InputPeerClass, sourceName string) error {
	// This makes an API call to messages.getHistory
	res, err := s.api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
		Peer:  peer,
		Limit: 10,
	})
	if err != nil {
		return err
	}

	messagesCount := 0
	switch history := res.(type) {
	case *tg.MessagesMessages:
		messagesCount = len(history.Messages)
		s.processMessages(ctx, history.Messages, sourceName)
	case *tg.MessagesMessagesSlice:
		messagesCount = len(history.Messages)
		s.processMessages(ctx, history.Messages, sourceName)
	case *tg.MessagesChannelMessages:
		messagesCount = len(history.Messages)
		s.processMessages(ctx, history.Messages, sourceName)
	default:
		return fmt.Errorf("unexpected history type")
	}

	log.Printf("Fetched %d messages from peer.", messagesCount)
	return nil
}

// processMessages takes raw Telegram messages and sends them to the Processor
func (s *ScraperService) processMessages(ctx context.Context, messages []tg.MessageClass, sourceName string) {
	for _, rawMsg := range messages {
		msg, ok := rawMsg.(*tg.Message)
		if !ok || msg.Message == "" {
			continue
		}

		// 1. Deduplication: Check if we've already handled this message
		processed, err := s.db.IsReportProcessed(ctx, msg.ID)
		if err != nil {
			log.Printf("Error checking deduplication for msg %d: %v", msg.ID, err)
			continue
		}
		if processed {
			continue
		}

		log.Printf("Processing NEW message from %s (ID: %d): %.50s...", sourceName, msg.ID, msg.Message)

		// 2. Intelligence Extraction: Ask Gemini what this means
		ext, err := s.processor.ProcessRawText(ctx, msg.Message, sourceName)
		if err != nil {
			log.Printf("Gemini extraction failed for msg %d: %v", msg.ID, err)
			continue
		}

		// 3. Status Update: Apply logic to the Figure and persist the report
		// Inject a temporary timestamp if missing
		tsCtx := context.WithValue(ctx, "timestamp", time.Now().Format(time.RFC3339))
		err = s.casualty.ProcessNewReport(tsCtx, ext, msg.Message, sourceName, msg.ID)
		if err != nil {
			log.Printf("Failed to process report for msg %d: %v", msg.ID, err)
		}
	}
}
