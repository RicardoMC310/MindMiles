package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Erro ao ler .env: %v", err)
	}

	return os.Getenv(key)
}