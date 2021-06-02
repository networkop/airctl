package cli

import (
	"io"
	"os"

	"github.com/networkop/airctl/pkg/air"
	"github.com/spf13/cobra"
)

type Cli struct {
	Air air.Client
	Out io.Writer
	Err io.Writer
}

type Resource struct {
	Name      string
	Getter    func() *cobra.Command
	Creater   func() *cobra.Command
	Destroyer func() *cobra.Command
	Setter    func() *cobra.Command
}

type CliOption func(cli *Cli) error

func NewCli(opts ...CliOption) (*Cli, error) {
	cli := Cli{
		Out: os.Stdout,
		Err: os.Stderr,
	}

	if err := cli.Apply(opts...); err != nil {
		return nil, err
	}

	rest, err := air.NewClient()
	if err != nil {
		return nil, err
	}

	cli.Air = *rest

	return &cli, nil
}

func (c *Cli) Apply(opts ...CliOption) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}
