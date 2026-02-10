package main

import (
	"fmt"
	"os"
	"url-shortening-service/cmd/api"
	"url-shortening-service/cmd/cli"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: url-shortening-service <serve function> <command> [flags]")
	}

	run := os.Args[1]

	switch run {
	case "cli":
		cli.Run(os.Args[2:])
	case "api":
		api.Run()
	default:
		break
	}

}
