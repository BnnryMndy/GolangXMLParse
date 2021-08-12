package main

import (
	"log"

	xmlparse "github.com/BnnryMndy/GolangXMLParse"
	"github.com/BnnryMndy/GolangXMLParse/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(xmlparse.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting server: %s", err.Error())
	}
}
