package url

import (
	"url-shortening-service/internal/domain"
	service "url-shortening-service/internal/service/connection"

	"gorm.io/gorm"
)

func Connection(DSN string, Driver string) (*gorm.DB, error) {

	connection := domain.Connection{
		Driver: Driver,
		DSN:    DSN,
	}

	connectionService := service.NewService(connection)

	db, err := connectionService.NewConnection(connection.Driver, connection.DSN)

	if err != nil {
		return nil, err
	}

	return db, nil
}
