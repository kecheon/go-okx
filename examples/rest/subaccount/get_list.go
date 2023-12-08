package main

import (
	"log"

	"github.com/kecheon/go-okx/examples/rest"
	"github.com/kecheon/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetListParam{}
	req, resp := subaccount.NewGetList(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.GetListResponse))
}
