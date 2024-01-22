package utils

import (
	"github.com/zcubbs/grill/cmd/cli/config"
	"github.com/zcubbs/grill/internal/grpcclient"
)

type Ctx struct {
	Cfg        *config.Config
	GrpcClient grpcclient.Querier
}

func NewCtx() *Ctx {
	cfg := config.Load()
	grpcClient := grpcclient.New(&grpcclient.Config{
		Host: cfg.GrpcClient.Host,
		Auth: &grpcclient.AuthData{
			AccessToken:           cfg.GrpcClient.AccessToken,
			RefreshToken:          cfg.GrpcClient.RefreshToken,
			AccessTokenExpiresAt:  cfg.GrpcClient.AccessTokenExpiresAt,
			RefreshTokenExpiresAt: cfg.GrpcClient.RefreshTokenExpiresAt,
		},
		RefreshTokenHook: OnRefreshToken,
	})
	return &Ctx{
		Cfg:        cfg,
		GrpcClient: grpcClient,
	}
}
