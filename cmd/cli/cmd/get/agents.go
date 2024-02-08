package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zcubbs/grill/cmd/cli/utils"
	zt "github.com/zcubbs/x/time"
	"time"
)

// agentsCmd represents the os command
var agentsCmd = &cobra.Command{
	Use:     "agents",
	Aliases: []string{"agent"},
	Short:   "get agents",
	Long:    `get agents`,
	Hidden:  false,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flag("verbose").Value.String() == "true"
		err := getAgents()
		utils.CheckNoError(err)
	},
}

func getAgents() (err error) {
	ctx := utils.NewCtx()

	agents, err := ctx.GrpcClient.GetAgents()
	if err != nil {
		return err
	}

	if len(agents) == 0 {
		fmt.Println("No agents found")
		return nil
	}

	now := time.Unix(time.Now().Unix(), 0).UTC()

	cols := []interface{}{"ID", "Name", "Active", "Group", "Token", "Scopes", "Created", "Last Connection"}
	rows := make([][]interface{}, len(agents))
	for _, agent := range agents {
		rows = append(rows, []interface{}{
			agent.Id,
			agent.Name,
			agent.IsActive,
			agent.Group,
			agent.Token,
			agent.Scopes,
			zt.TimeUntil(now, agent.CreatedAt.AsTime()),
			zt.TimeUntil(now, agent.LastConnection.AsTime()),
		})
	}

	utils.PrintTable(cols, rows)

	return nil
}

func init() {
	Cmd.AddCommand(agentsCmd)
}
