package ports

import (
	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"google.golang.org/grpc"
)

type ConfluenceClientPort interface {
	GetPage(int, string, string, string) (*domain.Page, error)
	CreatePage(string, string, string, string, string, string, string) (*domain.Page, error)
}

type GRPCClientPort interface {
	GetConnection(int) *grpc.ClientConn
}
