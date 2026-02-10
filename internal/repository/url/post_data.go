package url

import "url-shortening-service/internal/domain"

func (r *Repository) PostData(ShortUrl domain.ApiResponde) error {

	err := r.DB.Table("urls").Create(&ShortUrl).Error

	if err != nil {
		return err
	}

	return nil
}
