package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/utils"
)

// usersCmd represents the os command
var usersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"user"},
	Short:   "get users",
	Long:    `get users`,
	Hidden:  false,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flag("verbose").Value.String() == "true"
		err := getUsers()
		utils.CheckNoError(err)
	},
}

func getUsers() error {
	ctx := utils.NewCtx()

	users, err := ctx.GrpcClient.GetUsers()
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("No users found")
		return nil
	}

	cols := []interface{}{"Username", "Role", "Email", "Full Name"}
	rows := make([][]interface{}, len(users))
	for _, user := range users {
		rows = append(rows, []interface{}{
			user.Username,
			user.Role,
			user.Email,
			user.FullName,
		})
	}

	utils.PrintTable(cols, rows)

	return nil
}

func init() {
	Cmd.AddCommand(usersCmd)
}
