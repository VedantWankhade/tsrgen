package clients

import (
	"log"

	"google.golang.org/grpc"
)

type grpcClient struct{}

func NewGRPCClient() *grpcClient {
	return &grpcClient{}
}

func (c *grpcClient) GetConnection(serviceLocation string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.NewClient(serviceLocation, opts...)
	if err != nil {
		log.Fatalf("error starting grpc connection on %s: err: %v", serviceLocation, err)
	}
	return conn
}
