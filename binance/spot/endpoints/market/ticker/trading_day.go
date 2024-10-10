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

type TradingDay interface {
	Call(ctx context.Context) (body []*tradingDayResponse, err error)
}

// 参数名	类型	是否必需	描述
// symbol	STRING	YES	symbol 或者 symbols 必须提供之一
//
// symbols 可以接受的格式:
// ["BTCUSDT","BNBUSDT"]
// 或者
// %5B%22BTCUSDT%22,%22BNBUSDT%22%5D
//
// symbols 最多可以发送100个.
// symbols
// timeZone	STRING	NO	Default: 0 (UTC)
// type	ENUM	NO	可接受值: FULL or MINI.
// 默认值: FULL
type tradingDayRequest struct {
	*binance.Client
	symbols  []string
	timeZone string
	_type    enums.TickerType
}

// NewTradingDay 交易日行情(Ticker)
func NewTradingDay(client *binance.Client, symbols []string, timeZone string, _type enums.TickerType) TradingDay {
	return &tradingDayRequest{Client: client, symbols: symbols, timeZone: timeZone, _type: _type}
}

type tradingDayResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int    `json:"count"`
}

func (t *tradingDayRequest) Call(ctx context.Context) (body []*tradingDayResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketTickerTradingDay,
	}
	if len(t.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(t.symbols, `","`))
		req.SetParam("symbols", result)
	}
	req.SetParam("timeZone", t.timeZone)
	req.SetParam("type", t._type.String())
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*tradingDayResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiTradingDay interface {
	binance.WsApi[WsApiTradingDayResponse]
	SetSymbols(symbols []string) WsApiTradingDay
	SetTimeZone(timeZone string) WsApiTradingDay
	SetType(_type enums.TickerType) WsApiTradingDay
}
type WsApiTradingDayResponse struct {
	binance.WsApiResponse
	Result []*tradingDayResponse `json:"result"`
}

// NewWsApiTradingDay 交易日行情(Ticker)
// 交易日价格变动统计。
//
// 权重:
//
// 每个交易对占用4个权重.
//
// 当请求中的交易对数量超过50，此请求的权重将限制在200。
func NewWsApiTradingDay(c *binance.Client) WsApiTradingDay {
	return &tradingDayRequest{
		Client: c,
	}
}
func (t *tradingDayRequest) SetSymbols(symbols []string) WsApiTradingDay {
	t.symbols = symbols
	return t
}

func (t *tradingDayRequest) SetType(_type enums.TickerType) WsApiTradingDay {
	t._type = _type
	return t
}

// SetTimeZone 注意:
//
// timeZone支持的值包括：
// 小时和分钟（例如 -1:00，05:45）
// 仅小时（例如 0，8，4）
func (t *tradingDayRequest) SetTimeZone(timeZone string) WsApiTradingDay {
	t.timeZone = timeZone
	return t
}
func (t *tradingDayRequest) Receive(handler binance.Handler[WsApiTradingDayResponse], exception binance.ErrorHandler) error {
	return binance.WsHandler(t.Client, t.BaseURL, handler, exception)
}

func (t *tradingDayRequest) Send() error {
	req := &binance.Request{Path: "ticker.tradingDay"}
	if len(t.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(t.symbols, `","`))
		req.SetParam("symbols", result)
	}
	req.SetOptionalParam("timeZone", t.timeZone)
	req.SetOptionalParam("type", t._type.String())
	return t.SendMessage(req)
}
