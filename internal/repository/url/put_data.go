package url

import (
	"time"
	"url-shortening-service/internal/domain"
)

func (r *Repository) PutData(shortCode string, newUrl string, newShort string) error {

	err := r.DB.Table("urls").
		Model(domain.ApiResponde{}).
		Where("short_code = ?", shortCode).
		Updates(map[string]interface{}{
			"url":        newUrl,
			"short_code": newShort,
			"updated_at": time.Now(),
		}).Error

	if err != nil {
		return err
	}

	return nil
}
