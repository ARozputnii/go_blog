package config

import "github.com/joho/godotenv"

// LoadEnv загружає змінні оточення з файлу .env.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load the .env file")
	}
}
