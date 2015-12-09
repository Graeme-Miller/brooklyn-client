package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/codegangsta/cli"
	"fmt"
	"strings"
	"github.com/brooklyncentral/brooklyn-cli/command"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type List struct {
	network *net.Network
    listCommands map[string]command.Command
}

func NewList(network *net.Network) (cmd *List) {
	cmd = new(List)
	cmd.network = network
	cmd.listCommands = map[string]command.Command{
//		ListApplicationCommand: NewApplications(cmd.network),
//		ListEntityCommand: NewEntities(cmd.network),
//		ListSensorCommand: NewSensors(cmd.network),
//		ListEffectorCommand: NewEffectors(cmd.network),
	}
	return
}

const ListApplicationCommand = "application"
const ListEntityCommand = "entities"
const ListSensorCommand = "sensors"
const ListEffectorCommand = "effectors"

var listCommands = []string {
	ListApplicationCommand,
	ListEntityCommand,
	ListSensorCommand,
	ListEffectorCommand,
}
var listCommandsUsage = strings.Join(listCommands, " | ")

func (cmd *List) SubCommandNames() ([]string) {
	return listCommands
}

func (cmd *List) SubCommand(name string) command.Command {
	return cmd.listCommands[name]
}

func (cmd *List) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "list",
		Description: "List details for a variety of operands",
		Usage:       "BROOKLYN_NAME [ SCOPE ] list (" + listCommandsUsage + ")",
		Flags:       []cli.Flag{},
		Operands:    []command_metadata.CommandMetadata {
			cmd.SubCommand(ListApplicationCommand).Metadata(),
			cmd.SubCommand(ListEntityCommand).Metadata(),
			cmd.SubCommand(ListSensorCommand).Metadata(),
			cmd.SubCommand(ListEffectorCommand).Metadata(),
		},
	}
}

func (cmd *List) Run(scope scope.Scope, c *cli.Context) {
	fmt.Printf( "Unrecognised item for list, please use one of (%s)\n", listCommandsUsage)
}
