package api

import (
	github "url-shortening-service/cmd/api/Handler/Github"
	handler "url-shortening-service/cmd/api/Handler/urls"
	user "url-shortening-service/cmd/api/Handler/user"
	"url-shortening-service/cmd/api/router"
	"url-shortening-service/pkg/config"
	repo "url-shortening-service/pkg/repository/url"
	userepo "url-shortening-service/pkg/repository/user"
	githubService "url-shortening-service/pkg/service/github"
	service "url-shortening-service/pkg/service/url"
	userService "url-shortening-service/pkg/service/user"
)

func Run() {

	ctg := config.LoadConfig()

	db, err := repo.Connection(ctg.DSN, ctg.Driver)

	if err != nil {
		panic(err)
	}

	repository := repo.NewRepository(db, ctg.UrlTable)
	repository.Migrate()

	userRepo := userepo.NewRepository(db, ctg.UserTable)
	userRepo.Migrate()

	githubService := githubService.NewService(repository)
	GithubHandler := github.NewHandler(githubService)

	userService := userService.NewService(userRepo)
	userhandler := user.NewHandler(userService, ctg.Host)

	service := service.NewService(repository)
	handler := handler.NewHandler(service, userService, ctg.Host)

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

	router.NewRouter(addr, handler, GithubHandler, userhandler)

}
