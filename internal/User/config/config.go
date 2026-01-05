package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	if err := godotenv.Load("app.env"); err != nil {
		return fmt.Errorf("Error loading .env file: %w", err)
	}
	return nil
}
