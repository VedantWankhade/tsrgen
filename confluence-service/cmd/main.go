package main

import (
	"fmt"
	"log"

	"github.com/vedantwankhade/tsrgen/confluence-service/config"
	core "github.com/vedantwankhade/tsrgen/confluence-service/internal/application/core/api"
)

func main() {
	if config.GetAppRunMode() == "dev" {
		page, err := core.GetPageFromID(6651905, config.GetConfluenceInstance(), config.GetConfluenceUsername(), config.GetConfluenceToken())
		if err != nil {
			log.Fatalf("error getting page: %v", err)
		}
		fmt.Println(page)
	}
}
