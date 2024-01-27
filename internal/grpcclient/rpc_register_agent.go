package grpcclient

import (
	"context"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
)

func (s *Service) RegisterAgent(ctx context.Context, in *agentPb.RegisterAgentRequest) (*agentPb.RegisterAgentResponse, error) {
	response, err := s.client.RegisterAgent(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, nil
}
