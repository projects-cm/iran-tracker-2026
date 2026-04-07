package dal

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestDB_IsReportProcessed(t *testing.T) {
	// 1. Setup temporary test database
	dbPath := "test_tracker.db"
	defer os.Remove(dbPath)

	db, err := NewDB("file:"+dbPath, "")
	if err != nil {
		t.Fatalf("Failed to init database: %v", err)
	}

	ctx := context.Background()

	// 2. Add a figure for testing
	figID, err := db.CreateFigure(ctx, "Test Leader", "لیدر تست", 1, "Alive", "General", nil, []string{"Test"})
	if err != nil {
		t.Fatalf("Failed to create figure: %v", err)
	}

	// 3. Verify deduplication
	report := Report{
		MessageID:       12345,
		Source:          "Test Source",
		Headline:        "Test Headline",
		RawText:         "Test Text",
		ConfidenceLevel: 90,
		Status:          "Dead",
		PreviousStatus:  "Alive",
		Tier:            1,
		Timestamp:       "2026-04-06T20:00:00Z",
		EntityID:        int(figID),
	}

	// Message should NOT be processed yet (using a unique ID for the test)
	uniqueMsgID := 99999 + int(time.Now().UnixNano()%100000)
	processed, err := db.IsReportProcessed(ctx, uniqueMsgID, "Test Source")
	if err != nil || processed {
		t.Errorf("Expected msg %d to NOT be processed, got: %v, %v", uniqueMsgID, processed, err)
	}

	report.MessageID = uniqueMsgID

	// Add the report
	if err := db.AddReport(ctx, report); err != nil {
		t.Fatalf("Failed to add report: %v", err)
	}

	// Message should now BE processed
	processed, err = db.IsReportProcessed(ctx, 12345, "Test Source")
	if err != nil || !processed {
		t.Errorf("Expected msg 12345 to BE processed, got: %v, %v", processed, err)
	}
}

func TestDB_GetFigures(t *testing.T) {
	dbPath := "test_figures.db"
	defer os.Remove(dbPath)

	db, err := NewDB("file:"+dbPath, "")
	if err != nil {
		t.Fatalf("Failed to init database: %v", err)
	}

	ctx := context.Background()

	// Insert test figures
	_, _ = db.CreateFigure(ctx, "Leader A", "A", 1, "Alive", "Role", nil, nil)
	_, _ = db.CreateFigure(ctx, "Leader B", "B", 2, "Dead", "Role", nil, nil)

	figures, err := db.GetFigures(ctx)
	if err != nil {
		t.Fatalf("Failed to fetch figures: %v", err)
	}

	if len(figures) != 2 {
		t.Errorf("Expected 2 figures, got %d", len(figures))
	}
}
