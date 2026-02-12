package url

import (
	"time"
	"url-shortening-service/pkg/domain"
)

func (s *Service) CreateShortUrl(url string) (string, error) {

	normalized := normalizeURL(url)
	shortcode := s.GenerateShortCode(normalized)

	existing, err := s.Repo.FetchData(shortcode)
	if err != nil {
		return "", err
	}

	if existing != nil {
		return existing.ShortCode, nil
	}

	u := domain.ApiResponse{
		Url:       normalized,
		ShortCode: shortcode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.Repo.PostData(u)

	if err != nil {
		return "", err
	}

	return shortcode, nil
}
