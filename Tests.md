# Testing Plan: Iranian Leadership Casualty Tracker

This document outlines the validation and testing strategy to ensure the reliability and accuracy of the tracker.

## 1. Unit Testing (Go Backend)
- **Database Logic**: Verify migrations and CRUD operations for `reports`, `figures`, and `aliases`.
- **Entity Resolution**: Test the LLM-assisted name mapping logic with mock Persian/Hebrew/English inputs.
- **Synthesis Engine**: Mock LLM responses to verify headline transformation and status extraction.

## 2. Integration Testing
- **Telegram Client**: Verify that the MTProto connection (`gotd`) can successfully join and listen to the target channels (`AmitSegal`, `Abualiexpress`).
- **Pagination Logic**: Test the "Glean" engine's ability to fetch and process history in batches of 100 without exceeding rate limits.
- **API to DB**: Ensure the Gin server correctly serves data from the SQLite database to the frontend.

## 3. End-to-End (E2E) Testing
- **Visual Validation**: Confirm the React dashboard accurately reflects the latest "Confirmed" status of high-value targets.
- **Mock Data Flow**: Inject a mock "Targeted" report into the database and verify it triggers a dashboard update and appears in the target's timeline.

## 4. Manual Verification
- **Human Mimicry**: Verify that logs show staggered delays (10-30s per message) during the backfilling process.
- **Multimodal Visuals**: Manually check that "Martyr" posters are correctly analyzed and summarized by the LLM.
