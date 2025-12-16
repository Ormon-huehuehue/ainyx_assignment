# Go Backend Development Task

A RESTful API in Go managing user data with dynamic age calculation.

## Tech Stack

- **Go**: Language
- **Fiber**: Web Framework
- **PostgreSQL**: Database
- **SQLC**: Type-safe SQL generator
- **Zap**: Logging
- **Validator**: Input validation
- **Docker**: Containerization

## Project Structure

```
.
├── cmd/server/main.go       # Entry point
├── internal
│   ├── handler              # HTTP Handlers
│   ├── repository           # Data access layer
│   ├── service              # Business logic
│   ├── routes               # Route definitions
│   ├── middleware           # Custom middleware
│   ├── models               # Domain models
│   └── logger               # Zap logger setup
├── db                       # Database migrations and queries
├── config                   # Configuration
└── docker-compose.yml       # Docker services
```

## Setup & Run

### Prerequisites

- Docker & Docker Compose
- Go 1.23+ (optional, if running locally without Docker)

### Using Docker (Recommended)

1.  **Start the services:**
    ```bash
    docker-compose up --build
    ```
    The API will be available at `http://localhost:3000`.

### Running Locally

1.  **Start PostgreSQL:**
    Ensure PostgreSQL is running and update `.env` with credentials.

    ```bash
    docker-compose up -d db
    ```

2.  **Run Migrations:**
    (You might need a migration tool like `migrate` or manually apply `db/migrations/001_users.sql`)

3.  **Run the application:**
    ```bash
    go run cmd/server/main.go
    ```

## API Endpoints

| Method   | Endpoint            | Description                                 |
| :------- | :------------------ | :------------------------------------------ |
| `POST`   | `/api/v1/users`     | Create a new user                           |
| `GET`    | `/api/v1/users/:id` | Get a user by ID                            |
| `PUT`    | `/api/v1/users/:id` | Update a user                               |
| `DELETE` | `/api/v1/users/:id` | Delete a user                               |
| `GET`    | `/api/v1/users`     | List users (Pagination: `?page=1&limit=10`) |

## Testing

Run unit tests:

```bash
go test ./internal/service/...
```
# ainyx_assignment
