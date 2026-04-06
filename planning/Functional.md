# Functional Specification: Iranian Leadership Casualty Tracker

## 1. Project Overview
**Goal**: Build a high-velocity news aggregation site that tracks the status (Alive/Dead/Missing/Wounded) of Iranian political and military figures during the current 2026 conflict.
**Primary Source**: Telegram Channels (User-Account based scraping).
**Core Logic**: Multi-tiered confidence scoring and semantic deduplication via Gemini 1.5 Pro.

## 2. Information Architecture (The "Glean" Engine)
- **Source Input**: The system monitors specific Telegram channels (State Media, IRGC-affiliated, and verified OSINT).
- **Mode of Access**: Use a personal Telegram account (UserBot via `gotd/td`) to "read" messages, accessing both public and private group data.
- **Backfill Requirement**: On startup, the system must identify the last stored Message ID and retrieve all missed content between the last sync and the present.
- **Target Channels**: `@AmitSegal`, `@Abualiexpress`.

## 3. Casualty Tracking Tiers
Figures are categorized to allow for filtered views:
- **Tier 1**: National Leadership (Supreme Leader, President, SNSC).
- **Tier 2**: IRGC Command (Quds Force, Aerospace, Intelligence).
- **Tier 3**: Strategic Assets (Nuclear scientists, Defense research).
- **Tier 4**: Regional Proxies (IRGC advisors in Lebanon/Syria/Yemen).

## 4. Tracked Figures (Current)
| Name | Role | Current Status | Tier |
|------|------|----------------|------|
| Ali Khamenei | Supreme Leader | Dead | 1 |
| Mojtaba Khamenei | Son of Supreme Leader | Missing | 2 |
| Hossein Salami | Commander of IRGC | Alive | 2 |
| Ahmad Vahidi | Secretary of SNSC | Critically Wounded | 2 |
| Esmail Qaani | Commander Quds Force | Presumed Dead | 3 |
| Amir Ali Hajizadeh | Commander Aerospace Force | Dead | 3 |
| Masoud Pezeshkian | President | Missing | 3 |

## 5. Confidence & Verification Logic
Every report must be assigned a Confidence Level (1-4) based on the source tier and evidence type:
- **Level 1 (Rumor)**: Single unverified local source.
- **Level 2 (Reported)**: Multiple independent OSINT/Aggregator mentions.
- **Level 3 (High Probability)**: Visual evidence of strike + "Missing" status.
- **Level 4 (Confirmed)**: Official state funeral notice or "Martyr" poster from Tier A/B sources.

Reports with confidence below 50% are automatically filtered out by the CasualtyService.

## 6. Data Refinement & Deduplication
- **Semantic Check**: Use Gemini 1.5 Pro with structured JSON output to identify if a new message is a duplicate of a recent headline.
- **Update Logic**: If a message is a "Confirmation" or "Correction" of an earlier report, the system must update the existing database record rather than creating a new entry.
- **Headline Synthesis**: Transform raw, messy Telegram text (emojis, Hebrew/Persian/English mix) into a clean, professional English headline.

## 7. Dashboard UI
- **Layout**: Hierarchical node network (React Flow) showing the military/political chain of command.
- **Status Indicators**: Color-coded glow borders (Green=Alive, Orange=Wounded, Red=Dead, Gray=Missing).
- **Header**: Live stats showing target count, alive count, and KIA count.
- **Interaction**: Nodes are fixed in position; the canvas supports zoom and pan.

## 8. Future Roadmap (Post-Initial Release)
- **Multi-Source Expansion**: Integrate information from verified news websites and specialized OSINT platforms.
- **Dynamic Scrapers**: Develop adaptive scraping engines for structured and unstructured web data.
- **Cross-Reference Engine**: Automatically flag discrepancies between Telegram reports and official web updates.
- **Intel Feed Sidebar**: Live scrolling feed of incoming reports with confidence badges.
- **Click-to-Expand Node Detail**: Full audit trail timeline when clicking a figure's node.
