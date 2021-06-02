package cmd

import (
	"fmt"
	"os"

	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var debug bool

func Execute() error {

	logrus.SetOutput(os.Stdout)

	cli, err := cli.NewCli()
	if err != nil {
		logrus.Infof("Error initializing CLI: %s", err)
		os.Exit(1)
	}

	root := &cobra.Command{
		Use:   "airctl [command]",
		Short: "unofficial CLI client for air.nvidia.com",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				logrus.SetLevel(logrus.DebugLevel)
			}

		},
		Version: fmt.Sprintf("alpha"),
	}

	cobra.EnableCommandSorting = false
	addGlobalFlags(root.PersistentFlags())

	root.AddCommand(
		NewGetCommand(cli),
		NewSetCommand(cli),
		NewCreateCommand(cli),
		NewDeleteCommand(cli),
		NewLoginCommand(os.Stdout),
	)

	return root.Execute()
}

func addGlobalFlags(fs *pflag.FlagSet) {
	fs.BoolVarP(&debug, "debug", "d", debug, "Enable debug-level logging")
}

func NewGetCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Display one or many resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	for _, r := range GetResources(cli) {
		if r.Getter() != nil {
			cmd.AddCommand(r.Getter())

		}
	}

	return cmd
}

func NewSetCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set",
		Aliases: []string{"s"},
		Short:   "Update existing resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	for _, r := range GetResources(cli) {
		if r.Setter() != nil {
			cmd.AddCommand(r.Setter())

		}
	}

	return cmd
}

func NewCreateCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a resource from a file",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	for _, r := range GetResources(cli) {
		if r.Creater() != nil {
			cmd.AddCommand(r.Creater())
		}
	}

	return cmd
}

func NewDeleteCommand(cli *cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete resources by name or ID",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	for _, r := range GetResources(cli) {
		if r.Destroyer() != nil {
			cmd.AddCommand(r.Destroyer())
		}
	}

	return cmd
}
