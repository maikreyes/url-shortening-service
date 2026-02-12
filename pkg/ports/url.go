package ports

import "url-shortening-service/pkg/domain"

type UrlRepository interface {
	FetchData(shortCode string) (*domain.ApiResponse, error)
	PostData(shortUrl domain.ApiResponse) error
	PutData(shortCode string, newUrl string, newShort string) error
	DeleteData(shortCode string) error
	AddCount(data domain.ApiResponse) error
}

type UrlService interface {
	GetShortUrl(shortCode string) (*domain.ApiResponse, error)
	CreateShortUrl(url string) (string, error)
	UpdateShortUrl(shortCode string, newUrl string) (string, error)
	DeleteShortUrl(shortCode string) error
	AddCount(data domain.ApiResponse) error
}
