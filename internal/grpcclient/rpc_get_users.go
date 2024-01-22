package grpcclient

import (
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Service) GetUsers() ([]*pb.User, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}
	response, err := s.client.GetUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}

	return response.Users, nil
}
