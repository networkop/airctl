package account

import (
	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/networkop/airctl/internal/utils"
	"github.com/spf13/cobra"
)

func NewResource(c *cli.Cli) *cli.Resource {

	resource := &cli.Resource{
		Name: "account",
		Getter: func() *cobra.Command {
			return newGetCommand(c)
		},
		Setter: func() *cobra.Command {
			return nil
		},
		Creater: func() *cobra.Command {
			return nil
		},
		Destroyer: func() *cobra.Command {
			return nil
		},
	}

	return resource
}

func newGetCommand(c *cli.Cli) *cobra.Command {
	var quiet bool
	cmd := &cobra.Command{
		Use:     "account",
		Aliases: []string{"acc", "accs", "accounts"},
		Short:   "Get accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return utils.ProcessError(c.Air.ListAccounts(quiet))
			}

			return utils.ProcessError(c.Air.GetAccount(args[0], quiet))
		},
	}

	cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Only output UUIDs")
	return cmd
}
