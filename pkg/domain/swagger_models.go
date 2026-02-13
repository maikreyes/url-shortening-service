package domain

// Modelos auxiliares para documentación Swagger (swaggo).
// No cambian el comportamiento del runtime.

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Token TokenResponse `json:"token"`
}

type ShortenResponse struct {
	Url string `json:"url"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
