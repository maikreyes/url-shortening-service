package handler

import "url-shortening-service/internal/ports"

type Handler struct {
	Service ports.UrlService
}

func NewHandler(service ports.UrlService) *Handler {
	return &Handler{
		Service: service,
	}
}
