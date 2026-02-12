package connection

import (
	"url-shortening-service/internal/domain"
)

type Service struct {
	Connection domain.Connection
}

func NewService(connection domain.Connection) *Service {
	return &Service{
		Connection: connection,
	}
}
