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
## 🐦 API Overview

**Base URL:** `http://localhost:8080`

---

### 🛠️ Admin & System Endpoints

| Method | Path | Description | Authentication | Request Body | Response Status & Body |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **POST** | `/admin/reset` | Reset the database (dev only). | None | None | `200` |
| **GET** | `/ready` | Readiness probe. | None | None | `200` |
| **GET** | `/admin/metrics` | Prometheus metrics. | None | None | `200` |

---

### 👤 Users & Authentication Endpoints

| Method | Path | Description | Authentication | Request Body | Response Status & Body |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **POST** | `/api/users` | Create a new user. | None | `{ "email": "...", "password": "..." }` | `201` + `{ id, email }` |
| **POST** | `/api/login` | Log in user. | None | `{ "email": "...", "password": "..." }` | `200` + `{ token, email }` |
| **POST** | `/api/refresh` | Exchange refresh token for new token. | None | `{ "token": "<refresh-token>" }` | `200` + `{ token }` |
| **POST** | `/api/revoke` | Revoke a refresh token. | None | `{ "token": "<refresh-token>" }` | `200` |
| **PUT** | `/api/users` | Update authenticated user. | `Bearer <token>` | Updated user fields | `200` + updated user |

---

### 💬 Chirps Endpoints

| Method | Path | Description | Authentication | Request Body | Response Status & Body |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **POST** | `/api/chirps` | Create chirp. | `Bearer <token>` | `{ "body": "text" }` | `201` + created chirp |
| **GET** | `/api/chirps` | List chirps. | None | None | `200` + list of chirps |
| **GET** | `/api/chirps/{chirpID}` | Get single chirp. | None | None | `200` or `404` |
| **DELETE** | `/api/chirps/{chirpID}` | Delete chirp. | `Bearer <token>` | None | `204` or `403` or `404` |

> **Query Parameter for GET /api/chirps:** Use `author_id=<userID>` to filter the list.

---

### 🔗 Webhooks Endpoints

| Method | Path | Description | Authentication | Request Body | Response Status & Body |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **POST** | `/api/polka/webhooks` | Handle Polka webhook events. | None | Webhook payload | `200` |

---

## Development Notes
- Uses **sqlc** for DB queries.
- Handlers are simple Go functions using standard `net/http`.
- Creating, deleting, and updating resources require authentication.

---

## License
Open-source. Modify and use freely.

