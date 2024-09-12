package ticker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

// Ticker
// 注意: 此接口和 GET /api/v3/ticker/24hr 有所不同.
//
// 此接口统计的时间范围比请求的windowSize多不超过59999ms.
//
// 接口的 openTime 是某一分钟的起始，而结束是当前的时间. 所以实际的统计区间会比请求的时间窗口多不超过59999ms.
//
// 比如, 结束时间 closeTime 是 1641287867099 (January 04, 2022 09:17:47:099 UTC) , windowSize 为 1d. 那么开始时间 openTime 则为 1641201420000 (January 3, 2022, 09:17:00 UTC)
type Ticker interface {
	Call(ctx context.Context) (body []*tickerResponse, err error)
	SetMinute(m uint8) Ticker
	SetHour(h uint8) Ticker
	SetDay(d uint8) Ticker
}
type tickerRequest struct {
	*binance.Client
	symbols []string
	//默认为 1d
	//windowSize 支持的值:
	//如果是分钟: 1m,2m....59m
	//如果是小时: 1h, 2h....23h
	//如果是天: 1d...7d
	windowSize string
	_type      enums.TickerType //可接受的参数: FULL or MINI. 如果不提供, 默认值为 FULL
}

// 滚动窗口价格变动统计
type tickerResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`        // 价格变化
	PriceChangePercent string `json:"priceChangePercent"` // 价格变化百分比
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"` // 此k线内所有交易的price(价格) x volume(交易量)的总和
	OpenTime           int64  `json:"openTime"`    // ticker的开始时间
	CloseTime          int64  `json:"closeTime"`   // ticker的结束时间
	FirstId            int    `json:"firstId"`     // 统计时间内的第一笔trade id
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"` // 统计时间内交易笔数
}

func NewTicker(client *binance.Client, symbols []string, _type enums.TickerType) Ticker {
	return &tickerRequest{Client: client, symbols: symbols, _type: _type}
}

func (t *tickerRequest) SetMinute(m uint8) Ticker {
	if m > 59 {
		m = 59
	} else if m < 1 {
		m = 1
	}
	t.windowSize = fmt.Sprintf("%dm", m)
	return t
}
func (t *tickerRequest) SetHour(h uint8) Ticker {
	if h > 23 {
		h = 23
	} else if h < 1 {
		h = 1
	}
	t.windowSize = fmt.Sprintf("%dh", h)
	return t
}
func (t *tickerRequest) SetDay(d uint8) Ticker {
	if d > 7 {
		d = 7
	} else if d < 1 {
		d = 1
	}
	t.windowSize = fmt.Sprintf("%dd", d)
	return t
}

func (t *tickerRequest) Call(ctx context.Context) (body []*tickerResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketTicker,
	}
	if len(t.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(t.symbols, `","`))
		req.SetParam("symbols", result)
	}
	req.SetParam("windowSize", t.windowSize)
	req.SetParam("type", t._type.String())
	res, err := t.Client.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Debugf("ReadAll err:%v", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", bytes)
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
