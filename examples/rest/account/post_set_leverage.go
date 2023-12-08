package main

import (
	"log"

	"github.com/kecheon/go-okx/examples/rest"
	"github.com/kecheon/go-okx/rest/api"
	"github.com/kecheon/go-okx/rest/api/account"
)

func main() {
	param := &account.PostSetLeverageParam{
		InstId:  "BTC-USDT",
		Lever:   "5",
		MgnMode: api.MgnModeCross,
	}
	req, resp := account.NewPostSetLeverage(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.PostSetLeverageResponse))
}
