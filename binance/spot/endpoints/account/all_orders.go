package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type AllOrders interface {
	SetSymbol(symbol string) *allOrdersRequest
	SetOrderId(orderId int64) *allOrdersRequest
	SetLimit(limit enums.LimitType) *allOrdersRequest
	SetStartTime(startTime uint64) *allOrdersRequest
	SetEndTime(endTime uint64) *allOrdersRequest
	Call(ctx context.Context) (body []*allOrdersResponse, err error)
}

// AllOrdersRequest 查询所有订单（包括历史订单） (USER_DATA)
// 注意:
// 如设置 orderId , 订单量将 >= orderId。否则将返回最新订单。
// 一些历史订单 cummulativeQuoteQty < 0, 是指数据此时不存在。
// 如果设置 startTime 和 endTime, orderId 就不需要设置。
// startTime和endTime之间的时间不能超过 24 小时。
type allOrdersRequest struct {
	*binance.Client
	symbol    string
	orderId   *int64
	startTime *uint64
	endTime   *uint64
	limit     enums.LimitType
}

type allOrdersResponse struct {
	Symbol                  string                `json:"symbol"`                  // 交易对
	OrderId                 int                   `json:"fromId"`                  // 系统的订单ID
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

// NewAllOrders 账户订单历史 (USER_DATA)
func NewAllOrders(client *binance.Client, symbol string, limit enums.LimitType) AllOrders {
	return &allOrdersRequest{Client: client, symbol: symbol, limit: limit}
}
func (o *allOrdersRequest) SetSymbol(symbol string) *allOrdersRequest {
	o.symbol = symbol
	return o
}

func (o *allOrdersRequest) SetLimit(limit enums.LimitType) *allOrdersRequest {
	o.limit = limit
	return o
}

func (o *allOrdersRequest) SetOrderId(orderId int64) *allOrdersRequest {
	o.orderId = &orderId
	return o
}

func (o *allOrdersRequest) SetStartTime(startTime uint64) *allOrdersRequest {
	o.startTime = &startTime
	return o
}

func (o *allOrdersRequest) SetEndTime(endTime uint64) *allOrdersRequest {
	o.endTime = &endTime
	return o
}

func (o *allOrdersRequest) Call(ctx context.Context) (body []*allOrdersResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTradingAllOrders,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("limit", o.limit)
	req.SetOptionalParam("orderId", o.orderId)
	req.SetOptionalParam("startTime", o.startTime)
	req.SetOptionalParam("endTime", o.endTime)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("queryOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*allOrdersResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiAllOrders interface {
	binance.WsApi[*WsApiAllOrdersResponse]
	AllOrders
}
type WsApiAllOrdersResponse struct {
	binance.WsApiResponse
	Result []*allOrdersResponse `json:"result"`
}

// NewWsApiAllOrders 账户订单历史 (USER_DATA)
// 获取所有账户订单； 有效，已取消或已完成。按时间范围过滤。
func NewWsApiAllOrders(c *binance.Client) WsApiAllOrders {
	return &allOrdersRequest{Client: c}
}

func (o *allOrdersRequest) Send(ctx context.Context) (*WsApiAllOrdersResponse, error) {
	req := &binance.Request{Path: "allOrders"}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("orderId", o.orderId)
	req.SetOptionalParam("startTime", o.startTime)
	req.SetOptionalParam("endTime", o.endTime)
	req.SetOptionalParam("limit", o.limit)
	return binance.WsApiHandler[*WsApiAllOrdersResponse](ctx, o.Client, req)
}
