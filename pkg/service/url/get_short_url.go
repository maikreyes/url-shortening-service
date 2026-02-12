package url

import (
	"fmt"
	"url-shortening-service/pkg/domain"
)

func (s *Service) GetShortUrl(shortCode string) (*domain.ApiResponse, error) {

	data, err := s.Repo.FetchData(shortCode)

	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, fmt.Errorf("short code %s not found", shortCode)
	}

	return data, nil
}
