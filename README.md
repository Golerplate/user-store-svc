# Faceit - User Store Service

## Overview

In this exercise, I had to build a Golang service that manage Users and exposes them through a gPRC API.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.22 or later)
- [Docker](https://www.docker.com/products/docker-desktop) (for containerization)
- [Docker Compose](https://docs.docker.com/compose/) (optional, for managing multi-container Docker applications)
- [PostgreSQL](https://www.postgresql.org/) (for local development)
- [Goose](https://github.com/pressly/goose) (to run migration)

```
faceit-user-store-svc/
├── cmd
│   └── main.go - Define how the server/database/redis runs and start the gRPC server.
├── internal
│   ├── config - Contains all the config that the microservice needs.
│   ├── database - Contains all the migrations and the postgres logics.
│   ├── entities - Defines all the entities needed.
│   ├── handlers - Defines the HTTP server and all the handler methods.
│   └── service - Service layer; handlers call service methods, and service calls database methods.
├── pkg - Contains all the libraries used in different micro-services. These are non-product libraries.
└── contracts - Defines any communication between services, ensuring payload consistency and easier versioning.
```

## Getting Started

Make sure you have goose installed on your computer

```go install github.com/pressly/goose/v3/cmd/goose@latest```

### Update environments variables

**docker-compose.yml** is already configured but feel free to update environments variables.

### Run the project

```make run```

### Connnect to the database

```postgresql://root:root@127.0.0.1/faceit-user-store-svc-db?tLSMode=0```

### Run test coverage

```make unit-test```

