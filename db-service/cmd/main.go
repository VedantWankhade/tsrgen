package main

import (
	"log"

	"github.com/vedantwankhade/tsrgen/db-service/config"
	"github.com/vedantwankhade/tsrgen/db-service/internal/adapters/db"
	"github.com/vedantwankhade/tsrgen/db-service/internal/adapters/grpc"
	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/api"
)

func main() {
	db, err := db.NewDBAdapter(config.GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	app := api.NewApplication(db)
	server := grpc.NewServer(app, 80)
	server.Run()
}
