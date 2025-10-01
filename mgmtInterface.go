package gounity

import (
	"encoding/json"
	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
	"time"
)

type MgmtInterfaceInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content MgmtInterfaceContent `json:"content"`
	} `json:"entries"`
}

type MgmtInterfaceContent struct {
	Id         string                        `json:"id,omitempty"`
	ConfigMode types.InterfaceConfigModeEnum `json:"configMode,omitempty"`
	//EthernetPort EthernetPortContent `json:"ethernetPort,omitempty"`
	ProtocolVersion types.IpProtocolVersionEnum `json:"protocolVersion,omitempty"`
	IpAddress       string                      `json:"ipAddress,omitempty"`
	Netmask         string                      `json:"netmask,omitempty"`
	V6PrefixLength  int32                       `json:"v6PrefixLength,omitempty"`
	Gateway         string                      `json:"gateway,omitempty"`
}

func (_c *UnisphereClient) GetMgmtInterfaceInstances(fields []string, filter string) (*MgmtInterfaceInstances, error) {
	req, err := api.UnityAPIMgmtInterfaceInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	api.UnityAPIMgmtInterfaceInstances.WithFields(fields, req)
	api.UnityAPIMgmtInterfaceInstances.WithFilter(filter, req)
	_c.addHeader("GET", req)

	var body []byte
	body, err = _c.send(req)

	var data MgmtInterfaceInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
