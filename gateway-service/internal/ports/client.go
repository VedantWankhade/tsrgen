package ports

import (
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"google.golang.org/grpc"
)

type ConfluenceClientPort interface {
	GetPage(int, string, string, string) (*domain.Page, error)
	CreatePage(string, string, string, string, string, string, string) (*domain.Page, error)
}

type StatsClientPort interface {
	GetHTML([]*domain.Issue) (string, error)
}

type DBClientPort interface {
	SaveEntry(domain.DBPageSaveReq) (int, error)
}
type JiraClientPort interface {
	GetIssues(string, string, string, string) ([]*domain.Issue, error)
}

type GRPCClientPort interface {
	GetConnection(int) *grpc.ClientConn
}
