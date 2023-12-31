package sheetdb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
)

// Call calls the API and returns the response.
type APICaller interface {
	Call(
		ctx context.Context,
		method string,
		url string,
		opt Option,
		body interface{},
		response interface{},
	) error
}

// APICallerImpl is an implementation of APICaller that uses an
// HTTP client to call the API.
type APICallerImpl struct {
	HTTPClient *http.Client
}

func (a *APICallerImpl) Call(ctx context.Context, method string, url string, opt Option, body interface{}, response interface{}) error {
	requestBody := []byte("")

	var request *http.Request
	var err error

	isBodyNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())

	if !isBodyNil {
		requestBody, err = json.Marshal(body)

		if err != nil {
			return err
		}
	}

	request, err = http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(requestBody),
	)

	if err != nil {
		return err
	}

	request.SetBasicAuth(opt.Username, opt.Password)
	request.Header.Set("Content-Type", "application/json")

	return a.do(request, response)
}

func (a *APICallerImpl) do(request *http.Request, response interface{}) error {
	httpResponse, err := a.HTTPClient.Do(request)

	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	responseBody, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		return err
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		err := errors.New("HTTP ERROR: " + string(responseBody))

		return err
	}

	if err := json.Unmarshal(responseBody, &response); err != nil {
		return err
	}

	return nil
}
