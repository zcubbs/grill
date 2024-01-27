package grpcclient

import userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"

func (s *Service) CreateUser(user *userPb.CreateUserRequest) (*userPb.CreateUserResponse, error) {
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
