package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type historyTradesRequest struct {
	*binance.Client
	symbol string
	limit  enums.LimitType //Default 500; max 1000.
	fromId *uint64         //从哪一条成交id开始返回. 缺省返回最近的成交记录
}

type HistoryTrades interface {
	Call(ctx context.Context) (body []*tradesResponse, err error)
	SetFromId(fromId uint64) *historyTradesRequest
	SetSymbol(symbol string) *historyTradesRequest
	SetLimit(limit enums.LimitType) *historyTradesRequest
}

func NewHistoryTrades(client *binance.Client, symbol string, limit enums.LimitType) HistoryTrades {
	return &historyTradesRequest{Client: client, symbol: symbol, limit: limit}
}

func (t *historyTradesRequest) SetFromId(fromId uint64) *historyTradesRequest {
	t.fromId = &fromId
	return t
}

func (t *historyTradesRequest) SetSymbol(symbol string) *historyTradesRequest {
	t.symbol = symbol
	return t
}

func (t *historyTradesRequest) SetLimit(limit enums.LimitType) *historyTradesRequest {
	t.limit = limit
	return t
}

// Call 查询历史成交
// 权重: 25
// 名称	类型	是否必需	描述
// symbol	STRING	YES
// limit	INT	NO	Default 500; max 1000.
// fromId	LONG	NO	从哪一条成交id开始返回. 缺省返回最近的成交记录
func (t *historyTradesRequest) Call(ctx context.Context) (body []*tradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketHistoricalTrades,
	}
	req.SetParam("symbol", t.symbol)
	req.SetParam("limit", t.limit)
	req.SetOptionalParam("fromId", t.fromId)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*tradesResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiHistoryTrades interface {
	binance.WsApi[*WsApiHistoryTradesResponse]
	HistoryTrades
}
type WsApiHistoryTradesResponse struct {
	binance.WsApiResponse
	Result []*tradesResponse `json:"result"`
}

func NewWsApiHistoryTrades(c *binance.Client) WsApiHistoryTrades {
	return &historyTradesRequest{
		Client: c,
	}
}

func (t *historyTradesRequest) Send(ctx context.Context) (*WsApiHistoryTradesResponse, error) {
	req := &binance.Request{Path: "trades.historical"}
	req.SetParam("symbol", t.symbol)
	req.SetParam("fromId", t.fromId)
	req.SetParam("limit", t.limit)
	return binance.WsApiHandler[*WsApiHistoryTradesResponse](ctx, t.Client, req)
}
