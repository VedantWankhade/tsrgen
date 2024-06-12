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

func GetJiraInstance() string {
	return getEnv("JIRA_INSTANCE")
}

func GetJiraUsername() string {
	return getEnv("JIRA_USERNAME")
}

func GetJiraToken() string {
	return getEnv("JIRA_TOKEN")
}

func GetAppRunMode() string {
	return getEnv("MODE")
}

func init() {
	godotenv.Load()
}
