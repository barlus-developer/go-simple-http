# go-simple-http

A small HTTP API built with Gin, Zap, Viper, and a DDD-style package layout. The service exposes one endpoint:

```http
GET /
```

```json
{
  "status": "ok",
  "message": "Hello, World!!!"
}
```

## Tech Stack

- Gin for HTTP routing and middleware.
- Zap for structured logging.
- Viper for configuration from defaults, files, and environment variables.
- Standard Go `net/http` server with graceful shutdown.

## Project Structure

```text
.
├── cmd/server              # Application entrypoint
├── .env.example            # Example environment variable overrides
├── config.example.yaml     # Example local configuration
├── internal/application    # Use cases and application services
├── internal/bootstrap      # Dependency wiring
├── internal/domain         # Domain models
├── internal/infrastructure # Config, logging, and external adapters
└── internal/interfaces     # HTTP handlers, middleware, and routers
```

## Requirements

- Go 1.26.5 or compatible with the module version in `go.mod`.

## Run

```sh
make
```

By default, the server listens on `0.0.0.0:8080`.

## Test

```sh
make test
```

## Build

```sh
make build
```

## Format

```sh
make fmt
```

## Try the API

```sh
curl http://localhost:8080/
```

Expected response:

```json
{"status":"ok","message":"Hello, World!!!"}
```

## Configuration

The application has built-in defaults, so local configuration is optional. You can override settings with either `config.yaml`, environment variables, or a local `.env` file.

For YAML configuration, copy the example file:

```sh
cp config.example.yaml config.yaml
```

Example config:

```yaml
app:
  environment: development

server:
  host: 0.0.0.0
  port: 8080
```

For environment-based configuration, copy the example environment file:

```sh
cp .env.example .env
```

Example `.env`:

```env
APP_APP_ENVIRONMENT=development
APP_SERVER_HOST=0.0.0.0
APP_SERVER_PORT=8080
```

Environment variables use the `APP_` prefix and replace dots with underscores:

```sh
APP_APP_ENVIRONMENT=production APP_SERVER_PORT=3000 go run ./cmd/server
```

Real environment variables take precedence over values from `.env`. The local `.env` file is ignored by Git; commit changes to `.env.example` when adding new supported variables.

## Logging

Each request is logged by the HTTP middleware with structured fields:

- `method`
- `path`
- `status`
- `body_size`
- `client_ip`

Production mode uses Zap production logging. Other environments use Zap development logging.

## Architecture

See [ARCHITECTURE.md](./ARCHITECTURE.md) for the DDD package layout and request-flow diagrams.
