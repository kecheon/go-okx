package main

import (
	"log"

	"github.com/iaping/go-okx/examples"
	"github.com/iaping/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetBillsParam{}
	req, resp := subaccount.NewGetBills(param)
	if err := examples.Client.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*subaccount.GetBillsResponse))
}