package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/entity_policies"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type StartPolicy struct {
	network *net.Network
}

func NewStartPolicy(network *net.Network) (cmd *StartPolicy) {
	cmd = new(StartPolicy)
	cmd.network = network
	return
}

func (cmd *StartPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "start-policy",
		Description: "Start or resume a policy",
		Usage:       "BROOKLYN_NAME SCOPE start-policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *StartPolicy) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	spec, err := entity_policies.StartPolicy(cmd.network, scope.Application, scope.Entity, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(spec)
}
