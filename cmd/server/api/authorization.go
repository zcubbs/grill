package api

import (
	"context"
	"fmt"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
	"github.com/zcubbs/grill/internal/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type ContextValue string

const (
	authorizationHeader ContextValue = "authorization"
	authorizationBearer ContextValue = "bearer"
	authorizationApiKey ContextValue = "api-key"
)

func (s *Server) requireUser(ctx context.Context) (*token.Payload, error) {
	authPayload, err := s.getPayload(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	u, err := s.store.GetUserByUsername(ctx, authPayload.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	if u.Role != userPb.Role_ROLE_USER.String() && u.Role != userPb.Role_ROLE_ADMIN.String() {
		return nil, status.Errorf(codes.PermissionDenied, "user lacks 'user' role")
	}

	return authPayload, nil
}

func (s *Server) requireAdmin(ctx context.Context) (*token.Payload, error) {
	authPayload, err := s.getPayload(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	u, err := s.store.GetUserByUsername(ctx, authPayload.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	if u.Role != userPb.Role_ROLE_ADMIN.String() {
		return nil, status.Errorf(codes.PermissionDenied, "user is not an admin")
	}

	return authPayload, nil
}

func (s *Server) authorizeAgent(ctx context.Context) (*db.Agent, error) {
	// Get metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no auth metadata found in request")
	}

	// Read from metadata
	values := md.Get(string(authorizationHeader))
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	// Get API Key from header
	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) != 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if (ContextValue)(authType) != authorizationApiKey {
		return nil, fmt.Errorf("unsupported authorization type %s", authType)
	}

	apiKey := fields[1]

	agent, err := s.store.GetAgentByToken(ctx, apiKey)
	if err != nil {
		return nil, fmt.Errorf("invalid api key: %w", err)
	}

	return &agent, nil
}

func (s *Server) getPayload(ctx context.Context) (*token.Payload, error) {
	// Get metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no auth metadata found in request")
	}

	// Read from metadata
	values := md.Get(string(authorizationHeader))
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	// Get token from header
	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) != 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if (ContextValue)(authType) != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type %s", authType)
	}

	accessToken := fields[1]
	payload, err := s.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	return payload, nil
}
