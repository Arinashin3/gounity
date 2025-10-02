package gounity

import (
	"encoding/json"
	"time"

	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
)

type AlertInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content AlertContent `json:"content"`
	} `json:"entries"`
}

type AlertContent struct {
	Id             string                   `json:"id,omitempty"`
	Timestamp      time.Time                `json:"timestamp,omitempty"`
	Severity       types.SeverityEnum       `json:"severity,omitempty"`
	Component      types.ResourceRefContent `json:"component,omitempty"`
	MessageId      string                   `json:"messageId"`
	Message        string                   `json:"message,omitempty"`
	DescriptionId  string                   `json:"descriptionId"`
	Description    string                   `json:"description"`
	ResolutionId   string                   `json:"resolutionId"`
	Resolution     string                   `json:"resolution,omitempty"`
	IsAcknowledged bool                     `json:"isAcknowledged"`
	DuplicateCount int                      `json:"duplicateCount"`
	State          types.AlertStateEnum     `json:"state"`
}

func (_c *UnisphereClient) GetAlertInstances(fields []string, filter []string) (*AlertInstances, error) {
	req, err := api.UnityAPIAlertInstances.NewRequest(_c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	api.UnityAPIAlertInstances.WithFields(fields, req)
	api.UnityAPIAlertInstances.WithFilter(filter, req)
	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)
	if err != nil {
		return nil, err
	}

	var data AlertInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
