package config

import (
	"errors"
	"os"
	"time"
)

type FiberConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewFiberConfig() (*FiberConfig, error) {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		return nil, errors.New("environment variable HTTP_PORT is not set")
	}

	return &FiberConfig{
		Port:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}, nil
}
