package user

import "url-shortening-service/pkg/domain"

func (s *Service) PutUser(username string, user domain.RegisterInput) error {

	err := s.Repo.PutUser(username, user)

	if err != nil {
		return err
	}

	return nil

}
