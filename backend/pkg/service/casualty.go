package service

import (
	"context"
	"fmt"
	"log"

	"iranian-tracker/backend/pkg/dal"
)

// CasualtyService handles business logic related to casualty statuses and reporting
type CasualtyService struct {
	db *dal.DB
}

// NewCasualtyService creates a new CasualtyService instance
func NewCasualtyService(db *dal.DB) *CasualtyService {
	return &CasualtyService{
		db: db,
	}
}

// GetTrackerDashboard returns the aggregated figures and latest reports for the UI
func (s *CasualtyService) GetTrackerDashboard(ctx context.Context) ([]dal.Figure, error) {
	figures, err := s.db.GetFigures(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch figures: %w", err)
	}
	return figures, nil
}

// ProcessNewReport records a new intel update and changes the entity status if necessary
func (s *CasualtyService) ProcessNewReport(ctx context.Context, ext *ExtractionResult, originalText, source string, msgID int) error {
	// If the model returns 0, it means no relevant figure was identified.
	if ext.EntityID <= 0 {
		return nil
	}

	figures, err := s.db.GetFigures(ctx)
	if err != nil {
		return err
	}

	var currentFig *dal.Figure
	for _, f := range figures {
		if f.ID == ext.EntityID {
			currentFig = &f
			break
		}
	}

	if currentFig == nil {
		return fmt.Errorf("entity ID %d not found in database", ext.EntityID)
	}

	// We only upgrade the confidence status logically: e.g., if already Dead,
	// we shouldn't revert to missing unless there's an overwhelming new report
	// For simplicity in this iteration, we just accept the latest report as truth
	// if confidence > 50

	if ext.Confidence < 50 {
		log.Printf("Ignoring low confidence (%d) report for %s: %s", ext.Confidence, currentFig.CanonicalName, ext.Status)
		return nil
	}

	newReport := dal.Report{
		MessageID:       msgID,
		Source:          source,
		Headline:        ext.Headline,
		RawText:         originalText,
		ConfidenceLevel: ext.Confidence,
		Status:          ext.Status,
		PreviousStatus:  currentFig.CurrentStatus,
		Tier:            currentFig.Tier,
		Timestamp:       fmt.Sprintf("%v", ctx.Value("timestamp")), // Simplified for now
		EntityID:        ext.EntityID,
	}

	err = s.db.AddReport(ctx, newReport)
	if err != nil {
		return fmt.Errorf("failed to save report for %s: %w", currentFig.CanonicalName, err)
	}

	log.Printf("Updated %s status to %s based on %s.", currentFig.CanonicalName, ext.Status, source)
	return nil
}
