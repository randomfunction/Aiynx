# User API

A production-grade REST API in Go for managing users. Built with Fiber, PostgreSQL, SQLC, Zap, and Validator.

## Tech Stack

- **Language**: Go 1.23+
- **Framework**: [Fiber v2](https://github.com/gofiber/fiber) (Fastest Go HTTP framework)
- **Database**: PostgreSQL
- **Driver**: [pgx v5](https://github.com/jackc/pgx)
- **ORM/Query Builder**: [SQLC](https://sqlc.dev/) (Type-safe SQL)
- **Logger**: [Zap](https://github.com/uber-go/zap) (High performance structure logging)
- **Validation**: [Validator v10](https://github.com/go-playground/validator)

## Setup

### Prerequisites

- Go 1.23+
- PostgreSQL
- Make (optional)

### Environment Variables

Create a `.env` file or export these variables:

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable"
export PORT="3000"
```

### Database Setup

Run the migrations or the schema setup manually:

```bash
# Using psql
psql $DATABASE_URL -f db/sqlc/schema.sql
```

## How to Run

```bash
# Install dependencies
go mod tidy

# Run server
go run cmd/server/main.go
```

## API Examples

### Create User
`POST /users`
```json
{
  "name": "Alice",
  "dob": "1995-05-20"
}
```

### Get User
`GET /users/1`

### List Users
`GET /users`

### Update User
`PUT /users/1`
```json
{
  "name": "Alice Wonderland",
  "dob": "1995-05-20"
}
```

### Delete User
`DELETE /users/1`
