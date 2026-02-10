# Contributing

Thanks for your interest in contributing!

## Getting Started

### Prerequisites

- Go (see `go.mod` for the version)
- A MySQL instance
- A `.env` file at the project root (the app loads it via `godotenv`)

Example `.env`:

```dotenv
CONNECTION_STRING="user:pass@tcp(127.0.0.1:3306)/UrlShorteningService?parseTime=true"
HOST="localhost"
PORT="8080"
```

### Run the API

```bash
go run ./cmd api
```

### Run the CLI

```bash
go run ./cmd cli fetch -code abc123
```

## Development Guidelines

### Code style

- Keep changes focused and avoid unrelated refactors.
- Prefer clear names and small functions.
- Keep behavior consistent with existing patterns in `cmd/api` and `internal/service`.

### Project structure

- API entrypoint: `cmd/api`
- CLI entrypoint: `cmd/cli`
- Domain/ports: `internal/domain`, `internal/ports`
- Repository: `internal/repository/url`
- Service logic: `internal/service/url`

### Database

- The API performs a basic migration on startup (creates the `urls` table if missing).
- If you change the schema, update the migration accordingly and keep it backwards-compatible when possible.

## Submitting Changes

### Branches

- Create a feature branch from `main`.

### Commit messages

- Use concise, descriptive messages (e.g., `api: add stats endpoint`).

### Pull Requests

A good PR includes:

- What changed and why
- How to test it (commands + example requests)
- Any behavior differences or breaking changes

If you add or change API behavior, please update `README.md` with the new routes and examples.

## Reporting Issues

When opening an issue, include:

- Steps to reproduce
- Expected vs. actual behavior
- Go version and OS
- Relevant logs/output
