package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(envVar string) string {
	v := os.Getenv(envVar)
	if len(v) == 0 {
		log.Fatalf("Environment variable %s not set", envVar)
	}
	return v
}

func GetDBURL() string {
	return getEnv("DB_URL")
}

func GetAppRunMode() string {
	return getEnv("MODE")
}

func init() {
	godotenv.Load()
}
