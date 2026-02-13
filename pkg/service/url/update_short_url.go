package url

import "fmt"

func (s *Service) UpdateShortUrl(shortCode, username string, newUrl string) (string, error) {
	normalized := normalizeURL(newUrl)

	const maxAttempts = 6
	for attempt := 0; attempt < maxAttempts; attempt++ {
		userSalt := username
		if attempt > 0 {
			userSalt = fmt.Sprintf("%s#%d", username, attempt)
		}
		newShort := s.GenerateShortCode(normalized, userSalt)

		err := s.Repo.PutData(shortCode, normalized, newShort)
		if err == nil {
			return newShort, nil
		}

		if isDuplicateShortCodeError(err) {
			continue
		}

		return "", err
	}

	return "", fmt.Errorf("could not generate unique short code")
}
