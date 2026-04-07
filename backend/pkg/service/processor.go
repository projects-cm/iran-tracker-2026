package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	// "iran-tracker/pkg/dal"
)

// ProcessorService handles transforming raw text into structured JSON via Gemini
type ProcessorService struct {
	client *genai.Client
	model  *genai.GenerativeModel
	// dal    *dal.DB
}

// ExtractionResult represents the desired JSON output from the LLM
type ExtractionResult struct {
	EntityID   int    `json:"entityId"`
	Confidence int    `json:"confidence"`
	Status     string `json:"status"` // "Alive", "Missing", "Critically Wounded", "Dead", "Presumed Dead"
	Headline   string `json:"headline"`
}

// NewProcessorService initializes the Gemini model connection
func NewProcessorService(client *genai.Client) *ProcessorService {
	// Configure the model to return JSON
	model := client.GenerativeModel("gemini-flash-latest")
	model.ResponseMIMEType = "application/json"
	
	// Create the JSON schema for structured output
	schema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"entityId": {
				Type: genai.TypeInteger,
				Description: "The ID of the tracked figure: " +
					"1: Ali Khamenei (Supreme Leader), " +
					"2: Mojtaba Khamenei (Son/Possible Successor), " +
					"3: Masoud Pezeshkian (President), " +
					"4: Ahmad Vahidi (SNSC), " +
					"5: Hossein Salami (IRGC), " +
					"6: Esmail Qaani (Quds Force), " +
					"7: Amir Ali Hajizadeh (Aerospace Force). " +
					"Use 0 if no figure is mentioned.",
			},
			"confidence": {
				Type:        genai.TypeInteger,
				Description: "Confidence score from 0 to 100 based on source reliability (e.g. state TV vs rumors)",
			},
			"status": {
				Type:        genai.TypeString,
				Description: "One of: Alive, Missing, Critically Wounded, Dead, Presumed Dead",
			},
			"headline": {
				Type:        genai.TypeString,
				Description: "A short 1-sentence English summary of the report",
			},
		},
		Required: []string{"entityId", "confidence", "status", "headline"},
	}
	model.ResponseSchema = schema

	return &ProcessorService{
		client: client,
		model:  model,
	}
}

// ProcessRawText sends the text to Gemini and parses the structured response
func (s *ProcessorService) ProcessRawText(ctx context.Context, text string, sourceName string) (*ExtractionResult, error) {
	prompt := fmt.Sprintf(`
	You are an expert intelligence analyst tracking Iranian leadership casualties.
	Analyze the following news report from '%s' and extract the status of any key figures.
	
	Report Text:
	"%s"
	`, sourceName, text)

	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("gemini generation failed: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("empty response from model")
	}

	part := resp.Candidates[0].Content.Parts[0]
	jsonText, ok := part.(genai.Text)
	if !ok {
		return nil, fmt.Errorf("expected text part from model")
	}

	var result ExtractionResult
	if err := json.Unmarshal([]byte(jsonText), &result); err != nil {
		log.Printf("Failed to unmarshal JSON from model. Raw JSON: %s", jsonText)
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	return &result, nil
}
