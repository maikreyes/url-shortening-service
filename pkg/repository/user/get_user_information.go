package user

import "url-shortening-service/pkg/domain"

func (r *Repository) GetUserInformation(username string) (*domain.User, error) {

	var u domain.User

	if err := r.DB.Table(r.TableName).Where("email = ?", username).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}
