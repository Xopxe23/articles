package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/xopxe23/articles/internal/config"
	"github.com/xopxe23/articles/internal/domain"
	"github.com/xopxe23/articles/internal/repository"
	"github.com/xopxe23/articles/internal/service"
	"github.com/xopxe23/articles/internal/transport/rabbitmq"
	"github.com/xopxe23/articles/internal/transport/rest"
	"github.com/xopxe23/articles/pkg/database"
	hasher "github.com/xopxe23/articles/pkg/hash"
)

// @title Articles App API
// @version 1.0
// @description API Server for Articles Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("config initialization error: %s", err)
	}
	db, err := database.NewPostgresConnection(database.PostgresInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("db connection error: %s", err)
	}

	rmqClient := rabbitmq.NewClient()

	hasher := hasher.NewSHA1Hasher("salt")

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo, hasher, []byte("secret"), rmqClient)

	articlesRepo := repository.NewArticlesRepository(db)
	articlesService := service.NewArticlesService(articlesRepo, rmqClient)

	handler := rest.NewHandler(authService, articlesService)

	srv := new(domain.Server)
	go func() {
		if err := srv.Run("8000", handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("SERVER STARTED")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Print("SERVER SHUNDOWN")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := rmqClient.CloseConnection(); err != nil {
		logrus.Errorf("error occured on rabbitMQ connection close: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
