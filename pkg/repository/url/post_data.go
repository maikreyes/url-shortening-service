package url

import (
	"url-shortening-service/pkg/domain"

	"github.com/google/uuid"
)

func (r *Repository) PostData(ShortUrl domain.ApiResponse) error {

	if ShortUrl.ID == uuid.Nil {
		ShortUrl.ID = uuid.New()
	}

	err := r.DB.Table(r.TableName).Create(&ShortUrl).Error

	if err != nil {
		return err
	}

	return nil
}
