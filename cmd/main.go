package main

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"

	"github.com/xopxe23/articles/internal/config"
	"github.com/xopxe23/articles/pkg/database"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config initialization error: %s", err)
	}
	fmt.Println(cfg)
	db, err := database.NewPostgresConnection(database.PostgresInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Fatalf("db connection error: %s", err)
	}
	fmt.Printf("db connected: %T\n", db)
}
