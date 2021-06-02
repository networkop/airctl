package cmd

import (
	"io"

	"github.com/networkop/airctl/pkg/air"
	"github.com/spf13/cobra"
)

func NewLoginCommand(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "login <token>",
		Short: "Authenticate with air.nvidia.com",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			token := args[0]
			return air.Login(token)
		},
	}
	return cmd
}
