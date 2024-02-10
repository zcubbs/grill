package api

import (
	"context"
	"github.com/zcubbs/go-pkg/random"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateAgent handles the creation of a new agent.
func (s *Server) CreateAgent(ctx context.Context, req *agentPb.CreateAgentRequest) (*agentPb.CreateAgentResponse, error) {
	_, err := s.requireAdmin(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	violations := validateCreateAgentRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	// Prepare parameters for the database function.
	params := db.CreateAgentParams{
		Name:     req.GetName(),
		Token:    random.RandomString(32),
		IsActive: req.GetIsActive(),
	}

	// Use the Store interface to save the user details to the database.
	agent, err := s.store.CreateAgent(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create agent: %v", err)
	}

	return &agentPb.CreateAgentResponse{
		Agent: convertAgentToPb(agent),
	}, nil
}

func validateCreateAgentRequest(_ *agentPb.CreateAgentRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	return violations
}
