package handler

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	githubHandler "url-shortening-service/cmd/api/Handler/Github"
	urlsHandler "url-shortening-service/cmd/api/Handler/urls"
	userHandler "url-shortening-service/cmd/api/Handler/user"
	"url-shortening-service/cmd/api/router"
	"url-shortening-service/pkg/config"
	repo "url-shortening-service/pkg/repository/url"
	userepo "url-shortening-service/pkg/repository/user"
	githubService "url-shortening-service/pkg/service/github"
	urlService "url-shortening-service/pkg/service/url"
	userService "url-shortening-service/pkg/service/user"
)

var (
	initOnce sync.Once
	engine   http.Handler
	initErr  error
)

func initServer() {
	defer func() {
		if r := recover(); r != nil {
			initErr = fmt.Errorf("init panic: %v", r)
			log.Printf("[vercel] initServer panic: %v", r)
		}
	}()

	ctg := config.LoadConfig()

	db, err := repo.Connection(ctg.DSN, ctg.Driver)
	if err != nil {
		initErr = err
		log.Printf("[vercel] initServer DB connection error: %v", err)
		return
	}

	userRepository := userepo.NewRepository(db, ctg.UserTable)
	userRepository.Migrate()

	repository := repo.NewRepository(db, ctg.UrlTable)
	repository.Migrate()

	ghSvc := githubService.NewService(repository)
	ghH := githubHandler.NewHandler(ghSvc)

	usrSvc := userService.NewService(userRepository)
	usrH := userHandler.NewHandler(usrSvc, ctg.Host)

	urlSvc := urlService.NewService(repository)
	urlH := urlsHandler.NewHandler(urlSvc, usrSvc, ctg.Host)

	engine = router.BuildRouter(urlH, ghH, usrH)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	initOnce.Do(initServer)

	if initErr != nil {
		http.Error(w, initErr.Error(), http.StatusInternalServerError)
		return
	}

	originalPath := r.URL.Query().Get("path")
	if originalPath != "" {
		cloned := r.Clone(r.Context())
		cloned.URL.Path = originalPath
		cloned.URL.RawPath = originalPath

		q := cloned.URL.Query()
		q.Del("path")
		cloned.URL.RawQuery = q.Encode()

		engine.ServeHTTP(w, cloned)
		return
	}

	engine.ServeHTTP(w, r)
}
