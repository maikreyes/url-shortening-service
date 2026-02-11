package api

import (
	handler "url-shortening-service/cmd/api/Handler"
	"url-shortening-service/cmd/api/router"
	"url-shortening-service/internal/config"
	repo "url-shortening-service/internal/repository/url"
	service "url-shortening-service/internal/service/url"
)

func Run() {

	ctg := config.LoadConfig()

	db, err := repo.Connection(ctg.DSN)

	if err != nil {
		panic(err)
	}

	repository := repo.NewRepository(db, ctg.Table)
	repository.Migrate()

	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	addr := ctg.Host + ":" + ctg.Port

	router.NewRouter(addr, handler)

}
