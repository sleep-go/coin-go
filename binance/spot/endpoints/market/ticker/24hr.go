package ticker

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Hr24 interface {
	Call(ctx context.Context) (body []*hr24Response, err error)
}

type hr24Request struct {
	*binance.Client
	symbols []string
	_type   enums.TickerType //可接受的参数: FULL or MINI. 如果不提供, 默认值为 FULL
}

type hr24Response struct {
	Symbol             string `json:"symbol"` // 交易对
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"` // 间隔收盘价
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`   // 间隔开盘价
	HighPrice          string `json:"highPrice"`   // 间隔最高价
	LowPrice           string `json:"lowPrice"`    // 间隔最低价
	Volume             string `json:"volume"`      // 总交易量 (base asset)
	QuoteVolume        string `json:"quoteVolume"` // 总交易量 (quote asset)
	OpenTime           int64  `json:"openTime"`    // ticker间隔的开始时间
	CloseTime          int64  `json:"closeTime"`   // ticker间隔的结束时间
	FirstId            int    `json:"firstId"`     // 统计时间内的第一笔trade id
	LastId             int    `json:"lastId"`      // 统计时间内的最后一笔trade id
	Count              int    `json:"count"`       // 统计时间内交易笔数
}

func NewHr24(client *binance.Client, symbols []string, _type enums.TickerType) Hr24 {
	return &hr24Request{
		Client:  client,
		symbols: symbols,
		_type:   _type,
	}
}

// Call 24hr价格变动情况
// 请注意，不携带symbol参数会返回全部交易对数据，不仅数据庞大，而且权重极高
func (hr *hr24Request) Call(ctx context.Context) (body []*hr24Response, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketTicker24Hr,
	}
	if len(hr.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(hr.symbols, `","`))
		req.SetParam("symbols", result)
	}
	req.SetParam("type", hr._type.String())
	res, err := hr.Do(ctx, req)
	if err != nil {
		hr.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*hr24Response](res)
}

// ****************************** Websocket Api *******************************

type WsApiHr24 interface {
	binance.WsApi[WsApiHr24Response]
	SetSymbols(symbols []string) WsApiHr24
	SetType(_type enums.TickerType) WsApiHr24
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

// SetSymbols 如果未指定交易对，则返回有关当前在交易所交易的所有交易对的信息。
func (hr *hr24Request) SetSymbols(symbols []string) WsApiHr24 {
	hr.symbols = symbols
	return hr
}

func (hr *hr24Request) SetType(_type enums.TickerType) WsApiHr24 {
	hr._type = _type
	return hr
}
func (hr *hr24Request) Receive(handler binance.Handler[WsApiHr24Response], exception binance.ErrorHandler) error {
	return binance.WsHandler(hr.Client, hr.BaseURL, handler, exception)
}

// Send 如果未指定交易对，则返回有关当前在交易所交易的所有交易对的信息。
func (hr *hr24Request) Send() error {
	req := &binance.Request{Path: "ticker.24hr"}
	if len(hr.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(hr.symbols, `","`))
		req.SetParam("symbols", result)
	}
	return hr.SendMessage(req)
}
