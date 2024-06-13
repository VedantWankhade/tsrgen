package main

import (
	"github.com/vedantwankhade/tsrgen/jira-service/config"
	"github.com/vedantwankhade/tsrgen/jira-service/internal/adapters/grpc"
	"github.com/vedantwankhade/tsrgen/jira-service/internal/application/core/api"
)

func main() {
	if config.GetAppRunMode() == "dev" {
		app := api.NewApplication()
		server := grpc.NewServer(app, 1235)
		server.Run()
	}
}
