package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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
	defer func() {
		if r := recover(); r != nil {
			initErr = fmt.Errorf("init panic: %v", r)
			log.Printf("[vercel] initServer panic: %v", r)
		}
	}()

	ctg := config.LoadConfig()

	if strings.TrimSpace(ctg.Driver) == "" {
		initErr = fmt.Errorf("missing env var DB_DRIVER")
		log.Printf("[vercel] initServer error: %v", initErr)
		return
	}
	if strings.TrimSpace(ctg.DSN) == "" {
		initErr = fmt.Errorf("missing env var CONNECTION_STRING")
		log.Printf("[vercel] initServer error: %v", initErr)
		return
	}
	if strings.TrimSpace(ctg.Table) == "" {
		initErr = fmt.Errorf("missing env var TABLE_NAME")
		log.Printf("[vercel] initServer error: %v", initErr)
		return
	}

	db, err := repo.Connection(ctg.DSN, ctg.Driver)
	if err != nil {
		initErr = err
		log.Printf("[vercel] initServer DB connection error: %v", err)
		return
	}

	repository := repo.NewRepository(db, ctg.Table)
	repository.Migrate()

	urlSvc := urlService.NewService(repository)
	urlH := urlsHandler.NewHandler(urlSvc, ctg.Host)

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
