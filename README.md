# Laundry Service Backend API

[![Go](https://img.shields.io/badge/Go-1.25-blue?logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-green)](LICENSE)
[![Swagger](https://img.shields.io/badge/Swagger-Enabled-brightgreen?logo=swagger)](http://localhost:5500/swagger/index.html)

This is the backend API for an online laundry service platform. It enables users to register, login, manage laundry shops, and perform CRUD operations on shops. The API supports secure authentication, pagination, and is designed for scalability. Built with Go, it follows RESTful principles and includes features like JWT-based authentication, database migrations, and Swagger documentation.

The platform aims to simplify ordering and tracking laundry services with location-based search (future enhancements planned).

## Table of Contents
- [Laundry Service Backend API](#laundry-service-backend-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Tech Stack](#tech-stack)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Application](#running-the-application)
  - [API Documentation](#api-documentation)
  - [Testing](#testing)
  - [Roadmap](#roadmap)
  - [Contributing](#contributing)
  - [License](#license)

## Features
- **User Management**: Register new users, login with JWT authentication, and retrieve paginated user lists (admin-only).
- **Shop Management**: Create, read, update, and delete laundry shops (admin-only). Includes details like location, contact, and payment info.
- **Authentication & Authorization**: JWT-based secure access with role-based controls (e.g., admin vs. user).
- **Pagination**: Efficient listing of users and shops with query parameters for page and limit.
- **Database Migrations**: Automated schema management using Golang Migrate.
- **Swagger Integration**: Interactive API docs for easy testing and exploration.
- **Health Check**: Simple `/ping` endpoint for server health monitoring.

Planned features (from [ROADMAP.md](roadmap_docs/ROADMAP.md)):
- Order management with real-time tracking.
- Payment integration (cash-on-delivery and MFS like Bkash/Nagad).
- Review system and notifications.

## Tech Stack
- **Language**: Go 1.25+
- **Database**: PostgreSQL (with sqlx for queries)
- **Auth**: JWT (with HMAC-SHA256)
- **API Docs**: Swagger (via swaggo)
- **Migrations**: Golang Migrate
- **Dependencies**: godotenv, google/uuid, joho/godotenv, lib/pq, etc. (see [go.mod](go.mod))
- **Testing**: Go's built-in testing with testify
- **Deployment**: Docker for DB (see [DB_docker_container.txt](DB_docker_container.txt)); Terraform (planned for infrastructure)

## Installation
1. **Clone the Repository**:
   ```
   git clone https://github.com/enghasib/laundry-service.git
   cd laundry-service
   ```

2. **Install Dependencies**:
   ```
   go mod tidy
   ```

3. **Set Up PostgreSQL**:
   - Use the provided Docker command to spin up a Postgres container:
     ```
     docker run -d \
       --name laundry_service \
       -e POSTGRES_USER=postgres \
       -e POSTGRES_PASSWORD=postgres \
       -e POSTGRES_DB=laundry_service \
       -v laundry_service_data:/var/lib/postgresql/data \
       -p 5432:5432 \
       postgres:16
     ```
   - Run migrations:
     ```
     go run main.go  # (Migrations are auto-run on startup via cmd/serve.go)
     ```

## Configuration
Copy the example env file and update with your details:
```
cp example.env .env
```

Key variables in `.env`:
- `VERSION=1.0.0`
- `SERVICE_NAME=laundry-service`
- `HTTP_PORT=5500`
- `JWT_SECRET_KEY=your_secret_key` (generate a strong key)
- `DB_USER_NAME=postgres`
- `DB_PASSWORD=postgres`
- `DB_HOST=localhost`
- `DB_PORT=5432`
- `DB_NAME=laundry_service`
- `AUTH_CONTEXT_KEY=auth_context_key`

## Running the Application
1. Build and run:
   ```
   go build -o laundry-service main.go
   ./laundry-service
   ```
   Or directly:
   ```
   go run main.go
   ```

2. The server starts on `http://localhost:5500`.
   - Health check: `GET /ping` → `{ "message": "pong" }`

## API Documentation
- Swagger UI: Access at `http://localhost:5500/swagger/index.html`.
- Key Endpoints:
  - **Users**:
    - `POST /users/register`: Create a new user.
    - `POST /users/login`: Authenticate and get JWT.
    - `GET /users`: List users (paginated, auth required).
  - **Shops**:
    - `POST /shops`: Create a shop (admin auth required).
    - `GET /shops`: List shops (paginated, auth required).
    - `GET /shops/{shop_id}`: Get single shop.
    - `PUT /shops/{shop_id}`: Update shop (admin).
    - `DELETE /shops/{shop_id}`: Delete shop (admin).

Use tools like Postman or the provided [client.http](client.http) for testing.

## Testing
Run unit tests:
```
go test ./...
```

Example: [create_user_test.go](test/create_user_test.go) covers user creation with mocks.

## Roadmap
See [ROADMAP.md](roadmap_docs/ROADMAP.md) for detailed phases:
- **MVP (v0.1)**: Core auth, users, and shops (current).
- **v0.2**: Orders, payments, and delivery management.
- **v0.3+**: Reviews, notifications, and analytics.

Features are detailed in [FEATURES.md](roadmap_docs/FEATURES.md).

## Contributing
We welcome contributions! Follow these steps:
1. Fork the repo.
2. Create a feature branch: `git checkout -b feature/YourFeature`.
3. Commit changes: `git commit -m 'Add YourFeature'`.
4. Push: `git push origin feature/YourFeature`.
5. Open a Pull Request.

Use the [feature_request.yml](.github/ISSUE_TEMPLATE/feature_request.yml) template for suggestions. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines (if available).

## License
This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

---

For questions, open an issue or contact via [API Support](https://github.com/NesoHQ/laundry-service). Happy coding! 🚀