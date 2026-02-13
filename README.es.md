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
- [Deploy en Vercel](#deploy-en-vercel)
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

- Registro y login de usuarios (JWT)
- Crear / obtener / actualizar / eliminar URLs cortas (protegido por JWT)
- Redirección desde un código corto hacia la URL original
- Contador de accesos (incrementa `accessCount` al redirigir)
- Endpoint de estadísticas (`/api/v2/shorten/:code/stats`) que retorna `accessCount`
- Listado de URLs por usuario (`/api/v3/:username/urls`) (protegido por JWT)
- Endpoint webhook de GitHub que reenvía eventos de GitHub hacia Discord (usando una URL de webhook de Discord almacenada)
- Swagger UI en `/swagger/*any`

Diferencias de comportamiento vs. el reto:

- La API CRUD está montada bajo `/api/v1` y requiere JWT (no en la raíz)
- Crear/actualizar reciben la URL vía un header HTTP llamado `url` (no JSON)
- Crear responde con una URL compuesta `{ "url": "<HOST>/<code>" }` (no el recurso completo)
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
- `HOST`: host/dirección usada para construir las URLs cortas retornadas (ej. `localhost:8080` o `tu-app.vercel.app`)
- `PORT`: puerto usado al ejecutar la API localmente (ej. `8080`)
- `URL_TABLE_NAME`: nombre de la tabla de URLs (requerido)
- `USER_TABLE_NAME`: nombre de la tabla de usuarios (requerido)
- `JWT_SECRET`: secreto HMAC para firmar/validar JWT (requerido para `/login` y cualquier ruta `/api/*`)
- `ENVIRONMENT`: usa `production` para defaults más amigables a producción

Notas:

- Para desplegar en Vercel, no uses `127.0.0.1` / `localhost` en `CONNECTION_STRING`. Usa una base de datos administrada accesible desde Vercel.
- En Serverless Functions de Vercel normalmente no necesitas `HOST`/`PORT` (las requests las maneja la plataforma).

Ejemplo (MySQL):

```dotenv
DB_DRIVER="mysql"
CONNECTION_STRING="user:pass@tcp(127.0.0.1:3306)/UrlShorteningService?parseTime=true"
HOST="localhost:8080"
PORT="8080"
URL_TABLE_NAME="api_responses"
USER_TABLE_NAME="users"
JWT_SECRET="change-me"
```

Ejemplo (Postgres):

```dotenv
DB_DRIVER="postgres"
CONNECTION_STRING="host=127.0.0.1 user=postgres password=postgres dbname=UrlShorteningService port=5432 sslmode=disable"
HOST="localhost:8080"
PORT="8080"
URL_TABLE_NAME="api_responses"
USER_TABLE_NAME="users"
JWT_SECRET="change-me"
```

## Inicio rápido

1. Inicia MySQL o Postgres (y crea la base de datos si hace falta).
1. Configura `.env`.
1. Ejecuta la API:

```bash
go run ./cmd api
```

El servicio inicia en `http://HOST:PORT` y realiza una migración básica al arrancar (crea la tabla indicada si no existe).

Notas:

- `DB_DRIVER` es requerido (`mysql` o `postgres`).
- `URL_TABLE_NAME` y `USER_TABLE_NAME` son requeridos. La migración crea esas tablas si no existen.
- `JWT_SECRET` es requerido para emitir/validar JWT.

## Deploy en Vercel

Este proyecto puede correr en Vercel usando Go Serverless Functions.

Cómo funciona:

- El entrypoint de la función está en `api/index.go`.
- Todas las rutas se reescriben a la función vía `vercel.json` (preserva el path original usando `?path=...`).
- El router HTTP se construye sin abrir un puerto (ver `cmd/api/router/router.go`).

Pasos:

1. En Vercel Dashboard → Project → Settings → Environment Variables, configura:

- `DB_DRIVER` (`mysql` o `postgres`)
- `CONNECTION_STRING` (un DSN remoto; no uses `127.0.0.1`)
- `URL_TABLE_NAME`
- `USER_TABLE_NAME`
- `JWT_SECRET`

1. Despliega.

Nota de implementación:

- El build de Go en Vercel envuelve la función, lo cual vuelve incómodas las reglas de import de `internal/` en este contexto. Por eso, el código reutilizable también está disponible en `pkg/`.

## API (tal como está implementada)

Base:

- Swagger UI: `/swagger/*any`
- Endpoints públicos: `/`, `/:code`, `/:code/webhook`, `/login`, `/register`
- Endpoints protegidos por JWT: `/api/*`

Nota: los paths en este README son relativos al host (por ejemplo, `http://localhost:8080`).

### Redirección

- `GET /:code` → redirige a la URL original (HTTP `301 Moved Permanently`) e incrementa `accessCount`.
  Nota: `301` puede quedar en caché en el navegador; para probar múltiples accesos, usa `curl` o una ventana privada.

### Bienvenida

- `GET /` → retorna un JSON simple.

### Auth

- `POST /register` → registra un usuario.
  - Body (JSON): `{ "username": "...", "email": "...", "password": "..." }`
  - Respuesta: `201 Created` (nota: la implementación actual retorna el password en claro).

- `POST /login` → autentica y retorna un JWT (también setea cookie `access_token`).
  - Body (JSON): `{ "email": "...", "password": "..." }`
  - Respuesta: `200 OK` con `{ "token": { "token": "<jwt>", "email": "<email>" } }`

Puedes enviar el JWT como:

- `Authorization: Bearer <token>`
- Cookie `access_token=<token>`

### CRUD

- `POST /api/v1/shorten`
  - Requiere JWT
  - Entrada: Header `url: <url>`
  - Opcional: Header `webhook: true|false`
  - Salida: `201 Created` con `{ "url": "<HOST>/<shortCode>" }`
  - Si `webhook: true`, la respuesta se convierte en `{ "url": "<HOST>/<shortCode>/webhook" }`
  - Nota: si `webhook` está presente, debe ser un boolean válido (`true`/`false`).

- `GET /api/v1/shorten/:code`
  - Requiere JWT
  - Salida: `200 OK` con `{ ID, url, shortCode, createdAt, updatedAt }`
  - Nota: valida que el recurso pertenezca al usuario autenticado.

- `PUT /api/v1/shorten/:code`
  - Requiere JWT
  - Entrada: Header `url: <new_url>`
  - Salida: `200 OK` con un mensaje que contiene el nuevo short code
  - Nota: la implementación actual no valida ownership para actualizar.

- `DELETE /api/v1/shorten/:code`
  - Requiere JWT
  - Salida: `200 OK` con un mensaje
  - Nota: la implementación actual no valida ownership para eliminar.

### Estadísticas

- `GET /api/v2/shorten/:code/stats`
  - Requiere JWT
  - Salida: `200 OK` con el recurso almacenado incluyendo `accessCount`.

### URLs por usuario

- `GET /api/v3/:username/urls`
  - Requiere JWT
  - `:username` debe coincidir con el usuario autenticado.

### Relay webhook GitHub → Discord

Este repo incluye un relay simple:

1. Guarda una URL de webhook de Discord creando un short code con `webhook: true`.
1. GitHub envía eventos a `POST /:code/webhook`.
1. El servicio convierte el evento de GitHub en un embed de Discord y lo publica en el webhook de Discord guardado.

Endpoint:

- `POST /:code/webhook`
  - Requiere header: `X-GitHub-Event: <event>`
  - Headers opcionales: `avatarUrl`, `informanteName`
  - Body: payload JSON de GitHub
  - Respuesta: `200 OK` con `{ "status": "sent", "avatar": "..." }` (o un JSON de error)

Eventos soportados (`X-GitHub-Event`):

- `ping`
- `issues`
- `create` (creación de branch)
- `push`
- `pull_request`

## Ejemplos (cURL)

Registrar:

```bash
curl -i -X POST \
  -H 'Content-Type: application/json' \
  -d '{"username":"demo","email":"demo@example.com","password":"secret123"}' \
  http://localhost:8080/register
```

Login (copia el token de la respuesta, o usa la cookie `access_token`):

```bash
curl -i -X POST \
  -H 'Content-Type: application/json' \
  -d '{"email":"demo@example.com","password":"secret123"}' \
  http://localhost:8080/login
```

Crear:

```bash
curl -i -X POST \
  -H 'Authorization: Bearer <token>' \
  -H 'url: https://www.example.com/some/long/url' \
  http://localhost:8080/api/v1/shorten
```

Crear un relay de webhook GitHub (guardar una URL de webhook de Discord):

```bash
curl -i -X POST \
  -H 'Authorization: Bearer <token>' \
  -H 'url: https://discord.com/api/webhooks/XXX/YYY' \
  -H 'webhook: true' \
  http://localhost:8080/api/v1/shorten
```

Luego configura GitHub para enviar webhooks a:

- `http://localhost:8080/<code>/webhook`

Obtener:

```bash
curl -i \
  -H 'Authorization: Bearer <token>' \
  http://localhost:8080/api/v1/shorten/abc123
```

Redirección:

```bash
curl -i http://localhost:8080/abc123
```

Estadísticas:

```bash
curl -i \
  -H 'Authorization: Bearer <token>' \
  http://localhost:8080/api/v2/shorten/abc123/stats
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
  -H 'Authorization: Bearer <token>' \
  -H 'url: https://www.example.com/some/updated/url' \
  http://localhost:8080/api/v1/shorten/abc123
```

Eliminar:

```bash
curl -i -X DELETE \
  -H 'Authorization: Bearer <token>' \
  http://localhost:8080/api/v1/shorten/abc123
```

Listar URLs de un usuario:

```bash
curl -i \
  -H 'Authorization: Bearer <token>' \
  http://localhost:8080/api/v3/demo/urls
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

- La generación del short code es determinística por usuario (SHA-256 + Base62, longitud 7) e incluye el username como “salt”.
- La normalización de URL agrega `https://` cuando falta el esquema.
- Al redirigir, si la URL guardada no tiene esquema, el servidor asume `https://`.
- Las tablas se configuran vía `URL_TABLE_NAME` y `USER_TABLE_NAME`.
- El driver de base de datos se selecciona vía `DB_DRIVER` (`mysql` o `postgres`).
