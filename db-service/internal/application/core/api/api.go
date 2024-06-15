package api

import (
	"fmt"

	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/db-service/internal/ports"
)

type application struct {
	db ports.DBPort
}

func (a *application) Save(page domain.Page) error {
	fmt.Println(page)
	return a.db.Save(page)
}

func NewApplication(db ports.DBPort) *application {
	return &application{db: db}
}
