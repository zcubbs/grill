package utils

import (
	"fmt"
	"github.com/zcubbs/grill/cmd/cli/config"
	"github.com/zcubbs/grill/internal/grpcclient"
)

var client *grpcclient.Client

func GetRpcClient(clientConfig *grpcclient.Config) (*grpcclient.Client, error) {
	return grpcclient.New(clientConfig, OnRefreshToken)
}

func OnRefreshToken(cc grpcclient.Config) error {
	cfg := config.Load()

	cfg.GrpcClient.AccessToken = cc.Auth.AccessToken
	cfg.GrpcClient.RefreshToken = cc.Auth.RefreshToken
	cfg.GrpcClient.AccessTokenExpiresAt = cc.Auth.AccessTokenExpiresAt
	cfg.GrpcClient.RefreshTokenExpiresAt = cc.Auth.RefreshTokenExpiresAt

	err := config.SaveConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to save config on refreshToken: %w", err)
	}
	return nil
}

func GetRpcService(clientConfig *grpcclient.Config) (*grpcclient.Client, error) {
	if client != nil {
		return client, nil
	}

	var err error
	client, err = GetRpcClient(clientConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
