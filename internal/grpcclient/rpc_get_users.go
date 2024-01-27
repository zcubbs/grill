package grpcclient

import (
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
)

func (s *Service) GetUsers() ([]*userPb.User, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}
	response, err := s.client.GetUsers(ctx, &userPb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}

	return response.Users, nil
}
