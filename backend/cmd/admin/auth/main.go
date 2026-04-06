package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"github.com/joho/godotenv"
)

// TerminalAuth implements auth.User for terminal-based login
type TerminalAuth struct {
}

func (a TerminalAuth) Phone(ctx context.Context) (string, error) {
	fmt.Print("📞 Enter Phone Number (e.g., +123456789): ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(s), nil
}

func (a TerminalAuth) Password(ctx context.Context) (string, error) {
	fmt.Print("🛡️  Enter 2FA Password (if enabled, else leave blank): ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(s), nil
}

func (a TerminalAuth) Code(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
	fmt.Print("🔢 Enter the Code sent to your Telegram/SMS: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(s), nil
}

func (a TerminalAuth) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return nil
}

func (a TerminalAuth) SignUp(ctx context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, fmt.Errorf("sign up not supported in this tool")
}

func main() {
	_ = godotenv.Load("../.env")

	fmt.Println("🚀 Iranian Casualty Tracker - Telegram Authentication Setup")
	fmt.Println("---------------------------------------------------------")

	apiIDStr := os.Getenv("TELEGRAM_API_ID")
	apiHash := os.Getenv("TELEGRAM_API_HASH")
	if apiIDStr == "" || apiHash == "" {
		log.Fatal("❌ Missing TELEGRAM_API_ID or TELEGRAM_API_HASH in .env")
	}

	apiID, _ := strconv.Atoi(apiIDStr)

	sessionDir := ".session"
	_ = os.MkdirAll(sessionDir, 0700)
	sessionFile := filepath.Join(sessionDir, "session.json")

	// 1. Initialize Client
	client := telegram.NewClient(apiID, apiHash, telegram.Options{
		SessionStorage: &session.FileStorage{Path: sessionFile},
	})

	// 2. Handle Authentication Flow
	// Using our custom TerminalAuth struct
	flow := auth.NewFlow(
		TerminalAuth{},
		auth.SendCodeOptions{},
	)

	fmt.Println("Attempting to connect and authenticate...")
	if err := client.Run(context.Background(), func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}
		
		status, err := client.Self(ctx)
		if err != nil {
			return err
		}
		
		fmt.Printf("✅ SUCCESS: Successfully authenticated as %s (%s)\n", status.FirstName, status.Username)
		fmt.Printf("📂 Session saved to: %s\n", sessionFile)
		return nil
	}); err != nil {
		log.Fatalf("❌ Authentication failed: %v", err)
	}
}
