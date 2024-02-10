package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/zcubbs/grill/cmd/smoke-agent/config"
	"github.com/zcubbs/grill/cmd/smoke-agent/ctx"
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

	s, err := newScheduler(grpcCtx)
	if err != nil {
		panic("failed to create scheduler: " + err.Error())
	}
	defer func() { _ = s.Shutdown() }()
	s.Start()

	select {}
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
			}, grpcCtx,
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
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := grpcCtx.GrpcClient.Ping(timeoutCtx)
	return err
}
