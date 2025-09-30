package gounity

import (
	"encoding/json"
	"errors"
	"gounity/api"
	"io"
	"net/http"
	"time"
)

type SystemCapacityInstances struct {
	Base    string    `json:"@base"`
	Updated time.Time `json:"updated"`
	Entries []struct {
		Content SystemCapacityContent `json:"content"`
	} `json:"entries"`
}

type SystemCapacityContent struct {
	Id                     string  `json:"id,omitempty"`
	SizeFree               int64   `json:"sizeFree,omitempty"`
	SizeTotal              int64   `json:"sizeTotal,omitempty"`
	SizeUsed               int64   `json:"sizeUsed,omitempty"`
	SizePreallocated       int64   `json:"sizePreallocated,omitempty"`
	DataReductionSizeSaved int64   `json:"dataReductionSizeSaved,omitempty"`
	DataReductionPercent   int64   `json:"dataReductionPercent,omitempty"`
	DataReductionRatio     float64 `json:"dataReductionRatio,omitempty"`
	SizeSubscribed         int64   `json:"sizeSubscribed,omitempty"`
	TotalLogicalSize       int64   `json:"totalLogicalSize,omitempty"`
	ThinSavingRatio        float64 `json:"thinSavingRatio,omitempty"`
	SnapsSavingsRatio      float64 `json:"snapsSavingsRatio,omitempty"`
	OverallEfficiencyRatio float64 `json:"overallEfficiencyRatio,omitempty"`
	Tiers                  []struct {
		TierType  int64 `json:"tierType,omitempty"`
		SizeFree  int64 `json:"sizeFree,omitempty"`
		SizeTotal int64 `json:"sizeTotal,omitempty"`
		SizeUsed  int64 `json:"sizeUsed,omitempty"`
	} `json:"titers,omitempty"`
}

func (_c *UnisphereClient) GetSystemCapacityInstances(fields []string) (*SystemCapacityInstances, error) {
	inst := api.UnityAPISystemCapacityInstances
	req, err := inst.NewRequest(_c.endpoint)
	inst.WithFields(fields, req)
	if err != nil {
		return nil, err
	}
	_c.addHeader("GET", req)

	resp, err := _c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		_c.logined = false
	case http.StatusUnprocessableEntity:
		var data StatusUnprocessableEntity
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(data.Error.Messages[0].EnUS)
	}

	var data *SystemCapacityInstances
	err = json.Unmarshal(body, &data)
	return data, err

}
