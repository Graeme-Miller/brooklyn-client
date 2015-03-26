package activities

import (
	"encoding/json"
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/models"
	"github.com/robertgmoss/brooklyn-cli/net"
)

func ActivityList(network *net.Network, application, entity string) []models.TaskSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var activityList []models.TaskSummary
	err = json.Unmarshal(body, &activityList)
	if err != nil {
		fmt.Println(err)
	}
	return activityList
}

func Activity(network *net.Network, activity string) models.TaskSummary {
	url := fmt.Sprintf("/v1/activities/%s", activity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var task models.TaskSummary
	err = json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
	}
	return task
}

func ActivityChildren(network *net.Network, activity string) []models.TaskSummary {
	url := fmt.Sprintf("/v1/activities/%s/children", activity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var tasks []models.TaskSummary
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		fmt.Println(err)
	}
	return tasks
}
