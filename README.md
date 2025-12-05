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

Base URL: http://localhost:8080

### Admin
POST /admin/reset  
- Reset the database (dev only).  
- Returns: 200

### Users & Auth
POST /api/users  
- Create a new user.  
- Body: { "email": "...", "password": "..." }  
- Returns: 201 + { id, email }

POST /api/login  
- Log in user.  
- Body: { "email": "...", "password": "..." }  
- Returns: 200 + { token, email }

POST /api/refresh  
- Exchange refresh token for new token.  
- Body: { "token": "<refresh-token>" }  
- Returns: 200 + { token }

POST /api/revoke  
- Revoke a refresh token.  
- Body: { "token": "<refresh-token>" }  
- Returns: 200

PUT /api/users  
- Update authenticated user.  
- Auth: Bearer <token>  
- Returns: 200 + updated user

### Chirps
POST /api/chirps  
- Create chirp.  
- Auth: Bearer <token>  
- Body: { "body": "text" }  
- Returns: 201 + created chirp

GET /api/chirps  
- List chirps.  
- Query: author_id=<userID>  
- Returns: 200 + list of chirps

GET /api/chirps/{chirpID}  
- Get single chirp.  
- Returns: 200 or 404

DELETE /api/chirps/{chirpID}  
- Delete chirp.  
- Auth: Bearer <token>  
- Returns: 204 or 403 or 404

### Webhooks
POST /api/polka/webhooks  
- Handle Polka webhook events.  
- Returns: 200

### System
GET /ready  
- Readiness probe.  
- Returns: 200

GET /admin/metrics  
- Prometheus metrics.  
- Returns: 200

## Development Notes
- Uses **sqlc** for DB queries.
- Handlers are simple Go functions using standard `net/http`.
- Creating, deleting, and updating resources require authentication.

## License
Open-source. Modify and use freely.

