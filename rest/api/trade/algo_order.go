package trade

import "github.com/kecheon/go-okx/rest/api"

func NewPostCancelAlgoOrder(param []*PostCancelAlgoOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/cancel-algos",
		Method: api.MethodPost,
		Param:  param,
	}, &PostCancelAlgoOrderResponse{}
}

type PostCancelAlgoOrderResponse struct {
	api.Response
	Data []CancelAlgoOrder `json:"data"`
}

type PostCancelAlgoOrderParam struct {
	AlgoId string `json:"algoId"`
	InstId string `json:"instId"`
}

type CancelAlgoOrder struct {
	AlgoId string `json:"algoId"`
	SCode  string `json:"sCode"`
	SMsg   string `json:"sMsg"`
}

func NewGetAlgoOrder(param *GetAlgoOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/order-algo",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}

type GetAlgoOrderParam struct {
	AlgoId      string `url:"algoId"`
	AlgoClOrdId string `url:"algoClOrdId,omitempty"`
}
