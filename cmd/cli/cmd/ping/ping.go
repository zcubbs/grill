package ping

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/utils"
	"github.com/zcubbs/x/pretty"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "ping server",
	Long:  `ping API server`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := utils.NewCtx()

		resp, err := ctx.GrpcClient.Ping()
		if err != nil {
			log.Fatal("failed to ping server", "error", err.Error())
		}

		log.Info("✔️ server is up and running")
		pretty.PrintJson(resp)
	},
}
