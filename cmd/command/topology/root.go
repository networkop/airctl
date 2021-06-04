package topology

import (
	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/networkop/airctl/internal/utils"
	"github.com/spf13/cobra"
)

func NewResource(c *cli.Cli) *cli.Resource {

	resource := &cli.Resource{
		Name: "topology",
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
		Use:     "topo",
		Aliases: []string{"topology"},
		Short:   "Get topologies",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return utils.ProcessError(c.Air.ListTopologies())
			}

			return utils.ProcessError(c.Air.GetTopology(args[0], quiet))
		},
	}
	cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Only output UUIDs")
	return cmd
}

func newCreateCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "topo FILE",
		Aliases: []string{"topology"},
		Args:    cobra.ExactValidArgs(1),
		Short:   "Create a topology from dot file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ProcessError(c.Air.CreateTopology(args[0]))
		},
	}
	return cmd
}

func newDelCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "topo <ID...>",
		Aliases: []string{"topology"},
		Args:    cobra.MinimumNArgs(1),
		Short:   "Delete topology by name or ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.ProcessError(c.Air.DelTopology(args))
		},
	}
	return cmd
}
