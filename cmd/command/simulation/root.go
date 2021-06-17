package simulation

import (
	"fmt"

	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/networkop/airctl/internal/utils"
	"github.com/networkop/airctl/pkg/air"
	"github.com/spf13/cobra"
)

func NewResource(c *cli.Cli) *cli.Resource {

	resource := &cli.Resource{
		Name: "simulation",
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
	var quiet bool
	cmd := &cobra.Command{
		Use:     "sim ( ID | Name )",
		Aliases: []string{"sims", "simulation", "simulations"},
		Short:   "Get simulations",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return utils.ProcessError(c.Air.ListSimulations(quiet))
			}

			return utils.ProcessError(c.Air.GetSimulation(args[0], quiet))
		},
	}

	cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Only output UUIDs")
	return cmd
}

func newSetCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sim ( ID | Name ) up | down",
		Aliases: []string{"simulation"},
		Short:   "Set the state of a simulation",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("Expecting 2 arguments")
			}

			if args[1] == "up" {
				return utils.ProcessError(c.Air.SetSimulation(args[0], air.SimulationState.Up))
			}

			if args[1] == "down" {
				return utils.ProcessError(c.Air.SetSimulation(args[0], air.SimulationState.Down))
			}

			return fmt.Errorf("Second argument must be 'up' or 'down', got '%s'", args[1])

		},
	}
	return cmd
}

func newCreateCommand(c *cli.Cli) *cobra.Command {
	var simID, topo string
	var citc bool
	cmd := &cobra.Command{
		Use:     "sim",
		Aliases: []string{"simulation"},
		//Args:  cobra.ExactValidArgs(1),
		Short: "Create simulation from topology or another sim",
		RunE: func(cmd *cobra.Command, args []string) error {
			if citc {
				return utils.ProcessError(c.Air.CreateCITC())
			}
			if simID != "" {
				return utils.ProcessError(c.Air.CloneSimulation(simID))
			}
			if topo != "" {
				return utils.ProcessError(c.Air.CreateSimulation(topo))
			}
			return fmt.Errorf("Either simID or topoID must be provided")
		},
	}
	cmd.PersistentFlags().StringVarP(&simID, "sim", "s", "", "ID of an existing sim to clone")
	cmd.PersistentFlags().StringVarP(&topo, "topo", "t", "", "ID of an existing topo to create")
	cmd.PersistentFlags().BoolVarP(&citc, "citc", "c", false, "Create a citc sim")
	return cmd
}

func newDelCommand(c *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sim <ID...>",
		Aliases: []string{"simulation"},
		Args:    cobra.MinimumNArgs(1),
		Short:   "Delete simulation",
		RunE: func(cmd *cobra.Command, args []string) error {

			return utils.ProcessError(c.Air.DeleteSim(args))

		},
	}

	return cmd
}
