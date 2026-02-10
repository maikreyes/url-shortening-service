package handler

import (
	"fmt"
)

func (h *Handler) FetchData(code string) error {
	url, err := h.service.GetShortUrl(code)

	if err != nil || url == nil {
		fmt.Printf("El codigo %s no existe\n", code)
		if err == nil {
			return fmt.Errorf("short code %s not found", code)
		}
		return err
	}

	fmt.Printf("id: %d\nurl: %s\nshortcode: %s\n", url.ID, url.Url, url.ShortCode)

	return nil
}
