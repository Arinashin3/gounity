package gounity

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/Arinashin3/gounity/api"
)

type MetricQueryResultInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries struct {
		Content MetricQueryResultContent `json:"content"`
	} `json:"entries"`
}

type MetricQueryResultContent struct {
	QueryId   int         `json:"queryId,omitempty"`
	Path      string      `json:"path,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Values    interface{} `json:"values,omitempty"`
}

func (_c *UnisphereClient) GetMetricQueryResultInstances(queryId int) (*MetricQueryResultInstances, error) {
	req, err := api.UnityAPIMetricQueryResultInstances.NewRequest(_c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	if queryId == 0 {
		return nil, errors.New("queryId is required")
	}
	var filters []string
	filters = append(filters, "queryId eq "+strconv.Itoa(queryId))

	api.UnityAPIMetricQueryResultInstances.WithFilter(filters, req)
	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)

	var data MetricQueryResultInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}
