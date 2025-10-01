package api

import (
	"net/http"
	"strings"
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
