package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/application"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/models"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type Tree struct {
	network *net.Network
}

func NewTree(network *net.Network) (cmd *Tree) {
	cmd = new(Tree)
	cmd.network = network
	return
}

func (cmd *Tree) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "tree",
		Description: "* Show the tree of all applications",
		Usage:       "BROOKLYN_NAME tree",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Tree) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	trees, err := application.Tree(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	cmd.printTrees(trees, "")
}

func (cmd *Tree) printTrees(trees []models.Tree, indent string) {
	for i, app := range trees {
		cmd.printTree(app, indent, i == len(trees)-1)
	}
}

func (cmd *Tree) printTree(tree models.Tree, indent string, last bool) {
	fmt.Println(indent+"|-", tree.Name)
	fmt.Println(indent+"+-", tree.Type)

	if last {
		indent = indent + "  "
	} else {
		indent = indent + "| "
	}
	cmd.printTrees(tree.Children, indent)
}
