package grpcclient

import (
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
)

func (s *Service) GetAgents() ([]*agentPb.Agent, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}

	response, err := s.client.GetAgents(ctx, &agentPb.GetAgentsRequest{})
	if err != nil {
		return nil, err
	}

	return response.Agents, nil
}
