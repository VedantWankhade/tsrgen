package ports

import "github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"

type APIPort interface {
	Save(page domain.Page) (int, error)
}
