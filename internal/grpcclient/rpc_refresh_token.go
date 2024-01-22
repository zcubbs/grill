package grpcclient

import pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"

func (s *Service) RefreshToken() (*pb.RefreshTokenResponse, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}

	response, err := s.client.RefreshToken(ctx, &pb.RefreshTokenRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}
