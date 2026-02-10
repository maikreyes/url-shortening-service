package handler

import "fmt"

func (h *Handler) DeleteData(code string) error {

	err := h.service.DeleteShortUrl(code)

	if err != nil {
		return err
	}

	fmt.Printf("El codigo %s ha sido eliminado\n", code)

	return nil
}
