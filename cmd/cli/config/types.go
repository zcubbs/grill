package config

import "time"

const (
	debugPropertyName                           = "debug"
	grpcClientPropertyName                      = "grpc_client"
	grpcClientHostPropertyName                  = "host"
	grpcClientSessionIdPropertyName             = "session_id"
	grpcClientAccessTokenPropertyName           = "access_token"
	grpcClientRefreshTokenPropertyName          = "refresh_token"
	grpcClientAccessTokenExpiresAtPropertyName  = "access_token_expires_at"
	grpcClientRefreshTokenExpiresAtPropertyName = "refresh_token_expires_at"
)

type Config struct {
	Debug      bool             `mapstructure:"debug"`
	GrpcClient GrpcClientConfig `mapstructure:"grpc_client"`
}

type GrpcClientConfig struct {
	Host                  string    `mapstructure:"host"`
	SessionId             string    `mapstructure:"session_id"`
	AccessToken           string    `mapstructure:"access_token"`
	RefreshToken          string    `mapstructure:"refresh_token"`
	AccessTokenExpiresAt  time.Time `mapstructure:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `mapstructure:"refresh_token_expires_at"`
}
