# Open Questions

Live decision deck for Bount.ing v2. Decisions move to [V2_THESIS.md](V2_THESIS.md) and the Decisions Log (bottom of this file) as they're settled.

---

## Q6 — How is `impact_score` defined concretely?

**Status:** Open — this is the v0.5 deliverable.

`impact_score` is a function over observable signals. v0.5 will derive it from the v0 hand-weights in `bounties.yaml`. **First signal set to try (locked 2026-05-14):**

| Signal | Source | Cost to compute | Weight intuition |
|---|---|---|---|
| Repo star count | GitHub API | trivial | log-scaled, modest |
| Issue 👍 reactions | GitHub API | trivial | linear, modest |
| Dependency-graph reach | npm / PyPI / crates / Cargo.toml | medium | log-scaled, high (network effect) |
| Blocks-downstream signal | issue body parsing (`Closes #X`, references) | hard but worth it | linear, high |
| Security relevance | label heuristics + CVE feed cross-reference | medium | binary boost |
| Maintainer responsiveness | GitHub API (time-to-first-reply) | medium | tiebreaker, modest |

**v0.5 deliverable:** fit weights against `bounties.yaml`, target ≥80% reproduction of hand-weights. Iterate. Move full formula to V2_THESIS.md when stable.

---

## Q1.7 — Host bount.ing itself on Open Collective Europe?

**Status:** Settled lean (locked 2026-05-14), execution-detail open.

**Decision:** Yes. Host bount.ing on Open Collective Europe (or equivalent OCE fiscal host under the OC umbrella).

**Why:**
- Cleaner legal shell than re-opening Bount.ing S.L. during the autónomo bleed.
- Listing fees + premium API + featured-placement revenue flows into the bount.ing OC project.
- Mía draws monthly autónomo invoice against the project balance (standard EUR invoice flow, matches Foqum/RugbyTour7 posture).
- OCE handles VAT, transparency ledger, contributor compensation.
- Host fee: typically 5–10% of inflows. Versus €6k SL setup + ongoing cuota societaria — vastly cheaper.

**Open execution sub-questions:**
- Which fiscal host under OC exactly? OCE is one option; there are others (Open Source Collective, etc.). Verify at v2 launch prep.
- Transparency ledger publishes every transaction publicly — fine for OSS values alignment, worth confirming the sponsor pool is comfortable being on it.

---

## Q8 — Curve implementation depth at v2 launch

**Status:** Open. Depends on Q2 settlement + engineering reality at launch time.

Two ships:

| Option | Description |
|---|---|
| **Full euro-curves** | Sponsor pre-funds curve_max (appreciation) or commits with decay-redirect policy (decay). Actual euro value moves over time per the curve. Solver receives curve-value-at-close. |
| **Flat-bounty + curve-as-ranking** | Bounty euro value is flat at contribution-time. Curve only affects queue ranking (decay = slides down the queue; appreciation = climbs the queue). Cleaner to implement, ships v2 faster. |

**Default:** ship flat + ranking-only at v2 launch if full euro-curves aren't engineering-ready. Full euro-curves arrive at v2.5. Better to ship flat-but-real than complex-but-broken.

**Decision trigger:** engineering readiness check 2 weeks before v2 launch. If full curves not ready, ship flat + ranking-only.

---

## Decisions log (settled)

### Q1 — Rail integration sequence (resolved 2026-05-14)

**Decision:** Multi-rail at v2 launch, provider-agnostic. Sequence by administrative/financial fit with current autónomo posture:

1. **Open Collective** — v2 launch. EUR invoices match existing Foqum/RugbyTour7 flow, fiscal-host model means OC carries pooling compliance. Zero new tax categories.
2. **GitHub Sponsors** — v2.1. Admin-trivial (pure read-API integration), but mechanism-poor (single-sponsor only).
3. **Drips Network** — v2.5, after autónomo books stabilize. Mechanism-fit (streaming = appreciation) and values-fit, but adds crypto-as-business-income reporting overhead (Modelo 720/721, IRPF treatment ambiguity, segregation of personal vs. business wallets). Not while RETA arrears are in aplazamiento.
4. **GNU Taler** — when (a) a production EU exchange settles to Spanish banks with non-trivial reach, AND (b) at least one v2 sponsor requests it, AND/OR (c) NLnet grant capture makes Taler-extension engineering feasible.

