# Step 04: Telegram Scraper (Glean Engine)

## Goal:
Configure and implement the Telegram UserBot (`gotd`) to monitor Amit Segal and Abu Ali English.

## Actions:
1.  **MTProto Implementation**:
    -   Configure `gotd/td` with the `TELEGRAM_API_ID` and `TELEGRAM_API_HASH`.
    -   Connect to the target channels: **@AmitSegal** and **@Abualiexpress**.
2.  **Pagination and Backfill**:
    -   Implement paged history fetching for history since the last stored `message_id`.
3.  **Human Mimicry Logic**:
    -   Staggered delays (10-30s per message) to prevent account flagged/throttled.

## Verification:
-   Verify that raw messages reach the backend logs.
-   Check that history backfilling logic and pagination work as expected.
