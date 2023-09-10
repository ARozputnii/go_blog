package config

import (
	"github.com/joho/godotenv"
	"os"
)

// EnvConfig represents the configuration structure for environment variables.
type EnvConfig struct {
	Port string
	// other environment variables
}

// LoadEnv loads environment variables into the EnvConfig struct from a .env file.
func LoadEnv() (*EnvConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	envConfig := &EnvConfig{
		Port: os.Getenv("PORT"),
		// Initialize other environment variables
	}

	return envConfig, nil
}
