# Technical Specification: Iranian Leadership Casualty Tracker

This document outlines the technical architecture and implementation details for the casualty tracker.

## 1. System Architecture
The system is built on a **Go (Golang)** server that coordinates the following components:
- **Glean Engine (Input)**: A `gotd` (Golang MTProto) based UserBot running as a persistent background process.
- **Refinement Engine (Logic)**: An LLM-powered module using the **Google Generative AI Go SDK** (Gemini 1.5 Pro) for synthesis and deduplication.
- **Dashboard API (Output)**: RESTful endpoints served by **Chi** router to the React frontend.

## 2. Technical Implementation Details

### 2.1 Backend Server Stack
- **Web Server**: **go-chi/chi** (lightweight router, replaces Gin).
- **Telegram Client**: **gotd/td** (Go implementation of Telegram MTProto).
- **LLM**: **google/generative-ai-go** (Gemini 1.5 Pro with structured JSON output).
- **Database**: **SQLite** (using `github.com/mattn/go-sqlite3`).
- **Language**: Go 1.26+
- **Config**: **godotenv** for environment variable management.

### 2.2 Backend Architecture (3-Layer)
```
cmd/server/
├── main.go        # Entry point, graceful shutdown
├── compose.go     # Dependency injection, route registration
└── clients.go     # External service initialization (DB, Telegram, Gemini)

pkg/
├── handler/       # Presentation layer (HTTP handlers)
│   └── casualty.go
├── service/       # Business logic layer
│   ├── scraper.go     # Telegram channel monitoring with jitter
│   ├── processor.go   # Gemini structured JSON extraction
│   └── casualty.go    # Orchestration, confidence checks, status updates
└── dal/           # Data Access Layer
    └── db.go          # SQLite schema, migrations, seeds, queries
```

### 2.3 LLM & Refinement Engine
- **Provider**: **Gemini 1.5 Pro**.
- **Output**: Structured JSON via `ResponseMIMEType: "application/json"` with enforced schema.
- **Entity mapping**: Multilingual resolution (Persian/Hebrew/English) via Gemini.
- **Fields**: `entityId`, `confidence` (0-100), `status`, `headline`.

### 2.4 Database Schema (SQLite)
- **Audit Trail Decision**: The `reports` table **stores all historical status updates** for each figure, allowing a full timeline view (e.g., Targeted -> Wounded -> Killed).

```sql
-- Audit Trail: EVERY update is stored here
CREATE TABLE reports (
    message_id INTEGER PRIMARY KEY,
    source TEXT NOT NULL,
    headline TEXT NOT NULL,
    raw_text TEXT,
    confidence_level INTEGER,
    status TEXT, -- Alive, Killed, Missing, Wounded, Target
    previous_status TEXT, -- For tracking history
    tier INTEGER,
    timestamp TEXT NOT NULL,
    entity_id INTEGER,
    FOREIGN KEY(entity_id) REFERENCES figures(id)
);

-- Current State
CREATE TABLE figures (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    canonical_name TEXT UNIQUE NOT NULL, -- English (e.g., Ali Khamenei)
    persian_name TEXT, -- (e.g., علی شمخانی)
    tier INTEGER,
    current_status TEXT,
    last_update_id INTEGER,
    FOREIGN KEY(last_update_id) REFERENCES reports(message_id)
);

-- Multilingual Map
CREATE TABLE aliases (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    entity_id INTEGER,
    alias TEXT UNIQUE NOT NULL,
    FOREIGN KEY(entity_id) REFERENCES figures(id)
);
```

### 2.5 Scraper Implementation
- **gotd/td**: Asynchronous channel monitoring with goroutines per target channel.
- **Pagination**: Sequential message fetching (10 per batch) via `messages.getHistory`.
- **Human Mimicry**: Randomized jitter delays (2-10s between fetches) to avoid bans.

## 3. Frontend Dashboard
- **Stack**: Vite + React + Tailwind CSS v4 + React Flow + Lucide Icons.
- **Package Manager**: Yarn.
- **Aesthetic**: Premium dark-mode glassmorphism with neon status glow effects.
- **Layout**: Hierarchical node network (React Flow) reflecting military/political structure.
- **Status Colors**: Green (Alive), Orange (Critically Wounded), Red (Dead/Presumed Dead), Gray (Missing).

## 4. Git Workflow
- **Branching**: Feature branches (e.g., `feat/frontend-infrastructure`).
- **Commits**: Conventional Commits (`feat:`, `fix:`, `chore:`).
- **Review**: Pull Requests on GitHub before merging to `main`.
- **Repository**: Private @ `chaim0m/iran-tracker-2026`.
