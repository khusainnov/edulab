package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/khusainnov/edulab"
	"github.com/khusainnov/edulab/pkg/handler"
	"github.com/khusainnov/edulab/pkg/repository"
	"github.com/khusainnov/edulab/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Infoln("Loading config from .env")
	if err := godotenv.Load("./config/.env"); err != nil {
		logrus.Errorf("Cannot load .env config, due to error: %v", err)
	}

	logrus.Infoln("Loading config from .yml")
	if err := initConfig(); err != nil {
		logrus.Errorf("Cannot load .yml config, due to error: %v", err)
	}

	logrus.Infoln("Initializing DB")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Errorf("Cannot initialize db, due to error: %s", err.Error())
	}

	logrus.Infoln("Initializing repository")
	repos := repository.NewRepository(db)

	logrus.Infoln("Initializing services")
	services := service.NewService(repos)

	logrus.Infoln("Getting handlers")
	handlers := handler.NewHandler(services)

	logrus.Infoln("Initializing server")
	server := new(edulab.Server)

	logrus.Infof("Starting server on port: %s", os.Getenv("PORT"))
	if err := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Errorf("Cannot run server, due to error: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("configs")

	return viper.ReadInConfig()
}
