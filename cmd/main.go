package main

import (
	"log"

	servers "github.com/BnnryMndy/GolangXMLParse"
	"github.com/BnnryMndy/GolangXMLParse/internal/handler"
	"github.com/BnnryMndy/GolangXMLParse/internal/repository"
	"github.com/BnnryMndy/GolangXMLParse/internal/service"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewServices(repos)
	handlers := handler.NewHandlers(serv)

	srv := new(servers.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting server: %s", err.Error())
	}
}
