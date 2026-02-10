# URL Shortening Service

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
- [API (as implemented)](#api-as-implemented)
- [Examples (cURL)](#examples-curl)
- [CLI](#cli)
- [Notes](#notes)

## Overview

The service lets you create short codes for long URLs and later resolve them back. It also supports updating and deleting existing short URLs.

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

Not implemented yet (gap vs. challenge):

- Statistics endpoint (`/shorten/:code/stats`) and access counter tracking

Also note these behavioral differences vs. the challenge spec:

- The CRUD API is mounted under `/api/v1` (not at the root)
- Create/update accept the URL via an HTTP header named `url` (not JSON body)
- Delete returns `200 OK` with a message (the challenge expects `204 No Content`)

## Tech Stack

- Go
- Gin (HTTP)
- GORM + MySQL
- `.env` loading via `godotenv`

## Configuration

Environment variables (see `.env`):

- `CONNECTION_STRING`: MySQL DSN
- `HOST`: server host (e.g., `localhost`)
- `PORT`: server port (e.g., `8080`)

Example:

```dotenv
CONNECTION_STRING="user:pass@tcp(127.0.0.1:3306)/UrlShorteningService?parseTime=true"
HOST="localhost"
PORT="8080"
```

## Quickstart

1) Start MySQL (and create the database if needed).
2) Configure `.env`.
3) Run the API:

```bash
go run ./cmd api
```

The service starts at `http://HOST:PORT` and creates the `urls` table if it does not exist.

## API (as implemented)

Base:

- CRUD endpoints: `/api/v1`
- Redirect endpoint: `/` (root)

### Redirect

- `GET /:code` → redirects to the original URL (HTTP `301 Moved Permanently`).

### CRUD

- `POST /api/v1/shorten`
  - Input: Header `url: <url>`
  - Output: `201 Created` with `{ "url": "<shortCode>" }`

- `GET /api/v1/shorten/:code`
  - Output: `200 OK` with `{ id, url, shortCode, createdAt, updatedAt }`

- `PUT /api/v1/shorten/:code`
  - Input: Header `url: <new_url>`
  - Output: `200 OK` with a message containing the new short code

- `DELETE /api/v1/shorten/:code`
  - Output: `200 OK` with a message

## Examples (cURL)

Create:

```bash
curl -i -X POST \
  -H 'url: https://www.example.com/some/long/url' \
  http://localhost:8080/api/v1/shorten
```

Retrieve:

```bash
curl -i http://localhost:8080/api/v1/shorten/abc123
```

Redirect:

```bash
curl -i http://localhost:8080/abc123
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
- When redirecting, if the stored URL has no scheme, the server assumes `https://`.
