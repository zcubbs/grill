package ctx

import (
	"github.com/zcubbs/grill/cmd/agent/config"
	"github.com/zcubbs/grill/internal/grpcclient"
)

type Ctx struct {
	Cfg        *config.AgentConfig
	GrpcClient grpcclient.Querier
}

func NewCtx(cfg *config.AgentConfig) *Ctx {
	grpcClient := grpcclient.New(&grpcclient.Config{
		Host: cfg.Host,
	})
	return &Ctx{
		Cfg:        cfg,
		GrpcClient: grpcClient,
	}
}
