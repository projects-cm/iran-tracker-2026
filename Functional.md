# Functional Specification: Iranian Leadership Casualty Tracker

## 1. Project Overview
**Goal**: Build a high-velocity news aggregation site that tracks the status (Alive/Killed) of Iranian political and military figures during the current 2026 conflict.
**Primary Source**: Telegram Channels (User-Account based scraping).
**Core Logic**: Multi-tiered confidence scoring and semantic deduplication.

## 2. Information Architecture (The "Glean" Engine)
- **Source Input**: The system must monitor specific Telegram channels (State Media, IRGC-affiliated, and verified OSINT).
- **Mode of Access**: Use a personal Telegram account (UserBot) to "read" messages, accessing both public and private group data.
- **Backfill Requirement**: On startup, the system must identify the last stored Message ID and retrieve all missed content between the last sync and the present.

## 3. Casualty Tracking Tiers
Figures must be categorized to allow for filtered views:
- **Tier 1**: National Leadership (Supreme Leader, President, SNSC).
- **Tier 2**: IRGC Command (Quds Force, Aerospace, Intelligence).
- **Tier 3**: Strategic Assets (Nuclear scientists, Defense research).
- **Tier 4**: Regional Proxies (IRGC advisors in Lebanon/Syria/Yemen).

## 4. Confidence & Verification Logic
Every report must be assigned a Confidence Level (1-4) based on the source tier and evidence type:
- **Level 1 (Rumor)**: Single unverified local source.
- **Level 2 (Reported)**: Multiple independent OSINT/Aggregator mentions.
- **Level 3 (High Probability)**: Visual evidence of strike + "Missing" status.
- **Level 4 (Confirmed)**: Official state funeral notice or "Martyr" poster from Tier A/B sources.

## 5. Data Refinement & Deduplication
- **Semantic Check**: Use an LLM or embedding-based comparison to identify if a new message is a duplicate of a recent headline.
- **Update Logic**: If a message is a "Confirmation" or "Correction" of an earlier report, the system must update the existing database record rather than creating a new entry.
- **Headline Synthesis**: Transform raw, messy Telegram text (emojis, Hebrew/Persian/English mix) into a clean, professional English headline.

## 6. Future Roadmap (Post-Initial Release)
- **Multi-Source Expansion**: Integrate information from verified news websites and specialized OSINT platforms.
- **Dynamic Scrapers**: Develop adaptive scraping engines for structured and unstructured web data (methodology pending).
- **Cross-Reference Engine**: Automatically flag discrepancies between Telegram reports and official web updates.
