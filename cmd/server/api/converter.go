package api

import (
	"github.com/google/uuid"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUserToPb(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
		Role:              pb.Role(pb.Role_value[user.Role]),
	}
}

func convertPbToUser(user *pb.User) db.User {
	return db.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt.AsTime(),
		CreatedAt:         user.CreatedAt.AsTime(),
		Role:              user.Role.String(),
	}
}

func convertAgentToPb(agent db.Agent) *pb.Agent {
	return &pb.Agent{
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

func convertPbToAgent(agent *pb.Agent) db.Agent {
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
