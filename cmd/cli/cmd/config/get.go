package config

import (
	"fmt"
	"github.com/spf13/cobra"
	vconf "github.com/zcubbs/grill/cmd/cli/config"
)

// getCmd represents the os command
var getCmd = &cobra.Command{
	Use:    "get",
	Short:  "print config",
	Long:   `print JSON formatted config`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flag("verbose").Value.String() == "true"
		err := printConfig()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func printConfig() error {
	cfg := vconf.Load()
	vconf.PrintConfig(cfg)

	return nil
}

func init() {
	Cmd.AddCommand(getCmd)
}
