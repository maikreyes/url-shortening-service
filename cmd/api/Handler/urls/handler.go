package handler

import "url-shortening-service/pkg/ports"

type Handler struct {
	UrlService  ports.UrlService
	UserService ports.UserService
	Host        string
}

func NewHandler(urlService ports.UrlService, userService ports.UserService, host string) *Handler {
	return &Handler{
		UrlService:  urlService,
		UserService: userService,
		Host:        host,
	}
}
