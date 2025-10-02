package gounity

import (
	"encoding/json"
	"time"

	"github.com/Arinashin3/gounity/api"
)

type MetricValueInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content MetricValueContent `json:"content"`
	} `json:"entries"`
}

type MetricValueContent struct {
	Path      string      `json:"path,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Interval  int         `json:"interval,omitempty"`
	Values    interface{} `json:"values,omitempty"`
}

func (_c *UnisphereClient) GetMetricValueInstances(path string) (*MetricValueInstances, error) {
	req, err := api.UnityAPIMetricValueInstances.NewRequest(_c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	filter := []string{
		"path eq \"" + path + "\"",
	}
	api.UnityAPIMetricValueInstances.WithFilter(filter, req)
	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)
	if err != nil {
		return nil, err
	}

	var data MetricValueInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
