package service

import (
	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/networkop/airctl/internal/utils"
	"github.com/spf13/cobra"
)

func NewResource(c *cli.Cli) *cli.Resource {

	resource := &cli.Resource{
		Name: "service",
		Getter: func() *cobra.Command {
			return newGetCommand(c)
		},
		Setter: func() *cobra.Command {
			return nil
		},
		Creater: func() *cobra.Command {
			return newCreateCommand(c)
		},
		Destroyer: func() *cobra.Command {
			return newDelCommand(c)
		},
	}

	return resource
}

func newGetCommand(c *cli.Cli) *cobra.Command {
	var quiet bool

	cmd := &cobra.Command{
		Use:     "svc",
		Aliases: []string{"service"},
		Short:   "Get services",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return utils.ProcessError(c.Air.ListServices(quiet))
			}

			return utils.ProcessError(c.Air.GetService(args[0], quiet))
		},
	}

	cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Only output UUIDs")
	return cmd
}

func newCreateCommand(c *cli.Cli) *cobra.Command {
	var name, simID string
	cmd := &cobra.Command{
		Use:     "svc <oob-mgmt-server:eth0>",
		Aliases: []string{"service"},
		Short:   "Create SSH service",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ProcessError(c.Air.CreateSSHService(args[0], simID, name))
		},
	}

	cmd.PersistentFlags().StringVarP(&simID, "sim", "s", "", "simulation ID")
	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "service name")
	return cmd
}

func newDelCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "svc <ID...>",
		Aliases: []string{"service"},
		Args:    cobra.MinimumNArgs(1),
		Short:   "Delete service by name or ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ProcessError(c.Air.DelService(args))
		},
	}
	return cmd
}
