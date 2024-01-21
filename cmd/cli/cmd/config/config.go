package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:    "config",
	Short:  "print config",
	Long:   `print config`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
		}
	},
}
