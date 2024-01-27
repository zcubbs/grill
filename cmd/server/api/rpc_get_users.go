package api

import (
	"context"
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUsers(ctx context.Context, req *userPb.GetUsersRequest) (*userPb.GetUsersResponse, error) {
	_, err := s.requireUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	violations := validateGetUsersRequest(req)
	if len(violations) > 0 {
		return nil, invalidArgumentError(violations)
	}

	users, err := s.store.GetAllUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get users: %v", err)
	}

	respUsers := make([]*userPb.User, len(users))
	for i, user := range users {
		respUsers[i] = convertUserToPb(user)
	}

	return &userPb.GetUsersResponse{Users: respUsers}, nil
}

func validateGetUsersRequest(_ *userPb.GetUsersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	return violations
}
