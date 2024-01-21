package login

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:    "login",
	Short:  "login to server",
	Long:   `login to API server`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		err := loginUser()
		if err != nil {
			fmt.Println(err)
		}
	},
}

type Credentials struct {
	Username              string
	SessionId             string
	Password              string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

func loginUser() error {

	return nil
}
