package url

import (
	"errors"

	"url-shortening-service/internal/domain"

	"gorm.io/gorm"
)

func (r *Repository) FetchData(shortCode string) (*domain.ApiResponde, error) {

	var url domain.ApiResponde

	err := r.DB.
		Table("urls").
		Where("short_code = ?", shortCode).
		First(&url).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil

}
