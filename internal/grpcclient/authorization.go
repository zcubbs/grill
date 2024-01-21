package grpcclient

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/grpc/metadata"
	"time"
)

func (c *Client) GetCtxWithAuth() (ctx context.Context, err error) {
	err = checkAuthOk(*c.cfg)
	if err != nil {
		return nil, err
	}

	now := time.Unix(time.Now().Unix(), 0).UTC()

	if c.cfg.Auth.AccessTokenExpiresAt.Before(now) && c.cfg.Auth.RefreshTokenExpiresAt.After(now) {
		log.Debug("Access token expired, refreshing...")
		err = c.refreshAccessToken()
		if err != nil {
			return nil, err
		}
	} else if c.cfg.Auth.RefreshTokenExpiresAt.Before(now) {
		return nil, fmt.Errorf("logged out from server, please login again")
	}

	ctx = metadata.AppendToOutgoingContext(
		context.Background(),
		"Authorization", "Bearer "+c.cfg.Auth.AccessToken,
	)

	return ctx, nil
}

func checkAuthOk(cfg Config) error {
	if cfg.Auth == nil {
		return fmt.Errorf("user not logged in")
	}

	if cfg.Auth.AccessToken == "" || cfg.Auth.RefreshToken == "" {
		return fmt.Errorf("user not logged in")
	}

	return nil
}

func (c *Client) refreshAccessToken() error {
	ctx := context.Background()
	response, err := c.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: c.cfg.Auth.RefreshToken,
	})
	if err != nil {
		return err
	}

	c.cfg.Auth.AccessToken = response.AccessToken
	c.cfg.Auth.AccessTokenExpiresAt = response.AccessTokenExpiresAt.AsTime()

	err = c.onTokenRefresh(*c.cfg)
	if err != nil {
		return err
	}

	return nil
}
