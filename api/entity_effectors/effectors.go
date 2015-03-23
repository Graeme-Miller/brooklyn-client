package entity_effectors

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func EffectorList(network *net.Network, application, entity string) []models.EffectorSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors", application, entity)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var effectorList []models.EffectorSummary
	err = json.Unmarshal(body, &effectorList)
	if err != nil{
		fmt.Println(err)
	}
	return effectorList
}