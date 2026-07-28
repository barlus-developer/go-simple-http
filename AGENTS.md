# AGENTS.md

AI-only operating instructions for coding agents in this repository. Optimize for accurate, minimal, well-tested changes. Keep the README human-friendly; put agent-specific workflow, constraints, and performance notes here.

## Fast Context

- Project: `go-simple-http`
- Language: Go
- Module: `github.com/barlus-developer/go-simple-http`
- Current API: `GET /`
- Expected response: `{"status":"ok","message":"Hello, World!!!"}`
- Main command: `make`
- Test command: `make test`
- Build command: `make build`
- Format command: `make fmt`

## Agent Priority Order

1. Preserve current behavior unless the user explicitly asks to change it.
2. Keep edits scoped to the request.
3. Follow existing package boundaries.
4. Add or update tests when behavior changes.
5. Run the narrowest useful verification first, then `make test` and `make build` before handoff when practical.
6. Do not overwrite unrelated local changes, local config, or generated documentation unless requested.

## Package Boundaries

- `cmd/server`: process entrypoint, `net/http` server setup, graceful shutdown.
- `internal/bootstrap`: dependency wiring only.
- `internal/domain`: domain data structures and domain behavior.
- `internal/application`: use cases and application services.
- `internal/interfaces/http/handler`: Gin handlers.
- `internal/interfaces/http/middleware`: HTTP middleware.
- `internal/interfaces/http/router`: route registration and Gin engine setup.
- `internal/infrastructure/config`: Viper, `config.yaml`, and system environment variable config loading.
- `internal/infrastructure/logger`: Zap logger construction.

Do not move behavior across these layers without a clear reason. HTTP should depend on application services; application should depend on domain; infrastructure should not contain request handling or business behavior.

## Current Runtime Contract

`GET /` returns HTTP 200 and this JSON body:

```json
{"status":"ok","message":"Hello, World!!!"}
```

Treat this as a compatibility contract. If a task changes it, update:

- Application tests in `internal/application`.
- Router/API tests in `internal/interfaces/http/router`.
- README examples.
- Any architecture notes affected by the new flow.

## Configuration Contract

Defaults are built into `internal/infrastructure/config`:

- `app.debug`: `false`
- `server.host`: `0.0.0.0`
- `server.port`: `8080`

Supported override sources:

- `config.yaml`
- `config/config.yaml`
- Real system environment variables with the `APP_` prefix (no `.env` file support)

Environment key mapping: `app.debug` becomes `APP_APP_DEBUG`; `server.port` becomes `APP_SERVER_PORT`.

Real environment variables take precedence over `config.yaml`. Do not commit secrets or machine-specific config changes.

## Testing Map

Place tests near the behavior:

- Config loading, precedence, and errors: `internal/infrastructure/config`.
- Logger construction: `internal/infrastructure/logger`.
- Application service behavior: `internal/application`.
- Handler behavior: `internal/interfaces/http/handler`.
- Route behavior and HTTP response contract: `internal/interfaces/http/router`.
- Middleware logging behavior: `internal/interfaces/http/middleware`.
- Process startup/shutdown: prefer testing extracted logic; avoid brittle signal-process tests unless necessary.

## Commands

Use these commands from the repository root:

```sh
make          # run service
make test     # go test ./...
make build    # go build ./cmd/server
make fmt      # gofmt all Go files outside .git
make docker   # docker build -t go-simple-http .
```

When only Go files changed, run `make fmt` before tests. For small changes, a package-level `go test ./path/...` is acceptable while iterating, but final verification should prefer `make test` and `make build`.

## Documentation Rules

- `README.md` is for humans: explain what the project does, how to run it, how to configure it, and where to look next.
- `AGENTS.md` is for AI agents: keep it direct, operational, and optimized for quick correct edits.
- `ARCHITECTURE.md` holds structural diagrams and request-flow details.
- If API behavior, configuration, commands, or package flow changes, update the relevant docs in the same change.

## Git Rules

- Keep commits focused.
- Do not revert or rewrite user changes unless explicitly asked.
- Do not delete generated project documentation unless requested.
- If an AI coding agent contributes to a commit, add a co-author trailer identifying that specific agent — do not assume which agent you are or use another agent's name/email. Examples:

```text
Co-authored-by: Codex <codex@openai.com>
Co-authored-by: Claude <noreply@anthropic.com>
```

## Performance Notes For Agents

- Read the smallest set of files needed to understand the change.
- Prefer `rg` and package-local tests while exploring.
- Avoid broad refactors in this small service.
- Prefer standard library and existing dependencies over new packages.
- Keep new abstractions rare; add them only when they remove real duplication or clarify a growing boundary.
- Preserve simple behavior and explicit wiring over clever indirection.
