package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vedantwankhade/tsrgen/jira-service/config"
	"github.com/vedantwankhade/tsrgen/jira-service/internal/ports"
	"github.com/vedantwankhade/tsrgen/jira-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	api  ports.APIPort
	port int
	service.UnimplementedJiraServer
}

func (s *server) GetIssues(ctx context.Context, req *service.IssueReq) (*service.IssuesRes, error) {
	// TODO)) input validation
	domainIssues, err := s.api.GetIssuesWithJQL(req.GetJql(), req.GetJiraInstance(), req.GetJiraUsername(), req.GetJiraToken())
	if err != nil {
		return nil, fmt.Errorf("error getting issues: %w", err)
	}
	var issues []*service.Issue
	for _, issue := range domainIssues {
		issues = append(issues, &service.Issue{
			Id:       issue.Id,
			Key:      issue.Key,
			TestType: issue.Fields.TestType,
		})
	}
	issuesRes := &service.IssuesRes{
		Total:  int64(len(domainIssues)),
		Issues: issues,
	}
	return issuesRes, nil
}

/* func (s *server) GetIssues(ctx context.Context, req *service.IssueReq) (*service.IssuesRes, error) {
	// TODO)) input validation
	domainIssues, err := s.api.getIssuesWithJQL(req.GetJql(), req.GetJiraInstance(), req.GetJiraUsername(), req.GetJiraToken())
	if err != nil {
		return nil, fmt.Errorf("error getting issues: %w", err)
	}
	var issues []*service.Issue
	for _, issue := range domainIssues {
		issues = append(issues, &service.Issue{
			Id:       issue.Id,
			Key:      issue.Key,
			TestType: issue.Fields.TestType,
		})
	}
	issuesRes := &service.IssuesRes{
		Total:  int64(len(domainIssues)),
		Issues: issues,
	}
	return issuesRes, nil
} */

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
	service.RegisterJiraServer(grpcServer, s)
	if config.GetAppRunMode() == "dev" {
		reflection.Register(grpcServer)
	}
	err = grpcServer.Serve(conn)
	log.Fatalf("error serving grpc request: %v", err)
}
