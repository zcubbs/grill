package api

import (
	"github.com/google/uuid"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserToPb(user db.User) *userPb.User {
	return &userPb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
		Role:              userPb.Role(userPb.Role_value[user.Role]),
	}
}

func convertPbToUser(user *userPb.User) db.User {
	return db.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt.AsTime(),
		CreatedAt:         user.CreatedAt.AsTime(),
		Role:              user.Role.String(),
	}
}

func convertAgentToPb(agent db.Agent) *agentPb.Agent {
	return &agentPb.Agent{
		Id:             agent.ID.String(),
		Name:           agent.Name,
		Active:         agent.Active,
		Group:          agent.Group,
		Token:          agent.Token,
		Scopes:         agent.Scopes,
		CreatedAt:      timestamppb.New(agent.CreatedAt),
		LastConnection: timestamppb.New(agent.LastConnection.Time),
	}
}

func convertPbToAgent(agent *agentPb.Agent) db.Agent {
	id, _ := uuid.Parse(agent.Id)
	return db.Agent{
		ID:     id,
		Name:   agent.Name,
		Group:  agent.Group,
		Token:  agent.Token,
		Scopes: agent.Scopes,
		Active: agent.Active,
	}
}
