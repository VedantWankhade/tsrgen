package ports

import "github.com/vedantwankhade/tsrgen/stats-service/internal/application/core/domain"

type APIPort interface {
	GetHTML(issues []*domain.Issue) (string, error)
}
