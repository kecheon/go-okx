package public

import (
	"github.com/kecheon/go-okx/ws"
)

type Public struct {
	C *ws.Client
}

func NewPublic(simulated bool) *Public {
	public := &Public{
		C: ws.DefaultClientPublic,
	}
	if simulated {
		public.C = ws.DefaultClientPublicSimulated
	}
	return public
}

func NewPublicBusiness(simulated bool) *Public {
	public := &Public{
		C: ws.DefaultClientPublicBusiness,
	}
	if simulated {
		public.C = ws.DefaultClientPublicSimulatedBusiness
	}
	return public
}

// subscribe
func (p *Public) Subscribe(args interface{}, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(subscribe, nil)
}
