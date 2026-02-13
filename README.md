# URL Shortening Service

Read this in: **English** · [Español](README.es.md)

![Go Version](https://img.shields.io/badge/go-1.25.5-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)
![CI](https://img.shields.io/badge/ci-not_configured-lightgrey?style=flat-square)

A URL shortening service with a REST API and a CLI.

This repository is part of the roadmap.sh project challenge:
<https://roadmap.sh/projects/url-shortening-service>

Quick links: [Contributing](CONTRIBUTING.md) · [License](LICENSE)

## Table of Contents

- [Overview](#overview)
- [Challenge Requirements (roadmapsh)](#challenge-requirements-roadmapsh)
- [Current Implementation Status (this repository)](#current-implementation-status-this-repository)
- [Tech Stack](#tech-stack)
- [Configuration](#configuration)
- [Quickstart](#quickstart)
- [Deploy to Vercel](#deploy-to-vercel)
- [API (as implemented)](#api-as-implemented)
- [Examples (cURL)](#examples-curl)
- [CLI](#cli)
- [Notes](#notes)

## Overview

The service lets you create short codes for long URLs and later resolve them back. It also supports updating and deleting existing short URLs.

Additionally, the same storage can be used to save a Discord webhook URL and expose a GitHub-compatible webhook endpoint that relays selected GitHub events to Discord.

## Challenge Requirements (roadmap.sh)

The challenge asks for a RESTful API that supports:

- Create a new short URL
- Retrieve an original URL from a short URL
- Update an existing short URL
- Delete an existing short URL
- Get statistics for a short URL (e.g., number of times accessed)

### Required Endpoints (as specified by the challenge)

1. Create short URL

- `POST /shorten`
- Request body:

```json
{
  "url": "https://www.example.com/some/long/url"
}
```

- Responses:
  - `201 Created` with the newly created short URL resource
  - `400 Bad Request` on validation errors

Example response (from the challenge statement):

```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

1. Retrieve original URL

- `GET /shorten/:code`
- Responses:
  - `200 OK` with the short URL resource
  - `404 Not Found` if the short URL does not exist

1. Update short URL

- `PUT /shorten/:code`
- Request body:

```json
{
  "url": "https://www.example.com/some/updated/url"
}
```

- Responses:
  - `200 OK` with the updated short URL resource
  - `400 Bad Request` on validation errors
  - `404 Not Found` if the short URL does not exist

1. Delete short URL

- `DELETE /shorten/:code`
- Responses:
  - `204 No Content` if deleted
  - `404 Not Found` if the short URL does not exist

1. Get statistics

- `GET /shorten/:code/stats`
- Responses:
  - `200 OK` with the short URL resource + `accessCount`
  - `404 Not Found` if the short URL does not exist

Example (from the challenge statement):

```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z",
  "accessCount": 10
}
```

## Current Implementation Status (this repository)

Implemented:

- Create / retrieve / update / delete short URLs
- Redirect from a short code to the original URL
- Access counter increment on redirect
- Statistics endpoint (`/api/v2/shorten/:code/stats`) returning `accessCount`
- GitHub webhook endpoint that relays GitHub events to Discord (using a stored Discord webhook URL)

Also note these behavioral differences vs. the challenge spec:

- The CRUD API is mounted under `/api/v1` (not at the root)
- Create/update accept the URL via an HTTP header named `url` (not JSON body)
- Delete returns `200 OK` with a message (the challenge expects `204 No Content`)

## Tech Stack

- Go
- Gin (HTTP)
- GORM + MySQL / Postgres
- `.env` loading via `godotenv`

## Configuration

Environment variables (see `.env`):

- `CONNECTION_STRING`: DSN for the selected driver
- `DB_DRIVER`: database driver (`mysql` or `postgres`)
- `HOST`: server host (e.g., `localhost`)
- `PORT`: server port (e.g., `8080`)
- `TABLE_NAME`: table used by the repository (default in this repo: `api_responses`)
- `ENVIRONMENT`: set to `production` to use production-friendly defaults

Notes:

- When deploying to Vercel, do not use `127.0.0.1` / `localhost` in `CONNECTION_STRING`. Use a managed database reachable from Vercel.
- In Vercel Serverless Functions you typically do not need `HOST`/`PORT` (requests are handled by the platform).

Example:

```dotenv
DB_DRIVER="mysql"
CONNECTION_STRING="user:pass@tcp(127.0.0.1:3306)/UrlShorteningService?parseTime=true"
HOST="localhost"
PORT="8080"
TABLE_NAME="api_responses"
```

Postgres example:

```dotenv
DB_DRIVER="postgres"
CONNECTION_STRING="host=127.0.0.1 user=postgres password=postgres dbname=UrlShorteningService port=5432 sslmode=disable"
HOST="localhost"
PORT="8080"
TABLE_NAME="api_responses"
```

## Quickstart

1) Start MySQL or Postgres (and create the database if needed).
2) Configure `.env`.
3) Run the API:

```bash
go run ./cmd api
```

The service starts at `http://HOST:PORT` and performs a basic migration on startup (creates the `api_responses` table if it does not exist).

Notes:

- `DB_DRIVER` is required (`mysql` or `postgres`).
- `TABLE_NAME` is required. The migration creates the table specified by `TABLE_NAME` if it does not exist.

## Deploy to Vercel

This project can run on Vercel using Go Serverless Functions.

How it works:

- The function entrypoint is in `api/index.go`.
- All routes are rewritten to the function via `vercel.json` (it preserves the original path via `?path=...`).
- The HTTP router is built without binding to a port (see `cmd/api/router/router.go`).

Steps:

1) In Vercel Dashboard → Project → Settings → Environment Variables, configure:

- `DB_DRIVER` (`mysql` or `postgres`)
- `CONNECTION_STRING` (a remote DSN; do not use `127.0.0.1`)
- `TABLE_NAME`

2) Deploy.

Implementation note:

- Vercel's Go build wraps the function, which makes Go's `internal/` import rules inconvenient for serverless builds. For this reason, the reusable code is also available under `pkg/`.

## API (as implemented)

Base:

- CRUD endpoints: `/api/v1`
- Redirect endpoint: `/` (root)
- GitHub webhook relay endpoint: `/` (root)

### Redirect

- `GET /:code` → redirects to the original URL (HTTP `301 Moved Permanently`) and increments `accessCount`.
  Note: `301` may be cached by browsers; for testing repeated hits, prefer `curl` or a private window.

### Health

- `GET /api/v1/` → returns a simple JSON message.

### CRUD

- `POST /api/v1/shorten`
  - Input: Header `url: <url>`
  - Optional: Header `webhook: true|false`
  - Output: `201 Created` with `{ "url": "<shortCode>" }`
  - If `webhook: true`, the response becomes `{ "url": "<shortCode>/webhook" }`
    (note: this is a path suffix, not a fully-qualified URL).

  Note: if `webhook` is present, it must be a valid boolean (`true`/`false`).

- `GET /api/v1/shorten/:code`
  - Output: `200 OK` with `{ id, url, shortCode, createdAt, updatedAt }`

- `PUT /api/v1/shorten/:code`
  - Input: Header `url: <new_url>`
  - Output: `200 OK` with a message containing the new short code
  - Note: the current implementation does not validate that `:code` existed before updating.

- `DELETE /api/v1/shorten/:code`
  - Output: `200 OK` with a message
  - Note: the current implementation does not validate that `:code` existed before deleting.

### GitHub → Discord webhook relay

This repository also includes a simple webhook relay:

1) Store a Discord webhook URL by creating a short code with `webhook: true`.
2) GitHub sends events to `POST /:code/webhook`.
3) The service converts the GitHub event into a Discord embed and posts it to the stored Discord webhook URL.

Endpoint:

- `POST /:code/webhook`
  - Requires header: `X-GitHub-Event: <event>`
  - Body: GitHub JSON payload (the handler extracts a few fields)
  - Response: `200 OK` with `{ "status": "sent" }` (or an error JSON)

Supported events (`X-GitHub-Event`):

- `ping`
- `issues`
- `create` (branch created)
- `push`
- `pull_request`

## Examples (cURL)

Create:

```bash
curl -i -X POST \
  -H 'url: https://www.example.com/some/long/url' \
  http://localhost:8080/api/v1/shorten
```

Create a GitHub webhook relay (store a Discord webhook URL):

```bash
curl -i -X POST \
  -H 'url: https://discord.com/api/webhooks/XXX/YYY' \
  -H 'webhook: true' \
  http://localhost:8080/api/v1/shorten
```

Then configure GitHub to send webhooks to:

- `http://localhost:8080/<code>/webhook`

Retrieve:

```bash
curl -i http://localhost:8080/api/v1/shorten/abc123
```

Redirect:

```bash
curl -i http://localhost:8080/abc123
```

Stats:

```bash
curl -i http://localhost:8080/api/v2/shorten/abc123/stats
```

Manually test the webhook endpoint (example `ping` event):

```bash
curl -i -X POST \
  -H 'X-GitHub-Event: ping' \
  -H 'Content-Type: application/json' \
  -d '{"repository":{"name":"demo"},"sender":{"login":"octocat"}}' \
  http://localhost:8080/<code>/webhook
```

Update:

```bash
curl -i -X PUT \
  -H 'url: https://www.example.com/some/updated/url' \
  http://localhost:8080/api/v1/shorten/abc123
```

Delete:

```bash
curl -i -X DELETE http://localhost:8080/api/v1/shorten/abc123
```

## CLI

The binary also supports a CLI mode that operates on the same database.

```bash
go run ./cmd cli post   -url  https://www.example.com
go run ./cmd cli fetch  -code abc123
go run ./cmd cli put    -code abc123 -url https://www.example.com/updated
go run ./cmd cli delete -code abc123
```

## Notes

- Short code generation is currently deterministic (SHA-256 + Base62, length 7). The same normalized URL produces the same short code.
- URL normalization adds `https://` when the scheme is missing.
- When redirecting, if the stored URL has no scheme, the server assumes `https://`.
- The database table used by the repository is configured via `TABLE_NAME` (the default `.env` uses `api_responses`). `TABLE_NAME` is required.
- The database driver is selected via `DB_DRIVER` (`mysql` or `postgres`).
