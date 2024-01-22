package api

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAgents(ctx context.Context, req *pb.GetAgentsRequest) (*pb.GetAgentsResponse, error) {
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

	respAgents := make([]*pb.Agent, len(agents))
	for i, agent := range agents {
		respAgents[i] = convertAgentToPb(agent)
	}

	return &pb.GetAgentsResponse{
		Agents: respAgents,
	}, nil
}

func validateGetAgentsRequest(_ *pb.GetAgentsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	return violations
}
