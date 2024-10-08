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
type AvgPriceRequest struct {
	*binance.Client
	symbol string
}

// NewAvgPrice 当前平均价格
func NewAvgPrice(client *binance.Client, symbol string) AvgPrice {
	return &AvgPriceRequest{Client: client, symbol: symbol}
}

type AvgPriceResponse struct {
	Mins      int    `json:"mins"`
	Price     string `json:"price"`
	CloseTime int64  `json:"closeTime"`
}

// Call 当前平均价格
func (k *AvgPriceRequest) Call(ctx context.Context) (body *AvgPriceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketAvgPrice,
	}
	req.SetParam("symbol", k.symbol)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*AvgPriceResponse](resp)
}

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
func NewStreamAvgPrice(c *binance.WsClient, symbols []string, handler binance.Handler[StreamAvgPriceEvent], exception binance.ErrorHandler) error {
	return avgPrice(c, symbols, handler, exception)
}

// NewWsAvgPrice 平均价格
// 平均价格流推送在固定时间间隔内的平均价格变动。
//
// Stream 名称: <symbol>@avgPrice
//
// 更新速度: 1000ms
func NewWsAvgPrice(c *binance.WsClient, symbols []string, handler binance.Handler[WsAvgPriceEvent], exception binance.ErrorHandler) error {
	return avgPrice(c, symbols, handler, exception)
}
func avgPrice[T WsAvgPriceEvent | StreamAvgPriceEvent](c *binance.WsClient, symbols []string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.Endpoint
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@avgPrice", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return binance.WsHandler(c, endpoint, handler, exception)
}
