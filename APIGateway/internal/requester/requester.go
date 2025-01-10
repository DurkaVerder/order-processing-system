package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Requester is an interface for sending requests
type Requester interface {
	SendRequest(url, method string, something interface{}) (*http.Response, error)
	UnmarshalResponse(resp *http.Response, something interface{}) error
}

// RequestManager is a struct that implements Requester interface
type RequestManager struct {
	httpClient *http.Client
}

// NewRequestManager is a constructor for RequestManager
func NewRequestManager() *RequestManager {
	return &RequestManager{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// RequestData sends a request to the url with the method and the data
func (rm *RequestManager) SendRequest(url, method string, something interface{}) (*http.Response, error) {
	jsonData, err := rm.dataMarshal(something)
	if err != nil {
		return nil, err
	}

	resp, err := rm.request(url, method, bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp, nil
}

// dataMarshal marshals the data to json
func (rm *RequestManager) dataMarshal(something interface{}) ([]byte, error) {
	if something == nil {
		return []byte{}, nil
	}
	return json.Marshal(something)
}

func (rm *RequestManager) UnmarshalResponse(resp *http.Response, something interface{}) error {
	if resp == nil {
		return errors.New("response is nil")
	}

	err := json.NewDecoder(resp.Body).Decode(something)
	if err != nil {
		return err
	}

	return nil
}

// request sends a request to the url with the method and the data
func (rm *RequestManager) request(url, method string, data *bytes.Buffer) (*http.Response, error) {

	req, err := rm.initRequest(url, method, data)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	resp, err = rm.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// initRequest initializes a request with the url, method and data
func (rm *RequestManager) initRequest(url, method string, data *bytes.Buffer) (*http.Request, error) {
	var err error
	var req *http.Request
	if method == "GET" {
		req, err = http.NewRequest("GET", url, nil)
	} else if method == "POST" {
		req, err = http.NewRequest("POST", url, data)
	} else if method == "PUT" {
		req, err = http.NewRequest("PUT", url, data)
	} else if method == "DELETE" {
		req, err = http.NewRequest("DELETE", url, nil)
	} else {
		err = errors.New("unknown method")
	}
	if err != nil {
		return nil, err
	}

	if rm.isSendMethod(method) {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// isSendMethod checks if the method is POST or PUT
func (rm *RequestManager) isSendMethod(method string) bool {
	return method == "POST" || method == "PUT"
}
