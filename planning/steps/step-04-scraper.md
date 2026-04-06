# Step 04: Telegram Scraper (Glean Engine) ✅

## Goal:
Configure and implement the Telegram UserBot (`gotd`) to monitor target channels.

## Actions:
1.  **MTProto Implementation** (`pkg/service/scraper.go`):
    -   Configured `gotd/td` client with `NewScraperService(client)`.
    -   `resolveUsername()` — resolves channel usernames to `InputPeer`.
    -   `fetchRecentMessages()` — fetches latest 10 messages via `messages.getHistory`.
    -   `processMessages()` — filters service/empty messages, logs raw text.
2.  **Target Channels**: **@AmitSegal** and **@Abualiexpress**.
3.  **Human Mimicry Logic**:
    -   Randomized jitter delays (2-10s between fetches) via `scrapeChannelRoutine()`.
    -   Each channel runs in its own goroutine.

## Pending:
-   Telegram API credentials (`TELEGRAM_API_ID`, `TELEGRAM_API_HASH`) not yet provided.
-   Integration with ProcessorService (forwarding raw text to Gemini).

## Status: CODE COMPLETE — awaiting credentials
