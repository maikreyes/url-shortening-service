package ports

import "url-shortening-service/pkg/domain"

type UserRepository interface {
	GetUser(user domain.LoginInput) (*domain.User, error)
	PostUser(user domain.RegisterInput) error
	PutUser(username string, user domain.RegisterInput) error
	DeleteUser(username string) error
}

type UserService interface {
	EncryptPassword(password string) (string, error)
	GetUser(user domain.LoginInput) (*domain.User, error)
	PostUser(user domain.RegisterInput) error
	PutUser(username string, user domain.RegisterInput) error
	DeleteUser(usernname string) error
}
