package api

import (
	"context"
	grillPb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendStatus(ctx context.Context, req *grillPb.SendStatusRequest) (*grillPb.SendStatusResponse, error) {
	// authorize agent
	agent, err := s.authorizeAgent(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized")
	}

	// update agent status
	updated, err := s.store.UpdateAgentLastConnection(ctx, agent.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update agent: %v", err)
	}

	return &grillPb.SendStatusResponse{
		Status:     "OK",
		PullConfig: false,
		PullTasks:  false,
	}, nil
}
