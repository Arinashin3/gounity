package gounity

import (
	"encoding/json"
	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
	"github.com/Arinashin3/gounity/units"
	"time"
)

type SystemCapacityInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content SystemCapacityContent `json:"content"`
	} `json:"entries"`
}

type SystemCapacityContent struct {
	Id                     string                      `json:"id,omitempty"`
	SizeFree               units.Bytes                 `json:"sizeFree,omitempty"`
	SizeTotal              units.Bytes                 `json:"sizeTotal,omitempty"`
	SizeUsed               units.Bytes                 `json:"sizeUsed,omitempty"`
	SizePreallocated       units.Bytes                 `json:"sizePreallocated,omitempty"`
	DataReductionSizeSaved units.Bytes                 `json:"dataReductionSizeSaved,omitempty"`
	DataReductionPercent   units.Percentage            `json:"dataReductionPercent,omitempty"`
	DataReductionRatio     units.Ratio                 `json:"dataReductionRatio,omitempty"`
	SizeSubscribed         units.Bytes                 `json:"sizeSubscribed,omitempty"`
	TotalLogicalSize       units.Bytes                 `json:"totalLogicalSize,omitempty"`
	ThinSavingRatio        units.Ratio                 `json:"thinSavingRatio,omitempty"`
	SnapsSavingsRatio      units.Ratio                 `json:"snapsSavingsRatio,omitempty"`
	OverallEfficiencyRatio units.Ratio                 `json:"overallEfficiencyRatio,omitempty"`
	Tiers                  []SystemTierCapacityContent `json:"tiers,omitempty"`
}

// SystemTierCapacityContent
// Capacity data for one tier, collected from all pools.
type SystemTierCapacityContent struct {
	TierType  types.TierTypeEnum `json:"tierType,omitempty"`
	SizeFree  units.Bytes        `json:"sizeFree,omitempty"`
	SizeTotal units.Bytes        `json:"sizeTotal,omitempty"`
	SizeUsed  units.Bytes        `json:"sizeUsed,omitempty"`
}

func (_c *UnisphereClient) GetSystemCapacityInstances(fields []string, filter string) (*SystemCapacityInstances, error) {
	req, err := api.UnityAPISystemCapacityInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	api.UnityAPISystemCapacityInstances.WithFields(fields, req)
	api.UnityAPISystemCapacityInstances.WithFilter(filter, req)
	_c.addHeader("GET", req)

	var body []byte
	body, err = _c.send(req)

	var data SystemCapacityInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
