package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vedantwankhade/tsrgen/stats-service/config"
	"github.com/vedantwankhade/tsrgen/stats-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/stats-service/internal/ports"
	"github.com/vedantwankhade/tsrgen/stats-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	api  ports.APIPort
	port int
	service.UnimplementedStatsServer
}

func (s *server) GetHTML(ctx context.Context, req *service.StatsReq) (*service.StatsRes, error) {
	// TODO)) input validation
	var issues []*domain.Issue
	for _, issue := range req.Issues {
		issues = append(issues, &domain.Issue{
			Id:  issue.Id,
			Key: issue.Key,
			Fields: domain.IssueFields{
				TestType: issue.TestType,
			},
		})
	}
	html, err := s.api.GetHTML(issues)
	if err != nil {
		return nil, fmt.Errorf("error getting issues: %w", err)
	}
	issuesRes := &service.StatsRes{
		Html: html,
	}
	return issuesRes, nil
}

func NewServer(api ports.APIPort, port int) *server {
	return &server{
		api:  api,
		port: port,
	}
}

func (s *server) Run() {
	log.Printf("Starting the grpc server on port %d\n", s.port)
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("error starting the grpc server: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service.RegisterStatsServer(grpcServer, s)
	if config.GetAppRunMode() == "dev" {
		reflection.Register(grpcServer)
	}
	err = grpcServer.Serve(conn)
	log.Fatalf("error serving grpc request: %v", err)
}
