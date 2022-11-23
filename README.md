# Code Scanner

A simple code scanning application that detects sensitive keywords in public git repos

## Prerequisite

1. PostgreSQL database
2. Kafka server

## How to Run

### Install Depedencies

```
sudo make install-deps
go mod download
```

### Database Schema Migration

1. Create database
2. Set environment variable

```
POSTGRESQL_URL='postgres://<user>:<password>@<host>:<port>/<database>?sslmode=disable'
```

3. Run migration script

```
make migrate-up
```

### Run Directly

1. Set environment variables

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=code-scanner
DB_USER=postgres
DB_PASSWORD=postgres

SERVER_PORT=8000

KAFKA_SERVERS=localhost
KAFKA_SCAN_REPO_TOPIC='code-scanner.repository.scan'
```

2. Run server

```
go run cmd/server/main.go
```

3. Run worker

```
go run cmd/worker/main.go
```

### Run via Docker Compose

1. Modify docker-compose.yaml to fit the environment

```
version: "3.7"

services:
  server:
    container_name: code-scanner-server
    network_mode: "host"
    build:
      context: .
      dockerfile: server.Dockerfile
    ports:
      - 9000:9000
    environment:
      - SERVER_PORT=9000
      - DB_PASSWORD=secret

  worker:
    container_name: code-scanner-worker
    network_mode: "host"
    build:
      context: .
      dockerfile: worker.Dockerfile
    environment:
      - DB_PASSWORD=secret
```

2. Run server & worker

```
docker-compose up
```

### Run Tests

```
go test --cover ./...
```

## Documentation

### API Spec

Link to API Spec: [Link](api/openapi.yaml)

### Slide

```
present
```

Link to slide: http://localhost:3999/documentation.slide#1
