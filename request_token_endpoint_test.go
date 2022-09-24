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
	url := "http://localhost:4010/yahoo-token"
	t.Run("success test", func(t *testing.T) {
		e := requestTokenEndpoint(url)
		if e != nil {
			t.Error(e)
		}
	})
}
