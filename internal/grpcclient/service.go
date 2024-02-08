package grpcclient

import (
    "context"
    "github.com/charmbracelet/log"
    agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
    grillPb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
    userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
)

type Querier interface {
    LoginUser(username, password string) (*userPb.LoginUserResponse, error)
    LogoutUser(sessionId string) error
    CreateUser(user *userPb.CreateUserRequest) (*userPb.CreateUserResponse, error)
    GetUsers() ([]*userPb.User, error)
    CreateAgent(agent *agentPb.CreateAgentRequest) (*agentPb.CreateAgentResponse, error)
    GetAgents() ([]*agentPb.Agent, error)
    RefreshToken() (*userPb.RefreshTokenResponse, error)
    Ping(context.Context) (*grillPb.PingResponse, error)
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
