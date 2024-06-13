package main

import (
	"github.com/vedantwankhade/tsrgen/stats-service/internal/adapters/grpc"
	"github.com/vedantwankhade/tsrgen/stats-service/internal/application/core/api"
)

func main() {
	app := api.NewApplication()
	server := grpc.NewServer(app, 80)
	server.Run()
}
