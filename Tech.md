# Technical Specification: Iranian Leadership Casualty Tracker (Go Edition)

This document outlines the technical architecture and implementation details for the casualty tracker, using a Go-based backend.

## 1. System Architecture
The system is built on a **Go (Golang)** server that coordinates the following components:
- **Glean Engine (Input)**: A `gotd` (Golang MTProto) based UserBot running as a persistent background process.
- **Refinement Engine (Logic)**: An LLM-powered module using the **Google Generative AI Go SDK** (Gemini 1.5 Pro) for synthesis and deduplication.
- **Dashboard API (Output)**: RESTful endpoints served by **Gin** to the React frontend.

## 2. Technical Implementation Details

### 2.1 Backend Server Stack
- **Web Server**: **Gin-gonic/gin**.
- **Telegram Client**: **gotd/td** (Go implementation of Telegram MTProto).
- **LLM**: **google/generative-ai-go** (Gemini 1.5 Pro).
- **Database**: **SQLite** (using `modernc.org/sqlite` or `github.com/mattn/go-sqlite3`).
- **Language**: Go 1.21+

### 2.2 LLM & Refinement Engine
- **Provider**: **Gemini 1.5 Pro**.
- **Entity mapping**: Multilingual resolution (Persian/Hebrew/English) via Gemini.
- **Synthesis**: Structured output for clean headlines and status updates.

### 2.3 Database Schema (SQLite)
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

### 2.4 Scraper Implementation
- **gotd/td**: Asynchronous event handling for real-time messages.
- **Pagination**: Sequential message fetching (100 per batch) with randomized delays (5s batch / 10-30s message).

## 3. Frontend Dashboard
- **Stack**: Vite + React + Vanilla CSS.
- **Aesthetic**: Premium Glassmorphism.
