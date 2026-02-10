package handler

import "fmt"

func (h *Handler) PostData(url string) error {

	shortcode, err := h.service.CreateShortUrl(url)

	if err != nil {
		return err
	}

	fmt.Printf("Your shorted url is: %s:%s/%s\n", h.Host, h.Port, shortcode)

	return nil
}
