# Step 03: Entity Seeding and Multi-lingual resolution

## Goal:
Pre-populate the database with the core Iranian leadership entities and their known multi-lingual aliases.

## Actions:
1.  **Initial Figures**:
    -   Insert **Ali Khamenei** (Former Supreme Leader).
    -   Insert **Mojtaba Khamenei** (Successor/Current Supreme Leader).
    -   Insert **Masoud Pezeshkian** (President).
    -   Insert **Ahmad Vahidi** (IRGC Commander).
    -   Insert **Hossein Salami** (Former IRGC Commander).
    -   Insert **Ebrahim Raisi** (Former President - Audit Trail context).
2.  **Alias Mapping**:
    -   Translate and insert Persian/English/Hebrew names for these entities to the `aliases` table.

## Verification:
-   Verify that the `figures` and `aliases` tables are correctly populated in the SQLite database.
-   Test simple search queries to ensure multi-lingual lookups work as expected.
