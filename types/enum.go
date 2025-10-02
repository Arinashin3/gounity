package types

type HealthEnum int

const (
	HealthEnumUnknown        HealthEnum = 0
	HealthEnumOk             HealthEnum = 5
	HealthEnumOkBut          HealthEnum = 7
	HealthEnumDegraded       HealthEnum = 10
	HealthEnumMinor          HealthEnum = 15
	HealthEnumMajor          HealthEnum = 20
	HealthEnumCritical       HealthEnum = 25
	HealthEnumNonRecoverable HealthEnum = 30
)

var Health = map[HealthEnum]string{
	HealthEnumUnknown:        "Unknown",
	HealthEnumOk:             "Ok",
	HealthEnumOkBut:          "OkBut",
	HealthEnumDegraded:       "Degraded",
	HealthEnumMinor:          "Minor",
	HealthEnumMajor:          "Major",
	HealthEnumCritical:       "Critical",
	HealthEnumNonRecoverable: "NonRecoverable",
}

func (_enum HealthEnum) String() string {
	return Health[_enum]
}

type NodeEnum int

const (
	NodeSPA NodeEnum = iota
	NodeSPB
	NodeUnknown NodeEnum = 2989
)

var Node = map[NodeEnum]string{
	NodeSPA:     "SPA",
	NodeSPB:     "SPB",
	NodeUnknown: "Unknown",
}

func (_enum NodeEnum) String() string {
	return Node[_enum]
}

type TierTypeEnum int

const (
	TierTypeNone               TierTypeEnum = iota
	TierTypeExtremePerformance TierTypeEnum = 10
	TierTypePerformance        TierTypeEnum = 20
	TierTypeCapacity           TierTypeEnum = 30
)

var TierType = map[TierTypeEnum]string{
	TierTypeNone:               "None",
	TierTypeExtremePerformance: "Extreme_Performance",
	TierTypePerformance:        "Performance",
	TierTypeCapacity:           "Capacity",
}

func (_enum TierTypeEnum) String() string {
	return TierType[_enum]
}

type TieringPolicyEnum int

const (
	TieringPolicyAutotierHigh TieringPolicyEnum = iota
	TieringPolicyAutotier
	TieringPolicyHighest
	TieringPolicyLowest
	TieringPolicyNoDataMovement
	TieringPolicyMixed TieringPolicyEnum = 65535
)

var TieringPolicy = map[TieringPolicyEnum]string{
	TieringPolicyAutotierHigh:   "Autotier_High",
	TieringPolicyAutotier:       "Autotier",
	TieringPolicyHighest:        "Highest",
	TieringPolicyLowest:         "Lowest",
	TieringPolicyNoDataMovement: "NoDataMovement",
	TieringPolicyMixed:          "Mixed",
}

func (_enum TieringPolicyEnum) String() string {
	return TieringPolicy[_enum]
}

type SPModelNameEnum int

const (
	SPModelNameSP300 SPModelNameEnum = iota + 1
	SPModelNameSP400
	SPModelNameSP500
	SPModelNameSP600
	SPModelNameSP350
	SPModelNameSP450
	SPModelNameSP550
	SPModelNameSP650
	SPModelNameVSA2Core
	SPModelNameVSA12Core
	SPModelNameSP380
	SPModelNameSP480
	SPModelNameSP680
	SPModelNameSP880
)

var SPModelName = map[SPModelNameEnum]string{
	SPModelNameSP300:     "SP300",
	SPModelNameSP400:     "SP400",
	SPModelNameSP500:     "SP500",
	SPModelNameSP600:     "SP600",
	SPModelNameSP350:     "SP350",
	SPModelNameSP450:     "SP450",
	SPModelNameSP550:     "SP550",
	SPModelNameSP650:     "SP650",
	SPModelNameVSA2Core:  "VSA2Core",
	SPModelNameVSA12Core: "VSA12Core",
	SPModelNameSP380:     "SP380",
	SPModelNameSP480:     "SP480",
	SPModelNameSP680:     "SP680",
	SPModelNameSP880:     "SP880",
}

func (_enum SPModelNameEnum) String() string {
	return SPModelName[_enum]
}

type StorageResourceTypeEnum int

const (
	StorageResourceTypeFilesystem StorageResourceTypeEnum = iota + 1
	StorageResourceTypeConsistencyGroup
	StorageResourceTypeVmwarefs
	StorageResourceTypeVmwareiscsi
	StorageResourceTypeLun StorageResourceTypeEnum = iota + 4
	StorageResourceTypeVVolDatastoreFS
	StorageResourceTypeVVolDatastoreIscsi
)

