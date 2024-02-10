package api

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
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
	_, err = s.store.UpdateAgentLastConnection(ctx, db.UpdateAgentLastConnectionParams{
		ID:      agent.ID,
		Version: convertToText(req.Version),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update agent: %v", err)
	}

	return &grillPb.SendStatusResponse{
		Status:     "OK",
		PullConfig: false,
		PullTasks:  false,
	}, nil
}

func convertToText(input string) pgtype.Text {
	var output pgtype.Text
	output.String = input
	return output
}
