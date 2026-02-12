package url

import (
	"time"
	"url-shortening-service/internal/domain"
)

func (r *Repository) AddCount(data domain.ApiResponse) error {

	err := r.DB.Table(r.TableName).
		Where("short_code = ?", data.ShortCode).
		Updates(map[string]interface{}{
			"access_count": data.AccessCount + 1,
			"updated_at":   time.Now(),
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}
