package config

import (
	"errors"
	"fmt"
	"os"
)

func GetPostgresConnectionString(host, port, user, password, name string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func NewDBConfig() (*DBConfig, error) {
	cfg := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if cfg.Host == "" {
		return nil, errors.New("DB_HOST is required")
	}
	if cfg.Port == "" {
		return nil, errors.New("DB_PORT is required")
	}
	if cfg.Name == "" {
		return nil, errors.New("DB_NAME is required")
	}
	if cfg.User == "" {
		return nil, errors.New("DB_USER is required")
	}

	return &cfg, nil
}
