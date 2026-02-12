package url

import "url-shortening-service/pkg/domain"

func (s *Service) AddCount(data domain.ApiResponse) error {

	err := s.Repo.AddCount(data)

	if err != nil {
		return err
	}

	return nil

}
