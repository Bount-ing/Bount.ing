# Open Questions

Live decision deck for Bount.ing v2. Decisions move to [V2_THESIS.md](V2_THESIS.md) as they're settled.

---

## Q1 — First rail to integrate at v2

**Status:** Open. Pending decision.

The v2 marketplace rides existing pooling rails rather than building its own (Path B, see V2_THESIS.md). The question is which rail to integrate with *first*. Eventually probably multiple; the first one shapes the v2 spec.

| Rail | Why first | Why not |
|---|---|---|
| **Drips Network** | Crypto-native. Already supports streaming + multi-sponsor pooling natively. Maps cleanly to the appreciating-bounty mechanic (continuous flow). Operator is crypto-fluent (BASE/ETH wallets exist). Cut-on-close kickback supported. | Crypto-only — restricts sponsor pool to crypto-comfortable. UX overhead for non-crypto sponsors. |
| **Open Collective** | Fiat-native. Broad maintainer adoption. Fiscal-host model handles compliance globally. | Single-pool-per-project, less granular per-issue than the mechanism design requires. May need creative use of "expenses" to attach bounties to specific issues. |
| Gitcoin Allo | Quadratic funding primitives match the crowdfunded mechanic well. | Crypto-native like Drips; less mature for per-issue (more matched-round-shaped). |
| GitHub Sponsors (issue-level) | Closest to existing developer workflow. | Single-sponsor flow only — doesn't support multi-sponsor pooling, which is differentiator #1. Disqualifies as primary rail. |

**Default lean (not yet locked):** Drips first, OC second. Reasoning: appreciation mechanic maps to streaming natively; crypto-fluent operator reduces integration overhead; later add OC for fiat reach.

**To decide:** confirm or flip. Locking this gates the v2 architecture spec.

---

## Q2 — Decay/appreciation curve presets

**Status:** Open. Needs concrete numbers.

Three presets named — `urgent` / `standard` / `patient` — but the actual curves aren't specified.

Sketch to react to:

| Preset | Shape | Example |
|---|---|---|
| `urgent` | Linear decay over 30 days from initial → floor (e.g. €100 → €25 in 30d) | Security patches, time-sensitive launches |
| `standard` | Flat for 30 days, then linear decay to floor over the next 60d | Default for typical feature work |
| `patient` | Linear appreciation, capped at 3× initial or T+90d (whichever first) | Long-tail OSS maintenance |

Open sub-questions:
- Is "floor" a fixed % of initial (e.g. 25%) or absolute (e.g. €10)?
- Does the appreciation cap reset if a new sponsor joins the pool, or is it pool-lifetime?
- What happens after the curve terminates? Bounty expires? Stays at floor/cap indefinitely?

---

## Q3 — Self-sponsored bounty policy

**Status:** Open.

Risk: maintainer sponsors their own issue, then picks which PR to merge. Open-competition mechanism breaks if the sponsor is also the adjudicator.

Three options:

1. **Block** — maintainer cannot sponsor own repos. Cleanest but excludes maintainers from funding their own roadmap.
2. **Flag** — allow it, but display "self-sponsored — sponsor is also merge authority" prominently on the issue. Lets the market decide.
3. **Third-party validator** — self-sponsored bounties require a designated reviewer who isn't the sponsor or the solver. Heavier, but actually preserves the mechanism.

**To decide:** which? Probably (2) at v1, escalate to (3) only if abuse observed.

---

## Q4 — Listing fee structure

**Status:** Open.

Revenue model under Path B includes a listing fee (sponsor pays to surface a bounty on the ranked board).

- Flat fee (e.g. €5 per listing)?
- Scaled fee (% of bounty amount, e.g. 3%)?
- Tiered (free below €X, then scaled)?
- Free to list, paid for featured placement only?

**Tension:** flat fee deters tiny bounties (bad — multi-sponsor needs small contributions to work). Scaled fee feels Stripe-like and might trigger payment-services questions. Free-to-list with paid featured placement is cleanest legally but reduces revenue floor.

**Default lean:** free to list (multi-sponsor flow shouldn't have entry friction); revenue comes from featured placement + premium API + rail kickback. Reconsider only if revenue doesn't materialize.

---

## Q5 — GitHub-only at v1, or design for GitLab/Codeberg from the start?

**Status:** Open. Likely "GitHub-only at v1" but worth naming.

The "merged PR = bounty closes" mechanic is GitHub-shaped. v1 board could ingest from GitLab and Codeberg with adapter work, but it adds complexity for a small slice of the OSS market.

**Default lean:** GitHub-only at v1. Revisit at v2. Document the adapter shape now so the data model doesn't lock us out later.

---

## Q6 — How is `impact_score` defined concretely?

**Status:** Open — this is the v0.5 deliverable.

`impact_score` is a function over observable signals. v0.5 will derive it from the v0 hand-weights in `bounties.yaml`. Candidate signals:

| Signal | Source | Cost to compute |
|---|---|---|
| Repo star count | GitHub API | trivial |
| Dependency-graph reach (who imports this) | npm/PyPI/crates.io/Cargo.toml/etc. | medium |
| Issue 👍 reactions | GitHub API | trivial |
| Maintainer responsiveness | GitHub API (time-to-first-reply) | medium |
| Blocks-downstream signal | issue body parsing (`Closes #X`, references) | hard |
| Security relevance | label heuristics + CVE feed cross-reference | medium |
| Effort estimate | comment volume / referenced files / LOC of context | hard |

**To do at v0.5:** pick 3–4 signals, fit weights against `bounties.yaml`, see if reproduces ≥80%. If not, iterate.

---

## Q7 — Repo organization for v2 build

**Status:** Open.

v1 code lives in this repo (`api/`, `wui/`, `dcdn/` submodules — Go backend, Vue 3 frontend, Docker Compose). v2 is a different product with different mechanism design.

Options:
1. **Rebuild in-place** — replace `api/` and `wui/` content. Lose v1 history in the working tree but keep it in git history.
2. **Branch** — `v1` branch preserves shipped state, `master` rebuilds for v2.
3. **Greenfield** — archive this repo, start fresh at a new repo. Preserves clean history; loses domain authority.

**Default lean:** option 2 (branch off `v1`, rebuild on `master`). Preserves authorship trail, keeps the repo name + GitHub stars + history, lets v1 stay reachable for archival reference.

---

## Decisions log (settled)

*(empty — move resolved questions here with the decision and date)*
