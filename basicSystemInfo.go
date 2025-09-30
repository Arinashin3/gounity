package gounity

import (
	"encoding/json"
	"errors"
	"gounity/api"
	"io"
	"net/http"
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

func (_c *UnisphereClient) GetBasicSystemInfoInstances() (*BasicSystemInfoInstances, error) {
	req, err := api.UnityAPIBasicSystemInfoInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	_c.addHeader("GET", req)

	resp, err := _c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		_c.logined = false
	case http.StatusUnprocessableEntity:
		var data StatusUnprocessableEntity
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(data.Error.Messages[0].EnUS)
	}

	var data *BasicSystemInfoInstances
	err = json.Unmarshal(body, &data)
	return data, err

}
