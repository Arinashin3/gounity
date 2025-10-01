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

func (h *HealthEnum) String() string {
	switch *h {
	case HealthEnumUnknown:
		return "Unknown"
	case HealthEnumOk:
		return "Ok"
	case HealthEnumOkBut:
		return "OkBut"
	case HealthEnumDegraded:
		return "Degraded"
	case HealthEnumMinor:
		return "Minor"
	case HealthEnumMajor:
		return "Major"
	case HealthEnumCritical:
		return "Critical"
	case HealthEnumNonRecoverable:
		return "NonRecoverable"
	default:
		return "Unknown"
	}
}

type NodeEnum int

const (
	NodeSPA NodeEnum = iota
	NodeSPB
	NodeUnknown NodeEnum = 2989
)

func (s NodeEnum) String() string {
	switch s {
	case NodeSPA:
		return "SPA"
	case NodeSPB:
		return "SPB"
	case NodeUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

type TierTypeEnum int

const (
	TierTypeNone               TierTypeEnum = iota
	TierTypeExtremePerformance TierTypeEnum = 10
	TierTypePerformance        TierTypeEnum = 20
	TierTypeCapacity           TierTypeEnum = 30
)

func (s TierTypeEnum) String() string {
	switch s {
	case TierTypeNone:
		return "None"
	case TierTypeExtremePerformance:
		return "Extreme_Performance"
	case TierTypePerformance:
		return "Performance"
	case TierTypeCapacity:
		return "Capacity"
	default:
		return "Unknown"
	}
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

func (s TieringPolicyEnum) String() string {
	switch s {
	case TieringPolicyAutotierHigh:
		return "Autotier_High"
	case TieringPolicyAutotier:
		return "Autotier"
	case TieringPolicyHighest:
		return "Highest"
	case TieringPolicyLowest:
		return "Lowest"
	case TieringPolicyNoDataMovement:
		return "No_Data_Movement"
	case TieringPolicyMixed:
		return "Mixed"
	default:
		return "Unknown"
	}
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

func (s SPModelNameEnum) String() string {
	switch s {
	case SPModelNameSP300:
		return "SP300"
	case SPModelNameSP400:
		return "SP400"
	case SPModelNameSP500:
		return "SP500"
	case SPModelNameSP600:
		return "SP600"
	case SPModelNameSP350:
		return "SP350"
	case SPModelNameSP450:
		return "SP450"
	case SPModelNameSP550:
		return "SP550"
	case SPModelNameSP650:
		return "SP650"
	case SPModelNameVSA2Core:
		return "VSA2Core"
	case SPModelNameVSA12Core:
		return "VSA12Core"
	case SPModelNameSP380:
		return "SP380"
	case SPModelNameSP480:
		return "SP480"
	case SPModelNameSP680:
		return "SP680"
	case SPModelNameSP880:
		return "SP880"
	default:
		return "Unknown"
	}
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

func (_n StorageResourceTypeEnum) String() string {
	switch _n {
	case StorageResourceTypeFilesystem:
		return "Filesystem"
	case StorageResourceTypeConsistencyGroup:
		return "ConsistencyGroup"
	case StorageResourceTypeVmwarefs:
		return "Vmwarefs"
	case StorageResourceTypeVmwareiscsi:
		return "Vmwareiscsi"
	case StorageResourceTypeLun:
		return "Lun"
	case StorageResourceTypeVVolDatastoreFS:
		return "VVolDatastoreFS"
	case StorageResourceTypeVVolDatastoreIscsi:
		return "VVolDatastoreIscsi"
	default:
		return "Unknown"
	}
}

type ReplicationTypeEnum int

const (
	ReplicationTypeNone ReplicationTypeEnum = iota
	ReplicationTypeLocal
	ReplicationTypeRemote
	ReplicationTypeMixed
)

func (_n ReplicationTypeEnum) String() string {
	switch _n {
	case ReplicationTypeNone:
		return "None"
	case ReplicationTypeLocal:
		return "Local"
	case ReplicationTypeRemote:
		return "Remote"
	case ReplicationTypeMixed:
		return "Mixed"
	default:
		return "Unknown"
	}
}

type ThinStatusEnum int

const (
	ThinStatusFalse ThinStatusEnum = iota
	ThinStatusTrue
	ThinStatusMixed ThinStatusEnum = 65535
)

func (s ThinStatusEnum) String() string {
	switch s {
	case ThinStatusFalse:
		return "False"
	case ThinStatusTrue:
		return "True"
	case ThinStatusMixed:
		return "Mixed"
	default:
		return "Unknown"
	}
}

type DedupStatusEnum int

const (
	DedupStatusDisabled DedupStatusEnum = iota
	DedupStatusEnabled
	DedupStatusMixed
)

func (s DedupStatusEnum) String() string {
	switch s {
	case DedupStatusDisabled:
		return "Disabled"
	case DedupStatusEnabled:
		return "Enabled"
	case DedupStatusMixed:
		return "Mixed"
	default:
		return "Unknown"
	}
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

func (s RaidTypeEnum) String() string {
	switch s {
	case RaidTypeNone:
		return "None"
	case RaidTypeRAID5:
		return "RAID5"
	case RaidTypeRAID0:
		return "RAID0"
	case RaidTypeRAID1:
		return "RAID1"
	case RaidTypeRAID3:
		return "RAID3"
	case RaidTypeRAID10:
		return "RAID10"
	case RaidTypeRAID6:
		return "RAID6"
	case RaidTypeMixed:
		return "Mixed"
	case RaidTypeAutomatic:
		return "Automatic"
	default:
		return "Unknown"
	}
}

type StoragePoolTypeEnum int

const (
	StoragePoolTypeTraditional StoragePoolTypeEnum = iota + 1
	StoragePoolTypeDynamic
)

func (s StoragePoolTypeEnum) String() string {
	switch s {
	case StoragePoolTypeTraditional:
		return "Traditional"
	case StoragePoolTypeDynamic:
		return "Dynamic"
	default:
		return "Unknown"
	}
}

type InterfaceConfigModeEnum int

const (
	InterfaceConfigModeDisabled InterfaceConfigModeEnum = iota
	InterfaceConfigModeStatic
	InterfaceConfigModeAuto
)

func (s InterfaceConfigModeEnum) String() string {
	switch s {
	case InterfaceConfigModeDisabled:
		return "Disabled"
	case InterfaceConfigModeStatic:
		return "Static"
	case InterfaceConfigModeAuto:
		return "Auto"
	default:
		return "Unknown"
	}
}

type IpProtocolVersionEnum int

const (
	IpProtocolVersionIPv4 IpProtocolVersionEnum = 4
	IpProtocolVersionIPv6 IpProtocolVersionEnum = 6
)

func (s IpProtocolVersionEnum) String() string {
	switch s {
	case IpProtocolVersionIPv4:
		return "IPv4"
	case IpProtocolVersionIPv6:
		return "IPv6"
	default:
		return "Unknown"
	}
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

func (s SeverityEnum) String() string {
	switch s {
	case SeverityEMERGENCY:
		return "EMERGENCY"
	case SeverityALERT:
		return "ALERT"
	case SeverityCRITICAL:
		return "CRITICAL"
	case SeverityERROR:
		return "ERROR"
	case SeverityWARNING:
		return "WARNING"
	case SeverityNOTICE:
		return "NOTICE"
	case SeverityINFO:
		return "INFO"
	case SeverityDEBUG:
		return "DEBUG"
	case SeverityOK:
		return "OK"
	default:
		return "Unknown"
	}
}

type EventCategoryEnum int

const (
	EventCategoryUser EventCategoryEnum = iota
	EventCategoryAudit
	EventCategoryAuthentication
)

func (s EventCategoryEnum) String() string {
	switch s {
	case EventCategoryUser:
		return "User"
	case EventCategoryAudit:
		return "Audit"
	case EventCategoryAuthentication:
		return "Authentication"
	default:
		return "Unknown"
	}
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

func (s MetricTypeEnum) String() string {
	switch s {
	case MetricType32BitsCounter:
		return "32BitsCounter"
	case MetricType64BitsCounter:
		return "64BitsCounter"
	case MetricTypeRate:
		return "Rate"
	case MetricTypeFact:
		return "Fact"
	case MetricTypeText:
		return "Text"
	case MetricType32BitsVirtualCounter:
		return "32BitsVirtualCounter"
	default:
		return "Unknown"
	}
}