**Rationale:** filtered through administrative/financial-fit criterion, not personal-affinity. OC's fiscal-host model slots into the existing autónomo invoicing flow with no new tax categories. Drips deferred because crypto-business-income reporting is the worst thing to add right now.

**Architectural implication:** bount.ing v2 is a **federation layer** over funding rails. Adapter pattern, one per rail. Mirrors Lethe's multi-provider LLM pattern.

### Q2 — Curve presets (resolved 2026-05-14)

**Decision:** Three presets, locked.

| Preset | Shape | Floor / cap |
|---|---|---|
| `urgent` | Linear decay over 30d | Floor = 25% of initial |
| `standard` | Flat 30d → linear decay 60d to floor | Floor = 25% of initial |
| `patient` | Linear appreciation | Cap = 3× initial OR T+90d, whichever first |

**Sub-decisions:**
- Floor type: **fixed %** of initial (scales with bounty size; absolute floors would break for €5 and €5000 bounties alike).
- Appreciation cap: **pool-lifetime, not reset on new sponsors** (else gameable — sponsors would add late contributions to extend the cap).
- Post-terminal behavior: bounty stays at floor/cap until close. **Hard expire at T+180d** → unresolved pool re-routes per sponsor's pre-set destination policy.

**Implementation against OC:**
- **Decay:** decayed portion is *re-routed*, not refunded (refunds = fund-custody = autónomo poison). Sponsor picks destination at creation: default = bount.ing platform fund / community treasury / other underfunded bounties.
- **Appreciation:** sponsor pre-funds the curve_max upfront. Solver receives curve-value-at-close. Difference re-routes per sponsor's pre-set destination policy (same dropdown as decay). Sponsor knows max liability at creation; no escrow gymnastics.
- See Q8 for v2-launch fallback if full euro-curves aren't engineering-ready.

### Q3 — Self-sponsored bounty policy (resolved 2026-05-14)

**Decision:** Flag at v1, escalate later if abused.

Maintainer can sponsor own issue. UI displays "self-sponsored — sponsor is also merge authority" on the issue. Solvers see it before attempting. Escalate to third-party validator only if abuse pattern observed.

### Q4 — Listing fee structure (resolved 2026-05-14)

**Decision:** Free to list. Revenue from three other layers:

1. **Featured placement** — sponsor pays for top-of-queue visibility within their tag/repo
2. **Premium API** — agent fleets, subscription-based
3. **Rail kickback** where supported (Drips natively; OC via host-fee-share if negotiable)

**Why:** entry friction breaks the multi-sponsor pooling mechanic (differentiator #1). Free-to-list keeps the crowdfunded flow frictionless. Reconsider only if revenue doesn't materialize.

### Q5 — GitHub-only at v1 (resolved 2026-05-14)

**Decision:** GitHub-only at v1. Data model supports adapter shape for GitLab / Codeberg / forgejo at v2+.

Don't ship multi-platform before the GitHub experience is solid. Document the adapter interface now so the data model doesn't lock us out later.

### Q7 — Repo organization for v2 build (resolved 2026-05-14)

**Decision:** Branch `v1` from current `master`, rebuild v2 on `master`.

- Preserves authorship trail + GitHub stars + repo history.
- v1 stays reachable for archival reference (PROCESS.md, TESTS.md, 2024 codebase).
- v2 ships on clean `master`.
- 112 v1 CVEs become "v1 branch is archival, security alerts not blocking v2."
- `v1` branch created as part of this PR (defensive — locks in v1 state before any v2 work touches `master`).
