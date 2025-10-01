package types

// HealthContent
// Health
type HealthContent struct {
	Value          HealthEnum `json:"value"`
	DescriptionIds []string   `json:"descriptionIds"`
	Descriptions   []string   `json:"descriptions"`
	ResolutionIds  []string   `json:"resolutionIds"`
	Resolutions    []string   `json:"resolutions"`
}
