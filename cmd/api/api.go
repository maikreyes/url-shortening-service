package api

import (
	github "url-shortening-service/cmd/api/Handler/Github"
	handler "url-shortening-service/cmd/api/Handler/urls"
	"url-shortening-service/cmd/api/router"
	"url-shortening-service/pkg/config"
	repo "url-shortening-service/pkg/repository/url"
	githubService "url-shortening-service/pkg/service/github"
	service "url-shortening-service/pkg/service/url"
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
	handler := handler.NewHandler(service, ctg.Host)

	githubService := githubService.NewService(repository)
	GithubHandler := github.NewHandler(githubService)

	var addr string

	if ctg.Environment == "production" {

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
