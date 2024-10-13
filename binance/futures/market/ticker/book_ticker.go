package ticker

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type BookTicker interface {
	Call(ctx context.Context, symbol string) (body *bookTickerResponse, err error)
	CallAll(ctx context.Context) (body []*bookTickerResponse, err error)
}

type bookTickerRequest struct {
	*binance.Client
}

type bookTickerResponse struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int64  `json:"time"`
}

func NewBookTicker(client *binance.Client) BookTicker {
	return &bookTickerRequest{Client: client}
}
func (b *bookTickerRequest) Call(ctx context.Context, symbol string) (body *bookTickerResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTickerBookTicker,
	}
	req.SetParam("symbol", symbol)
	res, err := b.Do(ctx, req)
	if err != nil {
		b.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*bookTickerResponse](res)
}
func (b *bookTickerRequest) CallAll(ctx context.Context) (body []*bookTickerResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTickerBookTicker,
	}
	res, err := b.Do(ctx, req)
	if err != nil {
		b.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*bookTickerResponse](res)
}

// ****************************** Websocket 行情推送 *******************************

type StreamBookTickerEvent struct {
	Stream string            `json:"stream"`
	Data   WsBookTickerEvent `json:"data"`
}
type WsBookTickerEvent struct {
	UpdateID     int64  `json:"u"`
	Symbol       string `json:"s"`
	BestBidPrice string `json:"b"`
	BestBidQty   string `json:"B"`
	BestAskPrice string `json:"a"`
	BestAskQty   string `json:"A"`
}

// NewWsBookTicker 按Symbol的最优挂单信息
// 实时推送指定交易对最优挂单信息 多个 <symbol>@bookTicker 可以订阅在一个WebSocket连接上
//
// Stream 名称: <symbol>@bookTicker
//
// 更新速度: 实时
func NewWsBookTicker(c *binance.Client, symbols []string, handler binance.Handler[WsBookTickerEvent], exception binance.ErrorHandler) error {
	return bookTicker(c, symbols, handler, exception)
}

// NewStreamBookTicker 按Symbol的最优挂单信息
// 实时推送指定交易对最优挂单信息 多个 <symbol>@bookTicker 可以订阅在一个WebSocket连接上
//
// Stream 名称: <symbol>@bookTicker
//
// 更新速度: 实时
func NewStreamBookTicker(c *binance.Client, symbols []string, handler binance.Handler[StreamBookTickerEvent], exception binance.ErrorHandler) error {
	return bookTicker(c, symbols, handler, exception)
}
func bookTicker[T WsBookTickerEvent | StreamBookTickerEvent](c *binance.Client, symbols []string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.BaseURL
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@bookTicker", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return binance.WsHandler(c, endpoint, handler, exception)
}

// ****************************** Websocket Api *******************************

type WsApiBookTicker interface {
	binance.WsApi[*WsApiBookTickerResponse]
}
type WsApiBookTickerResponse struct {
	binance.WsApiResponse
	Result []*bookTickerResponse `json:"result"`
}

// NewWsApiBookTicker 当前最优挂单
// 在订单薄获取当前最优价格和数量。
//
// 如果您需要访问实时订单薄 ticker 更新，请考虑使用 WebSocket Streams:
//
// <symbol>@bookTicker
func NewWsApiBookTicker(c *binance.Client) WsApiBookTicker {
	return &bookTickerRequest{Client: c}
}

func (b *bookTickerRequest) Send(ctx context.Context) (*WsApiBookTickerResponse, error) {
	req := &binance.Request{Path: "ticker.book"}
	return binance.WsApiHandler[*WsApiBookTickerResponse](ctx, b.Client, req)
}
