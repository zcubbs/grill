package api

import (
	"context"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAgents(ctx context.Context, req *agentPb.GetAgentsRequest) (*agentPb.GetAgentsResponse, error) {
	_, err := s.requireUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	violations := validateGetAgentsRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	agents, err := s.store.GetAllAgents(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get agents: %v", err)
	}

	respAgents := make([]*agentPb.Agent, len(agents))
	for i, agent := range agents {
		respAgents[i] = convertAgentToPb(agent)
	}

	return &agentPb.GetAgentsResponse{
		Agents: respAgents,
	}, nil
}

func validateGetAgentsRequest(_ *agentPb.GetAgentsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	return violations
}
