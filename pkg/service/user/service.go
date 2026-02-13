package user

import (
	"url-shortening-service/pkg/ports"
)

type Service struct {
	Repo ports.UserRepository
}

func NewService(repo ports.UserRepository) *Service {
	return &Service{
		Repo: repo,
	}
}
