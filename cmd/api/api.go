package api

import (
	github "url-shortening-service/cmd/api/Handler/Github"
	handler "url-shortening-service/cmd/api/Handler/urls"
	"url-shortening-service/cmd/api/router"
	"url-shortening-service/internal/config"
	repo "url-shortening-service/internal/repository/url"
	githubService "url-shortening-service/internal/service/github"
	service "url-shortening-service/internal/service/url"
)

func Run() {

	ctg := config.LoadConfig()

	db, err := repo.Connection(ctg.DSN, ctg.Driver)

	if err != nil {
		panic(err)
	}

	repository := repo.NewRepository(db, ctg.Table)
	repository.Migrate()

	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	githubService := githubService.NewService(repository)
	GithubHandler := github.NewHandler(githubService)

	var addr string

	if ctg.Environment == "production" {
		// En producción normalmente el provider expone solo PORT.
		// Si HOST viene vacío, escuchamos en todas las interfaces.
		if ctg.Port != "" {
			addr = ":" + ctg.Port
		} else {
			addr = ctg.Host
		}
	} else {
		if ctg.Host == "" {
			addr = ":" + ctg.Port
		} else {
			addr = ctg.Host + ":" + ctg.Port
		}
	}

	router.NewRouter(addr, handler, GithubHandler)

}
