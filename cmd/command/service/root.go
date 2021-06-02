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
			return newSetCommand(c)
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
	cmd := &cobra.Command{
		Use:     "svc",
		Aliases: []string{"service"},
		Short:   "Get services",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return utils.ProcessError(c.Air.ListServices())
			}

			return utils.ProcessError(c.Air.GetService(args[0]))
		},
	}
	return cmd
}

func newSetCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "inspect <ID>",
		Aliases: []string{"inspect"},

		Short: "Inspect simulation by its ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ProcessError(c.Air.GetSimulation(args[0]))
		},
	}
	return cmd
}

func newCreateCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "inspect <ID>",
		Aliases: []string{"inspect"},

		Short: "Inspect simulation by its ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ProcessError(c.Air.GetSimulation(args[0]))
		},
	}
	return cmd
}

func newDelCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "inspect <ID>",
		Aliases: []string{"inspect"},

		Short: "Inspect simulation by its ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ProcessError(c.Air.GetSimulation(args[0]))
		},
	}
	return cmd
}