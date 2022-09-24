package main

import (
	"net/http"
	"net/url"
	"strings"
)

type TokenParams struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	IdToken      string `json:"id_token"`
}

func requestTokenEndpoint(endpoint string) (res *http.Response, e error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("redirect_uri", "https://example.com")
	values.Add("code", "SxlOBeZQ")

	req, e := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(values.Encode()))
	if e != nil {
		return nil, e
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic SlAV32hkKG")

	client := &http.Client{}
	res, e = client.Do(req)
	if e != nil {
		return nil, e
	}

	return res, nil
}
