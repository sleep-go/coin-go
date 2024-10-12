package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type MyTrades interface {
	SetSymbol(symbol string) *myTradesRequest
	SetLimit(limit enums.LimitType) *myTradesRequest
	SetOrderId(orderId int64) *myTradesRequest
	SetStartTime(startTime uint64) *myTradesRequest
	SetEndTime(endTime uint64) *myTradesRequest
	SetFromId(fromId int64) *myTradesRequest
	Call(ctx context.Context) (body []*myTradesResponse, err error)
}

// 备注:
//
// 如果设置了fromId, 会返回ID大于此fromId的交易. 不然则会返回最近的交易.
// startTime和endTime设置的时间间隔不能超过24小时.
// 支持的所有参数组合:
// symbol
// symbol + orderId
// symbol + startTime
// symbol + endTime
// symbol + fromId
// symbol + startTime + endTime
// symbol+ orderId + fromId
type myTradesRequest struct {
	*binance.Client
	symbol    string
	orderId   *int64 //必须要和参数symbol一起使用.
	startTime *uint64
	endTime   *uint64
	fromId    *int64          //返回该fromId之后的成交，缺省返回最近的成交
	limit     enums.LimitType //Default 500; max 1000.
}

type myTradesResponse struct {
	Symbol          string `json:"symbol"`
	Id              int    `json:"id"`
	OrderId         int    `json:"fromId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

func NewMyTrades(client *binance.Client, symbol string, limit enums.LimitType) MyTrades {
	return &myTradesRequest{Client: client, symbol: symbol, limit: limit}
}
func (m *myTradesRequest) SetLimit(limit enums.LimitType) *myTradesRequest {
	m.limit = limit
	return m
}

func (m *myTradesRequest) SetSymbol(symbol string) *myTradesRequest {
	m.symbol = symbol
	return m
}
func (m *myTradesRequest) SetOrderId(orderId int64) *myTradesRequest {
	m.orderId = &orderId
	return m
}

func (m *myTradesRequest) SetStartTime(startTime uint64) *myTradesRequest {
	m.startTime = &startTime
	return m
}

func (m *myTradesRequest) SetEndTime(endTime uint64) *myTradesRequest {
	m.endTime = &endTime
	return m
}

func (m *myTradesRequest) SetFromId(fromId int64) *myTradesRequest {
	m.fromId = &fromId
	return m
}

func (m *myTradesRequest) Call(ctx context.Context) (body []*myTradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountMyTrades,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetOptionalParam("fromId", m.orderId)
	req.SetOptionalParam("startTime", m.startTime)
	req.SetOptionalParam("endTime", m.endTime)
	req.SetOptionalParam("fromId", m.fromId)
	req.SetOptionalParam("limit", m.limit)
	resp, err := m.Do(ctx, req)
	if err != nil {
		m.Debugf("myTradesRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*myTradesResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiMyTrades interface {
	binance.WsApi[*WsApiMyTradesResponse]
	MyTrades
}
type WsApiMyTradesResponse struct {
	binance.WsApiResponse
	Result []*myTradesResponse `json:"result"`
}

// NewWsApiMyTrades 账户成交历史 (USER_DATA)
// 获取账户指定交易对的成交历史，按时间范围过滤。
// 备注：
//
// 如果指定了 fromId，则返回的交易将是 交易ID >= fromId。
//
// 如果指定了 startTime 和/或 endTime，则交易按执行时间（time）过滤。
//
// fromId 不能与 startTime 和 endTime 一起使用。
//
// 如果指定了 orderId，则只返回与该订单相关的交易。
//
// startTime 和 endTime 不能与 orderId 一起使用。
//
// 如果不指定条件，则返回最近的交易。
//
// startTime和endTime之间的时间不能超过 24 小时。
func NewWsApiMyTrades(c *binance.Client) WsApiMyTrades {
	return &myTradesRequest{Client: c}
}

func (m *myTradesRequest) Send(ctx context.Context) (*WsApiMyTradesResponse, error) {
	req := &binance.Request{Path: "myTrades"}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetOptionalParam("fromId", m.orderId)
	req.SetOptionalParam("startTime", m.startTime)
	req.SetOptionalParam("endTime", m.endTime)
	req.SetOptionalParam("fromId", m.fromId)
	req.SetOptionalParam("limit", m.limit)
	return binance.WsApiHandler[*WsApiMyTradesResponse](ctx, m.Client, req)
}
