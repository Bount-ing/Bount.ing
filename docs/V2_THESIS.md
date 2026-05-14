# Bount.ing v2 — thesis

> ⚠️ The code in this repo is **v1 (2024)** — single-sponsor, claim-based, static-bounty marketplace operated by Bount.ing S.L. (now dissolved). v2 is a mechanism-design rebuild, captured in this doc and the OPEN_QUESTIONS deck. Don't assume v1 code reflects v2 design.

## Thesis

**Use cashflow to create a priority, impact-weighted to-do list of open source development.**

Bount.ing v2 is the impact-weighted ranking + routing layer for OSS work. Cashflow drives priority. Two views of the same algorithm:

- **Solvers** see issues ranked by bounty (pay-first queue).
- **Sponsors** see issues ranked by `impact ÷ current_funding` ("your euros go far here") — surfaces underfunded high-impact work, prevents single-sponsor capture.

Results-based: pay on close (PR merged, issue resolved), not on attempt or time.

## What separates v2 from the graveyard

Three mechanism-design innovations. None of Gitcoin / IssueHunt / Bountysource / v1-Bount.ing combined these.

### 1. Multi-sponsor bounties

Many sponsors pool funds against the same issue.

- Pool view: sponsor list + running total. Social proof — "47 sponsors, €1,240 pooled" reads differently from a flat number.
- Withdrawal locked once contributed (else gameable).
- **Anti-coordination unlock:** queue position hidden until a minimum-viable-bounty threshold is reached; prevents whale-waiting / free-riding.

**Solves:** funding-supply problem. Small sponsors who care but can't justify €500 solo can pool €10 each.

### 2. Time-sensitive bounties

Bounty value follows a curve, not a flat number.

| Curve | Effect | Use case |
|---|---|---|
| **Decay** (€ drops over time) | Forces solver urgency. Dutch-auction shape. | Security patches, business-critical bugs, time-sensitive launches |
| **Appreciation** (€ grows over time) | Solves the neglected-good-issue problem — eventually stale issues become attractive | Long-tail OSS maintenance, refactors, "nice but not urgent" features |

**Three presets** (locked 2026-05-14):

| Preset | Shape | Floor / cap |
|---|---|---|
| `urgent` | Linear decay over 30d | Floor = 25% of initial |
| `standard` | Flat 30d → linear decay 60d to floor | Floor = 25% of initial |
| `patient` | Linear appreciation | Cap = 3× initial OR T+90d, whichever first |

Design constraints:
- 3 presets only at creation. No custom curves (predatory shapes would game solvers).
- Floor is a **percentage** of initial, not absolute (scales with bounty size).
- Appreciation cap is **pool-lifetime**, not reset on new sponsors (else gameable).
- Pool-wide clock (not per-contribution).
- Hard expire at T+180d → unresolved pool re-routes per sponsor's pre-set destination policy.

**Implementation against the integrated rail (OC at v2 launch):**
- **Decay:** decayed portion is *re-routed*, not refunded (refunds = fund-custody = autónomo poison). Sponsor picks destination at creation: bount.ing platform fund / community treasury / other underfunded bounties.
- **Appreciation:** sponsor pre-funds the curve_max upfront. Solver receives curve-value-at-close. Difference re-routes per sponsor's pre-set destination policy. Sponsor knows max liability at creation.
- **v2-launch fallback:** if full euro-curves aren't engineering-ready, ship flat + curve-as-ranking-only (curve affects queue position but not euro value). Full euro-curves at v2.5. See OPEN_QUESTIONS.md Q8.

**Solves:** static-bounty staleness. None of the graveyard platforms have this.

### 3. Open competition (no claim)

Anyone can attempt any open bounty without permission. First merged PR wins.

- Replaces the claim-based assignment pattern that kills bounty platforms ("I'm working on this" → 3-week lock → nothing ships → issue stale). Note: v1 in this repo was claim-based — see `PROCESS.md`. **v2 explicitly drops claims.**
- Agent-fleet friendly: parallel attempts without coordination overhead.
- Optional safety valve, post-v1 if needed: staked claim — €5 deposit for 48h exclusivity, refunded on merge, forfeit on no-show (added to pool).

**Solves:** participation friction + agent compatibility.

### The combination is the product

Multi-sponsor + appreciating bounty + open competition = **self-amplifying market for OSS maintenance.** Underfunded issues accrue both funds and value over time, attracting more sponsors AND more solvers. Solved issues clear the queue.

## Economic model

| Mechanic | Signal | Example |
|---|---|---|
| **Crowdfunded bounty** | Many small contributions aggregate → bounty size = breadth of impact | 500 devs each chip in €10 → €5k pool |
| **Business-critical bounty** | One large contribution from one stakeholder → bounty size = severity to that stakeholder | A scaleup loses €50k/day to a memory leak → drops €15k single bounty |

Bounty size encodes impact for the funded queue (market-priced). The platform's `impact_score` is the **discovery layer** for unfunded issues — sponsor-side view of "what should be funded but isn't yet."

## Architecture — autónomo-compatible (Path B)

**Bount.ing v2 is the ranking + routing layer. It does NOT hold third-party funds.**

```
┌──────────────────────────────────────────────────────────────┐
│                       bount.ing                              │
│  ┌──────────────┐  ┌────────────┐  ┌──────────────────────┐  │
│  │   Ranking    │  │   Impact   │  │  Routing / API to    │  │
│  │  algorithm   │◄─┤  algorithm │  │  external rails      │  │
│  └──────┬───────┘  └────────────┘  └──────────┬───────────┘  │
└─────────┼─────────────────────────────────────┼──────────────┘
          ▼                                     ▼
    Solver view                          Pooling rails
    (ranked queue)                       ┌─────────────────┐
                                         │ OC → GH Sponsors│
    Sponsor view                         │ → Drips →       │
    (impact ÷ funding)                   │ Taler (phased)  │
                                         └─────────────────┘
                                                ▼
                                         Maintainers / Solvers
```

