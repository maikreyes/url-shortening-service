package user

import (
	"errors"
	"url-shortening-service/pkg/domain"

	"gorm.io/gorm"
)

func (r *Repository) GetUser(user domain.LoginInput) (*domain.User, error) {

	var u domain.User

	err := r.DB.Table(r.TableName).
		Where("email = ?", user.Email).
		First(&u).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil

}
