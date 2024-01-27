package grpcclient

import (
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
)

// CreateAgent creates a new agent
func (s *Service) CreateAgent(agent *agentPb.CreateAgentRequest) (*agentPb.CreateAgentResponse, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}

	response, err := s.client.CreateAgent(ctx, agent)
	if err != nil {
		return nil, err
	}

	return response, nil
}
