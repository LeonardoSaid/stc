package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	errorHandler "github.com/leonardosaid/stc/accounts/pkg/stc-sdk/session/error"
)

type Requests interface {
	Request(requestData interface{}, uri, method string, expectedStatus int) ([]byte, error)
}

type requests struct {
	httpClient *http.Client
}

func NewRequests(client *http.Client) Requests {
	return &requests{
		httpClient: client,
	}
}

func (r *requests) Request(requestData interface{}, uri, method string, expectedStatus int) ([]byte, error) {
	requestJSON, _ := json.Marshal(requestData)

	req, err := http.NewRequest(method, uri, bytes.NewReader(requestJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if expectedStatus != res.StatusCode {
		responseError := errorHandler.ResponseError{
			Address:   uri,
			Status:    res.StatusCode,
			ErrorBody: string(body),
		}
		return nil, &responseError
	}

	return body, nil
}
