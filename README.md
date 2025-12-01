# Shortener — URL Shortening Service with Analytics

Shortener is a mini URL shortening service. It allows you to generate short links, redirect users to the original URLs, and collect analytics on link usage (who clicked, when, and from which device).  

---

## Features

- **Shorten URLs**: Generate short URLs for any long link.  
- **Redirect**: Redirect users from short URLs to original URLs.  
- **Analytics**: Track visits

---

## Project Structure

```

.
├── backend/                 # Backend service
│   ├── cmd/                 # Application entry points
│   ├── config/              # Configuration files
│   ├── internal/            # Internal application packages
│   │   ├── api/             # HTTP handlers, router, server
│   │   ├── config/          # Config parsing logic
│   │   ├── middlewares/     # HTTP middlewares
│   │   ├── models/           # Data models
│   │   ├── repository/      # Database repositories
│   │   ├── service/         # Business logic
│   ├── migrations/          # Database migrations
│   ├── go.mod
│   └── go.sum
├── frontend/                # Frontend application
├── .env.example             # Example environment variables
├── docker-compose.yml       # Multi-service Docker setup
├── Makefile                 # Development commands
└── README.md

````
---

## Running the Project
Copy and update .env:

```
cp .env.example .env
```

Build and run services via Docker:

```
make up
```

The backend will be available at:

- Backend API → http://localhost:8080/api/url-shortener
- Frontend UI → http://localhost:3000

To stop services:

```
make down
```

---

## API Endpoints

| Method | Endpoint                | Description                        |
| ------ | ----------------------- | ---------------------------------- |
| POST   | `/api/url-shortener/shorten`          | Create a new short URL             |
| GET    | `/api/url-shortener/s:short_url`      | Redirect to the original URL       |
| GET    | `/api/analytics/:short_url`           | Retrieve analytics for a short URL |

---

## Example Requests

### **1. Create Short URL**

**Request**

```http
POST /api/shorten
Content-Type: application/json

{
  "url": "https://example.com/long-url",
}
```

**Response**

```json
{
  "url": "sd3ffjaldsjfa",
}
```

---

### **2. Redirect Short URL**

Access via browser or HTTP client:

```
GET /api/url-shortener/:sd3ffjaldsjfa
```

Redirects to: `https://example.com/long-url`

---

### **3. Get Analytics**

**Request**

```
GET /api/url-shortener/analytics/sd3ffjaldsjfa
```

**Response**

```json
{
    "result": [
        {
            "id": 1,
            "short_url": "8wsnEQcYgivMx",
            "ip": "",
            "user_agent": "",
            "time": "2025-10-19T17:09:14.924592Z"
        },
        {
            "id": 2,
            "short_url": "8wsnEQcYgivMx",
            "ip": "",
            "user_agent": "",
            "time": "2025-10-19T17:09:16.918236Z"
        }
    ]
}
```

## Summary
- Backend (Go + PostgreSQL + Redis) → runs on port 8080
- Frontend → runs on port 3000
- URL can be shorten via API or UI