var StorageResourceType = map[StorageResourceTypeEnum]string{
	StorageResourceTypeFilesystem:         "Filesystem",
	StorageResourceTypeConsistencyGroup:   "ConsistencyGroup",
	StorageResourceTypeVmwarefs:           "Vmwarefs",
	StorageResourceTypeVmwareiscsi:        "Vmwareiscsi",
	StorageResourceTypeLun:                "Lun",
	StorageResourceTypeVVolDatastoreFS:    "VVolDatastoreFS",
	StorageResourceTypeVVolDatastoreIscsi: "VVolDatastoreIscsi",
}

func (_enum StorageResourceTypeEnum) String() string {
	return StorageResourceType[_enum]
}

type ReplicationTypeEnum int

const (
	ReplicationTypeNone ReplicationTypeEnum = iota
	ReplicationTypeLocal
	ReplicationTypeRemote
	ReplicationTypeMixed
)

var ReplicationType = map[ReplicationTypeEnum]string{
	ReplicationTypeNone:   "None",
	ReplicationTypeLocal:  "Local",
	ReplicationTypeRemote: "Remote",
	ReplicationTypeMixed:  "Mixed",
}

func (_enum ReplicationTypeEnum) String() string {
	return ReplicationType[_enum]
}

type ThinStatusEnum int

const (
	ThinStatusFalse ThinStatusEnum = iota
	ThinStatusTrue
	ThinStatusMixed ThinStatusEnum = 65535
)

var ThinStatus = map[ThinStatusEnum]string{
	ThinStatusFalse: "False",
	ThinStatusTrue:  "True",
	ThinStatusMixed: "Mixed",
}

func (_enum ThinStatusEnum) String() string {
	return ThinStatus[_enum]
}

type DedupStatusEnum int

const (
	DedupStatusDisabled DedupStatusEnum = iota
	DedupStatusEnabled
	DedupStatusMixed
)

var DedupStatus = map[DedupStatusEnum]string{
	DedupStatusDisabled: "Disabled",
	DedupStatusEnabled:  "Enabled",
	DedupStatusMixed:    "Mixed",
}

func (_enum DedupStatusEnum) String() string {
	return DedupStatus[_enum]
}

type RaidTypeEnum int

const (
	RaidTypeNone RaidTypeEnum = iota
	RaidTypeRAID5
	RaidTypeRAID0
	RaidTypeRAID1
	RaidTypeRAID3
	RaidTypeRAID10    RaidTypeEnum = 7
	RaidTypeRAID6     RaidTypeEnum = 10
	RaidTypeMixed     RaidTypeEnum = 12
	RaidTypeAutomatic RaidTypeEnum = 48879
)

var RaidType = map[RaidTypeEnum]string{
	RaidTypeNone:      "None",
	RaidTypeRAID5:     "RAID5",
	RaidTypeRAID0:     "RAID0",
	RaidTypeRAID1:     "RAID1",
	RaidTypeRAID3:     "RAID3",
	RaidTypeRAID10:    "RAID10",
	RaidTypeRAID6:     "RAID6",
	RaidTypeMixed:     "RAID6",
	RaidTypeAutomatic: "RAID6",
}

func (_enum RaidTypeEnum) String() string {
	return RaidType[_enum]
}

type StoragePoolTypeEnum int

const (
	StoragePoolTypeTraditional StoragePoolTypeEnum = iota + 1
	StoragePoolTypeDynamic
)

var StoragePoolType = map[StoragePoolTypeEnum]string{
	StoragePoolTypeTraditional: "Traditional",
	StoragePoolTypeDynamic:     "Dynamic",
}

func (_enum StoragePoolTypeEnum) String() string {
	return StoragePoolType[_enum]
}

type InterfaceConfigModeEnum int

const (
	InterfaceConfigModeDisabled InterfaceConfigModeEnum = iota
	InterfaceConfigModeStatic
	InterfaceConfigModeAuto
)

var InterfaceConfigMode = map[InterfaceConfigModeEnum]string{
	InterfaceConfigModeDisabled: "Disabled",
	InterfaceConfigModeStatic:   "Static",
	InterfaceConfigModeAuto:     "Auto",
}

func (_enum InterfaceConfigModeEnum) String() string {
	return InterfaceConfigMode[_enum]
}

type IpProtocolVersionEnum int

