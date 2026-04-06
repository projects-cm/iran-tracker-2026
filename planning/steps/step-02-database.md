# Step 02: Database and Schema Setup ✅

## Goal:
Establish the SQLite database structure to support the casualty tracker and audit trail.

## Actions:
1.  **Database Migration**:
    -   Implemented `pkg/dal/db.go` using `github.com/mattn/go-sqlite3`.
    -   Created tables: `reports`, `figures`, and `aliases`.
2.  **Schema Support**:
    -   Status tracking (Alive, Dead, Missing, Critically Wounded, Presumed Dead).
    -   `reports` table stores all historical entries (the "Audit Trail").
    -   `figures` table tracks current state with `last_update_id` foreign key.
3.  **Repository Methods**:
    -   `GetFigures(ctx)` — retrieves all tracked figures.
    -   `AddReport(ctx, report)` — transactional insert of report + status update.

## Status: COMPLETE
