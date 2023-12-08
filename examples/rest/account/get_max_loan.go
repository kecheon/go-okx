package main

import (
	"log"

	"github.com/kecheon/go-okx/examples/rest"
	"github.com/kecheon/go-okx/rest/api"
	"github.com/kecheon/go-okx/rest/api/account"
)

func main() {
	param := &account.GetMaxLoanParam{
		InstId:  "BTC-USDT",
		MgnMode: api.MgnModeCross,
		MgnCcy:  "USDT",
	}
	req, resp := account.NewGetMaxLoan(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetMaxLoanResponse))
}
