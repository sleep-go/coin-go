package market

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/spot/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type AggTrades interface {
	Call(ctx context.Context) (body []*aggTradesResponse, err error)
	SetSymbol(symbol string) *aggTradesRequest
	SetLimit(limit enums.LimitType) *aggTradesRequest
	SetFromId(fromId uint64) *aggTradesRequest
	SetStartTime(startTime int64) *aggTradesRequest
	SetEndTime(endTime int64) *aggTradesRequest
}

type aggTradesRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType //默认 500; 最大 1000.
	fromId    *uint64         //从包含fromID的成交开始返回结果
	startTime *int64          //从该时刻之后的成交记录开始返回结果
	endTime   *int64          //返回该时刻为止的成交记录
}

func (a *aggTradesRequest) SetSymbol(symbol string) *aggTradesRequest {
	a.symbol = symbol
	return a
}

func (a *aggTradesRequest) SetLimit(limit enums.LimitType) *aggTradesRequest {
	a.limit = limit
	return a
}

type aggTradesResponse struct {
	AggTradeID            int    `json:"a"` //归集成交ID
	Price                 string `json:"p"` // 成交价
	Quantity              string `json:"q"` // 成交量
	FirstBreakdownTradeID int    `json:"f"` // 被归集的首个成交ID
	LastBreakdownTradeID  int    `json:"l"` // 被归集的末个成交ID
	TradeTime             int64  `json:"T"` // 成交时间
	IsBuyerMaker          bool   `json:"m"` // 是否为主动卖出单
	Placeholder           bool   `json:"M"` // 是否为最优撮合单(可忽略，目前总为最优撮合)
}

func NewAggTrades(client *binance.Client, symbol string, limit enums.LimitType) AggTrades {
	return &aggTradesRequest{Client: client, symbol: symbol, limit: limit}
}

func (a *aggTradesRequest) SetFromId(fromId uint64) *aggTradesRequest {
	a.fromId = &fromId
	return a
}

func (a *aggTradesRequest) SetStartTime(startTime int64) *aggTradesRequest {
	a.startTime = &startTime
	return a
}

func (a *aggTradesRequest) SetEndTime(endTime int64) *aggTradesRequest {
	a.endTime = &endTime
	return a
}

// Call 近期成交(归集)
// 与trades的区别是，同一个taker在同一时间同一价格与多个maker的成交会被合并为一条记录
func (a *aggTradesRequest) Call(ctx context.Context) (body []*aggTradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketAggTrades,
	}
	req.SetParam("symbol", a.symbol)
	req.SetParam("limit", a.limit)
	req.SetOptionalParam("fromId", a.fromId)
	req.SetOptionalParam("startTime", a.startTime)
	req.SetOptionalParam("endTime", a.endTime)
	resp, err := a.Do(ctx, req)
	if err != nil {
		a.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*aggTradesResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

type StreamAggTradeEvent struct {
	Stream string          `json:"stream"`
	Data   WsAggTradeEvent `json:"data"`
}
type WsAggTradeEvent struct {
	Event  string `json:"e"` // 事件类型
	Time   int64  `json:"E"` // 事件时间
	Symbol string `json:"s"` // 交易对
	aggTradesResponse
}

// NewWsAggTrade 归集交易
// 归集交易与逐笔交易的区别在于，同一个taker在同一价格与多个maker成交时，会被归集为一笔成交。
// Stream 名称: <symbol>@aggTrade
// 更新速度: 实时
func NewWsAggTrade(c *binance.Client, symbols []string, handler binance.Handler[WsAggTradeEvent], exception binance.ErrorHandler) error {
	return wsAggTrade(c, symbols, handler, exception)
}

// NewStreamAggTrade 归集交易
// 归集交易与逐笔交易的区别在于，同一个taker在同一价格与多个maker成交时，会被归集为一笔成交。
// Stream 名称: <symbol>@aggTrade
// 更新速度: 实时
func NewStreamAggTrade(c *binance.Client, symbols []string, handler binance.Handler[StreamAggTradeEvent], exception binance.ErrorHandler) error {
	return wsAggTrade(c, symbols, handler, exception)
}
func wsAggTrade[T WsAggTradeEvent | StreamAggTradeEvent](c *binance.Client, symbols []string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.BaseURL
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@aggTrade", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return binance.WsHandler(c, endpoint, handler, exception)
}

// ****************************** Websocket Api *******************************

type WsApiAggTrades interface {
	binance.WsApi[*WsApiAggTradesResponse]
	AggTrades
}
type WsApiAggTradesResponse struct {
	binance.WsApiResponse
	Result []*aggTradesResponse `json:"result"`
}

// NewWsApiAggTrades 归集交易
// 一个 归集交易 (aggtrade) 代表一个或多个单独的交易。 同时间，同 taker 订单和同价格的执行交易会被聚合为一条归集交易。
//
// 如果需要访问实时交易活动，请考虑使用 WebSocket Streams：
//
// <symbol>@aggTrade
// 如果需要历史总交易数据，可以使用 data.binance.vision。
func NewWsApiAggTrades(c *binance.Client) WsApiAggTrades {
	return &aggTradesRequest{Client: c}
}

// Send 归集交易
// 备注：
//
// 如果指定了 fromId，则返回归集交易 ID >= fromId 的 aggtrades。
//
// 使用 fromId 和 limit 会对所有 aggtrades 进行分页。
//
// 如果指定了 startTime 和/或 endTime，响应中的 aggtrades 会按照执行时间 (T) 过滤。
//
// fromId 不能与 startTime 和 endTime 一起使用。
//
// 如果未指定条件，则返回最近的归集交易。
func (a *aggTradesRequest) Send(ctx context.Context) (*WsApiAggTradesResponse, error) {
	req := &binance.Request{Path: "trades.aggregate"}
	req.SetParam("symbol", a.symbol)
	req.SetParam("limit", a.limit)
	req.SetOptionalParam("fromId", a.fromId)
	req.SetOptionalParam("startTime", a.startTime)
	req.SetOptionalParam("endTime", a.endTime)
	handler, err := binance.WsApiHandler[WsApiAggTradesResponse](ctx, a.Client, req)
	if err != nil {
		return nil, err
	}
	return &handler, nil
}
