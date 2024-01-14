package utils

import (
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetGRPCClient Create Grill gRPC client
func GetGRPCClient(serverAddr string) (pb.GrillServiceClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewGrillServiceClient(conn)
	return client, nil
}
