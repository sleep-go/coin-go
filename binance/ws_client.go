package binance

import (
	"github.com/gorilla/websocket"
	"github.com/sleep-go/coin-go/binance/consts"
	"log"
	"os"
	"os/signal"
)

type messageHandler func(messageType int, msg []byte)
type ErrorHandler func(messageType int, err error)

type Handler[T any] func(event *T)
type WsClient struct {
	Endpoint   string
	IsCombined bool
	conn       *websocket.Conn
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

func (c *WsClient) WsServe(endpoint string, handler messageHandler, exception ErrorHandler) error {
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.SetReadLimit(655350)
	c.conn = conn
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				exception(mt, err)
				log.Println("read:", err)
				return
			}
			handler(mt, message)
		}
	}()
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	for {
		select {
		case <-done:
			log.Println("websocket closed")
			return nil
		case <-stopCh:
			log.Println("stopCh")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				exception(websocket.CloseMessage, err)
				return err
			}
		}
	}
}
