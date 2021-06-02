package cmd

import (
	"github.com/networkop/airctl/cmd/command/account"
	"github.com/networkop/airctl/cmd/command/cli"
	"github.com/networkop/airctl/cmd/command/service"
	"github.com/networkop/airctl/cmd/command/simulation"
	"github.com/networkop/airctl/cmd/command/topology"
	"github.com/networkop/airctl/cmd/command/training"
)

func GetResources(c *cli.Cli) map[string]*cli.Resource {
	resources := make(map[string]*cli.Resource)

	resources["sim"] = simulation.NewResource(c)
	resources["topo"] = topology.NewResource(c)
	resources["service"] = service.NewResource(c)
	resources["account"] = account.NewResource(c)
	resources["training"] = training.NewResource(c)

	return resources
}