const (
	IpProtocolVersionIPv4 IpProtocolVersionEnum = 4
	IpProtocolVersionIPv6 IpProtocolVersionEnum = 6
)

var IpProtocolVersion = map[IpProtocolVersionEnum]string{
	IpProtocolVersionIPv4: "IPv4",
	IpProtocolVersionIPv6: "IPv6",
}

func (_enum IpProtocolVersionEnum) String() string {
	return IpProtocolVersion[_enum]
}

type SeverityEnum int

const (
	SeverityEMERGENCY SeverityEnum = iota
	SeverityALERT
	SeverityCRITICAL
	SeverityERROR
	SeverityWARNING
	SeverityNOTICE
	SeverityINFO
	SeverityDEBUG
	SeverityOK
)

var Severity = map[SeverityEnum]string{
	SeverityEMERGENCY: "EMERGENCY",
	SeverityALERT:     "ALERT",
	SeverityCRITICAL:  "CRITICAL",
	SeverityERROR:     "ERROR",
	SeverityWARNING:   "WARNING",
	SeverityNOTICE:    "NOTICE",
	SeverityINFO:      "INFO",
	SeverityDEBUG:     "DEBUG",
	SeverityOK:        "OK",
}

func (_enum SeverityEnum) String() string {
	return Severity[_enum]
}

type EventCategoryEnum int

const (
	EventCategoryUser EventCategoryEnum = iota
	EventCategoryAudit
	EventCategoryAuthentication
)

var EventCategory = map[EventCategoryEnum]string{
	EventCategoryUser:           "User",
	EventCategoryAudit:          "Audit",
	EventCategoryAuthentication: "Authentication",
}

func (_enum EventCategoryEnum) String() string {
	return EventCategory[_enum]
}

type MetricTypeEnum int

const (
	MetricType32BitsCounter MetricTypeEnum = iota + 2
	MetricType64BitsCounter
	MetricTypeRate
	MetricTypeFact
	MetricTypeText
	MetricType32BitsVirtualCounter
	MetricType64BitsVirtualCounter
)

var MetricType = map[MetricTypeEnum]string{
	MetricType32BitsCounter:        "32BitsCounter",
	MetricType64BitsCounter:        "64BitsCounter",
	MetricTypeRate:                 "Rate",
	MetricTypeFact:                 "Fact",
	MetricTypeText:                 "Text",
	MetricType32BitsVirtualCounter: "32BitsVirtualCounter",
	MetricType64BitsVirtualCounter: "64BitsVirtualCounter",
}

func (_enum MetricTypeEnum) String() string {
	return MetricType[_enum]
}

type FilesystemTypeEnum int

const (
	FilesystemTypeFilesystem FilesystemTypeEnum = iota + 1
	FilesystemTypeVMware
)

var FilesystemType = map[FilesystemTypeEnum]string{
	FilesystemTypeFilesystem: "Filesystem",
	FilesystemTypeVMware:     "VMware",
}

func (_enum FilesystemTypeEnum) String() string {
	return FilesystemType[_enum]
}

type FLRVersionEnum int

const (
	FLRVersionOff FLRVersionEnum = iota
	FLRVersionEnterprise
	FLRVersionCompliance
)

var FLRVersion = map[FLRVersionEnum]string{
	FLRVersionOff:        "Off",
	FLRVersionEnterprise: "Enterprise",
	FLRVersionCompliance: "Compliance",
}

func (_enum FLRVersionEnum) String() string {
	return FLRVersion[_enum]
}

type FSSupportedProtocolEnum int

const (
	FSSupportedProtocolNFS = iota
	FSSupportedProtocolCIFS
	FSSupportedProtocolMultiprotocol
)

var FSSupportedProtocol = map[FSSupportedProtocolEnum]string{
	FSSupportedProtocolNFS:           "NFS",
	FSSupportedProtocolCIFS:          "CIFS",
	FSSupportedProtocolMultiprotocol: "Multiprotocol",
}

func (_enum FSSupportedProtocolEnum) String() string {
	return FSSupportedProtocol[FSSupportedProtocolNFS]
}

type AccessPolicyEnum int

const (
	AccessPolicyNative AccessPolicyEnum = iota
	AccessPolicyUnix
	AccessPolicyWindows
)

var AccessPolicy = map[AccessPolicyEnum]string{
	AccessPolicyNative:  "Native",
	AccessPolicyUnix:    "Unix",
	AccessPolicyWindows: "Windows",
}

func (_enum AccessPolicyEnum) String() string {
	return AccessPolicy[_enum]
}
