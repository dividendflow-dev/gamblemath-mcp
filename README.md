# GambleMath MCP Server

[![Smithery Compatible](https://img.shields.io/badge/Smithery-compatible-orange.svg)](https://smithery.ai)
[![Glama Compatible](https://img.shields.io/badge/Glama-compatible-blue.svg)](https://glama.ai/mcp/servers)
[![Protocol](https://img.shields.io/badge/MCP-streamable--http-green.svg)](https://modelcontextprotocol.io)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Official Model Context Protocol (MCP) server for **[GambleMath](https://gamblemath.org/)**. Provides real-time quantitative gambling analytics, theoretical RTP variance telemetry across 4,362 audited casino slots, official UKGC regulatory register lookups, and closed-form Monte Carlo probability calculators for AI assistants (Claude Desktop, Cursor, Cline, Windsurf, OpenCode).

## Key Resources & Hubs

- **[Online Casino RTP Database (4,362 Slots)](https://gamblemath.org/slot-xray)** — Compare default 96.5% RTP sheets against 94%, 91.5%, and 84% operator settings.
- **[Free HTML5 Slot Demos (1,971 Games)](https://gamblemath.org/demos)** — Instant sandbox access to studio client builds.
- **[UKGC Casino Licence Check (489 Brands)](https://gamblemath.org/autopsy)** — Sister-site graph & account verification against the official UK Gambling Commission register.
- **[Casino Bonus EV Calculator](https://gamblemath.org/bonus-lab)** — Closed-form expected value ($ EV) and completion probability across wagering requirements.
- **[Sharp Sportsbook & Parlay Math](https://gamblemath.org/sharp-desk)** — Multi-leg compounded hold, SGP correlation tax, and Early Cash-Out equity calculators.
- **[Crash Games Probability Lab](https://gamblemath.org/crash-lab)** — Multiplier survival curves and ruin probability.
- **[100-Buy Monte Carlo Simulators](https://gamblemath.org/simulators)** — Empirical variance distributions for bonus buys.
- **[Open Science & CSV API](https://gamblemath.org/open-science)** — Machine-readable datasets (`https://gamblemath.org/api/dataset.csv`) & OpenAPI 3.1 (`https://gamblemath.org/openapi.json`).

---

## Endpoints & Protocols

- **Remote Streamable HTTP:** `https://gamblemath.org/mcp`
- **OpenAPI 3.1 Specification:** `https://gamblemath.org/openapi.json`
- **Catalog Dataset CSV API:** `https://gamblemath.org/api/dataset.csv`
- **Stdio Bridge:** Available via `npx gamblemath-mcp` or `node index.mjs`

---

## Quick Client Setup

### Claude Desktop (`claude_desktop_config.json`)

#### Option A: Direct Remote HTTP (Recommended)
```json
{
  "mcpServers": {
    "gamblemath": {
      "url": "https://gamblemath.org/mcp"
    }
  }
}
```

#### Option B: Stdio CLI Runner
```json
{
  "mcpServers": {
    "gamblemath": {
      "command": "npx",
      "args": ["-y", "gamblemath-mcp"]
    }
  }
}
```

### Cursor (`.cursor/mcp.json`)
```json
{
  "mcpServers": {
    "gamblemath": {
      "url": "https://gamblemath.org/mcp"
    }
  }
}
```

### Cline / Roo Code (`cline_mcp_settings.json`)
```json
{
  "mcpServers": {
    "gamblemath": {
      "url": "https://gamblemath.org/mcp",
      "disabled": false,
      "autoApprove": [
        "audit_slot_rtp",
        "simulate_bonus_ev",
        "audit_parlay_and_cashout",
        "lookup_ukgc_domain"
      ]
    }
  }
}
```

---

## Available MCP Tools

| Tool Name | Required Args | Description |
| :--- | :--- | :--- |
| `audit_slot_rtp` | `slug` (string) | Retrieve technical specs, published RTP profile(s), max win multiplier, and bankroll half-life for any of 4,362 catalog slots (e.g. `sweet-bonanza`, `gates-of-olympus`, `book-of-dead`). |
| `simulate_bonus_ev` | `depositUsd`, `bonusUsd`, `wageringX`, `includeDepositInWager`, `slotRtpPct` | Calculate closed-form Cash EV ($) and completion probability (%) for a casino deposit bonus. |
| `audit_parlay_and_cashout` | `legs`, `stakeUsd`, `isSameGameParlay`, `ticketPayoutUsd`, `cashoutOfferUsd`, `lastLegFairOdds` | Calculate multi-leg parlay compounded bookmaker hold and Early Cash-Out vs manual hedge equity. |
| `lookup_ukgc_domain` | `domain` (string) | Verify a casino domain against the UKGC Public Register snapshot (account #, licensee entity, domain status). |

---

## License & Citation

Released under the **MIT License**.  
Principal Investigator: **Julian Mercer** ([ORCID: 0009-0003-1162-170X](https://orcid.org/0009-0003-1162-170X))  
Harvard Dataverse Archive: [https://dataverse.harvard.edu/dataverse/gamblemath](https://dataverse.harvard.edu/dataverse/gamblemath) (`doi:10.7910/DVN/WVA37E`, `doi:10.7910/DVN/BGYEMI`)
