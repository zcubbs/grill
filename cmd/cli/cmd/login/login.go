package login

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/config"
	"github.com/zcubbs/grill/cmd/cli/utils"
	"github.com/zcubbs/grill/internal/grpcclient"
	"golang.org/x/term"
	"os"
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
	ctx := utils.NewCtx()

	c, err := ReadCredentials(ctx.Cfg)

	if err != nil || c.AccessToken == "" || c.RefreshToken == "" {
		c.Username, c.Password = promptForCredentials()
		err := makeAPICall(ctx.GrpcClient, c, ctx.Cfg.Debug)
		if err != nil {
			return err
		}

		ctx.Cfg.GrpcClient.SessionId = c.SessionId
		ctx.Cfg.GrpcClient.AccessToken = c.AccessToken
		ctx.Cfg.GrpcClient.RefreshToken = c.RefreshToken
		ctx.Cfg.GrpcClient.AccessTokenExpiresAt = c.AccessTokenExpiresAt
		ctx.Cfg.GrpcClient.RefreshTokenExpiresAt = c.RefreshTokenExpiresAt

		if ctx.Cfg.Debug {
			fmt.Println("Access Token:", ctx.Cfg.GrpcClient.AccessToken)
			fmt.Println("Refresh Token:", ctx.Cfg.GrpcClient.RefreshToken)
			fmt.Println("Access Token Expires At:", ctx.Cfg.GrpcClient.AccessTokenExpiresAt)
			fmt.Println("Refresh Token Expires At:", ctx.Cfg.GrpcClient.RefreshTokenExpiresAt)
		}

		err = config.SaveConfig(ctx.Cfg)
		if err != nil {
			return err
		}
	}

	fmt.Println("Successfully authenticated!")

	return nil
}

func promptForCredentials() (string, string) {
	var username string
	var password string

	// Prompt for username
	fmt.Print("Enter Username: ")
	_, err := fmt.Scan(&username)
	if err != nil {
		fmt.Println("Error reading username:", err)
		return "", ""
	}

	// Prompt for password without echoing input
	fmt.Print("Enter Password: ")
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
		return "", ""
	}
	password = string(passwordBytes)
	fmt.Println() // Move to a new line after password input

	return username, password
}

// ReadCredentials reads the stored authentication details from the YAML file
func ReadCredentials(cfg *config.Config) (*Credentials, error) {
	var c Credentials
	c.AccessToken = cfg.GrpcClient.AccessToken
	c.RefreshToken = cfg.GrpcClient.RefreshToken

	return &c, nil
}

// makeAPICall simulates an API call for authentication
func makeAPICall(client grpcclient.Querier, credentials *Credentials, debug bool) (err error) {
	resp, err := client.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		return err
	}

	if debug {
		fmt.Println("Response:", resp)
	}

	credentials.AccessToken = resp.AccessToken
	credentials.RefreshToken = resp.RefreshToken
	credentials.AccessTokenExpiresAt = resp.AccessTokenExpiresAt.AsTime()
	credentials.RefreshTokenExpiresAt = resp.RefreshTokenExpiresAt.AsTime()
	credentials.SessionId = resp.SessionId

	return nil
}
