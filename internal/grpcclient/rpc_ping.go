package grpcclient

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Service) Ping(ctx context.Context) (*pb.PingResponse, error) {
	return s.client.Ping(ctx, &pb.PingRequest{})
}
