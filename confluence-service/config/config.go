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

func GetConfluenceInstance() string {
	return getEnv("CONFLUENCE_INSTANCE")
}

func GetConfluenceUsername() string {
	return getEnv("CONFLUENCE_USERNAME")
}

func GetConfluenceToken() string {
	return getEnv("CONFLUENCE_TOKEN")
}

func GetAppRunMode() string {
	return getEnv("MODE")
}

func init() {
	godotenv.Load()
}
