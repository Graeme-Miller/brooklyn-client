package entity_policy_config

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/models"
	"github.com/apache/brooklyn-client/net"
)

func CurrentState(network *net.Network, application, entity, policy string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/current-state", application, entity, policy)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func GetConfigValue(network *net.Network, application, entity, policy, config string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/%s", application, entity, policy, config)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

// WIP
func SetConfigValue(network *net.Network, application, entity, policy, config string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/%s", application, entity, policy, config)
	body, err := network.SendEmptyPostRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func GetAllConfigValues(network *net.Network, application, entity, policy string) ([]models.PolicyConfigList, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config", application, entity, policy)
	var policyConfigList []models.PolicyConfigList
	body, err := network.SendGetRequest(url)
	if nil != err {
		return policyConfigList, err
	}
	err = json.Unmarshal(body, &policyConfigList)
	return policyConfigList, err
}
