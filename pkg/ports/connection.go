package ports

import "gorm.io/gorm"

type ConnectionService interface {
	NewConnection(driver string, dsn string) (*gorm.DB, error)
}
