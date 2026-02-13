package user

import (
	"log"
	"strings"
	"url-shortening-service/pkg/domain"
)

func (r *Repository) Migrate() {
	table := strings.TrimSpace(r.TableName)
	if table == "" {
		panic("TABLE_NAME is required")
	}

	log.Println("[USER_REPOSITORY] Migrating user table...")

	if r.DB.Migrator().HasTable(table) {
		return
	}

	if err := r.DB.Table(table).AutoMigrate(&domain.User{}); err != nil {
		panic(err)
	}

}
