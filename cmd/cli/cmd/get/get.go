package get

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:   "get",
	Short: "get agents & users",
	Long:  `get agents & users`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Println(err)
		}
	},
}
