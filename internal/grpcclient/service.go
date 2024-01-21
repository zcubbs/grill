package grpcclient

import (
	"github.com/charmbracelet/log"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
)

type Querier interface {
	LoginUser(username, password string) (*pb.LoginUserResponse, error)
	LogoutUser(sessionId string) error
	CreateUser(user *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUsers() ([]*pb.User, error)
	CreateAgent(agent *pb.CreateAgentRequest) (*pb.CreateAgentResponse, error)
	GetAgents() ([]*pb.Agent, error)
	RefreshToken() (*pb.RefreshTokenResponse, error)
	Ping() (*pb.PingResponse, error)
}

type Service struct {
	client *Client
}

func New(cfg *Config) Querier {
	client, err := newClient(cfg)
	if err != nil {
		log.Fatal("failed to create client", "error", err.Error())
	}
	return &Service{
		client: client,
	}
}
