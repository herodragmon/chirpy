# Chirpy

**Chirpy** — a simple Twitter-like microblogging API written in Go.

This repo implements a minimal backend service that supports user registration/login and CRUD operations for short posts ("chirps"). It's a great showcase of Go backend fundamentals: routing, request handling, JSON APIs, persistence, and modular code structure.

---

## Features

* User registration and login (basic auth/session/token — depends on implementation)
* Create, read, update, delete "chirps" (posts)
* List recent chirps and fetch by ID
* Clean handler-based code organization
* SQL-backed persistence (adaptable to SQLite/Postgres)

---

## Tech stack

* Go (Golang)
* net/http
* SQL database (Postgres)
* JSON-based REST API

---

## Quick start

> These instructions give you a fast way to run the service locally.

### Prerequisites

* Go 1.20+ installed
* Git
* (Optional) Docker & Docker Compose if you prefer containerized runs

### Clone & build

```bash
git clone https://github.com/herodragmon/chirpy.git
cd chirpy
# Edit config/env if required
go build ./...
./chirpy    # or go run ./cmd/chirpy
```

### Environment variables

Create a `.env` or export the following before running (examples):

```bash
# Postgres (production / local Postgres)
export DB_URL="postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable"
export PORT=8080
export PLATFORM=dev            # set to "dev" to enable admin endpoints locally
export SECRET_KEY="<your-jwt-secret>"
export POLKA_KEY="<your-polka-webhook-key>"  # used to validate incoming webhooks
```

---

## Running with Docker (optional)

You can containerize the app by creating a simple `Dockerfile` and `docker-compose.yml`. Example `docker-compose.yml` snippet:

```yaml
version: '3.8'
services:
  chirpy:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - DATABASE_URL=sqlite3:///data/chirpy.db
    volumes:
      - ./data:/data
```

---

## API Reference

> Replace `:PORT` with your configured port, e.g. `8080`.

### Users / Authentication

* `POST /api/users`

  * Create a new user
  * Body example: `{ "email": "test@example.com", "password": "secret" }`
  * Response: `201 Created` with user JSON

* `PUT /api/users`

  * Update user email/password (auth required)
  * Body example: `{ "email": "new@example.com" }`
  * Response: `200 OK`

* `POST /api/login`

  * User login, returns access & refresh tokens
  * Body: `{ "email": "test@example.com", "password": "secret" }`
  * Response: `200 OK` with tokens

* `POST /api/refresh`

  * Exchange refresh token for new access token
  * Response: `200 OK`

* `POST /api/revoke`

  * Revoke refresh token
  * Response: `200 OK`

### Chirps (posts)

* `GET /api/chirps`

  * List recent chirps
  * Response: `200 OK` with `[{"id":"...","user_id":"...","text":"...","created_at":"..."}, ...]`

* `POST /api/chirps`

  * Create a new chirp (auth required)
  * Body: `{ "text": "hello world" }`
  * Response: `201 Created` with created chirp

* `GET /api/chirps/{id}`

  * Get a single chirp by id
  * Response: `200 OK` with chirp JSON

* `PUT /api/chirps/{id}`

  * Update an existing chirp (auth + owner)
  * Body: `{ "text": "updated text" }`
  * Response: `200 OK` with updated chirp

* `DELETE /api/chirps/{id}`

  * Delete a chirp (auth + owner)
  * Response: `204 No Content`

### Example curl usage

```bash
# List chirps
curl http://localhost:8080/api/chirps

# Create chirp (replace TOKEN)
curl -X POST http://localhost:8080/api/chirps \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"text":"first chirp!"}'
```

---

## Database

This project uses **Postgres** in production and can use **SQLite** for quick local development. The SQL schema and migrations live in the `sql/` folder.

### Migrations (Goose)

This project uses **Goose** for database migrations. Migration files live in the `sql/schema` directory.

#### Running migrations

```bash
# Install Goose
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations
goose -dir sql/schema postgres "$DB_URL" up
```

To roll back the last migration:

```bash
goose -dir sql/schema postgres "$DB_URL" down
```

To view migration status:

```bash
goose -dir sql/schema postgres "$DB_URL" status
```

> Note: Goose supports Postgres, SQLite, and others. Since this project uses Postgres in main.go (`sql.Open("postgres", dbURL)`), Goose should be run with the `postgres` driver.

---

## Admin endpoints & platform guard

Admin routes such as `/admin/metrics` and `/admin/reset` are only enabled when the `PLATFORM` env var equals `dev`. This prevents accidental exposure of admin capabilities in production.

If `PLATFORM` is not `dev`, admin endpoints will reject requests / not be registered.

---

## Webhooks (Polka)

The app exposes a webhook endpoint:

* `POST /api/polka/webhooks`

  * This endpoint expects incoming requests from the Polka service.
  * The app validates incoming webhook requests using the `POLKA_KEY` env var. Set `POLKA_KEY` to a secret value and provide the same secret to the Polka service so you can verify requests.

Example validation flow (conceptual):

1. Polka sends a POST to `/api/polka/webhooks` with a signature or key in the headers/body.
2. Your handler compares that value to the `POLKA_KEY` value loaded from env.
3. If they match, the webhook is accepted and processed; otherwise it is rejected with `401 Unauthorized`.

Make sure to keep `POLKA_KEY` private and do not commit it to the repository.

---

## Tests

For running auth tests:

```bash
go test ./...
```

Add unit and integration tests for handlers, DB access, and auth flows as you expand the project.


