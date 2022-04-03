package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectOption func(t *testing.T, response *http.Response)

func ExpectHttpStatus(statusCode int) ExpectOption {
	return func(t *testing.T, response *http.Response) {
		assert.Equal(t, statusCode, response.StatusCode)
	}
}

func NewHttpRequest(t *testing.T, method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("unable to create http request %v", err)
		t.FailNow()
	}

	request.Header.Add("Accept", "application/json")
	return request
}

func DoHttpRequest(t *testing.T, httpHandler http.Handler, request *http.Request, opts ...ExpectOption) {
	recorder := httptest.NewRecorder()
	httpHandler.ServeHTTP(recorder, request)

	result := recorder.Result()
	defer result.Body.Close()

	for _, opt := range opts {
		opt(t, result)
	}
}

func DoHttpRequestWithResponse[T any](t *testing.T, httpHandler http.Handler, request *http.Request, response T, opts ...ExpectOption) T {
	recorder := httptest.NewRecorder()
	httpHandler.ServeHTTP(recorder, request)

	result := recorder.Result()
	defer result.Body.Close()

	decoder := json.NewDecoder(result.Body)
	err := decoder.Decode(&response)
	if err != nil {
		t.Errorf("unable to successfully invoke http request %v", err)
		t.FailNow()
	}

	for _, opt := range opts {
		opt(t, result)
	}

	return response
}

func EncodeToJSON(t *testing.T, subject any) io.Reader {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(subject)
	if err != nil {
		t.Errorf("unable to encode into JSON %v", err)
		t.FailNow()
	}

	return &buf
}
