package url

import (
	"url-shortening-service/pkg/domain"

	"github.com/google/uuid"
)

func (r *Repository) GetUserUrls(uid uuid.UUID) ([]domain.ApiResponse, error) {

	var urls []domain.ApiResponse
	err := r.DB.Table(r.TableName).Where("user_id = ?", uid).Find(&urls).Error

	if err != nil {
		return nil, err
	}

	return urls, nil
}
