package main

import (
	"strconv"

	"github.com/vedantwankhade/tsrgen/gateway-service/config"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/adapters/clients"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/adapters/server"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/api"
)

func main() {
	confluenceClient := clients.NewConfluenceClient(clients.NewGRPCClient().GetConnection("dns:confluence-service:80"))
	jiraClient := clients.NewJiraClient(clients.NewGRPCClient().GetConnection("dns:jira-service:80"))
	statsClient := clients.NewStatsClient(clients.NewGRPCClient().GetConnection("dns:stats-service:80"))
	dbClient := clients.NewDBClient(clients.NewGRPCClient().GetConnection("dns:db-service:80"))
	app := api.NewApplication(confluenceClient, jiraClient, statsClient, dbClient)
	server := server.NewServer(app)
	port, _ := strconv.Atoi(config.GetServerPort())
	server.Run(port)
}
