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
PORT=8080
PLATFORM = "dev"
JWT_SECRET=your_jwt_secret_here
POKA_KEY = your polka key
``


## API Overview
### Auth
- `POST /login` — returns access + refresh tokens  
- `POST /refresh` — returns new access token  
- `POST /users` — create a new user  
- `PUT /users/{id}` — update user  

### Chirps
- `GET /chirps` — list chirps  
- `POST /chirps` — create chirp  
- `DELETE /chirps/{id}` — delete chirp  

### System
- `GET /ready` — readiness probe  
- `/metrics` — Prometheus metrics  

## Development Notes
- Uses **sqlc** for DB queries.
- Handlers are simple Go functions using standard `net/http`.
- Delete & create routes require authentication.

## License
Open-source. Modify and use freely.

