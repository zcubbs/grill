// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	BlockSession(ctx context.Context, id uuid.UUID) (Session, error)
	CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetAgent(ctx context.Context, id uuid.UUID) (Agent, error)
	GetAgentByToken(ctx context.Context, token string) (Agent, error)
	GetAllAgents(ctx context.Context) ([]Agent, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSessionByUserId(ctx context.Context, userID uuid.UUID) (Session, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
