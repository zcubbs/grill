package grpcclient

import pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"

func (s *Service) CreateUser(user *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}

	response, err := s.client.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return response, nil
}
