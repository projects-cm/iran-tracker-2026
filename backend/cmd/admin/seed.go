package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"iranian-tracker/pkg/dal"
)

func main() {
	// Load .env from root
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Failed to load .env:", err)
	}

	url := os.Getenv("TURSO_DATABASE_URL")
	token := os.Getenv("TURSO_AUTH_TOKEN")

	if url == "" || token == "" {
		log.Fatal("TURSO_DATABASE_URL and TURSO_AUTH_TOKEN are required")
	}

	fmt.Println("=== Iranian Leadership Tracker: Seeding Tool ===")
	fmt.Printf("Target Database: %s\n\n", url)

	db, err := dal.NewDB(url, token)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	ctx := context.Background()

	figures := []struct {
		Name    string
		Persian string
		Tier    int
		Aliases []string
	}{
		{"Ali Khamenei", "علی خامنه‌ای", 1, []string{"Khamenei", "Supreme Leader"}},
		{"Mojtaba Khamenei", "مجتبی خامنه‌ای", 2, []string{"Mojtaba"}},
		{"Masoud Pezeshkian", "مسعود پزشکیان", 2, []string{"Pezeshkian"}},
		{"Ahmad Vahidi", "احمد وحیدی", 2, []string{"Vahidi"}},
		{"Hossein Salami", "حسین سلامی", 2, []string{"Salami"}},
		{"Esmail Qaani", "اسماعیل قاآنی", 3, []string{"Qaani", "Ghaani"}},
		{"Amir Ali Hajizadeh", "امیرعلی حاجی‌زاده", 3, []string{"Hajizadeh"}},
	}

	for _, f := range figures {
		id, err := db.CreateFigure(ctx, f.Name, f.Persian, f.Tier, "Alive", f.Aliases)
		if err != nil {
			log.Printf("⚠️ Error seeding %s: %v", f.Name, err)
			continue
		}
		fmt.Printf("✅ Seeded %-25s (ID: %d)\n", f.Name, id)
	}

	fmt.Println("\n✅ Seeding complete.")
}
