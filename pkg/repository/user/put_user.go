package user

import (
	"time"
	"url-shortening-service/pkg/domain"
)

func (r *Repository) PutUser(username string, user domain.RegisterInput) error {

	err := r.DB.Table(r.TableName).
		Where("username = ?", username).
		Updates(map[string]interface{}{
			"username":   user.Username,
			"email":      user.Email,
			"password":   user.Password,
			"updated_at": time.Now(),
		}).Error

	if err != nil {
		return err
	}

	return nil

}
