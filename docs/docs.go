package docs

import "github.com/swaggo/swag"

// Package docs provee una especificación Swagger mínima para que el proyecto compile
// y Swagger UI funcione incluso antes de generar documentación completa con `swag init`.
//
// Si luego ejecutas `swag init`, esta carpeta típicamente será re-generada.

const swaggerTemplate = `{
  "swagger": "2.0",
  "info": {
	"description": "API para acortar URLs, autenticar usuarios y consultar estadísticas.",
	"title": "URL Shortening Service API",
	"version": "1.0"
  },
  "basePath": "/",
  "schemes": ["http", "https"],
  "paths": {}
}`

type swaggerInfo struct{}

func (swaggerInfo) ReadDoc() string {
	return swaggerTemplate
}

func init() {
	swag.Register(swag.Name, swaggerInfo{})
}























