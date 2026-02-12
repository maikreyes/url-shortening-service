package cli

import (
	"flag"
	handler "url-shortening-service/cmd/cli/Handler"
	"url-shortening-service/internal/config"
	repo "url-shortening-service/internal/repository/url"
	service "url-shortening-service/internal/service/url"
)

func Run(args []string) {
	if len(args) < 1 {
		panic("command is required (fetch)")
	}

	action := args[0]

	fs := flag.NewFlagSet("cli", flag.ExitOnError)
	code := fs.String("code", "", "The shortcode for search")
	url := fs.String("url", "", "Url to convert")

	fs.Parse(args[1:])

	if action == "fetch" && *code == "" {
		panic("code flag is required")
	}

	ctg := config.LoadConfig()

	db, err := repo.Connection(ctg.DSN, ctg.Driver)

	if err != nil {
		panic(err)
	}

	repository := repo.NewRepository(db, ctg.Table)
	service := service.NewService(repository)
	handler := handler.NewHandler(service, ctg.Host, ctg.Port)

	repository.Migrate()

	switch action {
	case "fetch":
		handler.FetchData(*code)
	case "post":
		handler.PostData(*url)
	case "put":
		handler.PutData(*code, *url)
	case "delete":
		handler.DeleteData(*code)
	default:
		panic("unknown command: " + action)
	}
}
