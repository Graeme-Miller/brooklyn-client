package locations

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func LocationList(network *net.Network) []models.LocationSummary {
	url := "/v1/locations"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var locationList []models.LocationSummary
	err = json.Unmarshal(body, &locationList)
	if err != nil {
		fmt.Println(err)
	}
	return locationList
}
