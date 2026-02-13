package url

import (
	"crypto/sha256"
)

func (s *Service) GenerateShortCode(url, username string) string {
	// Debe variar por usuario para permitir que varios usuarios acorten la misma URL
	// sin colisionar en el índice único de short_code.
	h := sha256.Sum256([]byte(url + "|" + username))

	const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 7

	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = base62[int(h[i])%len(base62)]
	}

	return string(code)

}
