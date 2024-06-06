package ports

import "github.com/vedantwankhade/tsrgen/confluence-service/internal/application/core/domain"

type APIPort interface {
	GetPageFromID(int, string, string, string) (*domain.Page, error)
}
