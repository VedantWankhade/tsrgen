package main

import (
	"fmt"
	"log"

	"github.com/vedantwankhade/tsrgen/db-service/internal/adapters/db"
	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/api"
	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"
)

func main() {
	db, err := db.NewDBAdapter("postgres://postgres:password@localhost:5432/pages")
	if err != nil {
		log.Fatal(err)
	}
	app := api.NewApplication(db)
	err = app.Save(domain.Page{
		Id:    "TEST",
		Title: "TESTBOR",
		Link:  "https://googlv.eomd",
	})
	fmt.Println(err)
}
