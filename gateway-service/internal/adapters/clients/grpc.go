package clients

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type grpcClient struct{}

func NewGRPCClient() *grpcClient {
	return &grpcClient{}
}

func (c *grpcClient) GetConnection(port int) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.NewClient(fmt.Sprintf(":%d", port), opts...)
	if err != nil {
		log.Fatalf("error starting grpc connection on %d: err: %v", port, err)
	}
	return conn
}
