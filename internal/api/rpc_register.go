package api

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Server) RegisterAgent(ctx context.Context, in *pb.RegisterAgentRequest) (*pb.RegisterAgentResponse, error) {
	// Read token from request
	//token := in.GetToken()
	//
	//// Check if token is valid
	//_, err := s.store.GetToken(ctx, token)

	return &pb.RegisterAgentResponse{}, nil
}
