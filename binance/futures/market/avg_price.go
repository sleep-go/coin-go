package market

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type AvgPrice interface {
	Call(ctx context.Context) (body *AvgPriceResponse, err error)
}
type avgPriceRequest struct {
	*binance.Client
	symbol string
}

func (a *avgPriceRequest) SetSymbol(symbol string) *avgPriceRequest {
	a.symbol = symbol
	return a
}

// NewAvgPrice 当前平均价格
// 获取交易对的当前平均价格
func NewAvgPrice(client *binance.Client, symbol string) AvgPrice {
	return &avgPriceRequest{Client: client, symbol: symbol}
}

type AvgPriceResponse struct {
	Mins      int    `json:"mins"`
	Price     string `json:"price"`
	CloseTime int64  `json:"closeTime"`
}

// Call 当前平均价格
func (a *avgPriceRequest) Call(ctx context.Context) (body *AvgPriceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketAvgPrice,
	}
	req.SetParam("symbol", a.symbol)
	resp, err := a.Do(ctx, req)
	if err != nil {
		a.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*AvgPriceResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

type StreamAvgPriceEvent struct {
	Stream string          `json:"stream"`
	Data   WsAvgPriceEvent `json:"data"`
}
type WsAvgPriceEvent struct {
	Event    string `json:"e"`
	Time     int64  `json:"E"`
	Symbol   string `json:"s"`
	Interval string `json:"i"`
	AvgPrice string `json:"w"`
	EndTime  int64  `json:"T"`
}

// NewStreamAvgPrice 平均价格
// 平均价格流推送在固定时间间隔内的平均价格变动。
//
// Stream 名称: <symbol>@avgPrice
//
// 更新速度: 1000ms
func NewStreamAvgPrice(c *binance.Client, symbols []string, handler binance.Handler[StreamAvgPriceEvent], exception binance.ErrorHandler) error {
	return avgPrice(c, symbols, handler, exception)
}

// NewWsAvgPrice 平均价格
// 平均价格流推送在固定时间间隔内的平均价格变动。
//
// Stream 名称: <symbol>@avgPrice
//
// 更新速度: 1000ms
func NewWsAvgPrice(c *binance.Client, symbols []string, handler binance.Handler[WsAvgPriceEvent], exception binance.ErrorHandler) error {
	return avgPrice(c, symbols, handler, exception)
}
func avgPrice[T WsAvgPriceEvent | StreamAvgPriceEvent](c *binance.Client, symbols []string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.BaseURL
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@avgPrice", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return binance.WsHandler(c, endpoint, handler, exception)
}

// ****************************** Websocket Api *******************************

type WsApiAvgPrice interface {
	binance.WsApi[*WsApiAvgPriceResponse]
	SetSymbol(symbol string) *avgPriceRequest
}
type WsApiAvgPriceResponse struct {
	binance.WsApiResponse
	Result *AvgPriceResponse `json:"result"`
}

// NewWsApiAvgPrice 获取交易对的当前平均价格
func NewWsApiAvgPrice(c *binance.Client) WsApiAvgPrice {
	return &avgPriceRequest{Client: c}
}

func (a *avgPriceRequest) Send(ctx context.Context) (*WsApiAvgPriceResponse, error) {
	req := &binance.Request{Path: "avgPrice"}
	req.SetParam("symbol", a.symbol)
	return binance.WsApiHandler[*WsApiAvgPriceResponse](ctx, a.Client, req)
}
