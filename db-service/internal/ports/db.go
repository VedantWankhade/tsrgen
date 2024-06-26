package ports

import "github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"

type DBPort interface {
	Save(page domain.Page) (int, error)
	GetAll() ([]*domain.Entry, error)
}
