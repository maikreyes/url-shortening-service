package url

import "url-shortening-service/pkg/domain"

func (r *Repository) PostData(ShortUrl domain.ApiResponse) error {

	err := r.DB.Table(r.TableName).Create(&ShortUrl).Error

	if err != nil {
		return err
	}

	return nil
}
