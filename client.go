package gounity

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
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
	AccessTime time.Time

	client *http.Client
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

func (_c *UnisphereClient) send(req *http.Request) ([]byte, error) {
	resp, err := _c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = _c.checkHttpCode(resp.StatusCode, body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (_c *UnisphereClient) checkHttpCode(code int, body []byte) error {
	var err error
	switch code {
	case http.StatusUnauthorized:
		_c.logined = false
		return errors.New("Unauthorized")
	case http.StatusUnprocessableEntity:
		var data StatusUnprocessableEntity
		err = json.Unmarshal(body, &data)
		if err != nil {
			return err
		}
		return errors.New(data.Error.Messages[0].EnUS)
	default:
		return nil
	}

}
