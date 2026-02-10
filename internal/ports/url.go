package ports

import "url-shortening-service/internal/domain"

type UrlRepository interface {
	FetchData(shortCode string) (*domain.ApiResponde, error)
	PostData(shortUrl domain.ApiResponde) error
	PutData(shortCode string, newUrl string, newShort string) error
	DeleteData(shortCode string) error
}

type UrlService interface {
	GetShortUrl(shortCode string) (*domain.ApiResponde, error)
	CreateShortUrl(url string) (string, error)
	UpdateShortUrl(shortCode string, newUrl string) (string, error)
	DeleteShortUrl(shortCode string) error
}
