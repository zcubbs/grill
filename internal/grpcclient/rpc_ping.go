package grpcclient

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Service) Ping() (*pb.PingResponse, error) {
	return s.client.Ping(context.Background(), &pb.PingRequest{})
}
