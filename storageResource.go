package gounity

import (
	"encoding/json"
	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
	"github.com/Arinashin3/gounity/units"
	"time"
)

type StorageResourceInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content StorageResourceContent `json:"content"`
	} `json:"entries"`
}

type StorageResourceContent struct {
	Id                       string                        `json:"id,omitempty"`
	Health                   types.HealthContent           `json:"health,omitempty"`
	Name                     string                        `json:"name,omitempty"`
	Description              string                        `json:"description,omitempty"`
	Type                     types.StorageResourceTypeEnum `json:"type,omitempty"`
	IsReplicationDestination bool                          `json:"isReplicationDestination,omitempty"`
	ReplicationType          types.ReplicationTypeEnum     `json:"replicationType,omitempty"`
	SyncReplicationType      types.ReplicationTypeEnum     `json:"syncReplicationType,omitempty"`
	SizeTotal                units.Bytes                   `json:"sizeTotal,omitempty"`
	SizeUsed                 units.Bytes                   `json:"sizeUsed,omitempty"`
	SizeAllocated            units.Bytes                   `json:"SizeAllocated,omitempty"`
	SizePreallocated         units.Bytes                   `json:"sizePreallocated,omitempty"`
	SizeAllocatedTotal       units.Bytes                   `json:"SizeAllocatedTotal,omitempty"`
	ThinStatus               types.ThinStatusEnum          `json:"thinStatus,omitempty"`
	DataReductionSizeSaved   units.Bytes                   `json:"dataReductionSizeSaved,omitempty"`
	DataReductionPercent     units.Percentage              `json:"dataReductionPercent,omitempty"`
	DataReductionRatio       units.Ratio                   `json:"dataReductionRatio,omitempty"`
	PerTierSizeUsed          []int64                       `json:"perTierSizeUsed,omitempty"`
	IsThinEnabled            bool                          `json:"isThinEnabled,omitempty"`
	IsDataReductionEnabled   bool                          `json:"isDataReductionEnabled,omitempty"`
	IsAdvancedDedupEnabled   bool                          `json:"isAdvancedDedupEnabled,omitempty"`
	Wwn                      string                        `json:"wwn,omitempty"`
}

func (_c *UnisphereClient) GetStorageResourceInstances(fields []string, filter string) (*StorageResourceInstances, error) {
	req, err := api.UnityAPIStorageResourceInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	api.UnityAPIStorageResourceInstances.WithFields(fields, req)
	api.UnityAPIStorageResourceInstances.WithFilter(filter, req)
	_c.addHeader("GET", req)

	var body []byte
	body, err = _c.send(req)

	var data StorageResourceInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
