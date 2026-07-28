# go-simple-http

A tiny Go HTTP service with one job: respond to `GET /` with a simple JSON message.

```json
{"status":"ok","message":"<random meme message, changes each request>"}
```

The project is intentionally small, but it is laid out like a production service so it can grow without turning into a single-package demo.

## What Is Inside

- Gin for HTTP routing and middleware.
- Zap for structured logs.
- Viper for configuration, sourced from `config.yaml` or system environment variables.
- A standard `net/http` server with graceful shutdown.
- A DDD-style `internal` layout with separate application, domain, HTTP, bootstrap, and infrastructure packages.

## Requirements

- Go `1.26.5`, or a compatible version for the `go.mod` setting.
- Docker, only if you want to build or run the container image.

## Run It

Start the API:

```sh
make
```

The server listens on `0.0.0.0:8080` by default.

Try it from another terminal:

```sh
curl http://localhost:8080/
```

Expected response:

```json
{"status":"ok","message":"<random meme message, changes each request>"}
```

## Common Commands

```sh
make          # run the server
make test     # run all tests
make build    # build the server binary
make fmt      # format Go files
make docker   # build the Docker image
```

## Configuration

You can run the service without any local config because defaults are built in:

```yaml
app:
  debug: false

server:
  host: 0.0.0.0
  port: 8080
```

To override those values, use `config.yaml` or real system environment variables.

Environment variables use the `APP_` prefix. Dots in config keys become underscores:

```sh
APP_APP_DEBUG=true APP_SERVER_PORT=3000 go run ./cmd/server
```

Real environment variables win over `config.yaml`. Set `APP_APP_DEBUG=true` (or `app.debug: true` in `config.yaml`) for verbose, human-readable development logs and Gin's debug mode. Leave it unset/`false` for production-style JSON logs and Gin's release mode.

## Docker

Build the image:

```sh
make docker
```

Run it:

```sh
docker run --rm -p 8080:8080 go-simple-http
```

Run it on another port:

```sh
docker run --rm -p 3000:3000 -e APP_SERVER_PORT=3000 go-simple-http
```

## Project Layout

```text
cmd/server              process entrypoint and graceful shutdown
internal/bootstrap      dependency wiring
internal/domain         domain models
internal/application    use cases and application services
internal/interfaces     HTTP handlers, middleware, and router setup
internal/infrastructure config, logging, and technical adapters
```

For more detail, see [ARCHITECTURE.md](./ARCHITECTURE.md).

## Notes For Contributors

Keep the public API small unless the project goal changes. When behavior changes, update or add tests near the package you touched, then run:

```sh
make test
make build
```
