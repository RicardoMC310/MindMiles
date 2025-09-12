package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	if os.Getenv("RENDER") == "" { // só carrega .env localmente
		_ = godotenv.Load(".env")
	}

	return os.Getenv(key)
}
