package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vedantwankhade/tsrgen/confluence-service/config"
	"github.com/vedantwankhade/tsrgen/confluence-service/internal/ports"
	"github.com/vedantwankhade/tsrgen/confluence-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	api  ports.APIPort
	port int
	service.UnimplementedConfluenceServer
}

func (s *server) CreatePage(ctx context.Context, req *service.PageCreateReq) (*service.PageRes, error) {
	// TODO)) input validation
	page, err := s.api.CreatePage(req.GetHtmlContent(), req.GetTitle(), req.GetParentPageId(), req.GetSpaceId(), req.GetConfluenceInstance(), req.GetConfleunceUsername(), req.GetConfluenceToken())
	if err != nil {
		return nil, fmt.Errorf("page creation request failed: %w", err)
	}
	pageRes := &service.PageRes{
		Id:    page.ID,
		Html:  page.Body.HTML["value"],
		Title: page.Title,
	}
	return pageRes, nil
}

func (s *server) GetPage(ctx context.Context, req *service.PageReq) (*service.PageRes, error) {
	// TODO)) better input validation
	if req.GetId() == 0 || req.GetConfluenceToken() == "" || req.GetConfleunceUsername() == "" || req.GetConfluenceToken() == "" {
		return nil, status.Errorf(codes.NotFound, "Invalid page id %d", req.GetId())
	}
	page, err := s.api.GetPageFromID(int(req.GetId()), req.GetConfluenceInstance(), req.GetConfleunceUsername(), req.GetConfluenceToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not get the page %d: %v", req.GetId(), err)
	}
	pageRes := &service.PageRes{
		Id:    page.ID,
		Title: page.Title,
		Html:  page.Body.HTML["value"],
	}
	return pageRes, nil
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
	service.RegisterConfluenceServer(grpcServer, s)
	if config.GetAppRunMode() == "dev" {
		reflection.Register(grpcServer)
	}
	err = grpcServer.Serve(conn)
	log.Fatalf("error serving grpc request: %v", err)
}
