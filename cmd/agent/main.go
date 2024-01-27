package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/zcubbs/grill/cmd/agent/config"
	"github.com/zcubbs/grill/cmd/agent/ctx"
	pbAgent "github.com/zcubbs/grill/gen/proto/go/agent/v1"
	"time"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var (
	serverAddr = flag.String("server", "localhost:9000", "The server address in the format of host:port")
	token      = flag.String("token", "", "The token to use for authentication")
)

func main() {
	flag.Parse()

	fmt.Println("Version: ", Version)
	fmt.Println("Commit: ", Commit)
	fmt.Println("Date: ", Date)

	fmt.Println("serverAddr: ", *serverAddr)
	fmt.Println("token: ", *token)

	grpcCtx := ctx.NewCtx(&config.AgentConfig{
		Host:  *serverAddr,
		Token: *token,
	})

	err := register(grpcCtx)
	if err != nil {
		panic("register failed: " + err.Error())
	}

	s, err := newScheduler(grpcCtx)
	if err != nil {
		panic("failed to create scheduler: " + err.Error())
	}
	defer func() { _ = s.Shutdown() }()
	s.Start()

	select {}
}

func register(grpcCtx *ctx.Ctx) error {
	fmt.Println("registering agent...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := grpcCtx.GrpcClient.RegisterAgent(ctx, &pbAgent.RegisterAgentRequest{
		Token: grpcCtx.Cfg.Token,
	})
	if err != nil {
		return err
	}

	return nil
}

func newScheduler(grpcCtx *ctx.Ctx) (gocron.Scheduler, error) {
	s, err := gocron.NewScheduler(
		gocron.WithStopTimeout(10 * time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create scheduler: %w", err)
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			5*time.Second,
		),
		gocron.NewTask(
			func(grpcCtx *ctx.Ctx) {
				err := heartbeat(grpcCtx)
				if err != nil {
					fmt.Println("ping failed: ", err)
				}
			}, grpcCtx.GrpcClient,
		),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
		gocron.WithName("PingServerJob"),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to instanciate job: %v", err.Error())
	}

	fmt.Printf("Ready to run %d jobs!\n", len(s.Jobs()))
	for _, job := range s.Jobs() {
		fmt.Printf("      job: %s\n", job.Name())
	}

	return s, nil
}

func heartbeat(grpcCtx *ctx.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := grpcCtx.GrpcClient.Ping(ctx)
	return err
}
