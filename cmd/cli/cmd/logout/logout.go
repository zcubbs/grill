package logout

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/config"
	"github.com/zcubbs/grill/cmd/cli/utils"
	"time"
)

// Cmd represents the os command
var Cmd = &cobra.Command{
	Use:    "logout",
	Short:  "logout",
	Long:   `logout user`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		err := logoutUser()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to logout user: %w", err))
		}
	},
}

func logoutUser() error {
	c := utils.NewCtx()

	err := c.GrpcClient.LogoutUser(c.Cfg.GrpcClient.SessionId)
	if err != nil {
		return fmt.Errorf("failed to logout user: %w", err)
	}

	c.Cfg.GrpcClient.SessionId = ""
	c.Cfg.GrpcClient.AccessToken = ""
	c.Cfg.GrpcClient.RefreshToken = ""
	c.Cfg.GrpcClient.AccessTokenExpiresAt = time.Now().Add(-1 * time.Hour)
	c.Cfg.GrpcClient.RefreshTokenExpiresAt = time.Now().Add(-1 * time.Hour)

	err = config.ResetConfig(c.Cfg)

	if err != nil {
		return err
	}

	return nil
}
