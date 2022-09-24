package main

import (
	"net/http"
	"testing"
)

type OverrideRequest struct {
	Transport http.RoundTripper
	Overrides map[string]string
}

func TestRequestTokenEndpoint(t *testing.T) {
	url := "https://test-example.com"
	e := requestTokenEndpoint(url)
	if e != nil {
		t.Error(e)
	}
}
