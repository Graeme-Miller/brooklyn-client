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

type StopPolicy struct {
	network *net.Network
}

func NewStopPolicy(network *net.Network) (cmd *StopPolicy) {
	cmd = new(StopPolicy)
	cmd.network = network
	return
}

func (cmd *StopPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stop-policy",
		Description: "Suspends a policy",
		Usage:       "BROOKLYN_NAME SCOPE stop-policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *StopPolicy) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	spec, err := entity_policies.StopPolicy(cmd.network, scope.Application, scope.Entity, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(spec)
}
