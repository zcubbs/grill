package grpcclient

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Service) RegisterAgent(ctx context.Context, in *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	response, err := s.client.RegisterAgent(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}
