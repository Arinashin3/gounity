package gounity

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

type HealthStruct struct {
	Value          HealthEnum `json:"value"`
	DescriptionIds []string   `json:"descriptionIds"`
	Descriptions   []string   `json:"descriptions"`
	ResolutionIds  []string   `json:"resolutionIds"`
	Resolutions    []string   `json:"resolutions"`
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

func (s SPModelNameEnum) Description() string {
	switch s {
	case SPModelNameSP300:
		return "Unity 300 or Unity 300F"
	case SPModelNameSP400:
		return "Unity 400 or Unity 400F"
	case SPModelNameSP500:
		return "Unity 500 or Unity 500F"
	case SPModelNameSP600:
		return "Unity 600 or Unity 600F"
	case SPModelNameSP350:
		return "Unity 350"
	case SPModelNameSP450:
		return "Unity 450"
	case SPModelNameSP550:
		return "Unity 550"
	case SPModelNameSP650:
		return "Unity 650"
	case SPModelNameVSA2Core:
		return "UnityVSA 2Core"
	case SPModelNameVSA12Core:
		return "UnityVSA 12Core"
	case SPModelNameSP380:
		return "Unity 380 or Unity 380F"
	case SPModelNameSP480:
		return "Unity 480 or Unity 480F"
	case SPModelNameSP680:
		return "Unity 680 or Unity 680F"
	case SPModelNameSP880:
		return "Unity 880 or Unity 880F"
	default:
		return "Unknown"
	}
}
