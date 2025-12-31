package ws

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	EndpointPublic           = "wss://ws.okx.com:8443/ws/v5/public"
	EndpointPrivate          = "wss://ws.okx.com:8443/ws/v5/private"
	EndpointPublicSimulated  = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	EndpointPrivateSimulated = "wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999"

	EndpointPublicBusiness           = "wss://ws.okx.com:8443/ws/v5/business"
	EndpointPrivateBusiness          = "wss://ws.okx.com:8443/ws/v5/business"
	EndpointPublicSimulatedBusiness  = "wss://wspap.okx.com:8443/ws/v5/business?brokerId=9999"
	EndpointPrivateSimulatedBusiness = "wss://wspap.okx.com:8443/ws/v5/business?brokerId=9999"

	PingTimeout  = 20 * time.Second
	PingDeadline = 10 * time.Second
)

var (
	DefaultClientPublic           = NewClient(EndpointPublic)
	DefaultClientPrivate          = NewClient(EndpointPrivate)
	DefaultClientPublicSimulated  = NewClient(EndpointPublicSimulated)
	DefaultClientPrivateSimulated = NewClient(EndpointPrivateSimulated)

	DefaultClientPublicBusiness           = NewClient(EndpointPublicBusiness)
	DefaultClientPrivateBusiness          = NewClient(EndpointPrivateBusiness)
	DefaultClientPublicSimulatedBusiness  = NewClient(EndpointPublicSimulatedBusiness)
	DefaultClientPrivateSimulatedBusiness = NewClient(EndpointPrivateSimulatedBusiness)

	PingMessage = []byte("ping")
)

type OperateCallback func(*websocket.Conn) error

type Client struct {
	Endpoint string
	Dialer   *websocket.Dialer
	mu       *sync.Mutex
}

// new Client
func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
		mu:       &sync.Mutex{},
	}
}
func NewBusinessClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
		mu:       &sync.Mutex{},
	}
}

// operate
func (c *Client) Operate(operate *Operate, callback OperateCallback) error {
	conn, _, err := c.dial()
	if err != nil {
		return err
	}

	if callback != nil {
		if err := callback(conn); err != nil {
			return err
		}
	}

	if err := c.MessageOperate(conn, operate); err != nil {
		return err
	}

	if operate.Handler != nil {
		ticker := time.NewTicker(PingTimeout)
		go c.keepAlive(conn, ticker)
		go c.messageLoop(conn, operate)
	}

	return nil
}

// message operate
func (c *Client) MessageOperate(conn *websocket.Conn, operate *Operate) error {
	if operate.Request == nil {
		return nil
	}
	if err := conn.WriteJSON(operate.Request); err != nil {
		return err
	}
	if err := conn.ReadJSON(&operate.Response); err != nil {
		return err
	}
	return operate.Response.Error()
}

// loop websocket message
func (c *Client) messageLoop(conn *websocket.Conn, operate *Operate) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			operate.HandlerError(err)
			return
		}
		operate.Handler(message)
	}
}

// keep websocket alive
func (c *Client) keepAlive(conn *websocket.Conn, ticker *time.Ticker) {
	defer ticker.Stop()
	for {
		<-ticker.C
		// OKX V5 requires "ping" as a text message to receive a "pong" response
		if err := conn.WriteMessage(websocket.TextMessage, PingMessage); err != nil {
			return
		}
	}
}

// dial endpoint
func (c *Client) dial() (*websocket.Conn, *http.Response, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Dialer == nil {
		c.Dialer = websocket.DefaultDialer
	}
	return c.Dialer.Dial(c.Endpoint, nil)
}
