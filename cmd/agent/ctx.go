package main

import (
	"github.com/zcubbs/grill/internal/grpcclient"
)

type Ctx struct {
	Cfg        *AgentConfig
	GrpcClient grpcclient.Querier
}

func NewCtx(cfg *AgentConfig) *Ctx {
	grpcClient := grpcclient.New(&grpcclient.Config{
		Host: cfg.Host,
	})
	return &Ctx{
		Cfg:        cfg,
		GrpcClient: grpcClient,
	}
}
