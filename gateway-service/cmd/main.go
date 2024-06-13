package main

import (
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/adapters/clients"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/adapters/server"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/api"
)

func main() {
	confluenceClient := clients.NewConfluenceClient(clients.NewGRPCClient().GetConnection(1234))
	jiraClient := clients.NewJiraClient(clients.NewGRPCClient().GetConnection(1235))
	app := api.NewApplication(confluenceClient, jiraClient)
	server := server.NewServer(app)
	server.Run(8080)
}
