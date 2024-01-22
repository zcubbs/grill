package api

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zcubbs/grill/cmd/server/config"
	mockdb "github.com/zcubbs/grill/cmd/server/db/mock"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"github.com/zcubbs/grill/internal/token"
	"github.com/zcubbs/x/random"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizeUser_Success(t *testing.T) {
	// Setup
	storeCtrl := gomock.NewController(t)
	defer storeCtrl.Finish()

	store := mockdb.NewMockStore(storeCtrl)

	s := newTestServer(t, store)

	username := random.String(32)
	userId := uuid.UUID{}
	duration := time.Minute

	createToken, _, err := s.tokenMaker.CreateToken(username, userId, duration)
	assert.Nil(t, err)
	assert.NotEmpty(t, createToken)

	md := metadata.New(map[string]string{
		string(authorizationHeader): "Bearer " + createToken,
	})

	ctx := metadata.NewIncomingContext(context.Background(), md)

	user := db.User{
		ID:       uuid.New(),
		Username: "test",
		Role:     "ROLE_USER",
	}

	store.EXPECT().
		GetUserByUsername(gomock.Any(), gomock.Any()).
		Times(1).
		Return(user, nil)

	// Call the function
	payload, err := s.requireUser(ctx)

	// Assertions
	assert.Nil(t, err)
	assert.NotNil(t, payload)
}

func TestServer_authorizeAgent(t *testing.T) {
	type fields struct {
		UnimplementedLinkupServiceServer pb.UnimplementedGrillServiceServer
		store                            db.Store
		tokenMaker                       token.Maker
		cfg                              *config.Configuration
		embedOpts                        []EmbedAssetsOpts
	}
	type args struct {
		ctx context.Context
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *db.Agent
		wantErr assert.ErrorAssertionFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGrillServiceServer: tt.fields.UnimplementedLinkupServiceServer,
				store:                           tt.fields.store,
				tokenMaker:                      tt.fields.tokenMaker,
				cfg:                             tt.fields.cfg,
				embedOpts:                       tt.fields.embedOpts,
			}
			got, err := s.authorizeAgent(tt.args.ctx)
			if !tt.wantErr(t, err, fmt.Sprintf("authorizeAgent(%v)", tt.args.ctx)) {
				return
			}
			assert.Equalf(t, tt.want, got, "authorizeAgent(%v)", tt.args.ctx)
		})
	}
}
