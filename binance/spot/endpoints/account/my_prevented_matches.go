package account

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

// MyPreventedMatches 获取 Prevented Matches (USER_DATA)
type MyPreventedMatches interface {
	SetOrderId(orderId int64) MyPreventedMatches
	SetPreventedMatchId(preventedMatchId int64) MyPreventedMatches
	SetFromPreventedMatchId(fromPreventedMatchId int64) MyPreventedMatches
	SetRecvWindow(recvWindow int64) MyPreventedMatches
	SetTimestamp(timestamp int64) MyPreventedMatches
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
	recvWindow           int64
	timestamp            int64
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

func (m *myPreventedMatchesRequest) SetOrderId(orderId int64) MyPreventedMatches {
	m.orderId = &orderId
	return m
}

func (m *myPreventedMatchesRequest) SetPreventedMatchId(preventedMatchId int64) MyPreventedMatches {
	m.preventedMatchId = &preventedMatchId
	return m
}

func (m *myPreventedMatchesRequest) SetFromPreventedMatchId(fromPreventedMatchId int64) MyPreventedMatches {
	m.fromPreventedMatchId = &fromPreventedMatchId
	return m
}

func (m *myPreventedMatchesRequest) SetRecvWindow(recvWindow int64) MyPreventedMatches {
	m.recvWindow = recvWindow
	return m
}

func (m *myPreventedMatchesRequest) SetTimestamp(timestamp int64) MyPreventedMatches {
	m.timestamp = timestamp
	return m
}

// Call 获取 Prevented Matches (USER_DATA)
// 获取因 STP 而过期的订单列表。
//
// 这些是支持的组合：
//
// symbol + preventedMatchId
// symbol + orderId
// symbol + orderId + fromPreventedMatchId (limit 默认为 500)
// symbol + orderId + fromPreventedMatchId + limit
func (m *myPreventedMatchesRequest) Call(ctx context.Context) (body []*myPreventedMatchesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountMyPreventedMatches,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetParam("limit", m.limit)
	if m.preventedMatchId != nil {
		req.SetParam("preventedMatchId", m.preventedMatchId)
	}
	if m.orderId != nil {
		req.SetParam("orderId", m.orderId)
	}
	if m.fromPreventedMatchId != nil {
		req.SetParam("fromPreventedMatchId", m.fromPreventedMatchId)
	}
	if m.recvWindow > 0 {
		req.SetParam("recvWindow", m.recvWindow)
	}
	req.SetParam("timestamp", m.timestamp)
	resp, err := m.Do(ctx, req)
	if err != nil {
		m.Debugf("myPreventedMatchesRequest response err:%v", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var e *consts.ErrorResponse
		err = netutil.ParseHttpResponse(resp, &e)
		return nil, consts.Error(e.Code, e.Msg)
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		m.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
