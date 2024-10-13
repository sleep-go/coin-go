package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sleep-go/coin-go/pkg/errors"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sleep-go/coin-go/binance/consts"
)

type WsApi[T any] interface {
	Send(ctx context.Context) (T, error)
}
type WsApiResponse struct {
	Id         string         `json:"id"`
	Status     int            `json:"status"`
	Error      *errors.Status `json:"error"`
	RateLimits []RateLimits   `json:"rateLimits"`
}

func NewWsApiHMACClient(apiKey, secretKey string, baseURL ...string) *Client {
	// 将默认基本 URL 设置为生产 WS URL
	url := consts.WS_API2
	if len(baseURL) > 0 {
		url = baseURL[0]
	}
	prefix := "[INFO] "
	switch LogLevel {
	case os.Stderr:
		prefix = "[ERROR] "
	case os.Stdout:
		prefix = "[INFO] "
	}
	// 根据客户端是否用于组合流附加到 baseURL
	return &Client{
		BaseURL:   url,
		APIKey:    apiKey,
		SecretKey: secretKey,
		Logger:    log.New(LogLevel, prefix, log.LstdFlags),
		dialer: &websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		},
	}
}
func NewWsApiRSAClient(apiKey, privateKeyPath string, baseURL ...string) *Client {
	// 将默认基本 URL 设置为生产 WS URL
	url := consts.WS_API2
	if len(baseURL) > 0 {
		url = baseURL[0]
	}
	prefix := "[INFO] "
	switch LogLevel {
	case os.Stderr:
		prefix = "[ERROR] "
	case os.Stdout:
		prefix = "[INFO] "
	}
	privateKey, err := loadRsaPrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	// 根据客户端是否用于组合流附加到 baseURL
	return &Client{
		BaseURL:    url,
		APIKey:     apiKey,
		PrivateKey: privateKey,
		Logger:     log.New(LogLevel, prefix, log.LstdFlags),
		dialer: &websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		},
	}
}
func NewWsApiED25519Client(apiKey, privateKeyPath string, baseURL ...string) *Client {
	// 将默认基本 URL 设置为生产 WS URL
	url := consts.WS_API2
	if len(baseURL) > 0 {
		url = baseURL[0]
	}
	prefix := "[INFO] "
	switch LogLevel {
	case os.Stderr:
		prefix = "[ERROR] "
	case os.Stdout:
		prefix = "[INFO] "
	}
	privateKey, err := loadED25519PrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	cli := &Client{
		BaseURL:    url,
		APIKey:     apiKey,
		PrivateKey: privateKey,
		Logger:     log.New(LogLevel, prefix, log.LstdFlags),
		dialer: &websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		},
	}
	cli.connect()
	return cli
}
func (c *Client) connect() error {
	if c.dialer == nil {
		c.dialer = websocket.DefaultDialer
	}
	conn, _, err := c.dialer.Dial(c.BaseURL, nil)
	if err != nil {
		return err
	}
	conn.SetReadLimit(655350)
	c.conn = conn
	c.ReqResponseMap = make(map[string]chan []byte)
	go func() {
		for {
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("Error reading:", err)
				return
			}
			var response WsApiResponse
			err = json.Unmarshal(message, &response)
			if err != nil {
				log.Println("Error unmarshaling:", err)
				return
			}
			// Send the message to the corresponding request
			if channel, ok := c.ReqResponseMap[response.Id]; ok {
				channel <- message
			}
		}
	}()
	return nil
}
func (c *Client) sendWsApiMsg(ctx context.Context, r *Request) (res []byte, err error) {
	//获取 query url
	queryString := r.query.Encode()
	r.SetParam("apiKey", c.APIKey)
	if r.needSign {
		r.SetOptionalParam("recvWindow", c.TimeOffset)
		r.SetParam("timestamp", time.Now().UnixMilli())
		//获取 query url
		queryString = r.query.Encode()
		//设置签名参数
		raw := fmt.Sprintf("%s", queryString)
		if c.SecretKey != "" {
			r.SetParam("signature", signPayload(raw, c.SecretKey))
		} else if c.PrivateKey != nil {
			r.SetParam("signature", signPayload(raw, c.PrivateKey))
		} else {
			c.Println("signature is empty")
		}
	}
	params := make(map[string]any)
	for k, v := range r.query {
		params[k] = v[0]
	}
	msg := &WsReqMsg{
		Id:     uuid.New().String(),
		Method: r.Path,
		Params: params,
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	c.Debugf("%s", marshal)
	err = c.conn.WriteJSON(msg)
	if err != nil {
		return nil, err
	}
	defer delete(c.ReqResponseMap, msg.Id)
	messageCh := make(chan []byte)
	c.ReqResponseMap[msg.Id] = messageCh
	select {
	case response := <-messageCh:
		return response, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
func WsApiHandler[T any](ctx context.Context, c *Client, r *Request) (res T, err error) {
	msg, err := c.sendWsApiMsg(ctx, r)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(msg, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
