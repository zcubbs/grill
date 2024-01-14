package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"github.com/zcubbs/grill/internal/cmd/ping"
	"github.com/zcubbs/grill/internal/utils"
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

	client, err := utils.GetGRPCClient(*serverAddr)
	if err != nil {
		panic("failed to create client: " + err.Error())
	}

	err = register(client, *token)
	if err != nil {
		panic("register failed: " + err.Error())
	}

	s, err := newScheduler(client, *token)
	if err != nil {
		panic("failed to create scheduler: " + err.Error())
	}
	defer func() { _ = s.Shutdown() }()
	s.Start()

	select {}
}

func register(client pb.GrillServiceClient, token string) error {
	fmt.Println("registering agent...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.RegisterAgent(ctx, &pb.RegisterAgentRequest{
		Token: token,
	})
	if err != nil {
		return err
	}

	return nil
}

func newScheduler(client pb.GrillServiceClient, token string) (gocron.Scheduler, error) {
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
			func(c pb.GrillServiceClient, t string) {
				err := heartbeat(c, t)
				if err != nil {
					fmt.Println("ping failed: ", err)
				}
			}, client, token,
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

func heartbeat(client pb.GrillServiceClient, token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_ = token

	return ping.Ping(client, ctx)
}
