package main

import (
	"log"
)

type TokenEndPointParams struct {
	grantType   string
	redirectURI string
	code        string
}

func main() {
	endpoint := "https://auth.login.yahoo.co.jp/yconnect/v2/token"
	e := requestTokenEndpoint(endpoint)
	if e != nil {
		log.Fatalln(e)
	}
}
