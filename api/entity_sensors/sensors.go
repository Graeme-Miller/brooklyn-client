package entity_sensors

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func SensorList(network *net.Network, application, entity string) []models.SensorSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors", application, entity)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var sensorList []models.SensorSummary
	err = json.Unmarshal(body, &sensorList)
	if err != nil{
		fmt.Println(err)
	}
	return sensorList
}

func SensorValue(network *net.Network, application, entity, sensor string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}