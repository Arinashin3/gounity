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
)

/*
	Request

Endpoint = Unity Console URL (ex: https://127.0.0.1:443)
Params = (ex: Key->"fields", Value->"id,name")
Logined = if you want login, set true
Auth = when you want login, that will use this auth(encoded base64)
*/
type Request struct {
	Endpoint string
	Params   map[string]string
	Logined  bool
	Auth     string
}

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
