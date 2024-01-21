package create

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	vconf "github.com/zcubbs/grill/cmd/cli/config"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
	"github.com/zcubbs/grill/internal/grpcclient"
	"github.com/zcubbs/x/pretty"
)

var (
	agentName   string
	agentActive bool
	agentGroup  string
	agentScopes []string
)

// createAgentCmd represents the os command
var createAgentCmd = &cobra.Command{
	Use:    "agent",
	Short:  "create agent",
	Long:   `create agent`,
	Hidden: false,
	Run: func(cmd *cobra.Command, args []string) {
		verbose := cmd.Flag("verbose").Value.String() == "true"
		err := createAgent(verbose)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func createAgent(verbose bool) (err error) {
	// load config
	cfg := vconf.Load()

	client := grpcclient.New(&grpcclient.Config{
		Host: cfg.GrpcClient.Host,
	})

	if err != nil {
		log.Fatal("failed to init grpc client", "error", err.Error())
	}

	resp, err := client.CreateAgent(&pb.CreateAgentRequest{
		Name:     agentName,
		Group:    agentGroup,
		IsActive: agentActive,
	})

	if err != nil {
		return err
	}

	if verbose {
		pretty.PrintJson(resp)
	}

	return nil
}

func init() {
	createAgentCmd.Flags().StringVarP(&agentName, "name", "n", "", "Name")
	createAgentCmd.Flags().BoolVarP(&agentActive, "active", "a", false, "Active")
	createAgentCmd.Flags().StringVarP(&agentGroup, "group", "g", "", "Group")
	createAgentCmd.Flags().StringSliceVarP(&agentScopes, "scopes", "s", []string{}, "Scopes")

	_ = createAgentCmd.MarkFlagRequired("name")
	_ = createAgentCmd.MarkFlagRequired("group")

	Cmd.AddCommand(createAgentCmd)
}
