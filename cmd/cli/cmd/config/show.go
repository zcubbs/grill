package config

import (
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/config"
)

var showCmd = &cobra.Command{
	Use:    "show",
	Short:  "show config",
	Long:   `show config`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Load()

		config.PrintConfig(cfg)
	},
}

func init() {
	Cmd.AddCommand(showCmd)
}
