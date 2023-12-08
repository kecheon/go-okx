package main

import (
	"log"

	"github.com/kecheon/go-okx/examples/rest"
	"github.com/kecheon/go-okx/rest/api/account"
)

func main() {
	param := &account.GetPositionsParam{}
	req, resp := account.NewGetPositions(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetPositionsResponse))
}
