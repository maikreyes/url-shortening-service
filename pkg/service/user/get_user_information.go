package user

import "url-shortening-service/pkg/domain"

func (s *Service) GetUserInformation(username string) (*domain.User, error) {

	u, err := s.Repo.GetUserInformation(username)

	if err != nil {
		return nil, err
	}

	return u, nil

}
