package ping

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	vconf "github.com/zcubbs/grill/cmd/cli/config"
	"github.com/zcubbs/grill/internal/grpcclient"
	"github.com/zcubbs/x/pretty"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "ping server",
	Long:  `ping API server`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose := cmd.Flag("verbose").Value.String() == "true"
		// load config
		cfg := vconf.Load()

		if verbose {
			pretty.PrintJson(cfg)
		}

		if cfg.GrpcClient.Host == "" {
			fmt.Println("please provide a server address: grill config set grpc-client.host <host>")
			return
		}

		client := grpcclient.New(&grpcclient.Config{
			Host: cfg.GrpcClient.Host,
		})

		err := client.Ping()
		if err != nil {
			log.Error("failed to ping server", "error", err.Error())
		}

		log.Info("server is up and running")
	},
}
