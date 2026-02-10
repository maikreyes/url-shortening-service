package url

func (r *Repository) DeleteData(ShortCode string) error {

	err := r.DB.Table("urls").Where("short_code = ?", ShortCode).Delete(&ShortCode).Error

	if err != nil {
		return err
	}

	return nil
}
