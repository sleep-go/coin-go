package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
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
	AvgPrice                string                 `json:"avgPrice"`                // 平均成交价
	ClientOrderId           string                 `json:"clientOrderId"`           // 用户自定义的订单号
	CumQuote                string                 `json:"cumQuote"`                // 成交金额
	ExecutedQty             string                 `json:"executedQty"`             // 成交量
	OrderId                 int                    `json:"orderId"`                 // 系统订单号
	OrigQty                 string                 `json:"origQty"`                 // 原始委托数量
	OrigType                enums.OrderType        `json:"origType"`                // 触发前订单类型
	Price                   string                 `json:"price"`                   // 委托价格
	ReduceOnly              bool                   `json:"reduceOnly"`              // 是否仅减仓
	Side                    enums.SideType         `json:"side"`                    // 买卖方向
	PositionSide            enums.PositionSideType `json:"positionSide"`            // 持仓方向
	Status                  enums.StatusType       `json:"status"`                  // 订单状态
	StopPrice               string                 `json:"stopPrice"`               // 触发价，对`TRAILING_STOP_MARKET`无效
	ClosePosition           bool                   `json:"closePosition"`           // 是否条件全平仓
	Symbol                  string                 `json:"symbol"`                  // 交易对
	Time                    int64                  `json:"time"`                    // 订单时间
	TimeInForce             enums.TimeInForceType  `json:"timeInForce"`             // 有效方法
	Type                    enums.OrderType        `json:"type"`                    // 订单类型
	ActivatePrice           string                 `json:"activatePrice"`           // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	PriceRate               string                 `json:"priceRate"`               // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	UpdateTime              int64                  `json:"updateTime"`              // 更新时间
	WorkingType             enums.WorkingType      `json:"workingType"`             // 条件价格触发类型
	PriceProtect            bool                   `json:"priceProtect"`            // 是否开启条件单触发保护
	PriceMatch              enums.PriceMatchType   `json:"priceMatch"`              //盘口价格下单模式
	SelfTradePreventionMode enums.StpModeType      `json:"selfTradePreventionMode"` //订单自成交保护模式
	GoodTillDate            int                    `json:"goodTillDate"`            //订单TIF为GTD时的自动取消时间
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
		Path:   consts.FApiTradingAllOrders,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("limit", o.limit)
	req.SetOptionalParam("orderId", o.orderId)
	req.SetOptionalParam("startTime", o.startTime)
	req.SetOptionalParam("endTime", o.endTime)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("allOrdersRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*allOrdersResponse](resp)
}

// ****************************** Websocket Api *******************************
