package gounity

import (
	"encoding/json"
	"errors"
	"gounity/api"
	"io"
	"net/http"
	"time"
)

type LunInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content LunContent `json:"content"`
	} `json:"entries"`
}

type LunTypeEnum int

const (
	LunTypeGenericStorage LunTypeEnum = iota + 1
	LunTypeStandalone
	LunTypeVmWareISCSI
)

func (_e LunTypeEnum) String() string {
	switch _e {
	case LunTypeGenericStorage:
		return "GenericStorage"
	case LunTypeStandalone:
		return "Standalone"
	case LunTypeVmWareISCSI:
		return "VmWareISCSI"
	default:
		return "Unknown"
	}
}

type LunContent struct {
	Id                     string       `json:"id,omitempty"`
	Health                 HealthStruct `json:"health,omitempty"`
	Name                   string       `json:"name,omitempty"`
	Description            string       `json:"description,omitempty"`
	Type                   LunTypeEnum  `json:"type,omitempty"`
	SizeTotal              int64        `json:"sizeTotal,omitempty"`
	SizeUsed               int64        `json:"sizeUsed,omitempty"`
	SizeAllocated          int64        `json:"SizeAllocated,omitempty"`
	SizePreallocated       int64        `json:"sizePreallocated,omitempty"`
	SizeAllocatedTotal     int64        `json:"SizeAllocatedTotal,omitempty"`
	DataReductionSizeSaved int64        `json:"dataReductionSizeSaved,omitempty"`
	DataReductionPercent   int64        `json:"dataReductionPercent,omitempty"`
	DataReductionRatio     float64      `json:"dataReductionRatio,omitempty"`
	PerTierSizeUsed        []int64      `json:"perTierSizeUsed,omitempty"`
	IsThinEnabled          bool         `json:"isThinEnabled,omitempty"`
	IsDataReductionEnabled bool         `json:"isDataReductionEnabled,omitempty"`
	IsAdvancedDedupEnabled bool         `json:"isAdvancedDedupEnabled,omitempty"`
	Wwn                    string       `json:"wwn,omitempty"`
}

func (_c *UnisphereClient) GetLunInstances(fields []string) (*LunInstances, error) {
	inst := api.UnityAPILunInstances
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

	var data *LunInstances
	err = json.Unmarshal(body, &data)
	return data, err

}
