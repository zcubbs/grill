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
	CreateCluster(ctx context.Context, arg CreateClusterParams) (Cluster, error)
	CreateNode(ctx context.Context, arg CreateNodeParams) (Node, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetAgent(ctx context.Context, id uuid.UUID) (Agent, error)
	GetAgentByToken(ctx context.Context, token string) (Agent, error)
	GetAllAgents(ctx context.Context) ([]Agent, error)
	GetAllClusters(ctx context.Context) ([]Cluster, error)
	GetAllClustersActive(ctx context.Context) ([]Cluster, error)
	GetAllNodes(ctx context.Context) ([]Node, error)
	GetAllNodesActive(ctx context.Context) ([]Node, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetCluster(ctx context.Context, id uuid.UUID) (Cluster, error)
	GetClusterByName(ctx context.Context, name string) (Cluster, error)
	GetNode(ctx context.Context, id uuid.UUID) (Node, error)
	GetNodeByName(ctx context.Context, name string) (Node, error)
	GetNodesByCluster(ctx context.Context, clusterID uuid.UUID) ([]Node, error)
	GetNodesByClusterActive(ctx context.Context, clusterID uuid.UUID) ([]Node, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSessionByUserId(ctx context.Context, userID uuid.UUID) (Session, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateAgentLastConnection(ctx context.Context, arg UpdateAgentLastConnectionParams) (Agent, error)
	UpdateCluster(ctx context.Context, arg UpdateClusterParams) (Cluster, error)
	UpdateNode(ctx context.Context, arg UpdateNodeParams) (Node, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
