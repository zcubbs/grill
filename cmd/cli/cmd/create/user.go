package create

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/utils"
	userPb "github.com/zcubbs/grill/gen/proto/go/user/v1"
)

var (
	userUsername string
	userFullname string
	userEmail    string
	userPassword string
	isAdmin      bool
)

// Cmd represents the os command
var user = &cobra.Command{
	Use:    "user",
	Short:  "create user",
	Long:   `create user`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flag("verbose").Value.String() == "true"
		err := createUser()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func createUser() (err error) {
	ctx := utils.NewCtx()

	var userRole userPb.Role
	if isAdmin {
		userRole = userPb.Role_ROLE_ADMIN
	} else {
		userRole = userPb.Role_ROLE_USER
	}
	response, err := ctx.GrpcClient.CreateUser(&userPb.CreateUserRequest{
		Username: userUsername,
		FullName: userFullname,
		Email:    userEmail,
		Password: userPassword,
		Role:     userRole,
	})

	if err != nil {
		return err
	}

	log.Info("User created",
		"username", response.User.Username,
		"role", response.User.Role,
		"email", response.User.Email,
		"full_name", response.User.FullName,
	)

	return nil
}

func init() {
	user.Flags().StringVarP(&userUsername, "username", "u", "", "Username")
	user.Flags().StringVarP(&userFullname, "fullname", "f", "", "Full name")
	user.Flags().StringVarP(&userEmail, "email", "e", "", "Email")
	user.Flags().StringVarP(&userPassword, "password", "p", "", "Password")
	user.Flags().BoolVarP(&isAdmin, "admin", "a", false, "Is admin")

	_ = user.MarkFlagRequired("username")
	_ = user.MarkFlagRequired("fullname")
	_ = user.MarkFlagRequired("email")
	_ = user.MarkFlagRequired("password")

	Cmd.AddCommand(user)
}
