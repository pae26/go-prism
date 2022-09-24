package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func requestTokenEndpoint(endpoint string) error {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("redirect_uri", "https://example.com")
	values.Add("code", "SxlOBeZQ")

	req, e := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(values.Encode()))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic SlAV32hkKG")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()

	body, e := io.ReadAll(res.Body)
	if e != nil {
		return e
	}
	fmt.Printf("%+v\n", string(body))

	return e
}
