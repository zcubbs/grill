package grpcclient

import (
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Service) GetAgents() ([]*pb.Agent, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}

	response, err := s.client.GetAgents(ctx, &pb.GetAgentsRequest{})
	if err != nil {
		return nil, err
	}

	return response.Agents, nil
}
