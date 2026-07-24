# Architecture

This project uses a small DDD-style layout. The current domain is intentionally minimal because the API only returns a static health response, but the package boundaries are ready for additional use cases.

## Layers

- `cmd/server`: starts the process, creates the HTTP server, and handles graceful shutdown.
- `internal/bootstrap`: wires configuration, logging, services, handlers, middleware, and routes.
- `internal/domain`: contains business objects.
- `internal/application`: contains use cases and application services that operate on domain objects.
- `internal/interfaces/http`: contains Gin handlers, middleware, and router setup.
- `internal/infrastructure`: contains technical concerns such as configuration and logging.

## Package Relationship

```mermaid
flowchart TD
    Main[cmd/server] --> Bootstrap[internal/bootstrap]
    Bootstrap --> Config[internal/infrastructure/config]
    Bootstrap --> Logger[internal/infrastructure/logger]
    Bootstrap --> AppService[internal/application/health]
    Bootstrap --> Handler[internal/interfaces/http/handler]
    Bootstrap --> Router[internal/interfaces/http/router]
    Router --> Middleware[internal/interfaces/http/middleware]
    Router --> Handler
    Handler --> AppService
    AppService --> Domain[internal/domain/health]
```

## Request Flow Algorithm

```mermaid
sequenceDiagram
    participant Client
    participant Gin as Gin Router
    participant Log as Logger Middleware
    participant Handler as Health Handler
    participant Service as Health Service
    participant Domain as Domain Status

    Client->>Gin: GET /
    Gin->>Log: execute middleware chain
    Log->>Handler: c.Next()
    Handler->>Service: Status()
    Service->>Domain: create Status
    Domain-->>Service: status object
    Service-->>Handler: status object
    Handler-->>Client: 200 JSON response
    Log->>Log: record method, path, status, body_size, client_ip
```

## Startup Algorithm

```mermaid
flowchart TD
    Start([Process starts]) --> LoadConfig[Load config with Viper]
    LoadConfig --> CreateLogger[Create Zap logger]
    CreateLogger --> BuildService[Create health application service]
    BuildService --> BuildHandler[Create HTTP handler]
    BuildHandler --> BuildRouter[Create Gin router and middleware]
    BuildRouter --> StartServer[Start net/http server]
    StartServer --> WaitSignal[Wait for SIGINT or SIGTERM]
    WaitSignal --> Shutdown[Gracefully shutdown server]
    Shutdown --> Stop([Process exits])
```

## Dependency Direction

Dependencies point inward toward the application and domain layers:

```mermaid
flowchart LR
    HTTP[interfaces/http] --> Application[application]
    Application --> Domain[domain]
    Bootstrap[bootstrap] --> HTTP
    Bootstrap --> Application
    Bootstrap --> Infrastructure[infrastructure]
    Main[cmd/server] --> Bootstrap
```

The router knows about HTTP handlers and middleware. The handler knows about the application service interface. The application service returns a domain object. Infrastructure packages do not contain request handling or business behavior.
