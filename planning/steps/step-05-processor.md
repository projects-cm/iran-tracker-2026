# Step 05: Refinement Engine (Gemini 1.5 Pro Processor) ✅

## Goal:
Implement the LLM logic for synthesis, cross-language entity mapping, and deduplication using Gemini 1.5 Pro.

## Actions:
1.  **Refinement Logic** (`pkg/service/processor.go`):
    -   Configured **Gemini 1.5 Pro** with `ResponseMIMEType: "application/json"`.
    -   Enforced structured output via `ResponseSchema` with fields:
        -   `entityId` (int) — mapped to tracked figures.
        -   `confidence` (int, 0-100) — based on source reliability.
        -   `status` (string) — one of: Alive, Missing, Critically Wounded, Dead, Presumed Dead.
        -   `headline` (string) — clean 1-sentence English summary.
    -   `ProcessRawText(ctx, text, sourceName)` — sends prompt to Gemini and parses JSON response.
2.  **Deduplication**:
    -   Planned: semantic similarity comparison against recent `reports` (not yet implemented).

## Pending:
-   Gemini API key (`GEMINI_API_KEY`) not yet provided.
-   Deduplication logic against existing reports.

## Status: CODE COMPLETE — awaiting API key
