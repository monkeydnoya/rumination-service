package config

import (
	"os"

	"github.com/joho/godotenv"
)

var setup bool = false

func Config(key string) string {
	if !setup {
		godotenv.Load(".env") // nolint - ignore if file not exists
		setup = true
	}
	return os.Getenv(key)
}
