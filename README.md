<p align="center">
  <img src="https://nuxt.com/assets/design-kit/logo/icon-green.svg" alt="Nuxt Logo" width="80" />
</p>

<h1 align="center">Personal Portfolio</h1>

<p align="center">
  🌐 <a href="https://athulanilthomas.in">athulanilthomas.in</a>
</p>

<p align="center">
  <a href="https://nuxt.com"><img src="https://img.shields.io/badge/Nuxt-4.x-00DC82?logo=nuxt.js&logoColor=white" alt="Nuxt" /></a>
  <a href="https://go.dev"><img src="https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white" alt="Go" /></a>
</p>

---

## Tech Stack

| Layer    | Technology                                    |
|----------|-----------------------------------------------|
| Frontend | Nuxt 4, Vue 3, Nuxt UI v4, Tailwind CSS v4    |
| Backend  | Go (Gin, Uber FX)                             |
| Tooling  | pnpm, Docker                                  |

---

## Running Locally

### Frontend

```bash
cd apps/ui
pnpm install
pnpm dev        # http://localhost:3000
```

### Backend

```bash
cd apps/api
cp .env.example .env   # fill in required credentials
go run ./cmd/main.go   # http://localhost:8080
```

Or with Docker:

```bash
cd apps/api
docker compose up --build
```
