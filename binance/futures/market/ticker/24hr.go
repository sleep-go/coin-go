package ticker

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Hr24 interface {
	Call(ctx context.Context) (body *hr24Response, err error)
	CallAll(ctx context.Context) (body []*hr24Response, err error)
}

type hr24Request struct {
	*binance.Client
	symbol string
}

type hr24Response struct {
	Symbol             string `json:"symbol"`             // 交易对
	PriceChange        string `json:"priceChange"`        // 24小时价格变动
	PriceChangePercent string `json:"priceChangePercent"` // 24小时价格变动百分比
	WeightedAvgPrice   string `json:"weightedAvgPrice"`   // 加权平均价
	LastPrice          string `json:"lastPrice"`          // 最近一次成交价
	LastQty            string `json:"lastQty"`            // 最近一次成交额
	OpenPrice          string `json:"openPrice"`          // 24小时内第一次成交的价格
	HighPrice          string `json:"highPrice"`          // 24小时最高价
	LowPrice           string `json:"lowPrice"`           // 24小时最低价
	Volume             string `json:"volume"`             // 24小时成交量
	QuoteVolume        string `json:"quoteVolume"`        // 24小时成交额
	OpenTime           int64  `json:"openTime"`           // 24小时内，第一笔交易的发生时间
	CloseTime          int64  `json:"closeTime"`          // 24小时内，最后一笔交易的发生时间
	FirstId            int    `json:"firstId"`            // 首笔成交id
	LastId             int    `json:"lastId"`             // 末笔成交id
	Count              int    `json:"count"`              // 成交笔数
}

// NewHr24 24hr价格变动情况
func NewHr24(client *binance.Client, symbol string) Hr24 {
	return &hr24Request{
		Client: client,
		symbol: symbol,
	}
}

// Call 24hr价格变动情况
// 请注意，不携带symbol参数会返回全部交易对数据，不仅数据庞大，而且权重极高
func (hr *hr24Request) Call(ctx context.Context) (body *hr24Response, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTicker24Hr,
	}
	req.SetOptionalParam("symbol", hr.symbol)
	res, err := hr.Do(ctx, req)
	if err != nil {
		hr.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*hr24Response](res)
}

// CallAll 24hr价格变动情况
// 请注意，不携带symbol参数会返回全部交易对数据，不仅数据庞大，而且权重极高
func (hr *hr24Request) CallAll(ctx context.Context) (body []*hr24Response, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTicker24Hr,
	}
	res, err := hr.Do(ctx, req)
	if err != nil {
		hr.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*hr24Response](res)
}

// ****************************** Websocket Api *******************************

type WsApiHr24 interface {
	binance.WsApi[*WsApiHr24Response]
	SetSymbol(symbol string) WsApiHr24
}
type WsApiHr24Response struct {
	binance.WsApiResponse
	Result []*hr24Response `json:"result"`
}

// NewWsApiHr24 24hr 价格变动情况
// 24 小时滚动窗口价格变动数据。 如果您需要持续监控交易统计，请考虑使用 WebSocket Streams:
//
// <symbol>@ticker 或者 !ticker@arr
// <symbol>@miniTicker 或者 !miniTicker@arr
// 如果你想用不同的窗口数量，可以用 ticker 请求。
func NewWsApiHr24(c *binance.Client) WsApiHr24 {
	return &hr24Request{
		Client: c,
	}
}

// SetSymbol 如果未指定交易对，则返回有关当前在交易所交易的所有交易对的信息。
func (hr *hr24Request) SetSymbol(symbol string) WsApiHr24 {
	hr.symbol = symbol
	return hr
}

// Send 如果未指定交易对，则返回有关当前在交易所交易的所有交易对的信息。
func (hr *hr24Request) Send(ctx context.Context) (*WsApiHr24Response, error) {
	req := &binance.Request{Path: "ticker.24hr"}
	req.SetOptionalParam("symbol", hr.symbol)
	return binance.WsApiHandler[*WsApiHr24Response](ctx, hr.Client, req)
}
