package gounity

import (
	"encoding/json"
	"time"

	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
	"github.com/Arinashin3/gounity/units"
)

type FilesystemInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content FilesystemContent `json:"content"`
	} `json:"entries"`
}

type FilesystemContent struct {
	Id                       string                   `json:"id,omitempty"`
	Health                   types.HealthContent      `json:"health,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	Description              string                   `json:"description,omitempty"`
	Type                     types.FilesystemTypeEnum `json:"type,omitempty"`
	SizeTotal                units.Bytes              `json:"sizeTotal,omitempty"`
	SizeUsed                 units.Bytes              `json:"sizeUsed,omitempty"`
	SizeAllocated            units.Bytes              `json:"sizeAllocated,omitempty"`
	SizePreallocated         units.Bytes              `json:"sizePreallocated,omitempty"`
	MinSizeAllocated         units.Bytes              `json:"minSizeAllocated,omitempty"`
	IsReadOnly               bool                     `json:"isReadOnly,omitempty"`
	IsThinEnabled            bool                     `json:"isThinEnabled,omitempty"`
	IsDataReductionEnabled   bool                     `json:"isDataReductionEnabled,omitempty"`
	DataReductionSizeSaved   units.Bytes              `json:"dataReductionSizeSaved,omitempty"`
	DataReductionPercent     units.Percentage         `json:"dataReductionPercent,omitempty"`
	DataReductionRatio       units.Ratio              `json:"dataReductionRatio,omitempty"`
	IsAdvancedDedupEnabled   bool                     `json:"isAdvancedDedupEnabled,omitempty"`
	StorageResource          StorageResourceContent   `json:"storageResource,omitempty"`
	FlrVersion               types.FLRVersionEnum     `json:"flrVersion,omitempty"`
	IsFlrProtected           bool                     `json:"isFlrProtected,omitempty"`
	FlrSoftwareClock         string                   `json:"flrSoftwareClock,omitempty"`
	FlrProtectedMaxRetention string                   `json:"flrProtectedMaxRetention,omitempty"`
	IsFlrAutoLockEnabled     bool                     `json:"isFlrAutoLockEnabled,omitempty"`
	IsFlrAutoDeleteEnabled   bool                     `json:"isFlrAutoDeleteEnabled,omitempty"`
	FlrAutoLockDelay         string                   `json:"flrAutoLockDelay,omitempty"`
	ErrorThreshold           units.Percentage         `json:"errorThreshold,omitempty"`
	WarningThreshold         units.Percentage         `json:"warningThreshold,omitempty"`
	InfoThreshold            units.Percentage         `json:"infoThreshold,omitempty"`
	IsCIFSSyncWritesEnabled  bool                     `json:"isCIFSSyncWritesEnabled,omitempty"`
	Pool                     PoolContent              `json:"pool,omitempty"`
	IsCIFSOpLocksEnabled     bool                     `json:"isCIFSOpLocksEnabled,omitempty"`
	//NasServer
	IsCIFSNotifyOnWriteEnabled  bool                          `json:"isCIFSNotifyOnWriteEnabled,omitempty"`
	IsCIFSNotifyOnAccessEnabled bool                          `json:"isCIFSNotifyOnAccessEnabled,omitempty"`
	CifsNotifyOnChangeDirDepth  int                           `json:"cifsNotifyOnChangeDirDepth,omitempty"`
	TieringPolicy               types.TieringPolicyEnum       `json:"tieringPolicy,omitempty"`
	SupportedProtocols          types.FSSupportedProtocolEnum `json:"supportedProtocols,omitempty"`
	MetadataSize                units.Bytes                   `json:"metadataSize,omitempty"`
	MetadataSizeAllocated       units.Bytes                   `json:"metadataSizeAllocated,omitempty"`
	PerTierSizeUsed             []units.Bytes                 `json:"perTierSizeUsed,omitempty"`
	SnapsSize                   units.Bytes                   `json:"snapsSize,omitempty"`
	SnapsSizeAllocated          units.Bytes                   `json:"snapsSizeAllocated,omitempty"`
	SnapCount                   int                           `json:"snapCount,omitempty"`
	IsSMBCA                     bool                          `json:"isSMBCA,omitempty"`
	AccessPolicy                types.AccessPolicyEnum        `json:"accessPolicy,omitempty"`
	//FolderRenamePolicy
	//LockingPolicy
	//Format
	//HostIOSize
	//PoolFullPolicy
	//FileEventSettings
	//CifsShare
	//NfsShare
}

func (_c *UnisphereClient) GetFilesystemInstances(fields []string, filter []string) (*FilesystemInstances, error) {
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
