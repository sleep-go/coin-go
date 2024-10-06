package binance

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/sleep-go/coin-go/binance/consts"
)

type WsClient struct {
	Endpoint   string
	IsCombined bool
}

func (ws *WsClient) Depth() {

}

var done chan any

func ReceiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
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
		url += "/ws"
	}

	return &WsClient{
		Endpoint:   url,
		IsCombined: isCombined,
	}
}
