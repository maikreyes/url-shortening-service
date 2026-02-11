package url

import "url-shortening-service/internal/domain"

func (r *Repository) Migrate() {

	migrator := r.DB.Migrator()
	if migrator.HasTable(&domain.ApiResponse{}) {
		return
	}

	if err := migrator.CreateTable(&domain.ApiResponse{}); err != nil {
		panic(err)
	}
}
