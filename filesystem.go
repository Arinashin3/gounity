package gounity

import (
	"encoding/json"
	"time"

	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
)

type FilesystemInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content FilesystemContent `json:"content"`
	} `json:"entries"`
}

type FilesystemContent struct {
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

func (_c *UnisphereClient) GetFilesystemInstances(fields []string, filter []string) (*SystemInstances, error) {
	req, err := api.UnityAPIFilesystemInstances.NewRequest(_c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	api.UnityAPIFilesystemInstances.WithFields(fields, req)
	api.UnityAPIFilesystemInstances.WithFilter(filter, req)
	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)
	if err != nil {
		return nil, err
	}

	var data FilesystemInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
