package user

import (
	"time"
	"url-shortening-service/pkg/domain"

	"github.com/google/uuid"
)

func (r *Repository) PostUser(user domain.RegisterInput) error {

	u := domain.User{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      "user",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	err := r.DB.Table(r.TableName).Create(&u).Error

	if err != nil {
		return err
	}

	return nil

}
