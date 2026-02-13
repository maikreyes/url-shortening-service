package handler

import "fmt"

func (h *Handler) PutData(code string, url string) error {

	newCode, err := h.service.UpdateShortUrl(code, "", url)

	if err != nil {
		return err
	}

	fmt.Printf("Url shorter updated with this new shorted url: %s:%s/%s\n", h.Host, h.Port, newCode)

	return nil
}
