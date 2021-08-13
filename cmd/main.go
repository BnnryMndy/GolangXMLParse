package main

import (
	"log"

	servers "github.com/BnnryMndy/GolangXMLParse"
	handlers "github.com/BnnryMndy/GolangXMLParse/pkg/handler"
)

func main() {
	handlers := new(handlers.Handler)

	srv := new(servers.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting server: %s", err.Error())
	}
}
