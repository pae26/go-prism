package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func requestTokenEndpoint(url string) error {
	reqParams := &TokenEndPointParams{
		grantType:   "authorization_code",
		redirectURI: "https://example.com",
		code:        "SxlOBeZQ",
	}

	reqParamsJSON, e := json.Marshal(reqParams)
	if e != nil {
		return e
	}

	req, e := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqParamsJSON))
	if e != nil {
		return e
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic SlAV32hkKG")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return e
	}

	fmt.Printf("%+v", res.Body)

	return e
}
