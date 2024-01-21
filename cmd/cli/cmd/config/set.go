package config

import (
	"fmt"
	"github.com/spf13/cobra"
	vconf "github.com/zcubbs/grill/cmd/cli/config"
)

var (
	host  string
	debug bool
)

// setCmd represents the os command
var setCmd = &cobra.Command{
	Use:    "set",
	Short:  "set config value",
	Long:   `set config value. possible values: [host, debug]`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flag("verbose").Value.String() == "true"
		err := setConfig()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to set config: %w", err))
		}
	},
}

func setConfig() error {
	cfg := vconf.Load()

	if host != "" {
		cfg.GrpcClient.Host = host
	}

	if debug {
		cfg.Debug = debug
	}

	return vconf.SaveConfig(cfg)
}

func init() {
	setCmd.Flags().StringVarP(&host, "host", "H", "", "API server host")
	setCmd.Flags().BoolVarP(&debug, "debug", "d", false, "debug mode")

	Cmd.AddCommand(setCmd)
}
