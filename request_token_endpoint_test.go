package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestRequestTokenEndpoint(t *testing.T) {
	var tokenParams TokenParams
	url := "http://localhost:4010/yahoo-token"
	expectedAccessToken := "SlAV32hkKG"

	t.Run("success test", func(t *testing.T) {
		res, e := requestTokenEndpoint(url)
		if e != nil {
			t.Error(e)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("mock server return status code: %s", res.Status)
		}

		body, _ := io.ReadAll(res.Body)
		e = json.Unmarshal(body, &tokenParams)
		if e != nil {
			t.Error(e)
		}
		if tokenParams.AccessToken != "SlAV32hkKG" {
			t.Errorf("expected access_token is %s, but got %s", expectedAccessToken, tokenParams.AccessToken)
		}

		fmt.Printf("%+v\n", tokenParams)
	})
}
