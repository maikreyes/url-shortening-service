package ports

import (
	"url-shortening-service/pkg/domain"

	"github.com/google/uuid"
)

type UrlRepository interface {
	FetchData(shortCode string) (*domain.ApiResponse, error)
	PostData(shortUrl domain.ApiResponse) error
	PutData(shortCode string, newUrl string, newShort string) error
	DeleteData(shortCode string) error
	AddCount(data domain.ApiResponse) error
	GetUserUrls(uId uuid.UUID) ([]domain.ApiResponse, error)
}

type UrlService interface {
	GetShortUrl(shortCode string) (*domain.ApiResponse, error)
	CreateShortUrl(url, username string, isWebhook bool, uId uuid.UUID) (string, error)
	UpdateShortUrl(shortCode, username string, newUrl string) (string, error)
	DeleteShortUrl(shortCode string) error
	AddCount(data domain.ApiResponse) error
	GetUserUrls(uId uuid.UUID) ([]domain.ApiResponse, error)
}
