package url

import (
	"errors"

	"url-shortening-service/pkg/domain"

	"gorm.io/gorm"
)

func (r *Repository) FetchData(shortCode string) (*domain.ApiResponse, error) {

	var url domain.ApiResponse

	err := r.DB.
		Table(r.TableName).
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
