package main

import (
	"github.com/EgorMizerov/todo-app"
	"github.com/EgorMizerov/todo-app/pkg/handler"
	"github.com/EgorMizerov/todo-app/pkg/repository"
	"github.com/EgorMizerov/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
