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
