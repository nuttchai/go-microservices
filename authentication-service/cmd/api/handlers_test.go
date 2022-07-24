package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(res *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Test_Authenticate(t *testing.T) {
	jsonToReturn := `
		{
			"error": false,
			"message": "some message"
		}
	`

	// NOTE: we need to mock http client for logger function to get the response we want (status accepted, 200)
	client := NewTestClient(func(res *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	testApp.Client = client

	postBody := map[string]interface{}{
		"email":    "me@here.com",
		"password": "verysecret",
	}

	body, _ := json.Marshal(postBody)

	req, _ := http.NewRequest("POST", "/authenticate", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handlers := http.HandlerFunc(testApp.Authenticate)

	handlers.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("expected http.StatusAccepted but got %d", rr.Code)
	}
}
