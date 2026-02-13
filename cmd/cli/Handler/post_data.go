package handler

import (
	"fmt"
	"url-shortening-service/pkg/domain"

	"github.com/google/uuid"
)

func (h *Handler) PostData(url string) error {

	var u domain.User

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	isWebhook := false

	shortcode, err := h.service.CreateShortUrl(url, u.Username, isWebhook, u.ID)

	if err != nil {
		return err
	}

	fmt.Printf("Your shorted url is: %s:%s/%s\n", h.Host, h.Port, shortcode)

	return nil
}
