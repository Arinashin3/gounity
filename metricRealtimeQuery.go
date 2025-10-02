package gounity

import (
	"encoding/json"
	"time"

	"github.com/Arinashin3/gounity/api"
)

type MetricRealTimeQueryInstances struct {
	Base    string                     `json:"@base"`
	Updated time.Time                  `json:"updated"`
	Content MetricRealTimeQueryContent `json:"content"`
}

type MetricRealTimeQueryContent struct {
	Id         int64     `json:"id,omitempty"`
	Paths      []string  `json:"paths,omitempty"`
	Interval   int       `json:"interval,omitempty"`
	Expiration time.Time `json:"expiration,omitempty"`
}

func (_c *UnisphereClient) GetMetricRealTimeQueryInstances(fields []string, filter []string) (*MetricRealTimeQueryInstances, error) {
	req, err := api.UnityAPIMetricRealTimeQueryInstances.NewRequest(_c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	api.UnityAPIMetricRealTimeQueryInstances.WithFields(fields, req)
	api.UnityAPIMetricRealTimeQueryInstances.WithFilter(filter, req)
	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)
	if err != nil {
		return nil, err
	}

	var data MetricRealTimeQueryInstances
	err = json.Unmarshal(body, &data)
	return &data, err
}

type MetricRealTimeQueryInstancesCreated struct {
	Content struct {
		Id             int       `json:"id,omitempty"`
		Paths          []string  `json:"paths,omitempty"`
		Interval       int       `json:"interval,omitempty"`
		MaximumSamples int       `json:"maximumSamples,omitempty"`
		Expiration     time.Time `json:"expiration,omitempty"`
	} `json:"content"`
}

func (_c *UnisphereClient) PostMetricRealTimeQueryInstances(paths []string, interval time.Duration) (*MetricRealTimeQueryInstancesCreated, error) {
	var reqContent struct {
		Paths    []string `json:"paths"`
		Interval int      `json:"interval"`
	}
	reqContent.Paths = paths
	reqContent.Interval = int(interval / time.Second)
	reqBody, err := json.Marshal(reqContent)
	if err != nil {
		return nil, err
	}

	// s
	req, err := api.UnityAPIMetricRealTimeQueryInstances.NewRequest(_c.endpoint, reqBody)
	if err != nil {
		return nil, err
	}

	_c.addHeader(req)

	var body []byte
	body, err = _c.send(req)

	var data MetricRealTimeQueryInstancesCreated
	err = json.Unmarshal(body, &data)
	return &data, err
}
