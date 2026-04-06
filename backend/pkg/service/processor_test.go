package service

import (
	"encoding/json"
	"testing"
)

func TestExtractionResult_Unmarshaling(t *testing.T) {
	// 1. Simulate a successful JSON response from Gemini
	rawJSON := `{
		"entityId": 6,
		"confidence": 95,
		"status": "Dead",
		"headline": "Esmail Qaani confirmed killed in the latest strike."
	}`

	var result ExtractionResult
	if err := json.Unmarshal([]byte(rawJSON), &result); err != nil {
		t.Fatalf("Failed to unmarshal Gemini JSON: %v", err)
	}

	// 2. Verify fields
	if result.EntityID != 6 {
		t.Errorf("Expected EntityID 6, got %d", result.EntityID)
	}
	if result.Status != "Dead" {
		t.Errorf("Expected Status 'Dead', got '%s'", result.Status)
	}
	if result.Confidence != 95 {
		t.Errorf("Expected Confidence 95, got %d", result.Confidence)
	}
}

func TestExtractionResult_InvalidJSON(t *testing.T) {
	// 1. Simulate invalid JSON
	rawJSON := `{ "invalid": "json"`

	var result ExtractionResult
	if err := json.Unmarshal([]byte(rawJSON), &result); err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
