package url

import (
	"url-shortening-service/pkg/domain"

	"github.com/google/uuid"
)

func (s *Service) GetUserUrls(uid uuid.UUID) ([]domain.ApiResponse, error) {

	urls, err := s.Repo.GetUserUrls(uid)

	if err != nil {
		return nil, err
	}

	return urls, nil
}
