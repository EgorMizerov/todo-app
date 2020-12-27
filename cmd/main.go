package main

import (
	"context"
	"github.com/EgorMizerov/todo-app"
	"github.com/EgorMizerov/todo-app/pkg/handler"
	"github.com/EgorMizerov/todo-app/pkg/repository"
	"github.com/EgorMizerov/todo-app/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация логгера
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализация конфига
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Инициализация базы данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Создание зависимостей
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	// Создание и запуск сервера
	server := new(todo.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("Start listening server...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Server shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on server connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs") // Директория конфига
	viper.SetConfigName("config")  // Название файла
	viper.SetConfigType("yaml")    // Рассширение файла
	return viper.ReadInConfig()
}
