package requester

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
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
func NewRequestManager(client *http.Client) *RequestManager {
	return &RequestManager{
		httpClient: client,
	}
}

// RequestData sends a request to the url with the method and the data
func (rm *RequestManager) SendRequest(url, method string, something interface{}) (*http.Response, error) {
	jsonData, err := rm.dataMarshal(something)
	if err != nil {
		return nil, err
	}

	resp, err := rm.request(url, method, bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
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

// UnmarshalResponse unmarshals the response to the something
func (rm *RequestManager) UnmarshalResponse(resp *http.Response, something interface{}) error {
	if resp == nil {
		return errors.New("response is nil")
	}

	if resp.Body == nil {
		return errors.New("response body is nil")
	}
	defer resp.Body.Close()

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
	var req *http.Request
	var err error

	switch method {
	case http.MethodGet, http.MethodDelete:
		req, err = http.NewRequest(method, url, nil)
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		req, err = http.NewRequest(method, url, data)
	default:
		return nil, errors.New("unsupported method")
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
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		return true
	default:
		return false
	}
}
