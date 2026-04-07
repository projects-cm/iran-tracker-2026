---
description: How to run and test the Iranian Leadership Tracker on a Windows machine.
---

# Windows Development Workflow

To ensure stability and respect project standards in this Windows environment, follow these rules:

### 1. Package Management (Yarn)
- **Always use `yarn`** for the frontend. Do NOT fallback to `npm`.
- **PowerShell Bypass**: Windows often blocks `yarn.ps1`. Before running any `yarn` command, always execute:
  ```powershell
  Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass
  ```
- Alternatively, you can run yarn via `npx yarn` if the direct script is blocked.

### 2. Port Management
- **Backend**: Runs on port `8080` (via `make run-backend`).
- **Frontend**: Runs on port `5173` (via `make run-frontend`).
- If a port is blocked, find the PID and kill it:
  ```powershell
  netstat -ano | findstr :8080
  taskkill /F /PID <PID>
  ```

### 3. Makefile Authority
- The [Makefile](file:///c:/Users/Shira/Projects/iran-tracker-2026/Makefile) in the root directory is the single source of truth for:
  - `make install`: Downloads Go and Yarn dependencies.
  - `make run-backend`: Starts the Go server and Telegram scraper.
  - `make run-frontend`: Starts the Vite+TS dashboard.
  - `make test`: Runs Go unit tests.

### 4. Database Safety
- Local development uses `.db` files in the root or `backend/` directory.
- Clear cache by deleting `*.db` files if schema changes occur.
