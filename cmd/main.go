package main

import (
	"os"

	todo "github.com/POMBNK/restAPI"
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
	if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running server %s", err)
	}
}
