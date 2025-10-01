package gounity

import (
	"encoding/json"
	"time"
)

type EventInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content EventContent `json:"content"`
	} `json:"entries"`
}

type EventContent struct {
	Id           string                  `json:"id,omitempty"`
	Node         types.NodeEnum          `json:"node,omitempty"`
	CreationTime time.Time               `json:"creationTime,omitempty"`
	Severity     types.SeverityEnum      `json:"severity,omitempty"`
	MessageId    string                  `json:"messageId"`
	Arguments    []string                `json:"arguments,omitempty"`
	Message      string                  `json:"message,omitempty"`
	Username     string                  `json:"username,omitempty"`
	Category     types.EventCategoryEnum `json:"category,omitempty"`
	Source       string                  `json:"source,omitempty"`
}

func (_c *UnisphereClient) GetEventInstances(fields []string, filter string) (*EventInstances, error) {
	req, err := api.UnityAPIEventInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	api.UnityAPIEventInstances.WithFields(fields, req)
	api.UnityAPIEventInstances.WithFilter(filter, req)
	_c.addHeader("GET", req)

	var body []byte
	body, err = _c.send(req)

	var data EventInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
