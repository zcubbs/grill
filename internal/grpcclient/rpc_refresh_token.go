package grpcclient

import userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"

func (s *Service) RefreshToken() (*userPb.RefreshTokenResponse, error) {
	ctx, err := s.client.GetCtxWithAuth()
	if err != nil {
		return nil, err
	}

	response, err := s.client.RefreshToken(ctx, &userPb.RefreshTokenRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}
