package entity_effectors

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/brooklyn-client/models"
	"github.com/apache/brooklyn-client/net"
	"net/url"
	"strconv"
	"strings"
)

func EffectorList(network *net.Network, application, entity string) ([]models.EffectorSummary, error) {
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors", application, entity)
	var effectorList []models.EffectorSummary
	body, err := network.SendGetRequest(path)
	if err != nil {
		return effectorList, err
	}

	err = json.Unmarshal(body, &effectorList)
	return effectorList, err
}

func TriggerEffector(network *net.Network, application, entity, effector string, params []string, args []string) (string, error) {
	if len(params) != len(args) {
		return "", errors.New(strings.Join([]string{"Parameters not supplied:", strings.Join(params, ", ")}, " "))
	}
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors/%s", application, entity, effector)
	data := url.Values{}
	for i := range params {
		data.Set(params[i], args[i])
	}
	req := network.NewPostRequest(path, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	body, err := network.SendRequest(req)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
