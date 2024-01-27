package grpcclient

import (
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
	grillPb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	cfg *Config

	grillPb.GrillServiceClient
	userPb.UserServiceClient
	agentPb.AgentServiceClient
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
	grillClient, err := getGrillClient(cfg)
	if err != nil {
		return nil, err
	}

	userClient, err := getUserClient(cfg)
	if err != nil {
		return nil, err
	}

	agentClient, err := getAgentClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		cfg:                cfg,
		GrillServiceClient: grillClient,
		UserServiceClient:  userClient,
		AgentServiceClient: agentClient,
	}, nil
}

func getGrillClient(cfg *Config) (grillPb.GrillServiceClient, error) {
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

	client := grillPb.NewGrillServiceClient(conn)

	return client, nil
}

func getUserClient(cfg *Config) (userPb.UserServiceClient, error) {
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

	client := userPb.NewUserServiceClient(conn)

	return client, nil
}

func getAgentClient(cfg *Config) (agentPb.AgentServiceClient, error) {
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

	client := agentPb.NewAgentServiceClient(conn)

	return client, nil
}
