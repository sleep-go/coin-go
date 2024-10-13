package market

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Klines interface {
	Call(ctx context.Context) (body []*KlinesResponse, err error)
	// CallContinuousKlines 连续合约K线数据。
	// 连续合约K线数据
	CallContinuousKlines(ctx context.Context) (body []*KlinesResponse, err error)
	// CallIndexPriceKlines 价格指数K线数据
	CallIndexPriceKlines(ctx context.Context) (body []*KlinesResponse, err error)
	// CallMarkPriceKlines 价格指数K线数据
	CallMarkPriceKlines(ctx context.Context) (body []*KlinesResponse, err error)
	// CallPremiumIndexKlines 溢价指数K线数据
	CallPremiumIndexKlines(ctx context.Context) (body []*KlinesResponse, err error)
	SetContractType(contractType enums.ContractType) *klinesRequest
	SetInterval(interval enums.KlineIntervalType) *klinesRequest
	SetStartTime(startTime int64) *klinesRequest
	SetEndTime(endTime int64) *klinesRequest
	SetTimeZone(timeZone string) *klinesRequest
}
type klinesRequest struct {
	*binance.Client
	symbol    string
	interval  enums.KlineIntervalType //	请参考 K线间隔
	startTime *int64
	endTime   *int64
	timeZone  string
	limit     enums.LimitType //Default 500; max 1000.

	contractType enums.ContractType
}

// KlinesResponse [
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
type KlinesResponse [12]any

func NewKlines(client *binance.Client, symbol string, limit enums.LimitType) Klines {
	return &klinesRequest{Client: client, symbol: symbol, limit: limit}
}
func (k *klinesRequest) SetContractType(contractType enums.ContractType) *klinesRequest {
	k.contractType = contractType
	return k
}

func (k *klinesRequest) SetSymbol(symbol string) *klinesRequest {
	k.symbol = symbol
	return k
}

func (k *klinesRequest) SetLimit(limit enums.LimitType) *klinesRequest {
	k.limit = limit
	return k
}

// SetInterval k线间隔 必传
func (k *klinesRequest) SetInterval(interval enums.KlineIntervalType) *klinesRequest {
	k.interval = interval
	return k
}

func (k *klinesRequest) SetStartTime(startTime int64) *klinesRequest {
	k.startTime = &startTime
	return k
}

func (k *klinesRequest) SetEndTime(endTime int64) *klinesRequest {
	k.endTime = &endTime
	return k
}

func (k *klinesRequest) SetTimeZone(timeZone string) *klinesRequest {
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
func (k *klinesRequest) Call(ctx context.Context) (body []*KlinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketKLines,
	}
	req.SetParam("symbol", k.symbol)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*KlinesResponse](resp)
}

// CallContinuousKlines 连续合约K线数据
// 每根K线的开盘时间可视为唯一ID
func (k *klinesRequest) CallContinuousKlines(ctx context.Context) (body []*KlinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketContinuousKlines,
	}
	req.SetParam("pair", k.symbol)
	req.SetParam("contractType", k.contractType)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*KlinesResponse](resp)
}

// CallIndexPriceKlines 价格指数K线数据
func (k *klinesRequest) CallIndexPriceKlines(ctx context.Context) (body []*KlinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketIndexPriceKlines,
	}
	req.SetParam("pair", k.symbol)
	req.SetParam("contractType", k.contractType)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*KlinesResponse](resp)
}

// CallMarkPriceKlines 标记价格K线数据
func (k *klinesRequest) CallMarkPriceKlines(ctx context.Context) (body []*KlinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketMarkPriceKlines,
	}
	req.SetParam("pair", k.symbol)
	req.SetParam("contractType", k.contractType)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*KlinesResponse](resp)
}

// CallPremiumIndexKlines 溢价指数K线数据
func (k *klinesRequest) CallPremiumIndexKlines(ctx context.Context) (body []*KlinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketPremiumIndexKlines,
	}
	req.SetParam("symbol", k.symbol)
	req.SetParam("contractType", k.contractType)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*KlinesResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

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

func NewWsKline(c *binance.Client, symbolsInterval map[string]enums.KlineIntervalType, handler binance.Handler[WsKlineEvent], exception binance.ErrorHandler) error {
	return wsKline(c, symbolsInterval, handler, exception)
}
func NewStreamKline(c *binance.Client, symbolsInterval map[string]enums.KlineIntervalType, handler binance.Handler[StreamKlineEvent], exception binance.ErrorHandler) error {
	return wsKline(c, symbolsInterval, handler, exception)
}

// wsKLines UTC K线
// K线stream逐秒推送所请求的K线种类(最新一根K线)的更新。此更新是基于 UTC+0 时区的。
//
// 订阅Kline需要提供间隔参数，最短为分钟线，最长为月线。支持以下间隔:
//
// m -> 分钟; h -> 小时; d -> 天; w -> 周; M -> 月
func wsKline[T WsKlineEvent | StreamKlineEvent](c *binance.Client, symbolsInterval map[string]enums.KlineIntervalType, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.BaseURL
	for symbol, interval := range symbolsInterval {
		endpoint += fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	if c.Timezone != "" {
		endpoint = fmt.Sprintf("%s@%s", endpoint, c.Timezone)
	}
	return binance.WsHandler(c, endpoint, handler, exception)
}

// ****************************** Websocket Api *******************************

type WsApiKlines interface {
	binance.WsApi[*WsApiKlinesResponse]
	SetSymbol(symbol string) *klinesRequest
	SetLimit(limit enums.LimitType) *klinesRequest
	SendUI(ctx context.Context) (*WsApiKlinesResponse, error)
}
type WsApiKlinesResponse struct {
	binance.WsApiResponse
	Result []*KlinesResponse `json:"result"`
}

// NewWsApiKlines K线数据
// 获取K线数据。
//
// Klines 由其开盘时间和收盘时间为唯一标识。
//
// 如果您需要访问实时 kline 更新，请考虑使用 WebSocket Streams：
//
// <symbol>@kline_<interval>
// 如果需要历史K线数据，可以使用 data.binance.vision。
func NewWsApiKlines(c *binance.Client) WsApiKlines {
	return &klinesRequest{Client: c}
}

// Send 备注:
// method: klines or uiKlines
//
// 如果没有指定 startTime，endTime，则返回最近的klines。
// timeZone支持的值包括：
// 小时和分钟（例如 -1:00，05:45）
// 仅小时（例如 0，8，4）
// 接受的值范围严格为 [-12:00 到 +14:00]（包括边界）
// 如果提供了timeZone，K线间隔将在该时区中解释，而不是在UTC中。
// 请注意，无论timeZone如何，startTime和endTime始终以UTC时区解释。
func (k *klinesRequest) Send(ctx context.Context) (*WsApiKlinesResponse, error) {
	req := &binance.Request{Path: "klines"}
	req.SetParam("symbol", k.symbol)
	req.SetOptionalParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	handler, err := binance.WsApiHandler[WsApiKlinesResponse](ctx, k.Client, req)
	if err != nil {
		return nil, err
	}
	return &handler, nil
}
func (k *klinesRequest) SendUI(ctx context.Context) (*WsApiKlinesResponse, error) {
	req := &binance.Request{Path: "uiKlines"}
	req.SetParam("symbol", k.symbol)
	req.SetOptionalParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	req.SetOptionalParam("timeZone", k.timeZone)
	return binance.WsApiHandler[*WsApiKlinesResponse](ctx, k.Client, req)
}
