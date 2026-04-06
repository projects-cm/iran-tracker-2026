# Step 06: Dashboard API (Gin Server)

## Goal:
Implement the Gin-gonic web server and the REST API to serve the tracker data to the frontend.

## Actions:
1.  **API Implementation**:
    -   Create Gin routes for figures and reports.
    -   Implement endpoints: `GET /figures`, `GET /figures/{id}/reports`, `GET /stats`.
2.  **CORS & Static Files**:
    -   Configure CORS for React frontend.
    -   Configure static file serving if necessary.

## Verification:
-   Verify that the FastAPI API server starts correctly and serves the expected data.
-   Check individual API endpoints with terminal tools.
