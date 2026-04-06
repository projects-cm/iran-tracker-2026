# Step 05: Refinement Engine (Gemini 1.5 Pro Processor)

## Goal:
Implement the LLM logic for synthesis, cross-language entity mapping, and deduplication using Gemini 1.5 Pro.

## Actions:
1.  **Refinement Logic**:
    -   Configure **Gemini 1.5 Pro** SDK with the `GEMINI_API_KEY`.
    -   Implement the Refinement Engine to process raw Telegram messages.
2.  **Translation & Synthesis**:
    -   Translate and summarize raw Persian/Hebrew/English text into clean JSON headers.
    -   Extract figure names and verify them using the `aliases` table.
3.  **Deduplication**:
    -   Implement semantic similarity comparison against recent `reports`.

## Verification:
-   Verify that raw messages reach the Refinement Engine and result in clean JSON output.
-   Check that multi-lingual name mapping works as expected.
