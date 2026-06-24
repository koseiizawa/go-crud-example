# AGENTS.md

## Cursor Cloud specific instructions

This repo is a single Go (Gin + GORM) CRUD REST API backend. No frontend, no `*_test.go` files, no Dockerfile. Entry point: `cmd/server/main.go`, serves on port `8080`.

### Services
- **PostgreSQL 16** is installed natively (not via the `docker-compose.yml`, since Docker is not available in this environment). The update script does NOT start services, so at the start of a session you must start the cluster yourself: `sudo pg_ctlcluster 16 main start` (idempotent — ignore "already running"). A role/db matching `docker-compose.yml` is provisioned: user `star`, password `123123123`, db `test`. To (re)create if missing: `sudo -u postgres psql -c "CREATE ROLE star LOGIN PASSWORD '123123123';"` and `sudo -u postgres psql -c "CREATE DATABASE test OWNER star;"`.
- **Go API server**: run with `go run ./cmd/server` from the repo root.

### Config gotchas
- A `.env` file at the repo root is required and is **gitignored** (not committed), so it is NOT recreated by pulling latest. If it is missing, recreate it with these values: `PORT=8080`, `DB_HOST=localhost`, `DB_PORT=5432`, `DB_USER=star`, `DB_PASS=123123123`, `DB_NAME=test`, `SSL_MODE=disable`, `JWT_TOKEN=dev_secret_please_change`, `JWT_EXPIRE=24`.
- The code reads `JWT_TOKEN` (NOT `JWT_SECRET` as the README's example shows) and `JWT_EXPIRE` (hours). See `internal/config/config.go`.
- DB env vars used: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`, `SSL_MODE`. The README's example DB values (`postgres`/`cruddb`) do NOT match `docker-compose.yml`; use the `star`/`123123123`/`test` values that the environment is provisioned with.
- On startup, GORM auto-migrates `User` and `Admin` tables. The server `log.Fatal`s if Postgres is unreachable, so ensure the cluster is running first.

### Lint / build / run
- Format check: `gofmt -l .` (empty output = clean)
- Vet: `go vet ./...`
- Build: `go build ./...`
- No automated tests exist in the repo.

### E2E flow (no UI — use curl)
`/users/*` routes require a `Authorization: Bearer <jwt>` header. Get a JWT via signup + login:
1. `POST /signup` `{"email":"...","password":"..."}` (min password length 6)
2. `POST /login` → returns `{"message":"<jwt>"}`
3. `POST /users/` (with Bearer token) `{"name":"...","email":"...","password":"..."}`

Note: passwords are stored in plain text (no hashing) in `AdminRepositoryImpl`.
