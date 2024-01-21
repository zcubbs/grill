package create

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "create",
	Short: "create agents & users",
	Long:  `create agents & users`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	Cmd.AddCommand(createAgentCmd)
}
