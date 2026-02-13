package handler

import "url-shortening-service/pkg/ports"

type Handler struct {
	Service ports.UrlService
	Host    string
}

func NewHandler(service ports.UrlService, host string) *Handler {
	return &Handler{
		Service: service,
		Host:    host,
	}
}
