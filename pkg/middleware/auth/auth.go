package auth

import "os"

type JWTConfig struct {
	Secret []byte
}

func LoadJWTCtg() *JWTConfig {
	return &JWTConfig{
		Secret: []byte(os.Getenv("JWT_SECRET")),
	}
}
