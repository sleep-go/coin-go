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

type Trades interface {
	Call(ctx context.Context) (body []*tradesResponse, err error)
}

// tradesRequest 近期成交
// 名称	类型	是否必须	描述
// symbol	STRING	YES
// limit	INT	NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
// 如果 limit > 5000, 最多返回5000条数据.
type tradesRequest struct {
	*binance.Client
	symbol string
	limit  enums.LimitType //默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
}

func NewTrades(client *binance.Client, symbol string, limit enums.LimitType) Trades {
	return &tradesRequest{
		Client: client,
		symbol: symbol,
		limit:  limit,
	}
}

type tradesResponse struct {
	Id           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

// Call 获取近期成交
// 名称	类型	是否必需	描述
// symbol	STRING	YES
// limit	INT	NO	Default 500; max 1000.
func (t *tradesRequest) Call(ctx context.Context) (body []*tradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketTrades,
	}
	req.SetParam("symbol", t.symbol)
	req.SetParam("limit", t.limit)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*tradesResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

type StreamTradeEvent struct {
	Stream string       `json:"stream"`
	Data   WsTradeEvent `json:"data"`
}
type WsTradeEvent struct {
	Event    string `json:"e"` // 事件类型
	Time     int64  `json:"E"` // 事件时间
	Symbol   string `json:"s"` // 交易对
	TradeID  int64  `json:"t"` // 交易ID
	Price    string `json:"p"` // 成交价格
	Quantity string `json:"q"` // 成交数量
	//BuyerOrderId  int64  `json:"b"`
	//SellerOrderId int64  `json:"a"`
	TradeTime    int64 `json:"T"` // 成交时间
	IsBuyerMaker bool  `json:"m"` // 买方是否是做市方。如true，则此次成交是一个主动卖出单，否则是一个主动买入单。
	Placeholder  bool  `json:"M"` // 请忽略该字段
}

// NewWsTrade 逐笔交易
// 逐笔交易推送每一笔成交的信息。成交，或者说交易的定义是仅有一个吃单者与一个挂单者相互交易。
//
// Stream 名称: <symbol>@trade
//
// 更新速度: 实时
func NewWsTrade(c *binance.Client, symbols []string, handler binance.Handler[WsTradeEvent], exception binance.ErrorHandler) error {
	return wsTrade(c, symbols, handler, exception)
}

// NewStreamTrade 逐笔交易
// 逐笔交易推送每一笔成交的信息。成交，或者说交易的定义是仅有一个吃单者与一个挂单者相互交易。
//
// Stream 名称: <symbol>@trade
//
// 更新速度: 实时
func NewStreamTrade(c *binance.Client, symbols []string, handler binance.Handler[StreamTradeEvent], exception binance.ErrorHandler) error {
	return wsTrade(c, symbols, handler, exception)
}
func wsTrade[T WsTradeEvent | StreamTradeEvent](c *binance.Client, symbols []string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.BaseURL
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@trade", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return binance.WsHandler(c, endpoint, handler, exception)
}

// ****************************** Websocket Api *******************************

type WsApiTrades interface {
	binance.WsApi[WsApiTradesResponse]
	SetSymbol(symbol string) WsApiTrades
	SetLimit(limit enums.LimitType) WsApiTrades
}
type WsApiTradesResponse struct {
	binance.WsApiResponse
	Result []*tradesResponse `json:"result"`
}

// NewWsApiTrades 最近的交易
// 获取最近的交易
//
// 如果您需要访问实时交易活动，请考虑使用 WebSocket Streams：
//
// <symbol>@trade
func NewWsApiTrades(c *binance.Client) WsApiTrades {
	return &tradesRequest{
		Client: c,
	}
}
func (t *tradesRequest) SetSymbol(symbol string) WsApiTrades {
	t.symbol = symbol
	return t
}

func (t *tradesRequest) SetLimit(limit enums.LimitType) WsApiTrades {
	t.limit = limit
	return t
}
func (t *tradesRequest) Receive(handler binance.Handler[WsApiTradesResponse], exception binance.ErrorHandler) error {
	return binance.WsHandler(t.Client, t.BaseURL, handler, exception)
}
func (t *tradesRequest) Send() error {
	req := &binance.Request{Path: "trades.recent"}
	req.SetParam("symbol", t.symbol)
	req.SetParam("limit", t.limit)
	return t.SendMessage(req)
}
