package cmd

import (
	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/networkop/airctl/internal/utils"
	"github.com/spf13/cobra"
)

func NewLoginCommand(c *cli.Cli) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "login <token>",
		Short: "Authenticate with air.nvidia.com",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			return utils.ProcessError(c.Air.Login(args[0]))
		},
	}
	return cmd
}
