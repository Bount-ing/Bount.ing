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

Design constraints:
- 3 presets only at creation (`urgent` / `standard` / `patient`). No adversarial custom curves.
- Appreciation capped at 3× initial or sunsets at T+90d. Unbounded growth = hoarding.
- Pool-wide clock (not per-contribution). Simpler, prevents micro-gaming.

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
                                         │ Drips / OC /    │
    Sponsor view                         │ Gitcoin Allo /  │
    (impact ÷ funding)                   │ GH Sponsors     │
                                         └─────────────────┘
                                                ▼
                                         Maintainers / Solvers
```

Pooling happens on existing fiscal-host rails (Drips Network, Open Collective, Gitcoin Allo, GitHub Sponsors). Compliance is handled by the rails. Bount.ing owns the algorithm and the surface, not the money.

**Path A (own pooling rails, SL needed)** is deferred to v3 (2027+) — triggered by platform revenue > €1k/mo recurring and strategic need to hold funds.

Revenue model under Path B (all invoiceable as autónomo):
- **Listing fee** — sponsor pays €5–€20 to surface a bounty on the ranked board
- **Premium API** — agent fleets that consume the queue programmatically
- **Featured placement** — high-trust sponsors get priority surface
- **Cut-on-close** — where integrated rails support kickback (Drips does)

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

- **Maintainer-as-merger conflict of interest.** If a maintainer sponsors their own issue and picks the winning PR, open competition breaks. Mitigation: self-sponsored bounties require third-party validation or are flagged.
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
