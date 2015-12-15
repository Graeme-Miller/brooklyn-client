package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type Rename struct {
	network *net.Network
}

func NewRename(network *net.Network) (cmd *Rename) {
	cmd = new(Rename)
	cmd.network = network
	return
}

func (cmd *Rename) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "rename",
		Description: "Rename an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE rename NEW_NAME",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Rename) Run(scope scope.Scope, c *cli.Context) {
	rename, err := entities.Rename(cmd.network, scope.Application, scope.Entity, c.Args().First())
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(rename)
}
