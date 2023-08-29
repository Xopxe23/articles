package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DB Postgres
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Name     string
	SSLMode  string
	Password string
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	err := envconfig.Process("db", &cfg.DB)
	return cfg, err
}
