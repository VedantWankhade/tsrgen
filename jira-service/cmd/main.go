package main

import (
	"github.com/vedantwankhade/tsrgen/jira-service/config"
	"github.com/vedantwankhade/tsrgen/jira-service/internal/application/core/api"
)

func main() {
	if config.GetAppRunMode() == "dev" {
		app := api.NewApplication()
		app.GetIssuesWithJQL("created >= -30d order by created DESC", config.GetJiraInstance(), config.GetJiraUsername(), config.GetJiraToken())
	}
}
