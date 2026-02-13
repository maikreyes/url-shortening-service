package user

import "url-shortening-service/pkg/domain"

func (r *Repository) DeleteUser(username string) error {

	err := r.DB.Table(r.TableName).
		Where("username = ?", username).
		Delete(&domain.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
