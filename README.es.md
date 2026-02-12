# URL Shortening Service

Leer en: [English](README.md) · **Español**

![Go Version](https://img.shields.io/badge/go-1.25.5-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)
![CI](https://img.shields.io/badge/ci-not_configured-lightgrey?style=flat-square)

Servicio de acortamiento de URLs con una API REST y una CLI.

Este repositorio es parte del reto de roadmap.sh:
<https://roadmap.sh/projects/url-shortening-service>

Enlaces rápidos: [Contributing](CONTRIBUTING.md) · [License](LICENSE)

## Tabla de Contenidos

- [Resumen](#resumen)
- [Requisitos del reto (roadmapsh)](#requisitos-del-reto-roadmapsh)
- [Estado actual de implementación (este repositorio)](#estado-actual-de-implementación-este-repositorio)
- [Stack tecnológico](#stack-tecnológico)
- [Configuración](#configuración)
- [Inicio rápido](#inicio-rápido)
- [API (tal como está implementada)](#api-tal-como-está-implementada)
- [Ejemplos (cURL)](#ejemplos-curl)
- [CLI](#cli)
- [Notas](#notas)

## Resumen

El servicio te permite crear códigos cortos para URLs largas y luego resolverlos de vuelta al URL original. También soporta actualizar y eliminar URLs cortas existentes.

Además, el mismo almacenamiento puede usarse para guardar una URL de webhook de Discord y exponer un endpoint de webhook compatible con GitHub que reenvía eventos seleccionados de GitHub hacia Discord.

## Requisitos del reto (roadmap.sh)

El reto pide una API REST que soporte:

- Crear una nueva URL corta
- Obtener la URL original a partir de la URL corta
- Actualizar una URL corta existente
- Eliminar una URL corta existente
- Obtener estadísticas para una URL corta (por ejemplo, número de accesos)

### Endpoints requeridos (según el reto)

1. Crear URL corta

- `POST /shorten`
- Cuerpo de la solicitud:

```json
{
  "url": "https://www.example.com/some/long/url"
}
```

- Respuestas:
  - `201 Created` con el recurso de la URL corta recién creada
  - `400 Bad Request` en errores de validación

Respuesta ejemplo (según el enunciado del reto):

```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

1. Obtener URL original

- `GET /shorten/:code`
- Respuestas:
  - `200 OK` con el recurso
  - `404 Not Found` si no existe

1. Actualizar URL corta

- `PUT /shorten/:code`
- Cuerpo de la solicitud:

```json
{
  "url": "https://www.example.com/some/updated/url"
}
```

- Respuestas:
  - `200 OK` con el recurso actualizado
  - `400 Bad Request` en errores de validación
  - `404 Not Found` si no existe

1. Eliminar URL corta

- `DELETE /shorten/:code`
- Respuestas:
  - `204 No Content` si se elimina
  - `404 Not Found` si no existe

1. Obtener estadísticas

- `GET /shorten/:code/stats`
- Respuestas:
  - `200 OK` con el recurso + `accessCount`
  - `404 Not Found` si no existe

Ejemplo (según el enunciado del reto):

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

## Estado actual de implementación (este repositorio)

Implementado:

- Crear / obtener / actualizar / eliminar URLs cortas
- Redirección desde un código corto hacia la URL original
- Contador de accesos (incrementa `accessCount` al redirigir)
- Endpoint de estadísticas (`/api/v2/shorten/:code/stats`) que retorna `accessCount`
- Endpoint webhook de GitHub que reenvía eventos de GitHub hacia Discord (usando una URL de webhook de Discord almacenada)

Diferencias de comportamiento vs. el reto:

- La API CRUD está montada bajo `/api/v1` (no en la raíz)
- Crear/actualizar reciben la URL vía un header HTTP llamado `url` (no JSON)
- Eliminar devuelve `200 OK` con un mensaje (el reto espera `204 No Content`)

## Stack tecnológico

- Go
- Gin (HTTP)
- GORM + MySQL / Postgres
- Carga de `.env` vía `godotenv`

## Configuración

Variables de entorno (ver `.env`):

- `CONNECTION_STRING`: DSN para el driver seleccionado
- `DB_DRIVER`: driver de base de datos (`mysql` o `postgres`)
- `HOST`: host del servidor (ej. `localhost`)
- `PORT`: puerto del servidor (ej. `8080`)
- `TABLE_NAME`: tabla usada por el repositorio (por defecto en este repo: `api_responses`)

Ejemplo (MySQL):

```dotenv
DB_DRIVER="mysql"
CONNECTION_STRING="user:pass@tcp(127.0.0.1:3306)/UrlShorteningService?parseTime=true"
HOST="localhost"
PORT="8080"
TABLE_NAME="api_responses"
```

Ejemplo (Postgres):

```dotenv
DB_DRIVER="postgres"
CONNECTION_STRING="host=127.0.0.1 user=postgres password=postgres dbname=UrlShorteningService port=5432 sslmode=disable"
HOST="localhost"
PORT="8080"
TABLE_NAME="api_responses"
```

## Inicio rápido

1) Inicia MySQL o Postgres (y crea la base de datos si hace falta).
2) Configura `.env`.
3) Ejecuta la API:

```bash
go run ./cmd api
```

El servicio inicia en `http://HOST:PORT` y realiza una migración básica al arrancar (crea la tabla indicada si no existe).

Notas:

- `DB_DRIVER` es requerido (`mysql` o `postgres`).
- `TABLE_NAME` es requerido. La migración crea la tabla especificada por `TABLE_NAME` si no existe.

## API (tal como está implementada)

Base:

- Endpoints CRUD: `/api/v1`
- Endpoint de redirección: `/` (raíz)
- Endpoint de relay de webhook GitHub: `/` (raíz)

### Redirección

- `GET /:code` → redirige a la URL original (HTTP `301 Moved Permanently`) e incrementa `accessCount`.
  Nota: `301` puede quedar en caché en el navegador; para probar múltiples accesos, usa `curl` o una ventana privada.

### Health

- `GET /api/v1/` → retorna un JSON simple.

### CRUD

- `POST /api/v1/shorten`
  - Entrada: Header `url: <url>`
  - Opcional: Header `webhook: true|false`
  - Salida: `201 Created` con `{ "url": "<shortCode>" }`
  - Si `webhook: true`, la respuesta se convierte en `{ "url": "<shortCode>/webhook" }`
    (nota: es un sufijo de ruta, no una URL completamente calificada).
  - Nota: si `webhook` está presente, debe ser un boolean válido (`true`/`false`).

- `GET /api/v1/shorten/:code`
  - Salida: `200 OK` con `{ id, url, shortCode, createdAt, updatedAt }`

- `PUT /api/v1/shorten/:code`
  - Entrada: Header `url: <new_url>`
  - Salida: `200 OK` con un mensaje que contiene el nuevo short code
  - Nota: la implementación actual no valida que `:code` existiera antes de actualizar.

- `DELETE /api/v1/shorten/:code`
  - Salida: `200 OK` con un mensaje
  - Nota: la implementación actual no valida que `:code` existiera antes de eliminar.

### Relay webhook GitHub → Discord

Este repo incluye un relay simple:

1) Guarda una URL de webhook de Discord creando un short code con `webhook: true`.
2) GitHub envía eventos a `POST /:code/webhook`.
3) El servicio convierte el evento de GitHub en un embed de Discord y lo publica en el webhook de Discord guardado.

Endpoint:

- `POST /:code/webhook`
  - Requiere header: `X-GitHub-Event: <event>`
  - Body: payload JSON de GitHub
  - Respuesta: `200 OK` con `{ "status": "sent" }` (o un JSON de error)

Eventos soportados (`X-GitHub-Event`):

- `ping`
- `issues`
- `create` (creación de branch)
- `push`
- `pull_request`

## Ejemplos (cURL)

Crear:

```bash
curl -i -X POST \
  -H 'url: https://www.example.com/some/long/url' \
  http://localhost:8080/api/v1/shorten
```

Crear un relay de webhook GitHub (guardar una URL de webhook de Discord):

```bash
curl -i -X POST \
  -H 'url: https://discord.com/api/webhooks/XXX/YYY' \
  -H 'webhook: true' \
  http://localhost:8080/api/v1/shorten
```

Luego configura GitHub para enviar webhooks a:

- `http://localhost:8080/<code>/webhook`

Obtener:

```bash
curl -i http://localhost:8080/api/v1/shorten/abc123
```

Redirección:

```bash
curl -i http://localhost:8080/abc123
```

Estadísticas:

```bash
curl -i http://localhost:8080/api/v2/shorten/abc123/stats
```

Probar manualmente el endpoint webhook (ejemplo de evento `ping`):

```bash
curl -i -X POST \
  -H 'X-GitHub-Event: ping' \
  -H 'Content-Type: application/json' \
  -d '{"repository":{"name":"demo"},"sender":{"login":"octocat"}}' \
  http://localhost:8080/<code>/webhook
```

Actualizar:

```bash
curl -i -X PUT \
  -H 'url: https://www.example.com/some/updated/url' \
  http://localhost:8080/api/v1/shorten/abc123
```

Eliminar:

```bash
curl -i -X DELETE http://localhost:8080/api/v1/shorten/abc123
```

## CLI

El binario también soporta modo CLI que opera sobre la misma base de datos.

```bash
go run ./cmd cli post   -url  https://www.example.com
go run ./cmd cli fetch  -code abc123
go run ./cmd cli put    -code abc123 -url https://www.example.com/updated
go run ./cmd cli delete -code abc123
```

## Notas

- La generación del short code es determinística (SHA-256 + Base62, longitud 7). La misma URL normalizada produce el mismo short code.
- La normalización de URL agrega `https://` cuando falta el esquema.
- Al redirigir, si la URL guardada no tiene esquema, el servidor asume `https://`.
- La tabla usada por el repositorio se configura vía `TABLE_NAME` (el `.env` por defecto usa `api_responses`). `TABLE_NAME` es requerido.
- El driver de base de datos se selecciona vía `DB_DRIVER` (`mysql` o `postgres`).
