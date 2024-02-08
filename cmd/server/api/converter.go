package api

import (
	"github.com/google/uuid"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
	grillPb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
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
		IsActive:       agent.Active,
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
		Active: agent.IsActive,
	}
}

func convertClusterToPb(cluster db.Cluster) *grillPb.Cluster {
	return &grillPb.Cluster{
		Id:          cluster.ID.String(),
		Name:        cluster.Name,
		Description: cluster.Description,
		IsActive:    cluster.IsActive,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
}

func convertPbToCluster(cluster *grillPb.Cluster) db.Cluster {
	id, _ := uuid.Parse(cluster.Id)
	return db.Cluster{
		ID:          id,
		Name:        cluster.Name,
		Description: cluster.Description,
		IsActive:    cluster.IsActive,
	}
}

func convertNodeToPb(node db.Node) *grillPb.Node {
	return &grillPb.Node{
		Id:        node.ID.String(),
		Name:      node.Name,
		Ip:        node.Ip,
		Cpu:       node.Cpu,
		Memory:    node.Memory,
		Disk:      node.Disk,
		IsActive:  node.IsActive,
		ClusterId: node.ClusterID.String(),
		CreatedAt: timestamppb.New(node.CreatedAt),
		UpdatedAt: timestamppb.New(node.UpdatedAt),
	}
}

func convertPbToNode(node *grillPb.Node) db.Node {
	id, _ := uuid.Parse(node.Id)
	clusterId, _ := uuid.Parse(node.ClusterId)
	return db.Node{
		ID:        id,
		Name:      node.Name,
		Ip:        node.Ip,
		Cpu:       node.Cpu,
		Memory:    node.Memory,
		Disk:      node.Disk,
		ClusterID: clusterId,
		IsActive:  node.IsActive,
		CreatedAt: node.CreatedAt.AsTime(),
		UpdatedAt: node.UpdatedAt.AsTime(),
	}
}

func convertNodesToPb(nodes []db.Node) []*grillPb.Node {
	var pbNodes []*grillPb.Node
	for _, node := range nodes {
		pbNodes = append(pbNodes, convertNodeToPb(node))
	}
	return pbNodes
}

func convertPbToNodes(nodes []*grillPb.Node) []db.Node {
	var dbNodes []db.Node
	for _, node := range nodes {
		dbNodes = append(dbNodes, convertPbToNode(node))
	}
	return dbNodes
}
