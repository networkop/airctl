package training

import (
	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/networkop/airctl/internal/utils"
	"github.com/spf13/cobra"
)

func NewResource(c *cli.Cli) *cli.Resource {

	resource := &cli.Resource{
		Name: "training",
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

	cmd := &cobra.Command{
		Use:     "training",
		Aliases: []string{"trainings"},
		Short:   "Get trainings",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return utils.ProcessError(c.Air.GetTrainings())
		},
	}
	return cmd
}
