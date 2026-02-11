package github

import "url-shortening-service/internal/ports"

type Service struct {
	Repository ports.UrlRepository
}

func NewService(repo ports.UrlRepository) *Service {
	return &Service{
		Repository: repo,
	}
}
