package user

import "url-shortening-service/pkg/ports"

type Handler struct {
	Service ports.UserService
	Host    string
}

func NewHandler(service ports.UserService, host string) *Handler {
	return &Handler{
		Service: service,
		Host:    host,
	}
}
