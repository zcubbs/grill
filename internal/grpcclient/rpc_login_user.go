package grpcclient

import (
	"context"
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
)

func (s *Service) LoginUser(username, password string) (*userPb.LoginUserResponse, error) {
	response, err := s.client.LoginUser(context.Background(), &userPb.LoginUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}
