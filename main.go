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
	url := "https://auth.login.yahoo.co.jp/yconnect/v2/token"
	e := requestTokenEndpoint(url)
	if e != nil {
		log.Fatalln(e)
	}
}
