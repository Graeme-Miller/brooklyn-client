package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_effectors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"strings"
)

type Effectors struct {
	network *net.Network
}

func NewListEffector(network *net.Network) (cmd *Effectors) {
	cmd = new(Effectors)
	cmd.network = network
	return
}

func (cmd *Effectors) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "effector",
		Description: "Show the list of effectors for an application and entity",
		Usage:       "list effector APPLICATION ENTITY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Effectors) Run(c *cli.Context) {
	effectors := entity_effectors.EffectorList(cmd.network, c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Name", "Description", "Parameters"})
	for _, effector := range effectors {
		var parameters []string
		for _, parameter := range effector.Parameters {
			parameters = append(parameters, parameter.Name)
		}
		table.Add(effector.Name, effector.Description, strings.Join(parameters, ","))
	}
	table.Print()
}
