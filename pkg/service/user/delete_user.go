package user

func (s *Service) DeleteUser(username string) error {

	err := s.Repo.DeleteUser(username)

	if err != nil {
		return err
	}

	return nil

}
