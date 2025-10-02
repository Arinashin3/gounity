package gounity

import (
	"encoding/json"
	"time"

	"github.com/Arinashin3/gounity/api"
	"github.com/Arinashin3/gounity/types"
)

type MetricInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content MetricContent `json:"content"`
	} `json:"entries"`
}

type MetricContent struct {
	Id                    int                  `json:"id,omitempty"`
	Name                  string               `json:"name,omitempty"`
	Path                  string               `json:"path,omitempty"`
	Type                  types.MetricTypeEnum `json:"type,omitempty"`
	Description           string               `json:"description,omitempty"`
	IsHistoricalAvailable bool                 `json:"isHistoricalAvailable,omitempty"`
	IsRealtimeAvailable   bool                 `json:"isRealtimeAvailable,omitempty"`
	UnitDisplayString     string               `json:"unitDisplayString,omitempty"`
}

func (_c *UnisphereClient) GetMetricInstances(fields []string, filter []string) (*MetricInstances, error) {
	req, err := api.UnityAPIMetricInstances.NewRequest(_c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	api.UnityAPIMetricInstances.WithFields(fields, req)
	api.UnityAPIMetricInstances.WithFilter(filter, req)
	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)
	if err != nil {
		return nil, err
	}

	var data MetricInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
