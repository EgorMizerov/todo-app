package main

import (
	"github.com/EgorMizerov/todo-app"
	"github.com/EgorMizerov/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
