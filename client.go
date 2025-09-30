package gounity

import (
	"crypto/tls"
	"encoding/base64"
	"net/http"
	"net/http/cookiejar"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type UnisphereClient struct {
	endpoint   string
	auth       string
	token      string
	logined    bool
	accessTime time.Time

	client *http.Client
}

func NewClient(endpoint string, username string, password string, insecure bool) *UnisphereClient {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	return &UnisphereClient{
		endpoint: endpoint,
		auth:     auth,
		token:    "",
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		},
	}
}

func (_c *UnisphereClient) addHeader(method string, req *http.Request) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-EMC-REST-CLIENT", "true")

	switch method {
	case "GET":
		if !_c.logined {
			if time.Since(_c.accessTime).Minutes() < 1 {
				return
			}
			req.Header.Add("Authorization", "Basic "+_c.auth)
			_c.client.Jar, _ = cookiejar.New(nil)
		}
	case "POST":
		if !_c.logined {
			return
		}
		req.Header.Add("EMC-CSRF-TOKEN", _c.token)
	}
}

type StatusUnprocessableEntity struct {
	Error struct {
		ErrorCode      int `json:"errorCode"`
		HttpStatusCode int `json:"httpStatusCode"`
		Messages       []struct {
			EnUS string `json:"en-US"`
		} `json:"messages"`
		Created time.Time `json:"created"`
	} `json:"error"`
}
