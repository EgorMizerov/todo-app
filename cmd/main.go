package main

import (
	"github.com/EgorMizerov/todo-app"
	"github.com/EgorMizerov/todo-app/pkg/handler"
	"github.com/EgorMizerov/todo-app/pkg/repository"
	"github.com/EgorMizerov/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
