# Step 06: Dashboard API (Chi Server) ✅

## Goal:
Implement the Chi web server and the REST API to serve the tracker data to the frontend.

## Actions:
1.  **API Implementation** (`pkg/handler/casualty.go` + `cmd/server/compose.go`):
    -   Chi router with `middleware.Logger`, `middleware.Recoverer`, `middleware.Timeout`.
    -   Endpoints:
        -   `GET /health` — returns `{"status": "up"}`.
        -   `GET /api/v1/figures` — returns all tracked figures (currently serves dummy data).
        -   `GET /api/v1/figures/{id}/reports` — planned.
        -   `GET /api/v1/stats` — planned.
2.  **Dependency Injection** (`compose.go`):
    -   Wires DAL → Service → Handler layers.
    -   All dependencies initialized from `Clients` struct.
3.  **CORS & Proxy**:
    -   Vite dev server proxies `/api` requests to Go backend on port 8080.

## Pending:
-   Replace dummy data with live DAL queries.
-   Add reports and stats endpoints.

## Status: CODE COMPLETE — serving dummy data
