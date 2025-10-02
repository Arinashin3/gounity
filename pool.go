package gounity

import (
	"encoding/json"
	"time"

	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
	"github.com/Arinashin3/gounity/units"
)

type PoolInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content PoolContent `json:"content"`
	} `json:"entries"`
}

type PoolTypeEnum int

const (
	PoolTypeGenericStorage PoolTypeEnum = iota + 1
	PoolTypeStandalone
	PoolTypeVmWareISCSI
)

func (_e PoolTypeEnum) String() string {
	switch _e {
	case PoolTypeGenericStorage:
		return "GenericStorage"
	case PoolTypeStandalone:
		return "Standalone"
	case PoolTypeVmWareISCSI:
		return "VmWareISCSI"
	default:
		return "Unknown"
	}
}

type PoolContent struct {
	Id                          string              `json:"id,omitempty"`
	Health                      types.HealthContent `json:"health,omitempty"`
	Name                        string              `json:"name,omitempty"`
	Description                 string              `json:"description,omitempty"`
	RaidType                    types.RaidTypeEnum  `json:"raidType,omitempty"`
	SizeFree                    units.Bytes         `json:"sizeFree,omitempty"`
	SizeTotal                   units.Bytes         `json:"sizeTotal,omitempty"`
	SizeUsed                    units.Bytes         `json:"sizeUsed,omitempty"`
	SizePreallocated            units.Bytes         `json:"sizePreallocated,omitempty"`
	DataReductionSizeSaved      units.Bytes         `json:"dataReductionSizeSaved,omitempty"`
	DataReductionPercent        units.Bytes         `json:"dataReductionPercent,omitempty"`
	DataReductionRatio          float64             `json:"dataReductionRatio,omitempty"`
	FlashPercentage             units.Percentage    `json:"flashPercentage,omitempty"`
	SizeSubscribed              units.Bytes         `json:"sizeSubscribed,omitempty"`
	AlertThreshold              units.Percentage    `json:"alertThreshold,omitempty"`
	HasDataReductionEnabledLuns bool                `json:"hasDataReductionEnabledLuns,omitempty"`
	IsFASTCacheEnabled          bool                `json:"isFASTCacheEnabled,omitempty"`
	// Tiers PoolTiers
	CreationTime time.Time `json:"creationTime,omitempty"`
	IsEmpty      bool      `json:"isEmpty,omitempty"`
	//PoolFastVP
	IsHarvestEnabled bool `json:"isHarvestEnabled,omitempty"`
	//HarvestState UsageHarvestStateEnum `json:"harvestState,omitempty"`
	IsSnapHarvestEnabled          bool                      `json:"isSnapHarvestEnabled,omitempty"`
	PoolSpaceHarvestHighThreshold units.Percentage          `json:"poolSpaceHarvestHighThreshold,omitempty"`
	PoolSpaceHarvestLowThreshold  units.Percentage          `json:"poolSpaceHarvestLowThreshold,omitempty"`
	SnapSpaceHarvestHighThreshold units.Percentage          `json:"snapSpaceHarvestHighThreshold,omitempty"`
	SnapSpaceHarvestLowThreshold  units.Percentage          `json:"snapSpaceHarvestLowThreshold,omitempty"`
	MetadataSizeSubscribed        units.Bytes               `json:"metadataSizeSubscribed,omitempty"`
	SnapSizeSubscribed            units.Bytes               `json:"snapSizeSubscribed,omitempty"`
	NonBaseSizeSubscribed         units.Bytes               `json:"nonBaseSizeSubscribed,omitempty"`
	MetadataSizeUsed              units.Bytes               `json:"metadataSizeUsed,omitempty"`
	SnapSizeUsed                  units.Bytes               `json:"snapSizeUsed,omitempty"`
	NonBaseSizeUsed               units.Bytes               `json:"nonBaseSizeUsed,omitempty"`
	RebalanceProgress             units.Percentage          `json:"rebalanceProgress,omitempty"`
	Type                          types.StoragePoolTypeEnum `json:"type"`
	IsAllFlash                    bool                      `json:"isAllFlash,omitempty"`
}

func (_c *UnisphereClient) GetPoolInstances(fields []string, filter []string) (*PoolInstances, error) {
	req, err := api.UnityAPIPoolInstances.NewRequest(_c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	api.UnityAPIPoolInstances.WithFields(fields, req)
	api.UnityAPIPoolInstances.WithFilter(filter, req)
	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)
	if err != nil {
		return nil, err
	}

	var data PoolInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
