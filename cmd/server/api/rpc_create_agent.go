package api

import (
	"context"
	"github.com/zcubbs/go-pkg/random"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateAgent handles the creation of a new agent.
func (s *Server) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentResponse, error) {
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
		Name:   req.GetName(),
		Token:  random.RandomString(32),
		Group:  req.GetGroup(),
		Scopes: "all",
		Active: req.GetIsActive(),
	}

	// Use the Store interface to save the user details to the database.
	agent, err := s.store.CreateAgent(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create agent: %v", err)
	}

	return &pb.CreateAgentResponse{
		Agent: convertAgentToPb(agent),
	}, nil
}

func validateCreateAgentRequest(_ *pb.CreateAgentRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	return violations
}
