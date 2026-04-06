# Testing Plan: Iranian Leadership Casualty Tracker

This document outlines the validation and testing strategy to ensure the reliability and accuracy of the tracker.

## 1. Unit Testing (Go Backend)
- **Database Logic**: Verify migrations and CRUD operations for `reports`, `figures`, and `aliases` in `pkg/dal/db.go`.
- **Entity Resolution**: Test the Gemini-assisted name mapping logic with mock Persian/Hebrew/English inputs.
- **Processor Engine**: Mock Gemini responses to verify JSON schema extraction and status parsing in `pkg/service/processor.go`.
- **Casualty Service**: Test confidence threshold logic (reports < 50% confidence are ignored).

## 2. Integration Testing
- **Telegram Client**: Verify that the MTProto connection (`gotd`) can successfully resolve and listen to the target channels (`AmitSegal`, `Abualiexpress`).
- **Pagination Logic**: Test the scraper's ability to fetch message history in batches of 10 without exceeding rate limits.
- **API to DB**: Ensure the Chi server correctly serves data from the SQLite database via `/api/v1/figures`.
- **Frontend Proxy**: Verify that Vite's proxy configuration correctly forwards `/api` requests to the Go backend on port 8080.

## 3. End-to-End (E2E) Testing
- **Visual Validation**: Confirm the React Flow dashboard accurately renders the hierarchical node network with correct statuses.
- **Mock Data Flow**: Inject a mock report into the database and verify it triggers a node status change and appears in the figure's timeline.
- **Status Glow**: Verify that each status (Alive, Dead, Missing, Critically Wounded, Presumed Dead) renders the correct border color and glow effect.

## 4. Manual Verification
- **Human Mimicry**: Verify that logs show randomized jitter delays (2-10s) during channel polling.
- **Node Layout**: Confirm that the tier-based layout correctly positions Tier 1 at the top, Tier 2 in the middle, and Tier 3 at the bottom.
- **PR Workflow**: Verify that conventional commits and PRs are created correctly via `gh`.
