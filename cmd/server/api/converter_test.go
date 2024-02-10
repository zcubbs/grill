package api

import (
	grillPb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
)

func TestConvertUserToPb(t *testing.T) {
	// Setup
	dbUser := db.User{
		Username:          "testUser",
		FullName:          "Test User",
		Email:             "test@example.com",
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
		Role:              "ROLE_USER",
	}

	// Call the function
	pbUser := convertUserToPb(dbUser)

	// Assertions
	assert.Equal(t, dbUser.Username, pbUser.GetUsername())
	assert.Equal(t, dbUser.FullName, pbUser.GetFullName())
	assert.Equal(t, dbUser.Email, pbUser.GetEmail())
	assert.Equal(t, dbUser.PasswordChangedAt.UTC(), pbUser.GetPasswordChangedAt().AsTime())
	assert.Equal(t, dbUser.CreatedAt.UTC(), pbUser.GetCreatedAt().AsTime())
	assert.Equal(t, dbUser.Role, pbUser.GetRole().String())
}

func TestConvertPbToUser(t *testing.T) {
	// Setup
	pbUser := &userPb.User{
		Username:          "testUser",
		FullName:          "Test User",
		Email:             "test@example.com",
		PasswordChangedAt: timestamppb.Now(),
		CreatedAt:         timestamppb.Now(),
		Role:              userPb.Role_ROLE_USER,
	}

	// Call the function
	dbUser := convertPbToUser(pbUser)

	// Assertions
	assert.Equal(t, pbUser.GetUsername(), dbUser.Username)
	assert.Equal(t, pbUser.GetFullName(), dbUser.FullName)
	assert.Equal(t, pbUser.GetEmail(), dbUser.Email)
	assert.Equal(t, pbUser.GetPasswordChangedAt().AsTime(), dbUser.PasswordChangedAt.UTC())
	assert.Equal(t, pbUser.GetCreatedAt().AsTime(), dbUser.CreatedAt.UTC())
	assert.Equal(t, pbUser.GetRole().String(), dbUser.Role)
}

func TestConvertPbToAgent(t *testing.T) {
	// Setup
	pbAgent := &agentPb.Agent{
		Name:    "testAgent",
		Version: "testVersion",
		Token:   "testToken",
	}

	// Call the function
	dbAgent := convertPbToAgent(pbAgent)

	// Assertions
	assert.Equal(t, pbAgent.GetName(), dbAgent.Name)
	assert.Equal(t, pbAgent.GetVersion(), dbAgent.Version)
	assert.Equal(t, pbAgent.GetToken(), dbAgent.Token)
}

func TestConvertAgentToPb(t *testing.T) {
	// Setup
	dbAgent := db.Agent{
		Name: "testAgent",
	}

	// Call the function
	pbAgent := convertAgentToPb(dbAgent)

	// Assertions
	assert.Equal(t, dbAgent.Name, pbAgent.GetName())
}

func TestConvertClusterToPb(t *testing.T) {
	// Setup
	dbCluster := db.Cluster{
		Name: "testCluster",
	}

	// Call the function
	pbCluster := convertClusterToPb(dbCluster)

	// Assertions
	assert.Equal(t, dbCluster.Name, pbCluster.GetName())
}

func TestConvertPbToCluster(t *testing.T) {
	// Setup
	pbCluster := &grillPb.Cluster{
		Name: "testCluster",
	}

	// Call the function
	dbCluster := convertPbToCluster(pbCluster)

	// Assertions
	assert.Equal(t, pbCluster.GetName(), dbCluster.Name)

}
func TestConvertPbToNode(t *testing.T) {
	// Setup
	pbNode := &grillPb.Node{
		Name: "testNode",
	}

	// Call the function
	dbNode := convertPbToNode(pbNode)

	// Assertions
	assert.Equal(t, pbNode.GetName(), dbNode.Name)
}

func TestConvertNodeToPb(t *testing.T) {
	// Setup
	dbNode := db.Node{
		Name: "testNode",
	}

	// Call the function
	pbNode := convertNodeToPb(dbNode)

	// Assertions
	assert.Equal(t, dbNode.Name, pbNode.GetName())
}

func Test_convertPbToNodes(t *testing.T) {
	// Setup
	pbNodes := []*grillPb.Node{
		{
			Name: "testNode1",
		},
		{
			Name: "testNode2",
		},
	}

	// Call the function
	dbNodes := convertPbToNodes(pbNodes)

	// Assertions
	assert.Equal(t, pbNodes[0].GetName(), dbNodes[0].Name)
	assert.Equal(t, pbNodes[1].GetName(), dbNodes[1].Name)
}

func Test_convertNodesToPb(t *testing.T) {
	// Setup
	dbNodes := []db.Node{
		{
			Name: "testNode1",
		},
		{
			Name: "testNode2",
		},
	}

	// Call the function
	pbNodes := convertNodesToPb(dbNodes)

	// Assertions
	assert.Equal(t, dbNodes[0].Name, pbNodes[0].GetName())
	assert.Equal(t, dbNodes[1].Name, pbNodes[1].GetName())
}
