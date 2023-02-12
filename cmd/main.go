package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/POMBNK/restAPI"
	"github.com/POMBNK/restAPI/config"
	"github.com/POMBNK/restAPI/pkg/handler"
	"github.com/POMBNK/restAPI/pkg/repository"
	"github.com/POMBNK/restAPI/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg := config.GetCfg()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed to load env vars %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.Db.Host,
		Port:     cfg.Db.Port,
		Username: cfg.Db.Username,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   cfg.Db.DbName,
		SSLMode:  cfg.Db.SSLMode,
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while running server %s", err)
		}
	}()

	logrus.Println("Service started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Service shutting sown")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error ocured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error ocured on db connection closed: %s", err.Error())
	}
}
