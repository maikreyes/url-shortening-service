package handler

import (
	"url-shortening-service/internal/ports"
)

type Handler struct {
	service ports.UrlService
	Host    string
	Port    string
}

func NewHandler(service ports.UrlService, host string, port string) *Handler {
	return &Handler{
		service: service,
		Host:    host,
		Port:    port,
	}
}
