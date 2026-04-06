---
description: how to run the application locally
---

This workflow explains how to start the backend and frontend correctly on this machine, including the necessary environment paths.

### 🗄️ 0. Prerequisites (Windows Pathing)
If `go`, `npm`, or `yarn` are missing from the current terminal's `$env:PATH`, use these prefixes:

```powershell
$env:PATH += ";C:\Program Files\Go\bin;C:\Program Files\Git\cmd;C:\Program Files\nodejs"
```

### 🚀 1. Start the Backend
1. Open a terminal in the `backend/` directory.
2. Ensure `.env` is present in the root.
// turbo
3. Run: `go run ./cmd/server`
   - The API will be available at `http://localhost:8080/api/v1`

### 🎨 2. Start the Frontend
1. Open a terminal in the `frontend/` directory.
// turbo
2. Run: `& "C:\Users\Shira\AppData\Roaming\npm\yarn.cmd" dev`
   - The Dashboard will be available at `http://localhost:5173` (or the next available port)

### 🛠️ 3. Seed the Database
If you need to reset or re-seed the data:
// turbo
1. Run: `cd backend; go run cmd/admin/seed.go`
