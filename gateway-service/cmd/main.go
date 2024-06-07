package main

import (
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/adapters/clients"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/adapters/server"
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/api"
)

func main() {
	conn := clients.NewGRPCClient().GetConnection(1234)
	confluenceClient := clients.NewConfluenceClient(conn)
	app := api.NewApplication(confluenceClient)
	server := server.NewServer(app)
	server.Run(8080)
}
