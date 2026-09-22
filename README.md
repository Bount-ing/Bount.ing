# Bount.ing — archived (September 2026)

> **This project is no longer maintained.** Nothing runs at bount.ing any more, the code is kept as-is for reference, and the repository is read-only.

Bount.ing was a results-based marketplace for open-source contributions: pooled bounties on issues, first merged PR wins, a solver reputation built from merge history, and structured provenance disclosure on every PR. Two attempts were made, in 2024 and in May 2026. Both stalled on the same thing: distribution. The idea was never refuted by the market, because it never reached one. That is the honest cause of death.

## What is worth keeping

- [docs/V2_THESIS.md](docs/V2_THESIS.md) — the five-mechanism thesis (multi-sponsor bounties, time-sensitive bounties, open competition, solver reputation, provenance disclosure). It aged well: AI agents solving open-source issues at scale is now the present, and the trust layer it describes is the missing piece.
- [docs/OPEN_QUESTIONS.md](docs/OPEN_QUESTIONS.md) — the design questions that were resolved and the ones that were not.
- The Go API and Vue 3 front end below, as a worked example, not as a base to build on.

## If you want to pick this up

Start from the thesis, not the code. And start with a channel: who will list the first bounties and who will solve them. A third attempt that leads with mechanism design repeats the first two.

Revival: the domain is held. If it ever restarts, it restarts with this autopsy.

---

## Original README

# Bount.ing

Bount.ing is a gamified platform designed to incentivize and reward open source contributions. By integrating game mechanics, Bount.ing aims to make contributing to open source projects more engaging and rewarding.

## Technologies Used
Frontend: Vue 3 (Vite)
Backend: Golang
Containerization: Docker Compose

## Getting Started
These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
Make sure you have the following installed on your system:

- `docker`
- `docker-compose`

### Download the repo

```
git clone git@github.com:Bount-ing/Bount.ing.git &&
cd Bount.ing
```

### Set up Env

Fill the files in `api/.env` and `front/.env`

### Run the docker

`sudo docker-compose build  && sudo docker-compose up -d && sudo docker-compose logs -f`
