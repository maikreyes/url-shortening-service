package url

import (
	"crypto/sha256"
)

func (s *Service) generateShortCode(url string) string {

	h := sha256.Sum256([]byte(url))

	const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 7

	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = base62[int(h[i])%len(base62)]
	}

	return string(code)

}
