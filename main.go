package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func main() {
	var tokenParams TokenParams
	endpoint := "https://auth.login.yahoo.co.jp/yconnect/v2/token"
	res, e := requestTokenEndpoint(endpoint)
	if e != nil {
		log.Fatalln(e)
	}

	body, e := io.ReadAll(res.Body)
	if e != nil {
		log.Fatalln(e)
	}
	e = json.Unmarshal([]byte(body), &tokenParams)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Printf("%+v\n", tokenParams)
}
