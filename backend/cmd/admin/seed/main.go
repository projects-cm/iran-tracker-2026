package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"iranian-tracker/pkg/dal"
)

func ptr[T any](v T) *T {
	return &v
}

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

	dbConn, err := dal.NewDB(url, token)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	// RESET SCHEMA FOR DEV (One-time migration)
	fmt.Println("🔄 Resetting schema for new columns...")
	_, _ = dbConn.GetDB().Exec("DROP TABLE IF EXISTS aliases")
	_, _ = dbConn.GetDB().Exec("DROP TABLE IF EXISTS reports")
	_, _ = dbConn.GetDB().Exec("DROP TABLE IF EXISTS figures")
	
	// Re-run NewDB to trigger initSchema
	dbConn, err = dal.NewDB(url, token)
	if err != nil {
		log.Fatal("Failed to re-initialize schema:", err)
	}

	ctx := context.Background()

	figures := []struct {
		Name     string
		Persian  string
		Role     string
		Tier     int
		Status   string
		ParentID *int
		Aliases  []string
	}{
		{"Ali Khamenei", "علی خامنه‌ای", "Supreme Leader", 1, "Alive", nil, []string{"Khamenei", "Supreme Leader"}},
		{"Mojtaba Khamenei", "مجتبی خامنه‌ای", "Son of Supreme Leader", 2, "Alive", ptr(1), []string{"Mojtaba"}},
		{"Masoud Pezeshkian", "مسعود پزشکیان", "President", 2, "Alive", ptr(1), []string{"Pezeshkian"}},
		{"Ahmad Vahidi", "احمد وحیدی", "Secretary of SNSC", 2, "Critically Wounded", ptr(1), []string{"Vahidi"}},
		{"Hossein Salami", "حسین سلامی", "Commander of IRGC", 2, "Alive", ptr(1), []string{"Salami"}},
		{"Esmail Qaani", "اسماعیل قاآنی", "Commander Quds Force", 3, "Missing", ptr(5), []string{"Qaani", "Ghaani"}},
		{"Amir Ali Hajizadeh", "امیرعلی حاجی‌زاده", "Commander Aerospace Force", 3, "Dead", ptr(5), []string{"Hajizadeh"}},
	}

	for _, f := range figures {
		id, err := dbConn.CreateFigure(ctx, f.Name, f.Persian, f.Tier, f.Status, f.Role, f.ParentID, f.Aliases)
		if err != nil {
			log.Printf("⚠️ Error seeding %s: %v", f.Name, err)
			continue
		}
		fmt.Printf("✅ Seeded %-25s (ID: %d)\n", f.Name, id)
	}

	fmt.Println("\n✅ Seeding complete.")
}
