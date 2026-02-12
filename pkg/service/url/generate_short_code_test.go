package url_test

import (
	"strings"
	"testing"
	"url-shortening-service/pkg/service/url"
)

func TestGenerateShortCode(t *testing.T) {

	s := url.NewService(nil)

	const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	isBase62 := func(code string) bool {
		return code != "" && strings.IndexFunc(code, func(r rune) bool {
			return !strings.ContainsRune(base62, r)
		}) == -1
	}

	t.Run("devuelve 7 caracteres y solo base62", func(t *testing.T) {
		code := s.GenerateShortCode("https://example.com/some/path?x=1")

		if len(code) != 7 {
			t.Fatalf("len(code) = %d, esperado 7; code=%q", len(code), code)
		}
		if !isBase62(code) {
			t.Fatalf("code contiene caracteres fuera de base62: %q", code)
		}
	})

	t.Run("es determinista para la misma URL", func(t *testing.T) {
		u := "https://example.com/abc"
		c1 := s.GenerateShortCode(u)
		c2 := s.GenerateShortCode(u)

		if c1 != c2 {
			t.Fatalf("no determinista: c1=%q c2=%q", c1, c2)
		}
	})

	t.Run("soporta string vacío", func(t *testing.T) {
		code := s.GenerateShortCode("")
		if len(code) != 7 {
			t.Fatalf("len(code) = %d, esperado 7; code=%q", len(code), code)
		}
		if !isBase62(code) {
			t.Fatalf("code contiene caracteres fuera de base62: %q", code)
		}
	})

	t.Run("URLs distintas normalmente generan códigos distintos", func(t *testing.T) {

		c1 := s.GenerateShortCode("https://example.com/a")
		c2 := s.GenerateShortCode("https://example.com/b")

		if c1 == c2 {
			t.Fatalf("colisión inesperada para inputs distintos: c1=%q c2=%q", c1, c2)
		}
	})
}
