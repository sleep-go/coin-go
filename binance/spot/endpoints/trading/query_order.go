package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type QueryOrder interface {
	Call(ctx context.Context) (body *queryOrderResponse, err error)
	CallOpenOrders(ctx context.Context) (body []*queryOrderResponse, err error)
	SetSymbol(symbol string) *queryOrderRequest
	SetOrderId(orderId int64) *queryOrderRequest
	SetOrigClientOrderId(origClientOrderId string) *queryOrderRequest
}
type queryOrderRequest struct {
	*binance.Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
}

// queryOrderResponse 查询订单 (USER_DATA)
// 至少需要发送 orderId 与 origClientOrderId中的一个
// 某些订单中cummulativeQuoteQty<0，是由于这些订单是cummulativeQuoteQty功能上线之前的订单。
type queryOrderResponse struct {
	Symbol                  string                `json:"symbol"`                  // 交易对
	OrderId                 int                   `json:"orderId"`                 // 系统的订单ID
	OrderListId             int                   `json:"orderListId"`             // 除非此单是订单列表的一部分, 否则此值为 -1
	ClientOrderId           string                `json:"clientOrderId"`           // 客户自己设置的ID
	Price                   string                `json:"price"`                   // 订单价格
	OrigQty                 string                `json:"origQty"`                 // 用户设置的原始订单数量
	ExecutedQty             string                `json:"executedQty"`             // 交易的订单数量
	CummulativeQuoteQty     string                `json:"cummulativeQuoteQty"`     // 累计交易的金额
	Status                  enums.OrderStatusType `json:"status"`                  // 订单状态
	TimeInForce             enums.TimeInForceType `json:"timeInForce"`             // 订单的时效方式
	Type                    enums.OrderType       `json:"type"`                    // 订单类型， 比如市价单，现价单等
	Side                    enums.SideType        `json:"side"`                    // 订单方向，买还是卖
	StopPrice               string                `json:"stopPrice"`               // 止损价格
	IcebergQty              string                `json:"icebergQty"`              // 冰山数量
	Time                    int64                 `json:"time"`                    // 订单时间
	UpdateTime              int64                 `json:"updateTime"`              // 最后更新时间
	IsWorking               bool                  `json:"isWorking"`               // 订单是否出现在orderbook中
	WorkingTime             int64                 `json:"workingTime"`             // 订单添加到 order book 的时间
	OrigQuoteOrderQty       string                `json:"origQuoteOrderQty"`       // 原始的交易金额
	SelfTradePreventionMode enums.StpModeType     `json:"selfTradePreventionMode"` // 如何处理自我交易模式
	//订单响应中的特定条件时才会出现的字段
	PreventedMatchId  int64  `json:"preventedMatchId,omitempty"`
	PreventedQuantity string `json:"preventedQuantity,omitempty"`
	StrategyId        int64  `json:"strategyId,omitempty"`
	StrategyType      int64  `json:"strategyType,omitempty"`
	TrailingDelta     string `json:"trailingDelta,omitempty"`
	TrailingTime      int64  `json:"trailingTime,omitempty"`
}

func NewQueryOrder(client *binance.Client, symbol string) QueryOrder {
	return &queryOrderRequest{Client: client, symbol: symbol}
}

func (o *queryOrderRequest) SetSymbol(symbol string) *queryOrderRequest {
	o.symbol = symbol
	return o
}

func (o *queryOrderRequest) SetOrderId(orderId int64) *queryOrderRequest {
	o.orderId = &orderId
	return o
}

func (o *queryOrderRequest) SetOrigClientOrderId(origClientOrderId string) *queryOrderRequest {
	o.origClientOrderId = &origClientOrderId
	return o
}

// Call 查询订单 (USER_DATA)
// 至少需要发送 orderId 与 origClientOrderId中的一个
// 某些订单中cummulativeQuoteQty<0，是由于这些订单是cummulativeQuoteQty功能上线之前的订单。
func (o *queryOrderRequest) Call(ctx context.Context) (body *queryOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("orderId", o.orderId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("queryOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*queryOrderResponse](resp)
}
func (o *queryOrderRequest) CallOpenOrders(ctx context.Context) (body []*queryOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiOpenOrders,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("symbol", o.symbol)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("queryOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*queryOrderResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiQueryOrder interface {
	binance.WsApi[*WsApiQueryOrderResponse]
	QueryOrder
	SendOpenOrders(ctx context.Context) (*WsApiOpenOrdersResponse, error)
}
type WsApiQueryOrderResponse struct {
	binance.WsApiResponse
	Result *queryOrderResponse `json:"result"`
}
type WsApiOpenOrdersResponse struct {
	binance.WsApiResponse
	Result []*queryOrderResponse `json:"result"`
}

func NewWsApiQueryOrder(c *binance.Client) WsApiQueryOrder {
	return &queryOrderRequest{Client: c}
}

// Send 下新的订单 (TRADE)
func (o *queryOrderRequest) Send(ctx context.Context) (*WsApiQueryOrderResponse, error) {
	req := &binance.Request{Path: "order.status"}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("orderId", o.orderId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	return binance.WsApiHandler[*WsApiQueryOrderResponse](ctx, o.Client, req)
}

// SendOpenOrders 当前挂单 (USER_DATA)
func (o *queryOrderRequest) SendOpenOrders(ctx context.Context) (*WsApiOpenOrdersResponse, error) {
	req := &binance.Request{Path: "openOrders.status"}
	req.SetNeedSign(true)
	req.SetOptionalParam("symbol", o.symbol)
	return binance.WsApiHandler[*WsApiOpenOrdersResponse](ctx, o.Client, req)
}
