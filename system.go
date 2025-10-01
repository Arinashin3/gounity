package gounity

import (
	"encoding/json"
	"gounity/api"
	"gounity/types"
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
	Id                           string                  `json:"id,omitempty"`
	Health                       types.HealthContent     `json:"health,omitempty"`
	Name                         string                  `json:"name,omitempty"`
	Model                        string                  `json:"model,omitempty"`
	SerialNumber                 string                  `json:"serialNumber,omitempty"`
	UuidBase                     int32                   `json:"uuidBase,omitempty"`
	InternalModel                string                  `json:"internalModel,omitempty"`
	Platform                     string                  `json:"platform,omitempty"`
	IsAllFlash                   bool                    `json:"isAllFlash,omitempty"`
	MacAddress                   string                  `json:"macAddress,omitempty"`
	IsEULAAccount                bool                    `json:"isEULAAccount,omitempty"`
	IsUpgradeComplete            bool                    `json:"isUpgradeComplete,omitempty"`
	IsAutoFailbackEnabled        bool                    `json:"isAutoFailbackEnabled,omitempty"`
	CurrentPower                 int32                   `json:"currentPower,omitempty"`
	AvgPower                     int32                   `json:"avgPower,omitempty"`
	SupportedUpgradeModels       []types.SPModelNameEnum `json:"supportedUpgradeModels,omitempty"`
	IsRemoteSysInterfaceAutoPair bool                    `json:"isRemoteSysInterfaceAutoPair,omitempty"`
}

func (_c *UnisphereClient) GetSystemInstances(fields []string, filter string) (*SystemInstances, error) {
	req, err := api.UnityAPISystemInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	api.UnityAPISystemInstances.WithFields(fields, req)
	api.UnityAPISystemInstances.WithFilter(filter, req)
	_c.addHeader("GET", req)

	var body []byte
	body, err = _c.send(req)
	if err != nil {
		return nil, err
	}

	var data SystemInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
