package url

import "url-shortening-service/pkg/ports"

type Service struct {
	Repo ports.UrlRepository
}

func NewService(repo ports.UrlRepository) *Service {
	return &Service{
		Repo: repo,
	}
}
