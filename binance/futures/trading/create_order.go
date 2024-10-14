package trading

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type CreateOrder interface {
	SetSymbol(symbol string) *CreateOrderRequest
	SetSide(side enums.SideType) *CreateOrderRequest
	SetPositionSide(positionSide enums.PositionSideType) *CreateOrderRequest
	SetReduceOnly(reduceOnly bool) *CreateOrderRequest
	SetClosePosition(closePosition bool) *CreateOrderRequest
	SetActivationPrice(activationPrice string) *CreateOrderRequest
	SetCallbackRate(callbackRate string) *CreateOrderRequest
	SetWorkingType(workingType enums.WorkingType) *CreateOrderRequest
	SetPriceProtect(priceProtect string) *CreateOrderRequest
	SetPriceMatch(priceMatch enums.PriceMatchType) *CreateOrderRequest
	SetGoodTillDate(goodTillDate int) *CreateOrderRequest
	SetType(orderType enums.OrderType) *CreateOrderRequest
	SetTimeInForce(timeInForce enums.TimeInForceType) *CreateOrderRequest
	SetQuantity(quantity string) *CreateOrderRequest
	SetPrice(price string) *CreateOrderRequest
	SetNewClientOrderId(newClientOrderId string) *CreateOrderRequest
	SetStopPrice(stopPrice string) *CreateOrderRequest
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *CreateOrderRequest
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *CreateOrderRequest
	Call(ctx context.Context) (body *createOrderResponse, err error)
	CallTest(ctx context.Context) (body *createOrderResponse, err error)
	CallBatch(ctx context.Context, data []*CreateOrderRequest) (body []*createOrderResponse, err error)
}
type CreateOrderRequest struct {
	*binance.Client
	Symbol                  string                 `json:"symbol,omitempty"`
	Side                    enums.SideType         `json:"side,omitempty"`                    //订单方向
	PositionSide            enums.PositionSideType `json:"positionSide,omitempty"`            //持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	Type                    enums.OrderType        `json:"type,omitempty"`                    //订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	ReduceOnly              *bool                  `json:"reduceOnly,omitempty"`              //true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
	Quantity                *string                `json:"quantity"`                          //下单数量,使用closePosition不支持此参数。
	Price                   *string                `json:"price,omitempty"`                   //委托价格
	NewClientOrderId        *string                `json:"newClientOrderId,omitempty"`        //用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
	StopPrice               *string                `json:"stopPrice,omitempty"`               //触发价, 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	ClosePosition           *bool                  `json:"closePosition,omitempty"`           //true, false；触发后全部平仓，仅支持STOP_MARKET和TAKE_PROFIT_MARKET；不与quantity合用；自带只平仓效果，不与reduceOnly 合用
	ActivationPrice         *string                `json:"activationPrice,omitempty"`         //追踪止损激活价格，仅TRAILING_STOP_MARKET 需要此参数, 默认为下单当前市场价格(支持不同workingType)
	CallbackRate            *string                `json:"callbackRate,omitempty"`            //追踪止损回调比例，可取值范围[0.1, 5],其中 1代表1% ,仅TRAILING_STOP_MARKET 需要此参数
	TimeInForce             enums.TimeInForceType  `json:"timeInForce,omitempty"`             //有效方法
	WorkingType             enums.WorkingType      `json:"workingType,omitempty"`             //stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	PriceProtect            *string                `json:"priceProtect,omitempty"`            //条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	NewOrderRespType        enums.NewOrderRespType `json:"newOrderRespType,omitempty"`        //指定响应类型 ACK, RESULT, or FULL; MARKET 与 LIMIT 订单默认为FULL, 其他默认为ACK。
	PriceMatch              enums.PriceMatchType   `json:"priceMatch,omitempty"`              //OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	SelfTradePreventionMode enums.StpModeType      `json:"selfTradePreventionMode,omitempty"` //允许的 ENUM 取决于交易对的配置。
	GoodTillDate            *int                   `json:"goodTillDate,omitempty"`            //TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
}
type createOrderResponse struct {
	ClientOrderId           string                 `json:"clientOrderId"`
	CumQty                  string                 `json:"cumQty"`
	CumQuote                string                 `json:"cumQuote"`
	ExecutedQty             string                 `json:"executedQty"`
	OrderId                 int                    `json:"orderId"`
	AvgPrice                string                 `json:"avgPrice"`
	OrigQty                 string                 `json:"origQty"`
	Price                   string                 `json:"price"`
	ReduceOnly              bool                   `json:"reduceOnly"`
	Side                    enums.SideType         `json:"side"`
	PositionSide            enums.PositionSideType `json:"positionSide"`
	Status                  enums.StatusType       `json:"status"`
	StopPrice               string                 `json:"stopPrice"`
	ClosePosition           bool                   `json:"closePosition"`
	Symbol                  string                 `json:"symbol"`
	TimeInForce             enums.TimeInForceType  `json:"timeInForce"`
	Type                    enums.OrderType        `json:"type"`
	OrigType                enums.OrderType        `json:"origType"`
	ActivatePrice           string                 `json:"activatePrice"`
	PriceRate               string                 `json:"priceRate"`
	UpdateTime              int64                  `json:"updateTime"`
	WorkingType             enums.WorkingType      `json:"workingType"`
	PriceProtect            bool                   `json:"priceProtect"`
	PriceMatch              enums.PriceMatchType   `json:"priceMatch"`
	SelfTradePreventionMode enums.StpModeType      `json:"selfTradePreventionMode"`
	GoodTillDate            int64                  `json:"goodTillDate"`
}

