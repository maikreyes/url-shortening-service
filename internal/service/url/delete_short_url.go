package url

func (s *Service) DeleteShortUrl(shortCode string) error {

	err := s.Repo.DeleteData(shortCode)

	if err != nil {
		return err
	}

	return nil
}
