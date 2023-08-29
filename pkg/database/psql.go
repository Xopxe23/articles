package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresInfo struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

func NewPostgresConnection(psql PostgresInfo) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		psql.Host, psql.Port, psql.Username, psql.DBName, psql.Password, psql.SSLMode),
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
