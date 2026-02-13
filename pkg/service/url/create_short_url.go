package url

import (
	"fmt"
	"time"
	"url-shortening-service/pkg/domain"

	"github.com/google/uuid"
)

func (s *Service) CreateShortUrl(url, username string, isWebhook bool, uId uuid.UUID) (string, error) {

	normalized := normalizeURL(url)

	const maxAttempts = 6
	for attempt := 0; attempt < maxAttempts; attempt++ {
		userSalt := username
		if attempt > 0 {
			userSalt = fmt.Sprintf("%s#%d", username, attempt)
		}

		shortcode := s.GenerateShortCode(normalized, userSalt)
		u := domain.ApiResponse{
			Url:       normalized,
			UserID:    uId,
			ShortCode: shortcode,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsWebhook: isWebhook,
		}

		err := s.Repo.PostData(u)
		if err == nil {
			return shortcode, nil
		}

		if isDuplicateShortCodeError(err) {
			existing, fetchErr := s.Repo.FetchData(shortcode)
			if fetchErr != nil {
				return "", fetchErr
			}
			// Si ya existe para el mismo usuario+url, lo tratamos como idempotente.
			if existing != nil && existing.UserID == uId && existing.Url == normalized && existing.IsWebhook == isWebhook {
				return shortcode, nil
			}
			continue
		}

		return "", err
	}

	return "", fmt.Errorf("could not generate unique short code")
}