Pooling happens on existing fiscal-host rails (Open Collective, GitHub Sponsors, Drips Network, GNU Taler). Compliance is handled by the rails. Bount.ing owns the algorithm and the surface, not the money.

**Provider-agnostic by design.** Bount.ing v2 is a federation layer over funding rails — adapter pattern, one per rail. Sponsor picks the rail when funding a bounty. Mirrors Lethe's multi-provider LLM pattern.

**Rail integration sequence** (resolved 2026-05-14, see OPEN_QUESTIONS.md Decisions Log Q1):

1. **Open Collective** — v2 launch. Fiscal-host model matches existing autónomo invoicing flow; zero new tax categories.
2. **GitHub Sponsors** — v2.1. Admin-trivial; mechanism-poor (single-sponsor only).
3. **Drips Network** — v2.5, after autónomo books stabilize. Mechanism-fit (streaming = appreciation curve) but adds crypto-as-business-income reporting overhead.
4. **GNU Taler** — when EU exchange ecosystem matures and/or NLnet grant capture funds protocol-extension engineering.

**Path A (own pooling rails, SL needed)** is deferred to v3 (2027+) — triggered by platform revenue > €1k/mo recurring and strategic need to hold funds.

**Bount.ing as a hosted OC project.** The platform itself is hosted on Open Collective Europe (or equivalent OCE fiscal host). Listing fees + premium API + featured-placement revenue flows into the bount.ing OC project. Mía draws monthly autónomo invoice against project balance. OCE handles VAT + transparency ledger + contributor compensation. Host fee 5–10% — vastly cheaper than €6k SL setup + cuota societaria.

Revenue model under Path B (all invoiceable as autónomo via OCE):
- **Featured placement** — sponsor pays for top-of-queue visibility within their tag/repo
- **Premium API** — agent fleets, subscription-based
- **Rail kickback** — Drips natively; OC via host-fee-share if negotiable
- **Listing: free** — entry friction breaks the multi-sponsor pooling mechanic (differentiator #1)

## Phased roadmap

| Phase | Scope | Status / Trigger |
|---|---|---|
| **v0 — calibration set** | Hand-curated `bounties.yaml` on Mía's personal stack. Functions as labelled training data for the impact algorithm. | ✅ Seeded 2026-05-14, 19 bounties. Lives in `miam-knowledge-base/bounties.yaml`. |
| **v0.5 — derive the formula** | Reverse-engineer the impact algorithm from v0 hand-weights. Target: function over observable signals (stars, deps, 👍, blocks-graph, security labels) that reproduces ~80% of hand-weights. | Next. ETA 2026-06. |
| **v1 — public read-only board** | Dashboard at bount.ing. Apply impact algorithm to curated external OSS repos (5–10). No payments. | ETA 2026 Q3. |
| **v2 — marketplace integration** | Open submission + rail integration. Three mechanisms (multi-sponsor / time-curve / open competition) enforced at UX layer; fund-holding remains on the integrated rail. | ETA 2026 Q4 / 2027 Q1. |
| **v3 — own the rails** | SL re-incorporation + Stripe Connect / escrow / native pooling. | 2027+, revenue-gated. |

## Non-goals (load-bearing)

- **No payment intermediation before v3.** Funds never sit in a bount.ing-owned account until the SL exists and is justified by revenue. Path B until then.
- **No claim-based assignment.** Open competition is a feature, not a bug. If solvers want safety, optional staked claim — never mandatory pre-work locking.
- **No multipliers on already-funded issues for ranking.** Funded queue is sorted by bounty (market-priced). Impact is for discovery, not distortion.
- **No flat-bounty fallback.** Time-curve is the differentiator. 3 presets, no opt-out.

## Risks

- **Maintainer-as-merger conflict of interest.** If a maintainer sponsors their own issue and picks the winning PR, open competition breaks. Mitigation (locked 2026-05-14): self-sponsored bounties are **flagged** in the UI ("self-sponsored — sponsor is also merge authority") at v1; escalate to third-party validator only if abuse pattern observed.
- **Wasted effort from open competition.** N solvers attempt, 1 wins. Mitigation: show "active attempt count" per issue; optional staked claim as v1+ safety valve.
- **Decay-curve adversarial design.** Predatory curves (€100 → €1 in 24h) game solvers. Mitigation: 3 presets, no custom.
- **GitHub lock-in.** "Merged PR = closed" is GitHub-shaped. Locks out GitLab / Codeberg. Acceptable v1 cost; revisit at v2.
- **Rail dependency.** Drips/OC policy changes break v2 features. Mitigation: integrate multiple rails from v2.
- **Single-user calibration risk.** v0 yaml is one person's hand-weights — may not generalize. Mitigation: at v0.5, invite 1–2 maintainers to score the same issues, check correlation.

## Cross-references

- Full project notes & history: [miam-knowledge-base/docs/life/works/projects/bount-ing.md](../../Miam/miam-knowledge-base/docs/life/works/projects/bount-ing.md)
- v0 calibration set: [miam-knowledge-base/bounties.yaml](../../Miam/miam-knowledge-base/bounties.yaml)
- Live decisions: [OPEN_QUESTIONS.md](OPEN_QUESTIONS.md)
- v1 process (claim-based, dissolved): [PROCESS.md](PROCESS.md)
- v1 tests (Stripe Connect / GitHub OAuth, archival): [TESTS.md](TESTS.md)
- 2024 v1 terms (archival, not binding — SL dissolved): `/home/mia/Downloads/Bount.ing-terms-2025-02-04.pdf`
