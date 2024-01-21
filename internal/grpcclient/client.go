package grpcclient

import (
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	cfg *Config

	pb.GrillServiceClient
}

type Config struct {
	Host             string
	Auth             *AuthData
	RefreshTokenHook RefreshTokenHook
}

type AuthData struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

type RefreshTokenHook func(Config) error

func newClient(cfg *Config) (*Client, error) {
	client, err := getClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		cfg:                cfg,
		GrillServiceClient: client,
	}, nil
}

func getClient(cfg *Config) (pb.GrillServiceClient, error) {
	// TLS config
	//tc := credentials.NewTLS(&tls.Config{
	//	InsecureSkipVerify: true,
	//})

	conn, err := grpc.Dial(cfg.Host,
		//grpc.WithTransportCredentials(tc), // TLS
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	client := pb.NewGrillServiceClient(conn)

	return client, nil
}
