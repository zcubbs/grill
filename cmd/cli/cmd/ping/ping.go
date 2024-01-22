package ping

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/utils"
	"github.com/zcubbs/x/pretty"
	"time"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "ping server",
	Long:  `ping API server`,
	Run: func(cmd *cobra.Command, args []string) {
		c := utils.NewCtx()

		ctxWithTimeout, cancel := context.WithTimeout(cmd.Context(), 5*time.Second)
		defer cancel()

		resp, err := c.GrpcClient.Ping(ctxWithTimeout)
		if err != nil {
			log.Fatal("failed to ping server", "error", err.Error())
		}

		log.Info("✔️ server is up and running")
		pretty.PrintJson(resp)
	},
}
