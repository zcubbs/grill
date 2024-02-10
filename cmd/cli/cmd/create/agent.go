package create

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/utils"
	agentPb "github.com/zcubbs/grill/gen/proto/go/agent/v1"
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
		_ = cmd.Flag("verbose").Value.String() == "true"
		err := createAgent()
		utils.CheckNoError(err)
	},
}

func createAgent() (err error) {
	// init cli context
	ctx := utils.NewCtx()

	resp, err := ctx.GrpcClient.CreateAgent(&agentPb.CreateAgentRequest{
		Name:     agentName,
		Group:    agentGroup,
		IsActive: agentActive,
	})

	if err != nil {
		return err
	}

	log.Info("Agent created",
		"id", resp.Agent.Id,
		"name", resp.Agent.Name,
		"active", resp.Agent.IsActive,
		"token", resp.Agent.Token,
		"version", resp.Agent.Version,
	)

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
