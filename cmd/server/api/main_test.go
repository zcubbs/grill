package api

import (
	"github.com/stretchr/testify/require"
	"github.com/zcubbs/go-pkg/random"
	"github.com/zcubbs/grill/cmd/server/config"
	db "github.com/zcubbs/grill/cmd/server/db/sqlc"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	cfg := &config.Configuration{
		Auth: config.AuthConfig{
			TokenSymmetricKey:    random.RandomString(32),
			AccessTokenDuration:  time.Minute,
			RefreshTokenDuration: 5 * time.Minute,
		},
	}

	server, err := NewServer(store, cfg)
	require.NoError(t, err)

	return server
}
