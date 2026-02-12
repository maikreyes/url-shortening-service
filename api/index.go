package handler

import (
	"net/http"
	"sync"

	githubHandler "url-shortening-service/cmd/api/Handler/Github"
	urlsHandler "url-shortening-service/cmd/api/Handler/urls"
	"url-shortening-service/cmd/api/router"
	"url-shortening-service/pkg/config"
	repo "url-shortening-service/pkg/repository/url"
	githubService "url-shortening-service/pkg/service/github"
	urlService "url-shortening-service/pkg/service/url"
)

var (
	initOnce sync.Once
	engine   http.Handler
	initErr  error
)

func initServer() {
	ctg := config.LoadConfig()

	db, err := repo.Connection(ctg.DSN, ctg.Driver)
	if err != nil {
		initErr = err
		return
	}

	repository := repo.NewRepository(db, ctg.Table)
	repository.Migrate()

	urlSvc := urlService.NewService(repository)
	urlH := urlsHandler.NewHandler(urlSvc)

	ghSvc := githubService.NewService(repository)
	ghH := githubHandler.NewHandler(ghSvc)

	engine = router.BuildRouter(urlH, ghH)
}

// Handler es el entrypoint que Vercel invoca.
// Con el rewrite configurado como `/api?path=/<ruta_original>`, reconstruimos el path.
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
