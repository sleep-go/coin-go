package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/errors"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type DeleteOrder interface {
	SetSymbol(symbol string) *deleteOrderRequest
	SetOrderId(orderId int64) *deleteOrderRequest
	SetOrigClientOrderId(origClientOrderId string) *deleteOrderRequest
	Call(ctx context.Context) (body *deleteOrderResponse, err error)
	CallBatch(ctx context.Context, orderIdList []int64) (body []*deleteOrderResponse, err error)
	CallAllOpenOrders(ctx context.Context) (body *errors.Status, err error)
}

// deleteOrderRequest 撤销订单 (TRADE)
// orderId 与 origClientOrderId 必须至少发送一个.
type deleteOrderRequest struct {
	*binance.Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
}
type deleteOrderResponse struct {
	Code                    int                    `json:"code,omitempty"`
	Msg                     string                 `json:"msg,omitempty"`
	ClientOrderId           string                 `json:"clientOrderId"` // 用户自定义的订单号
	CumQty                  string                 `json:"cumQty"`
	CumQuote                string                 `json:"cumQuote"`                // 成交金额
	ExecutedQty             string                 `json:"executedQty"`             // 成交量
	OrderId                 int                    `json:"orderId"`                 // 系统订单号
	OrigQty                 string                 `json:"origQty"`                 // 原始委托数量
	Price                   string                 `json:"price"`                   // 委托价格
	ReduceOnly              bool                   `json:"reduceOnly"`              // 仅减仓
	Side                    enums.SideType         `json:"side"`                    // 买卖方向
	PositionSide            enums.PositionSideType `json:"positionSide"`            // 持仓方向
	Status                  enums.StatusType       `json:"status"`                  // 订单状态
	StopPrice               string                 `json:"stopPrice"`               // 触发价，对`TRAILING_STOP_MARKET`无效
	ClosePosition           bool                   `json:"closePosition"`           // 是否条件全平仓
	Symbol                  string                 `json:"symbol"`                  // 交易对
	TimeInForce             enums.TimeInForceType  `json:"timeInForce"`             // 有效方法
	OrigType                enums.OrderType        `json:"origType"`                // 触发前订单类型
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

func NewDeleteOrder(client *binance.Client, symbol string) DeleteOrder {
	return &deleteOrderRequest{Client: client, symbol: symbol}
}

func (d *deleteOrderRequest) SetSymbol(symbol string) *deleteOrderRequest {
	d.symbol = symbol
	return d
}
func (d *deleteOrderRequest) SetOrderId(orderId int64) *deleteOrderRequest {
	d.orderId = &orderId
	return d
}

func (d *deleteOrderRequest) SetOrigClientOrderId(origClientOrderId string) *deleteOrderRequest {
	d.origClientOrderId = &origClientOrderId
	return d
}

func (d *deleteOrderRequest) Call(ctx context.Context) (body *deleteOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.FApiOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	req.SetOptionalParam("orderId", d.orderId)
	req.SetOptionalParam("origClientOrderId", d.origClientOrderId)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("deleteOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*deleteOrderResponse](resp)
}

// CallBatch 批量撤销订单 (TRADE)
func (d *deleteOrderRequest) CallBatch(ctx context.Context, orderIdList []int64) (body []*deleteOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.FApiBatchOrders,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	req.SetOptionalParam("orderIdList", orderIdList)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("CallBatch response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*deleteOrderResponse](resp)
}

// CallAllOpenOrders 撤销全部订单(TRADE)
func (d *deleteOrderRequest) CallAllOpenOrders(ctx context.Context) (body *errors.Status, err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.FApiAllOpenOrders,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("CallAllOpenOrders response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*errors.Status](resp)
}

// ****************************** Websocket Api *******************************
