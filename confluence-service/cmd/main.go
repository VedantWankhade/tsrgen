package main

import (
	"github.com/vedantwankhade/tsrgen/confluence-service/config"
	core "github.com/vedantwankhade/tsrgen/confluence-service/internal/application/core/api"
)

func main() {
	if config.GetAppRunMode() == "dev" {
		core.GetPageFromID(6651905, config.GetConfluenceInstance(), config.GetConfluenceUsername(), config.GetConfluenceToken())
	}
}