func (c *CreateOrderRequest) SetSymbol(symbol string) *CreateOrderRequest {
	c.Symbol = symbol
	return c
}
func (c *CreateOrderRequest) SetSide(side enums.SideType) *CreateOrderRequest {
	c.Side = side
	return c
}
func (c *CreateOrderRequest) SetPositionSide(positionSide enums.PositionSideType) *CreateOrderRequest {
	c.PositionSide = positionSide
	return c
}
func (c *CreateOrderRequest) SetReduceOnly(reduceOnly bool) *CreateOrderRequest {
	c.ReduceOnly = &reduceOnly
	return c
}
func (c *CreateOrderRequest) SetClosePosition(closePosition bool) *CreateOrderRequest {
	c.ClosePosition = &closePosition
	return c
}
func (c *CreateOrderRequest) SetActivationPrice(activationPrice string) *CreateOrderRequest {
	c.ActivationPrice = &activationPrice
	return c
}
func (c *CreateOrderRequest) SetCallbackRate(callbackRate string) *CreateOrderRequest {
	c.CallbackRate = &callbackRate
	return c
}
func (c *CreateOrderRequest) SetWorkingType(workingType enums.WorkingType) *CreateOrderRequest {
	c.WorkingType = workingType
	return c
}
func (c *CreateOrderRequest) SetPriceProtect(priceProtect string) *CreateOrderRequest {
	c.PriceProtect = &priceProtect
	return c
}
func (c *CreateOrderRequest) SetPriceMatch(priceMatch enums.PriceMatchType) *CreateOrderRequest {
	c.PriceMatch = priceMatch
	return c
}
func (c *CreateOrderRequest) SetType(_type enums.OrderType) *CreateOrderRequest {
	c.Type = _type
	return c
}
func (c *CreateOrderRequest) SetTimeInForce(timeInForce enums.TimeInForceType) *CreateOrderRequest {
	c.TimeInForce = timeInForce
	return c
}
func (c *CreateOrderRequest) SetQuantity(quantity string) *CreateOrderRequest {
	c.Quantity = &quantity
	return c
}
func (c *CreateOrderRequest) SetPrice(price string) *CreateOrderRequest {
	c.Price = &price
	return c
}
func (c *CreateOrderRequest) SetNewClientOrderId(newClientOrderId string) *CreateOrderRequest {
	c.NewClientOrderId = &newClientOrderId
	return c
}
func (c *CreateOrderRequest) SetStopPrice(stopPrice string) *CreateOrderRequest {
	c.StopPrice = &stopPrice
	return c
}
func (c *CreateOrderRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *CreateOrderRequest {
	c.NewOrderRespType = newOrderRespType
	return c
}
func (c *CreateOrderRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *CreateOrderRequest {
	c.SelfTradePreventionMode = selfTradePreventionMode
	return c
}
func (c *CreateOrderRequest) SetGoodTillDate(goodTillDate int) *CreateOrderRequest {
	c.GoodTillDate = &goodTillDate
	return c
}

func NewOrder(client *binance.Client, symbol string) CreateOrder {
	return &CreateOrderRequest{Client: client, Symbol: symbol}
}
func (c *CreateOrderRequest) Call(ctx context.Context) (body *createOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.FApiOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.Symbol)
	req.SetParam("side", c.Side)
	req.SetOptionalParam("positionSide", c.PositionSide)
	req.SetParam("type", c.Type)
	req.SetOptionalParam("reduceOnly", c.ReduceOnly)
	req.SetOptionalParam("quantity", c.Quantity)
	req.SetOptionalParam("price", c.Price)
	req.SetOptionalParam("newClientOrderId", c.NewClientOrderId)
	req.SetOptionalParam("stopPrice", c.StopPrice)
	req.SetOptionalParam("closePosition", c.ClosePosition)
	req.SetOptionalParam("activationPrice", c.ActivationPrice)
	req.SetOptionalParam("callbackRate", c.CallbackRate)
	req.SetOptionalParam("timeInForce", c.TimeInForce)
	req.SetOptionalParam("workingType", c.WorkingType)
	req.SetOptionalParam("priceProtect", c.PriceProtect)
	req.SetOptionalParam("newOrderRespType", c.NewOrderRespType)
	req.SetOptionalParam("priceMatch", c.PriceMatch)
	req.SetOptionalParam("selfTradePreventionMode", c.SelfTradePreventionMode)
	req.SetOptionalParam("goodTillDate", c.GoodTillDate)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("createOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*createOrderResponse](resp)
}

// CallBatch 批量下单(TRADE)
func (c *CreateOrderRequest) CallBatch(ctx context.Context, data []*CreateOrderRequest) (body []*createOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.FApiBatchOrders,
	}
	req.SetNeedSign(true)
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req.SetParam("batchOrders", string(bytes))
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("CallBatch response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*createOrderResponse](resp)
}

func (c *CreateOrderRequest) CallTest(ctx context.Context) (body *createOrderResponse, err error) {
	// 没有 computeCommissionRates返回空
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.FApiTradingOrderTest,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.Symbol)
	req.SetParam("side", c.Side)
	req.SetOptionalParam("positionSide", c.PositionSide)
	req.SetParam("type", c.Type)
	req.SetOptionalParam("reduceOnly", c.ReduceOnly)
	req.SetOptionalParam("quantity", c.Quantity)
	req.SetOptionalParam("price", c.Price)
	req.SetOptionalParam("newClientOrderId", c.NewClientOrderId)
	req.SetOptionalParam("stopPrice", c.StopPrice)
	req.SetOptionalParam("closePosition", c.ClosePosition)
	req.SetOptionalParam("activationPrice", c.ActivationPrice)
	req.SetOptionalParam("callbackRate", c.CallbackRate)
	req.SetOptionalParam("timeInForce", c.TimeInForce)
	req.SetOptionalParam("workingType", c.WorkingType)
	req.SetOptionalParam("priceProtect", c.PriceProtect)
	req.SetOptionalParam("newOrderRespType", c.NewOrderRespType)
	req.SetOptionalParam("priceMatch", c.PriceMatch)
	req.SetOptionalParam("selfTradePreventionMode", c.SelfTradePreventionMode)
	req.SetOptionalParam("goodTillDate", c.GoodTillDate)

	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("createOrderTestRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*createOrderResponse](resp)
}

// ****************************** Websocket Api *******************************
