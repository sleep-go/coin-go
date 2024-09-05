package ticker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
	"github.com/sleep-go/exchange-go/binance/enums"
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
	res, err := hr.Client.Do(ctx, req)
	if err != nil {
		hr.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		hr.Debugf("ReadAll err:%v", err)
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
