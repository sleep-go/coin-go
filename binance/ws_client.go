package binance

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sleep-go/coin-go/binance/consts"
)

type messageHandler func(messageType int, msg []byte)
type ErrorHandler func(messageType int, err error)

type Handler[T any] func(event T)
type WsClient struct {
	Endpoint   string
	IsCombined bool
	IsFast     bool // 更新速度更快： 100ms
	Timezone   string
	apiKey     string
	conn       *websocket.Conn
	dialer     *websocket.Dialer
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func NewWsClient(apiKey string, isCombined, isFast bool, baseURL ...string) *Client {
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
	return &Client{
		BaseURL:    url,
		IsCombined: isCombined,
		IsFast:     isFast,
		APIKey:     apiKey,
		conn:       nil,
		dialer: &websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		},
	}
}

func (c *Client) serve(endpoint string, handler messageHandler, exception ErrorHandler) error {
	if c.dialer == nil {
		c.dialer = websocket.DefaultDialer
	}
	conn, _, err := c.dialer.Dial(endpoint, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.SetReadLimit(655350)
	c.conn = conn
	done := make(chan struct{})
	go func() {
		defer close(done)
		go c.keepAlive(time.Minute)
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				exception(mt, err)
				log.Println("read:", err)
				return
			}
			handler(mt, message)
			log.Println("read:", mt, string(message))
		}
	}()
	return c.waitClose(done, exception)
}
func (c *Client) waitClose(done chan struct{}, exception ErrorHandler) error {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	select {
	case <-done:
		log.Println("websocket closed")
		return nil
	case <-stopCh:
		log.Println("stopCh")
		err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			exception(websocket.CloseMessage, err)
			return err
		}
	}
	return nil
}
func (c *Client) keepAlive(timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.conn.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.conn.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				return
			}
		}
	}()
}
func (c *Client) SendMessage(msg any) error {
	return c.conn.WriteJSON(msg)
}
func (c *Client) makeParams(msg map[string]any) (params map[string]any) {
	for k, v := range msg {
		params[k] = v
	}
	params["timestamp"] = time.Now().UnixMilli()
	params["apiKey"] = c.APIKey
	return params
}
func WsHandler[T any](c *Client, endpoint string, handler Handler[T], exception ErrorHandler) error {
	log.Println(endpoint)
	h := func(mt int, msg []byte) {
		event := new(T)
		err := json.Unmarshal(msg, &event)
		if err != nil {
			exception(mt, err)
			return
		}
		handler(*event)
	}
	return c.serve(endpoint, h, exception)
}
