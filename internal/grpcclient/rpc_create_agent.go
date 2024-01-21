package grpcclient

import pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"

// CreateAgent creates a new agent
func (s *Service) CreateAgent(agent *pb.CreateAgentRequest) (*pb.CreateAgentResponse, error) {
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
