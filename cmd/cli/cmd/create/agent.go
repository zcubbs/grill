package create

import (
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/utils"
	pb "github.com/zcubbs/grill/gen/proto/go/grill/v1"
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
		utils.CheckNoError(err)
	},
}

func createAgent(verbose bool) (err error) {
	// init cli context
	ctx := utils.NewCtx()

	resp, err := ctx.GrpcClient.CreateAgent(&pb.CreateAgentRequest{
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
