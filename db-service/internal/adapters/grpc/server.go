package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/vedantwankhade/tsrgen/db-service/config"
	"github.com/vedantwankhade/tsrgen/db-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/db-service/internal/ports"
	"github.com/vedantwankhade/tsrgen/db-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	api  ports.APIPort
	port int
	service.UnimplementedDBServer
}

func (s *server) GetEntries(ctx context.Context, req *service.None) (*service.EntriesRes, error) {
	// TODO)) input validation
	entries, err := s.api.GetAll()
	var entriesRes service.EntriesRes
	if err != nil {
		return nil, fmt.Errorf("creating entry in db failed: %w", err)
	}
	for _, entry := range entries {
		entriesRes.Entries = append(entriesRes.Entries, &service.Entry{
			EntryId:   int64(entry.EntryId),
			PageId:    entry.PageId,
			PageTitle: entry.PageTitle,
			PageLink:  entry.PageLink,
		})
	}
	return &entriesRes, nil
}

func (s *server) SaveEntry(ctx context.Context, req *service.EntrySaveReq) (*service.EntrySaveRes, error) {
	// TODO)) input validation
	pageReq := domain.Page{
		Id:    req.PageId,
		Title: req.PageTitle,
		Link:  req.PageLink,
	}
	entryId, err := s.api.Save(pageReq)
	if err != nil {
		return nil, fmt.Errorf("creating entry in db failed: %w", err)
	}
	entryRes := &service.EntrySaveRes{
		EntryId: strconv.Itoa(entryId),
	}
	return entryRes, nil
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
	service.RegisterDBServer(grpcServer, s)
	if config.GetAppRunMode() == "dev" {
		reflection.Register(grpcServer)
	}
	err = grpcServer.Serve(conn)
	log.Fatalf("error serving grpc request: %v", err)
}
