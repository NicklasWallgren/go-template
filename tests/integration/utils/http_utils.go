package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectOption func(t *testing.T, response *http.Response)

func ExpectHTTPStatus(statusCode int) ExpectOption {
	return func(t *testing.T, response *http.Response) {
		t.Helper()

		assert.Equal(t, statusCode, response.StatusCode)
	}
}

func NewHTTPRequest(t *testing.T, method string, url string, body io.Reader) *http.Request {
	t.Helper()

	request, err := http.NewRequestWithContext(context.TODO(), method, url, body)
	if err != nil {
		t.Errorf("unable to create http request %v", err)
		t.FailNow()
	}

	request.Header.Add("Accept", "application/json")

	return request
}

func DoHTTPRequest(t *testing.T, httpHandler http.Handler, request *http.Request, opts ...ExpectOption) {
	t.Helper()

	recorder := httptest.NewRecorder()
	httpHandler.ServeHTTP(recorder, request)

	result := recorder.Result()
	defer result.Body.Close() // nolint:errcheck

	for _, opt := range opts {
		opt(t, result)
	}
}

func DoHTTPRequestWithResponse[T any](
	t *testing.T, httpHandler http.Handler, request *http.Request, response T, opts ...ExpectOption,
) T {
	t.Helper()

	recorder := httptest.NewRecorder()
	httpHandler.ServeHTTP(recorder, request)

	result := recorder.Result()
	defer result.Body.Close() // nolint:errcheck

	decoder := json.NewDecoder(result.Body)
	err := decoder.Decode(&response)
	if err != nil { // nolint: wsl
		t.Errorf("unable to successfully invoke http request %v", err)
		t.FailNow()
	}

	for _, opt := range opts {
		opt(t, result)
	}

	return response
}

func EncodeToJSON(t *testing.T, subject any) io.Reader {
	t.Helper()

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(subject)
	if err != nil {
		t.Errorf("unable to encode into JSON %v", err)
		t.FailNow()
	}

	return &buf
}
