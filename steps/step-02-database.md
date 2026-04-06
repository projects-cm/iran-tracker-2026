# Step 02: Database and Schema Setup

## Goal:
Establish the SQLite database structure to support the casualty tracker and audit trail.

## Actions:
1.  **Database Migration**:
    -   Implement `database/db.go` using a pure-Go SQLite driver.
    -   Create tables: `reports`, `figures`, and `aliases`.
2.  **Schema Support**:
    -   Add logic for status updates (Alive, Killed, Missing, etc.).
    -   Ensure the `reports` table stores all historical entries (the "Audit Trail").

## Verification:
-   Verify that `database.db` exists and has the correct schema.
-   Test simple insertions and retrievals using a test script.
