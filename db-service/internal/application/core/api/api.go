package api

import (
	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/db-service/internal/ports"
)

type application struct {
	db ports.DBPort
}

func (a *application) Save(page domain.Page) (int, error) {
	return a.db.Save(page)
}

func (a *application) GetAll() ([]*domain.Entry, error) {
	return a.db.GetAll()
}

func NewApplication(db ports.DBPort) *application {
	return &application{db: db}
}
