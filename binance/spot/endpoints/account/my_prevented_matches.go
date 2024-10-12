package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// MyPreventedMatches 获取 Prevented Matches (USER_DATA)
type MyPreventedMatches interface {
	SetSymbol(symbol string) *myPreventedMatchesRequest
	SetLimit(limit enums.LimitType) *myPreventedMatchesRequest
	SetOrderId(orderId int64) *myPreventedMatchesRequest
	SetPreventedMatchId(preventedMatchId int64) *myPreventedMatchesRequest
	SetFromPreventedMatchId(fromPreventedMatchId int64) *myPreventedMatchesRequest
	Call(ctx context.Context) (body []*myPreventedMatchesResponse, err error)
}

// myPreventedMatchesRequest 获取因 STP 而过期的订单列表。
// 这些是支持的组合：
//
// symbol + preventedMatchId
// symbol + orderId
// symbol + orderId + fromPreventedMatchId (limit 默认为 500)
// symbol + orderId + fromPreventedMatchId + limit
type myPreventedMatchesRequest struct {
	*binance.Client
	symbol               string
	preventedMatchId     *int64
	orderId              *int64
	fromPreventedMatchId *int64
	limit                enums.LimitType
}

type myPreventedMatchesResponse struct {
	Symbol                  string `json:"symbol"`
	PreventedMatchId        int    `json:"preventedMatchId"`
	TakerOrderId            int    `json:"takerOrderId"`
	MakerSymbol             string `json:"makerSymbol"`
	MakerOrderId            int    `json:"makerOrderId"`
	TradeGroupId            int    `json:"tradeGroupId"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	Price                   string `json:"price"`
	MakerPreventedQuantity  string `json:"makerPreventedQuantity"`
	TransactTime            int64  `json:"transactTime"`
}

func NewMyPreventedMatches(client *binance.Client, symbol string, limit enums.LimitType) MyPreventedMatches {
	return &myPreventedMatchesRequest{Client: client, symbol: symbol, limit: limit}
}

func (m *myPreventedMatchesRequest) SetSymbol(symbol string) *myPreventedMatchesRequest {
	m.symbol = symbol
	return m
}

func (m *myPreventedMatchesRequest) SetLimit(limit enums.LimitType) *myPreventedMatchesRequest {
	m.limit = limit
	return m
}
func (m *myPreventedMatchesRequest) SetOrderId(orderId int64) *myPreventedMatchesRequest {
	m.orderId = &orderId
	return m
}

func (m *myPreventedMatchesRequest) SetPreventedMatchId(preventedMatchId int64) *myPreventedMatchesRequest {
	m.preventedMatchId = &preventedMatchId
	return m
}

func (m *myPreventedMatchesRequest) SetFromPreventedMatchId(fromPreventedMatchId int64) *myPreventedMatchesRequest {
	m.fromPreventedMatchId = &fromPreventedMatchId
	return m
}

// Call 获取 Prevented Matches (USER_DATA)
// 获取因 STP 而过期的订单列表。
//
// 这些是支持的组合：
//
// symbol + preventedMatchId
// symbol + fromId
// symbol + fromId + fromPreventedMatchId (limit 默认为 500)
// symbol + fromId + fromPreventedMatchId + limit
func (m *myPreventedMatchesRequest) Call(ctx context.Context) (body []*myPreventedMatchesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountMyPreventedMatches,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetOptionalParam("limit", m.limit)
	req.SetOptionalParam("preventedMatchId", m.preventedMatchId)
	req.SetOptionalParam("fromId", m.orderId)
	req.SetOptionalParam("fromPreventedMatchId", m.fromPreventedMatchId)
	resp, err := m.Do(ctx, req)
	if err != nil {
		m.Debugf("myPreventedMatchesRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*myPreventedMatchesResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiMyPreventedMatches interface {
	binance.WsApi[*WsApiMyPreventedMatchesResponse]
	MyPreventedMatches
}
type WsApiMyPreventedMatchesResponse struct {
	binance.WsApiResponse
	Result []*myPreventedMatchesResponse `json:"result"`
}

// NewWsApiMyPreventedMatches 账户的 Prevented Matches (USER_DATA)
// 获取因 STP 而过期的订单列表。
//
// 这些是支持的组合：
//
// symbol + preventedMatchId
// symbol + orderId
// symbol + orderId + fromPreventedMatchId (limit 默认为 500)
// symbol + orderId + fromPreventedMatchId + limit
func NewWsApiMyPreventedMatches(c *binance.Client) WsApiMyPreventedMatches {
	return &myPreventedMatchesRequest{Client: c}
}

func (m *myPreventedMatchesRequest) Send(ctx context.Context) (*WsApiMyPreventedMatchesResponse, error) {
	req := &binance.Request{Path: "myTrades"}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetOptionalParam("limit", m.limit)
	req.SetOptionalParam("preventedMatchId", m.preventedMatchId)
	req.SetOptionalParam("fromId", m.orderId)
	req.SetOptionalParam("fromPreventedMatchId", m.fromPreventedMatchId)
	return binance.WsApiHandler[*WsApiMyPreventedMatchesResponse](ctx, m.Client, req)
}
