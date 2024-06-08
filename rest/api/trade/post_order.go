package trade

import (
	"github.com/google/uuid"
	"github.com/kecheon/go-okx/rest/api"
)

func NewPostOrder(param *PostOrderParam) (api.IRequest, api.IResponse) {
	param.ClOrdId = "0f597be3756fBCDE" + uuid.New().String()[:8]
	param.Tag = "0f597be3756fBCDE"
	return &api.Request{
		Path:   "/api/v5/trade/order",
		Method: api.MethodPost,
		Param:  param,
	}, &PostOrderResponse{}
}

type PostOrderParam struct {
	InstId     string `json:"instId"`
	TdMode     string `json:"tdMode"`
	Ccy        string `json:"ccy,omitempty"`
	ClOrdId    string `json:"clOrdId,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Side       string `json:"side"`
	PosSide    string `json:"posSide,omitempty"`
	OrdType    string `json:"ordType"`
	Sz         string `json:"sz"`
	Px         string `json:"px,omitempty"`
	ReduceOnly bool   `json:"reduceOnly,omitempty"`
	TgtCcy     string `json:"tgtCcy,omitempty"`
}

type PostOrderResponse struct {
	api.Response
	Data []PostOrder `json:"data"`
}

type PostOrder struct {
	OrdId       string `json:"ordId"`
	ClOrdId     string `json:"clOrdId"`
	Tag         string `json:"tag"`
	SCode       string `json:"sCode"`
	SMsg        string `json:"sMsg"`
	AlgoId      string `json:"algoId"`
	AlgoClOrdId string `json:"algoClOrdId"`
}

type PostAlgoOrderParam struct {
	InstId      string `json:"instId"`
	TdMode      string `json:"tdMode"`
	Ccy         string `json:"ccy,omitempty"`
	ClOrdId     string `json:"clOrdId,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Side        string `json:"side"`
	PosSide     string `json:"posSide,omitempty"`
	OrdType     string `json:"ordType"`
	Sz          string `json:"sz"`
	Px          string `json:"px,omitempty"`
	ReduceOnly  bool   `json:"reduceOnly,omitempty"`
	TgtCcy      string `json:"tgtCcy,omitempty"`
	TriggerPx   string `json:"triggerPx"`
	OrderPx     string `json:"orderPx"`
	TpTriggerPx string `json:"tpTriggerPx,omitempty"`
	SlTriggerPx string `json:"slTriggerPx,omitempty"`
	SlOrdPx     string `json:"slOrdPx,omitempty"`
}

func NewPostAlgoOrder(param *PostAlgoOrderParam) (api.IRequest, api.IResponse) {
	param.ClOrdId = "0f597be3756fBCDE" + uuid.New().String()[:8]
	param.Tag = "0f597be3756fBCDE"
	return &api.Request{
		Path:   "/api/v5/trade/order-algo",
		Method: api.MethodPost,
		Param:  param,
	}, &PostOrderResponse{}
}

/************** Algo order Request examples *************************
# Place Take Profit / Stop Loss Order
POST /api/v5/trade/order-algo
body
{
    "instId":"BTC-USDT",
    "tdMode":"cross",
    "side":"buy",
    "ordType":"conditional",
    "sz":"2",
    "tpTriggerPx":"15",
    "tpOrdPx":"18"
}

# Place Trigger Order
POST /api/v5/trade/order-algo
body
{
    "instId": "BTC-USDT-SWAP",
    "side": "buy",
    "tdMode": "cross",
    "posSide": "net",
    "sz": "1",
    "ordType": "trigger",
    "triggerPx": "25920",
    "triggerPxType": "last",
    "orderPx": "-1",
    "attachAlgoOrds": [{
        "attachAlgoClOrdId": "",
        "slTriggerPx": "100",
        "slOrdPx": "600",
        "tpTriggerPx": "25921",
        "tpOrdPx": "2001"
    }]
}

# Place Trailing Stop Order
POST /api/v5/trade/order-algo
body
{
    "instId": "BTC-USDT-SWAP",
    "tdMode": "cross",
    "side": "buy",
    "ordType": "move_order_stop",
    "sz": "10",
    "posSide": "net",
    "callbackRatio": "0.05",
    "reduceOnly": true
}

# Place TWAP Order
POST /api/v5/trade/order-algo
body
{
    "instId": "BTC-USDT-SWAP",
    "tdMode": "cross",
    "side": "buy",
    "ordType": "twap",
    "sz": "10",
    "posSide": "net",
    "szLimit": "10",
    "pxLimit": "100",
    "timeInterval": "10",
    "pxSpread": "10"
}
*********************************************/

/************* response ***************
{
  "code": "0",
  "msg": "",
  "data": [
    {
      "algoId": "12345689",
      "clOrdId": "",
      "algoClOrdId": "",
      "sCode": "0",
      "sMsg": ""
    }
  ]
}
******************************************/
