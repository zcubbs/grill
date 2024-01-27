package api

import (
	"context"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
)

func (s *Server) RegisterAgent(ctx context.Context, in *agentPb.RegisterAgentRequest) (*agentPb.RegisterAgentResponse, error) {
	// Read token from request
	//token := in.GetToken()
	//
	//// Check if token is valid
	//_, err := s.store.GetToken(ctx, token)

	return &agentPb.RegisterAgentResponse{}, nil
}
