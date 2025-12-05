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
* net/http or a lightweight router (mux/chi depending on repo)
* SQL database (SQLite/Postgres)
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
export PORT=8080
export DATABASE_URL=sqlite3://./chirpy.db   # or a Postgres URL
# any other app-specific env variables
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

### Authentication

This project may use simple token auth or session cookies. Update these examples to match your implementation.

* `POST /api/register`

  * Body: `{ "username": "alice", "password": "secret" }`
  * Response: `201 Created` with user JSON

* `POST /api/login`

  * Body: `{ "username": "alice", "password": "secret" }`
  * Response: `200 OK` with token

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

The project contains SQL schema under the `sql/` folder. By default the app can use SQLite for local dev and Postgres in production — change the `DATABASE_URL` accordingly.

If you use SQLite, an example to initialize DB (if the project provides migrations):

```bash
# run migrations (if you have a migrate tool)
# migrate -path sql -database "sqlite3://./chirpy.db" up
```

---

## Tests

If the project includes tests, run them like this:

```bash
go test ./...
```

Add unit and integration tests for handlers, DB access, and auth flows as you expand the project.

---

## Recommended improvements (to make repo stand out)

* Add a polished `README.md` (this file)
* Add usage examples and example requests/responses
* Add tests for handlers and database layer
* Add a Dockerfile and docker-compose for easy local deployment
* Add CI (GitHub Actions) to run tests on PRs
* Add basic API documentation (Swagger/OpenAPI or a simple markdown doc)

---

## License

Add a LICENSE file (MIT or Apache-2.0 recommended for student projects).

---

## Contact

If you'd like changes to this README or want a version tailored with exact commands from your project structure, tell me and I’ll update it to match your code.


