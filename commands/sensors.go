package commands

import(
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_sensors"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Sensors struct {
	network *net.Network
}

func NewSensors(network *net.Network) (cmd *Sensors){
	cmd = new(Sensors)
	cmd.network = network
	return
}

func (cmd *Sensors) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "sensors",
		Description: "show the sensors for an application and entity",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Sensors) Run(c *cli.Context) {
	sensors := entity_sensors.SensorList(cmd.network, c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Name", "Description", "Value"})
	for _, sensor := range sensors {
		value := entity_sensors.SensorValue(cmd.network, c.Args()[0], c.Args()[1], sensor.Name)
		table.Add(sensor.Name, sensor.Description, value)
	}
	table.Print()
}