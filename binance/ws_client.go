package binance

import (
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/sleep-go/coin-go/binance/consts"
)

type MessageHandler func(messageType int, msg []byte)
type ErrorHandler func(messageType int, err error)

type WsClient struct {
	Endpoint       string
	IsCombined     bool
	MessageHandler func(messageType int, msg []byte)
	ErrorHandler   func(messageType int, err error)
	conn           *websocket.Conn
}

func NewWsClient(isCombined bool, baseURL ...string) *WsClient {
	// 将默认基本 URL 设置为生产 WS URL
	url := consts.WS_STREAM

	if len(baseURL) > 0 {
		url = baseURL[0]
	}

	// 根据客户端是否用于组合流附加到 baseURL
	if isCombined {
		url += "/stream?streams="
	} else {
		url += "/ws/"
	}
	return &WsClient{
		Endpoint:   url,
		IsCombined: isCombined,
	}
}

func (c *WsClient) Connection() *WsClient {
	conn, _, err := websocket.DefaultDialer.Dial(c.Endpoint, nil)
	if err != nil {
		panic(err)
	}
	c.conn = conn
	return c
}
func (c *WsClient) WsDepth(symbols []string) (*WsClient, error) {
	endpoint := c.Endpoint
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		panic(err)
	}
	c.conn = conn
	mt, msg, err := c.conn.ReadMessage()
	if err != nil {
		c.ErrorHandler(mt, err)
		return nil, err
	}
	c.MessageHandler(mt, msg)
	return c, nil
}
