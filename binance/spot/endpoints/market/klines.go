package market

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

type Klines interface {
	Call(ctx context.Context) (body []*klinesResponse, err error)
	// CallUI 请求参数与响应和k线接口相同。
	// uiKlines 返回修改后的k线数据，针对k线图的呈现进行了优化。
	CallUI(ctx context.Context) (body []*klinesResponse, err error)
	SetInterval(interval enums.KlineIntervalType) Klines
	SetStartTime(startTime int64) Klines
	SetEndTime(endTime int64) Klines
	SetTimeZone(timeZone string) Klines
}
type klinesRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType         //Default 500; max 1000.
	interval  enums.KlineIntervalType //	请参考 K线间隔
	startTime *int64
	endTime   *int64
	timeZone  string
}

// [
//
//	[
//	  1499040000000,      // 开盘时间
//	  "0.01634790",       // 开盘价
//	  "0.80000000",       // 最高价
//	  "0.01575800",       // 最低价
//	  "0.01577100",       // 收盘价(当前K线未结束的即为最新价)
//	  "148976.11427815",  // 成交量
//	  1499644799999,      // 收盘时间
//	  "2434.19055334",    // 成交额
//	  308,                // 成交笔数
//	  "1756.87402397",    // 主动买入成交量
//	  "28.46694368",      // 主动买入成交额
//	  "17928899.62484339" // 请忽略该参数
//	]
//
// ]
type klinesResponse [12]any

func NewKlines(client *binance.Client, symbol string, limit enums.LimitType) Klines {
	return &klinesRequest{Client: client, symbol: symbol, limit: limit}
}

// SetInterval k线间隔 必传
func (k *klinesRequest) SetInterval(interval enums.KlineIntervalType) Klines {
	k.interval = interval
	return k
}

func (k *klinesRequest) SetStartTime(startTime int64) Klines {
	k.startTime = &startTime
	return k
}

func (k *klinesRequest) SetEndTime(endTime int64) Klines {
	k.endTime = &endTime
	return k
}

func (k *klinesRequest) SetTimeZone(timeZone string) Klines {
	k.timeZone = timeZone
	return k
}

// Call 请注意：
//
// 如果未发送startTime和endTime，将返回最近的K线数据。
// timeZone支持的值包括：
// 小时和分钟（例如 -1:00，05:45）
// 仅小时（例如 0，8，4）
// 接受的值范围严格为 [-12:00 到 +14:00]（包括边界）
// 如果提供了timeZone，K线间隔将在该时区中解释，而不是在UTC中。
// 请注意，无论timeZone如何，startTime和endTime始终以UTC时区解释。
func (k *klinesRequest) Call(ctx context.Context) (body []*klinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketKLines,
	}
	req.SetParam("symbol", k.symbol)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", string(k.interval))
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*klinesResponse](resp)
}
func (k *klinesRequest) CallUI(ctx context.Context) (body []*klinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketUIKLines,
	}
	req.SetParam("symbol", k.symbol)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", string(k.interval))
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*klinesResponse](resp)
}

type StreamKlineEvent struct {
	Stream string       `json:"stream"`
	Data   WsKlineEvent `json:"data"`
}
type WsKlineEvent struct {
	Event  string  `json:"e"`
	Time   int64   `json:"E"`
	Symbol string  `json:"s"`
	Kline  WsKline `json:"k"`
}
type WsKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Symbol               string `json:"s"`
	Interval             string `json:"i"`
	FirstTradeID         int64  `json:"f"`
	LastTradeID          int64  `json:"L"`
	Open                 string `json:"o"`
	Close                string `json:"c"`
	High                 string `json:"h"`
	Low                  string `json:"l"`
	Volume               string `json:"v"`
	TradeNum             int64  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

func NewWsKline(c *binance.WsClient, symbolsInterval map[string]enums.KlineIntervalType, handler binance.Handler[WsKlineEvent], exception binance.ErrorHandler) error {
	return wsKline(c, symbolsInterval, handler, exception)
}
func NewStreamKline(c *binance.WsClient, symbolsInterval map[string]enums.KlineIntervalType, handler binance.Handler[StreamKlineEvent], exception binance.ErrorHandler) error {
	return wsKline(c, symbolsInterval, handler, exception)
}

// wsKLines UTC K线
// K线stream逐秒推送所请求的K线种类(最新一根K线)的更新。此更新是基于 UTC+0 时区的。
//
// 订阅Kline需要提供间隔参数，最短为分钟线，最长为月线。支持以下间隔:
//
// m -> 分钟; h -> 小时; d -> 天; w -> 周; M -> 月
func wsKline[T WsKlineEvent | StreamKlineEvent](c *binance.WsClient, symbolsInterval map[string]enums.KlineIntervalType, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.Endpoint
	for symbol, interval := range symbolsInterval {
		endpoint += fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	if c.Timezone != "" {
		endpoint = fmt.Sprintf("%s@%s", endpoint, c.Timezone)
	}
	return wsHandler(c, endpoint, handler, exception)
}
