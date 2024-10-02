package trading

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

type AllOrders interface {
	SetOrderId(orderId int64) AllOrders
	SetRecvWindow(recvWindow int64) AllOrders
	SetTimestamp(timestamp int64) AllOrders
	Call(ctx context.Context) (body []*queryOrderResponse, err error)
}

// AllOrdersRequest 查询所有订单（包括历史订单） (USER_DATA)
// 注意:
// 如设置 orderId , 订单量将 >= orderId。否则将返回最新订单。
// 一些历史订单 cummulativeQuoteQty < 0, 是指数据此时不存在。
// 如果设置 startTime 和 endTime, orderId 就不需要设置。
// startTime和endTime之间的时间不能超过 24 小时。
type allOrdersRequest struct {
	*binance.Client
	symbol     string
	orderId    *int64
	startTime  *uint64
	endTime    *uint64
	limit      enums.LimitType
	recvWindow int64
	timestamp  int64
}

type allOrdersResponse struct {
	consts.ErrorResponse
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

func NewAllOrders(client *binance.Client, symbol string, limit enums.LimitType) AllOrders {
	return &allOrdersRequest{Client: client, symbol: symbol, limit: limit}
}
func (o *allOrdersRequest) SetStartTime(startTime uint64) {
	o.startTime = &startTime
}

func (o *allOrdersRequest) SetEndTime(endTime uint64) {
	o.endTime = &endTime
}

func (o *allOrdersRequest) SetOrderId(orderId int64) AllOrders {
	o.orderId = &orderId
	return o
}

func (o *allOrdersRequest) SetRecvWindow(recvWindow int64) AllOrders {
	o.recvWindow = recvWindow
	return o
}

func (o *allOrdersRequest) SetTimestamp(timestamp int64) AllOrders {
	o.timestamp = timestamp
	return o
}

func (o *allOrdersRequest) Call(ctx context.Context) (body []*queryOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTradingAllOrders,
	}
	req.SetNeedSign(true)
	if o.symbol != "" {
		req.SetParam("symbol", o.symbol)
	}
	if o.limit > 0 {
		req.SetParam("limit", o.limit)
	}
	if o.orderId != nil {
		req.SetParam("orderId", *o.orderId)
	}
	if o.startTime != nil {
		req.SetParam("startTime", *o.startTime)
	}
	if o.endTime != nil {
		req.SetParam("endTime", *o.endTime)
	}
	if o.recvWindow > 0 {
		req.SetParam("recvWindow", o.recvWindow)
	}
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("queryOrderRequest response err:%v", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var e *consts.ErrorResponse
		err = netutil.ParseHttpResponse(resp, &e)
		return nil, consts.Error(e.Code, e.Msg)
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		o.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
