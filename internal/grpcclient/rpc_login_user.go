package grpcclient

import (
	"context"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

func (s *Service) LoginUser(username, password string) (*pb.LoginUserResponse, error) {
	response, err := s.client.LoginUser(context.Background(), &pb.LoginUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}
	return response, nil
}
