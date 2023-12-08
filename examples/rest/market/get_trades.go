package main

import (
	"log"

	"github.com/kecheon/go-okx/examples/rest"
	"github.com/kecheon/go-okx/rest/api/market"
)

func main() {
	param := &market.GetTradesParam{
		InstId: "ETH-USDT",
	}
	req, resp := market.NewGetTrades(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetTradesResponse))
}
