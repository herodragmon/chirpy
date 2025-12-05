# Chirpy

Chirpy is a small Go-based microservice that lets users register, log in, and create short messages called *chirps*. It includes token-based authentication, CRUD operations for chirps, and a lightweight web UI.

## Features
- User registration & login (with hashed passwords)
- Access & refresh token authentication
- Create, list, and delete chirps
- SQL-backed storage (sqlc-generated DB layer)
- Webhook handler
- Readiness & metrics endpoints
- Simple front-end (`index.html`)


## Getting Started

### 1. Install dependencies
`go mod tidy`


### 2. (Optional) Generate SQL code using sqlc
`sqlc generate`


### 3. Run the server
`go run .`


The service runs on **http://localhost:8080** (or your configured port).

## Environment Variables
Create a `.env` or export variables:

``
DATABASE_URL=postgres://user:pass@localhost:5432/chirpy?sslmode=disable
PLATFORM = "dev"
JWT_SECRET=your_jwt_secret_here
POKA_KEY = your polka key
``
## API Overview

### Authentication
- **POST /api/users**  
  Create a new user account.

- **POST /api/login**  
  Log in with email + password.  
  Returns: `access_token` and `refresh_token`.

- **POST /api/refresh**  
  Exchange a valid refresh token for a new access token.

- **POST /api/revoke**  
  Revoke a refresh token (logout everywhere).

- **PUT /api/users**  
  Update the authenticated user's data.

---

### Chirps
- **GET /api/chirps**  
  List all chirps. Supports optional query params:  
  - `author_id`  
  - `sort` (if implemented)

- **POST /api/chirps**  
  Create a new chirp (requires access token).

- **GET /api/chirps/{chirpID}**  
  Get a single chirp by ID.

- **DELETE /api/chirps/{chirpID}**  
  Delete a chirp by ID (owner-only, authenticated).

---

### Webhooks
- **POST /api/polka/webhooks**  
  Handle incoming Polka webhook events.

---

### Admin / System
- **GET /admin/metrics**  
  Prometheus metrics endpoint.

- **POST /admin/reset**  
  Reset database (enabled only when `PLATFORM=dev`).

- **GET /ready**  
  Readiness probe for health checks.

## Development Notes
- Uses **sqlc** for DB queries.
- Handlers are simple Go functions using standard `net/http`.
- Creating, deleting, and updating resources require authentication.

## License
Open-source. Modify and use freely.

