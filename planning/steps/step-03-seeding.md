# Step 03: Entity Seeding and Multi-lingual Resolution ✅

## Goal:
Pre-populate the database with the core Iranian leadership entities and their known multi-lingual aliases.

## Actions:
1.  **Initial Figures** (seeded in `pkg/dal/db.go` via `SeedInitialData`):
    -   **Ali Khamenei** — Former Supreme Leader (Dead).
    -   **Mojtaba Khamenei** — Son of Supreme Leader (Missing).
    -   **Masoud Pezeshkian** — President (Missing).
    -   **Ahmad Vahidi** — Secretary of SNSC (Critically Wounded).
    -   **Hossein Salami** — Commander of IRGC (Alive).
    -   **Esmail Qaani** — Commander Quds Force (Presumed Dead).
    -   **Amir Ali Hajizadeh** — Commander Aerospace Force (Dead).
2.  **Alias Mapping**:
    -   Persian and English aliases inserted into the `aliases` table for multilingual lookup.

## Status: COMPLETE
