package api

import (
	"net/http"
	"strings"
)

const (
	UnityAPIPrefix      = "/api"
	UnityAPITypesPrefix = UnityAPIPrefix + "/types"
)

type UnityAPITypesInstances string

const (
	UnityAPIBasicSystemInfoInstances UnityAPITypesInstances = UnityAPITypesPrefix + "/basicSystemInfo/instances?compact=true"
	UnityAPISystemCapacityInstances  UnityAPITypesInstances = UnityAPITypesPrefix + "/systemCapacity/instances?compact=true"
	UnityAPISystemInstances          UnityAPITypesInstances = UnityAPITypesPrefix + "/system/instances?compact=true"
	UnityAPILunInstances             UnityAPITypesInstances = UnityAPITypesPrefix + "/lun/instances?compact=true"
	UnityAPIPoolInstances            UnityAPITypesInstances = UnityAPITypesPrefix + "/pool/instances?compact=true"
	UnityAPIStorageResourceInstances UnityAPITypesInstances = UnityAPITypesPrefix + "/storageResource/instances?compact=true"
	UnityAPIMgmtInterfaceInstances   UnityAPITypesInstances = UnityAPITypesPrefix + "/mgmtInterface/instances?compact=true"
	UnityAPIEventInstances           UnityAPITypesInstances = UnityAPITypesPrefix + "/event/instances?compact=true"
)

func (_u UnityAPITypesInstances) NewRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest("GET", string(_u), nil)
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
	if fields == nil {
		return
	}
	if req.URL.RawQuery != "" {
		req.URL.RawQuery += "&"
	}
	req.URL.RawQuery += "fields=" + strings.Join(fields, ",")
}

func (_u UnityAPITypesInstances) WithFilter(filter string, req *http.Request) {
	if filter == "" {
		return
	}
	if req.URL.RawQuery != "" {
		req.URL.RawQuery += "&"
	}
	req.URL.RawQuery += "filter=" + filter
}
