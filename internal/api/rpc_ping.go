package api

import (
	"context"
	grillpb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Server) Ping(_ context.Context, _ *grillpb.PingRequest) (*grillpb.PingResponse, error) {
	return &grillpb.PingResponse{
		Message:   "Pong",
		Version:   s.cfg.Version,
		Commit:    s.cfg.Commit,
		BuildTime: s.cfg.Date,
	}, nil
}
