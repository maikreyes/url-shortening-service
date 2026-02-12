package url

import "gorm.io/gorm"

type Repository struct {
	DB        *gorm.DB
	TableName string
}

func NewRepository(db *gorm.DB, table string) *Repository {

	return &Repository{
		DB:        db,
		TableName: table,
	}

}
