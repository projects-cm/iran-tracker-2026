# Step 01: Project Initialization

## Goal:
Set up the core development environment and project structure for the Go backend and React frontend.

## Actions:
1.  **Backend Initialization**:
    -   Create `backend/` directory.
    -   Run `go mod init iranian-tracker`.
    -   Create `main.go` and directory structure for `database`, `scraper`, and `processor`.
2.  **Environment Configuration**:
    -   Create a `.env.example` file with placeholders for `TELEGRAM_API_ID`, `TELEGRAM_API_HASH`, and `GEMINI_API_KEY`.
3.  **Frontend Initialization**:
    -   Create `frontend/` directory using Vite.
    -   Set up the basic React application.

## Verification:
-   Confirm that `go mod tidy` runs without errors.
-   Verify that the Vite dev server starts correctly.
