package gounity

import (
	"encoding/json"
	"errors"
	"gounity/api"
	"io"
	"net/http"
	"time"
)

type SystemInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content SystemContent `json:"content"`
	} `json:"entries"`
}

type SystemContent struct {
	Id                           string            `json:"id,omitempty"`
	Health                       HealthStruct      `json:"health,omitempty"`
	Name                         string            `json:"name,omitempty"`
	Model                        string            `json:"model,omitempty"`
	SerialNumber                 string            `json:"serialNumber,omitempty"`
	UuidBase                     int32             `json:"uuidBase,omitempty"`
	InternalModel                string            `json:"internalModel,omitempty"`
	Platform                     string            `json:"platform,omitempty"`
	IsAllFlash                   bool              `json:"isAllFlash,omitempty"`
	MacAddress                   string            `json:"macAddress,omitempty"`
	IsEULAAccount                bool              `json:"isEULAAccount,omitempty"`
	IsUpgradeComplete            bool              `json:"isUpgradeComplete,omitempty"`
	isAutoFailbackEnabled        bool              `json:"isAutoFailbackEnabled,omitempty"`
	CurrentPower                 int32             `json:"currentPower,omitempty"`
	AvgPower                     int32             `json:"avgPower,omitempty"`
	SupportedUpgradeModels       []SPModelNameEnum `json:"supportedUpgradeModels,omitempty"`
	IsRemoteSysInterfaceAutoPair bool              `json:"isRemoteSysInterfaceAutoPair,omitempty"`
}

func (_c *UnisphereClient) GetSystemInstances(fields []string) (*SystemInstances, error) {
	inst := api.UnityAPISystemInstances
	req, err := inst.NewRequest(_c.endpoint)
	inst.WithFields(fields, req)
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

	var data *SystemInstances
	err = json.Unmarshal(body, &data)
	return data, err

}
