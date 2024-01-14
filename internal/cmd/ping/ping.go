package ping

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"github.com/zcubbs/grill/internal/utils"
	"github.com/zcubbs/x/pretty"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "ping server",
	Long:  `ping API server`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flag("verbose").Value.String() == "true"
		serverAddress := cmd.Flag("server").Value.String()
		client, err := utils.GetGRPCClient(serverAddress)
		if err != nil {
			panic("failed to create client: " + err.Error())
		}
		err = Ping(client, cmd.Context())
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Ping(client pb.GrillServiceClient, ctx context.Context) error {
	response, err := client.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		return err
	}

	pretty.PrintJson(response)
	return nil
}
