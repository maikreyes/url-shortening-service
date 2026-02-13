package user

import "url-shortening-service/pkg/domain"

func (s *Service) GetUser(user domain.LoginInput) (*domain.User, error) {

	u, err := s.Repo.GetUser(user)

	if err != nil {
		return nil, err
	}

	return u, nil

}
