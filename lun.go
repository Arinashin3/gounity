package gounity

import (
	"encoding/json"
	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
	"github.com/Arinashin3/gounity/units"
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
	Id                       string                  `json:"id,omitempty"`
	Health                   types.HealthContent     `json:"health,omitempty"`
	Name                     string                  `json:"name,omitempty"`
	Description              string                  `json:"description,omitempty"`
	Type                     LunTypeEnum             `json:"type,omitempty"`
	SizeTotal                units.Bytes             `json:"sizeTotal,omitempty"`
	SizeUsed                 units.Bytes             `json:"sizeUsed,omitempty"`
	SizeAllocated            units.Bytes             `json:"SizeAllocated,omitempty"`
	SizePreallocated         units.Bytes             `json:"sizePreallocated,omitempty"`
	SizeAllocatedTotal       units.Bytes             `json:"SizeAllocatedTotal,omitempty"`
	DataReductionSizeSaved   units.Bytes             `json:"dataReductionSizeSaved,omitempty"`
	DataReductionPercent     int64                   `json:"dataReductionPercent,omitempty"`
	DataReductionRatio       float64                 `json:"dataReductionRatio,omitempty"`
	PerTierSizeUsed          []units.Bytes           `json:"perTierSizeUsed,omitempty"`
	IsThinEnabled            bool                    `json:"isThinEnabled,omitempty"`
	IsDataReductionEnabled   bool                    `json:"isDataReductionEnabled,omitempty"`
	IsAdvancedDedupEnabled   bool                    `json:"isAdvancedDedupEnabled,omitempty"`
	StorageResource          StorageResourceContent  `json:"storageResource,omitempty"`
	Pool                     PoolContent             `json:"pool,omitempty"`
	Wwn                      string                  `json:"wwn,omitempty"`
	TieringPolicy            types.TieringPolicyEnum `json:"tieringPolicy,omitempty"`
	DefaultNode              types.NodeEnum          `json:"defaultNode,omitempty"`
	IsReplicationDestination bool                    `json:"isReplicationDestination,omitempty"`
	CurrentNode              types.NodeEnum          `json:"currentNode,omitempty"`
	//SnapSchedule		SnapScheduleContent `json:"snapSchedule,omitempty"`
	IsSnapSchedulePaused bool `json:"isSnapSchedulePaused,omitempty"`
	//IoLimitPolicy IoLimitPolicyContent `json:"ioLimitPolicy,omitempty"`
	MetadataSize          units.Bytes `json:"metadataSize,omitempty"`
	MetadataSizeAllocated units.Bytes `json:"metadataSizeAllocated,omitempty"`
	SnapWwn               string      `json:"snapWwn,omitempty"`
	SnapsSize             units.Bytes `json:"snapsSize,omitempty"`
	SnapsSizeAllocated    units.Bytes `json:"snapsSizeAllocated,omitempty"`
	//HostAccess
	SnapCount   int64 `json:"snapCount,omitempty"`
	IsThinClone bool  `json:"isThinClone,omitempty"`
	//FamilyBaseLun
	//ParentSnap
	//OriginalParentLun
	FamilySnapCount          int64       `json:"familySnapCount,omitempty"`
	FamilyCloneCount         int64       `json:"familyCloneCount,omitempty"`
	NonBaseSizeAllocated     units.Bytes `json:"nonBaseSizeAllocated,omitempty"`
	FamilySizeAllocatedTotal units.Bytes `json:"familySizeAllocatedTotal,omitempty"`
	NonBaseSizeUsed          units.Bytes `json:"nonBaseSizeUsed,omitempty"`
	MigrationState           string      `json:"migrationState,omitempty"`
}

func (_c *UnisphereClient) GetLunInstances(fields []string, filter string) (*LunInstances, error) {
	req, err := api.UnityAPILunInstances.NewRequest(_c.endpoint)
	if err != nil {
		return nil, err
	}
	api.UnityAPILunInstances.WithFields(fields, req)
	api.UnityAPILunInstances.WithFilter(filter, req)
	_c.addHeader("GET", req)

	var body []byte
	body, err = _c.send(req)

	var data LunInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
