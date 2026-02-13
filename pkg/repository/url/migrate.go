package url

import (
	"strings"
	"url-shortening-service/pkg/domain"
)

func (r *Repository) Migrate() {
	table := strings.TrimSpace(r.TableName)
	if table == "" {
		panic("TABLE_NAME is required")
	}

	if r.DB == nil {
		panic("DB connection is nil")
	}

	if r.DB.Migrator().HasTable(&domain.ApiResponse{}) {
		return
	}

	if err := r.DB.AutoMigrate(&domain.ApiResponse{}); err != nil {
		panic(err)
	}

}
