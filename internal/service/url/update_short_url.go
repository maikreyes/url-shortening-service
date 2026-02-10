package url

func (s *Service) UpdateShortUrl(shortCode string, newUrl string) (string, error) {
	normalized := normalizeURL(newUrl)
	newShort := s.generateShortCode(normalized)

	err := s.Repo.PutData(shortCode, normalized, newShort)

	if err != nil {
		return "", err
	}

	return newShort, nil
}
