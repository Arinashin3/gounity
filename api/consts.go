package api

import (
	"bytes"
	"net/http"
	"strings"
)

const (
	UnityAPIPrefix      = "/api"
	UnityAPITypesPrefix = UnityAPIPrefix + "/types"
)

type UnityAPITypesInstances string

const (
	UnityAPIBasicSystemInfoInstances     UnityAPITypesInstances = UnityAPITypesPrefix + "/basicSystemInfo/instances?compact=true"
	UnityAPISystemCapacityInstances      UnityAPITypesInstances = UnityAPITypesPrefix + "/systemCapacity/instances?compact=true"
	UnityAPISystemInstances              UnityAPITypesInstances = UnityAPITypesPrefix + "/system/instances?compact=true"
	UnityAPILunInstances                 UnityAPITypesInstances = UnityAPITypesPrefix + "/lun/instances?compact=true"
	UnityAPIPoolInstances                UnityAPITypesInstances = UnityAPITypesPrefix + "/pool/instances?compact=true"
	UnityAPIStorageResourceInstances     UnityAPITypesInstances = UnityAPITypesPrefix + "/storageResource/instances?compact=true"
	UnityAPIMgmtInterfaceInstances       UnityAPITypesInstances = UnityAPITypesPrefix + "/mgmtInterface/instances?compact=true"
	UnityAPIEventInstances               UnityAPITypesInstances = UnityAPITypesPrefix + "/event/instances?compact=true"
	UnityAPIMetricInstances              UnityAPITypesInstances = UnityAPITypesPrefix + "/metric/instances?compact=true"
	UnityAPIMetricRealTimeQueryInstances UnityAPITypesInstances = UnityAPITypesPrefix + "/metricRealTimeQuery/instances?compact=true"
	UnityAPIMetricQueryResultInstances   UnityAPITypesInstances = UnityAPITypesPrefix + "/metricQueryResult/instances?compact=true"
	UnityAPIMetricValueInstances         UnityAPITypesInstances = UnityAPITypesPrefix + "/metricValue/instances?compact=true"
	UnityAPIFilesystemInstances          UnityAPITypesInstances = UnityAPITypesPrefix + "/filesystem/instances?compact=true"
)

func (_u UnityAPITypesInstances) NewRequest(endpoint string, body []byte) (*http.Request, error) {
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest("GET", string(_u), nil)
	} else {
		req, err = http.NewRequest("POST", string(_u), bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}

	req.URL, err = req.URL.Parse(endpoint + string(_u))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (_u UnityAPITypesInstances) WithFields(fields []string, req *http.Request) {
	if len(fields) == 0 {
		return
	}
	if req.URL.RawQuery != "" {
		req.URL.RawQuery += "&"
	}
	req.URL.RawQuery += "fields=" + strings.Join(fields, ",")
}

func (_u UnityAPITypesInstances) WithFilter(filter []string, req *http.Request) {
	if len(filter) == 0 {
		return
	}
	if req.URL.RawQuery != "" {
		req.URL.RawQuery += "&"
	}
	for i, s := range filter {
		s = strings.Replace(s, " ", "%20", -1)
		req.URL.RawQuery += "filter=" + s
		if i < len(filter)-1 {
			req.URL.RawQuery += "&"
		}
	}
}
