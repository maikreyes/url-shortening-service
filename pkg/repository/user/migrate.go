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

	if r.DB == nil {
		panic("DB connection is nil")
	}

	log.Println("[USER_REPOSITORY] Migrating user table...")

	if r.DB.Migrator().HasTable(&domain.User{}) {
		return
	}

	if err := r.DB.AutoMigrate(&domain.User{}); err != nil {
		panic(err)
	}

}
