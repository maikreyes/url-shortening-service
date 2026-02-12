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

	if r.DB.Migrator().HasTable(table) {
		return
	}

	if err := r.DB.Table(table).AutoMigrate(&domain.ApiResponse{}); err != nil {
		panic(err)
	}
}
