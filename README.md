# athulanilthomas.in

> Personal portfolio website — built as a monorepo with a Nuxt.js frontend and a Go backend.

[![Nuxt](https://img.shields.io/badge/Nuxt-4.x-00DC82?logo=nuxt.js&logoColor=white)](https://nuxt.com)
[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev)

---

## Tech Stack

| Layer    | Technology                                      |
|----------|-------------------------------------------------|
| Frontend | Nuxt 4, Vue 3, Nuxt UI v4, Tailwind CSS v4      |
| Content  | Nuxt Content v3 (YAML-driven, Zod-validated)    |
| Backend  | Go (Gin, Uber FX, MongoDB)                      |
| 3D / FX  | Three.js, OGL                                   |
| Tooling  | pnpm, ESLint, Docker                            |

---

## Project Structure

```
athulanilthomas.in/
├── apps/
│   ├── ui/     # Nuxt 4 frontend — portfolio pages, WebGL effects, SEO
│   └── api/    # Go backend — GitHub & Spotify integrations, REST API
```

---

## Features

- **Portfolio pages** — About, Experience, Education, Skills, Projects
- **Live integrations** — GitHub activity and currently-playing Spotify track via the Go API
- **WebGL effects** — Aurora (OGL) and Liquid Ether (Three.js) background animations, lazy-loaded for performance
- **YAML content layer** — All portfolio data lives in structured YAML files, validated with Zod
- **SEO-ready** — Sitemap, robots.txt, Schema.org `Person` identity, page prerendering / ISR
- **Page transitions** — Smooth `out-in` transitions between routes

---

## Getting Started

### Prerequisites

- [Node.js](https://nodejs.org) ≥ 20 and [pnpm](https://pnpm.io) ≥ 10
- [Go](https://go.dev) ≥ 1.25

### Frontend

```bash
cd apps/ui
pnpm install
pnpm dev        # http://localhost:3000
```

### Backend

```bash
cd apps/api
cp .env.example .env   # fill in GitHub App & Spotify credentials
go run ./cmd/main.go   # http://localhost:8080
```

Or with Docker:

```bash
cd apps/api
docker compose up --build
```

---

## Development Notes

- The monorepo uses independent package managers per app — pnpm for `ui`, Go modules for `api`.
- Frontend reads `NUXT_API_BASE` and `NUXT_API_SECRET` at runtime to proxy requests to the Go server.
- Heavy WebGL components are lazy-loaded (`LazyUiAurora`, `LazyUiLiquidEther`) to keep the initial bundle light.
- The Projects page uses ISR (revalidated every hour); all other routes are fully prerendered.

---

## Deployment

- **Frontend** — Deploy the Nuxt output to any Node-compatible host (Vercel, Fly.io, etc.).
- **Backend** — A `Dockerfile` and `compose.yaml` are included in `apps/api` for containerised deployment.

---

## Author

**Athul Anil Thomas** — Full Stack Developer

- 🌐 [athulanilthomas.in](https://athulanilthomas.in)
- 💼 [LinkedIn](https://www.linkedin.com/in/athul-anil-thomas)
- 🐙 [GitHub](https://github.com/athulanilthomas)
