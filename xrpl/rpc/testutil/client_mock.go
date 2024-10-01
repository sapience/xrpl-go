package testutil

import (
	"bytes"
	"io"
	"net/http"
)

type JsonRpcMockClient struct {
	DoFunc       func(req *http.Request) (*http.Response, error)
	Spy          *http.Request
	RequestCount int
}

func (m *JsonRpcMockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	// just in case you want default correct return value
	return &http.Response{}, nil
}

func MockResponse(resString string, statusCode int, m *JsonRpcMockClient) func(req *http.Request) (*http.Response, error) {
	return func(req *http.Request) (*http.Response, error) {
		m.Spy = req
		return &http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(bytes.NewReader([]byte(resString))),
		}, nil
	}
}
