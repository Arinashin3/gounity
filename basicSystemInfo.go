package gounity

import (
	"encoding/json"
	"gounity/api"
	"time"
)

type BasicSystemInfoInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content BasicSystemInfoContent `json:"content"`
	} `json:"entries"`
}

type BasicSystemInfoContent struct {
	Id                  string `json:"id,omitempty"`
	Model               string `json:"model,omitempty"`
	Name                string `json:"name,omitempty"`
	SoftwareVersion     string `json:"softwareVersion,omitempty"`
	SoftwareFullVersion string `json:"softwareFullVersion,omitempty"`
	ApiVersion          string `json:"apiVersion,omitempty"`
	EarliestApiVersion  string `json:"earliestApiVersion,omitempty"`
}

// GetBasicSystemInfoInstances
// Field를 추가하지 않아도 기본적으로 모든 필드를 제공한다.
func (_c *UnisphereClient) GetBasicSystemInfoInstances() (*BasicSystemInfoInstances, error) {
	req, err := api.UnityAPIBasicSystemInfoInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	_c.addHeader("GET", req)

	var body []byte
	body, err = _c.send(req)

	var data BasicSystemInfoInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
