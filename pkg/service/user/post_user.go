package user

import "url-shortening-service/pkg/domain"

func (s *Service) PostUser(user domain.RegisterInput) error {

	err := s.Repo.PostUser(user)

	if err != nil {
		return err
	}

	return nil

}
