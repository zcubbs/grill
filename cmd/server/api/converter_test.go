package api

import (
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
		Name:   "testAgent",
		Active: false,
		Group:  "testGroup",
		Token:  "testToken",
		Scopes: "testScopes",
	}

	// Call the function
	dbAgent := convertPbToAgent(pbAgent)

	// Assertions
	assert.Equal(t, pbAgent.GetName(), dbAgent.Name)
	assert.Equal(t, pbAgent.GetActive(), dbAgent.Active)
	assert.Equal(t, pbAgent.GetGroup(), dbAgent.Group)
	assert.Equal(t, pbAgent.GetToken(), dbAgent.Token)
	assert.Equal(t, pbAgent.GetScopes(), dbAgent.Scopes)
}
