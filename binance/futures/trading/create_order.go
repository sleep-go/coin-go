package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type CreateOrder interface {
	SetSymbol(symbol string) *createOrderRequest
	SetSide(side enums.SideType) *createOrderRequest
	SetPositionSide(positionSide enums.PositionSideType) *createOrderRequest
	SetReduceOnly(reduceOnly bool) *createOrderRequest
	SetClosePosition(closePosition bool) *createOrderRequest
	SetActivationPrice(activationPrice string) *createOrderRequest
	SetCallbackRate(callbackRate string) *createOrderRequest
	SetWorkingType(workingType enums.WorkingType) *createOrderRequest
	SetPriceProtect(priceProtect string) *createOrderRequest
	SetPriceMatch(priceMatch enums.PriceMatchType) *createOrderRequest
	SetGoodTillDate(goodTillDate int) *createOrderRequest
	SetType(orderType enums.OrderType) *createOrderRequest
	SetTimeInForce(timeInForce enums.TimeInForceType) *createOrderRequest
	SetQuantity(quantity string) *createOrderRequest
	SetPrice(price string) *createOrderRequest
	SetNewClientOrderId(newClientOrderId string) *createOrderRequest
	SetStopPrice(stopPrice string) *createOrderRequest
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *createOrderRequest
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *createOrderRequest
	Call(ctx context.Context) (body *createOrderResponse, err error)
	CallTest(ctx context.Context) (body *createOrderResponse, err error)
}
type createOrderRequest struct {
	*binance.Client
	symbol                  string
	side                    enums.SideType         //订单方向
	positionSide            enums.PositionSideType //持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	_type                   enums.OrderType        //订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	reduceOnly              *bool                  //true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
	quantity                *string                //下单数量,使用closePosition不支持此参数。
	price                   *string                //委托价格
	newClientOrderId        *string                //用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
	stopPrice               *string                //触发价, 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	closePosition           *bool                  //true, false；触发后全部平仓，仅支持STOP_MARKET和TAKE_PROFIT_MARKET；不与quantity合用；自带只平仓效果，不与reduceOnly 合用
	activationPrice         *string                //追踪止损激活价格，仅TRAILING_STOP_MARKET 需要此参数, 默认为下单当前市场价格(支持不同workingType)
	callbackRate            *string                //追踪止损回调比例，可取值范围[0.1, 5],其中 1代表1% ,仅TRAILING_STOP_MARKET 需要此参数
	timeInForce             enums.TimeInForceType  //有效方法
	workingType             enums.WorkingType      //stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	priceProtect            *string                //条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	newOrderRespType        enums.NewOrderRespType //指定响应类型 ACK, RESULT, or FULL; MARKET 与 LIMIT 订单默认为FULL, 其他默认为ACK。
	priceMatch              enums.PriceMatchType   //OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	selfTradePreventionMode enums.StpModeType      //允许的 ENUM 取决于交易对的配置。
	goodTillDate            *int                   //TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
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

func (c *createOrderRequest) SetSymbol(symbol string) *createOrderRequest {
	c.symbol = symbol
	return c
}
func (c *createOrderRequest) SetSide(side enums.SideType) *createOrderRequest {
	c.side = side
	return c
}
func (c *createOrderRequest) SetPositionSide(positionSide enums.PositionSideType) *createOrderRequest {
	c.positionSide = positionSide
	return c
}
func (c *createOrderRequest) SetReduceOnly(reduceOnly bool) *createOrderRequest {
	c.reduceOnly = &reduceOnly
	return c
}
func (c *createOrderRequest) SetClosePosition(closePosition bool) *createOrderRequest {
	c.closePosition = &closePosition
	return c
}
func (c *createOrderRequest) SetActivationPrice(activationPrice string) *createOrderRequest {
	c.activationPrice = &activationPrice
	return c
}
func (c *createOrderRequest) SetCallbackRate(callbackRate string) *createOrderRequest {
	c.callbackRate = &callbackRate
	return c
}
func (c *createOrderRequest) SetWorkingType(workingType enums.WorkingType) *createOrderRequest {
	c.workingType = workingType
	return c
}
func (c *createOrderRequest) SetPriceProtect(priceProtect string) *createOrderRequest {
	c.priceProtect = &priceProtect
	return c
}
func (c *createOrderRequest) SetPriceMatch(priceMatch enums.PriceMatchType) *createOrderRequest {
	c.priceMatch = priceMatch
	return c
}
func (c *createOrderRequest) SetType(_type enums.OrderType) *createOrderRequest {
	c._type = _type
	return c
}
func (c *createOrderRequest) SetTimeInForce(timeInForce enums.TimeInForceType) *createOrderRequest {
	c.timeInForce = timeInForce
	return c
}
func (c *createOrderRequest) SetQuantity(quantity string) *createOrderRequest {
	c.quantity = &quantity
	return c
}
func (c *createOrderRequest) SetPrice(price string) *createOrderRequest {
	c.price = &price
	return c
}
func (c *createOrderRequest) SetNewClientOrderId(newClientOrderId string) *createOrderRequest {
	c.newClientOrderId = &newClientOrderId
	return c
}
func (c *createOrderRequest) SetStopPrice(stopPrice string) *createOrderRequest {
	c.stopPrice = &stopPrice
	return c
}
func (c *createOrderRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *createOrderRequest {
	c.newOrderRespType = newOrderRespType
	return c
}
func (c *createOrderRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *createOrderRequest {
	c.selfTradePreventionMode = selfTradePreventionMode
	return c
}
func (c *createOrderRequest) SetGoodTillDate(goodTillDate int) *createOrderRequest {
	c.goodTillDate = &goodTillDate
	return c
}

func NewOrder(client *binance.Client, symbol string) CreateOrder {
	return &createOrderRequest{Client: client, symbol: symbol}
}
func (c *createOrderRequest) Call(ctx context.Context) (body *createOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.FApiOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("side", c.side)
	req.SetOptionalParam("positionSide", c.positionSide)
	req.SetParam("type", c._type)
	req.SetOptionalParam("reduceOnly", c.reduceOnly)
	req.SetOptionalParam("quantity", c.quantity)
	req.SetOptionalParam("price", c.price)
	req.SetOptionalParam("newClientOrderId", c.newClientOrderId)
	req.SetOptionalParam("stopPrice", c.stopPrice)
	req.SetOptionalParam("closePosition", c.closePosition)
	req.SetOptionalParam("activationPrice", c.activationPrice)
	req.SetOptionalParam("callbackRate", c.callbackRate)
	req.SetOptionalParam("timeInForce", c.timeInForce)
	req.SetOptionalParam("workingType", c.workingType)
	req.SetOptionalParam("priceProtect", c.priceProtect)
	req.SetOptionalParam("newOrderRespType", c.newOrderRespType)
	req.SetOptionalParam("priceMatch", c.priceMatch)
	req.SetOptionalParam("selfTradePreventionMode", c.selfTradePreventionMode)
	req.SetOptionalParam("goodTillDate", c.goodTillDate)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("createOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*createOrderResponse](resp)
}
func (c *createOrderRequest) CallTest(ctx context.Context) (body *createOrderResponse, err error) {
	// 没有 computeCommissionRates返回空
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.FApiTradingOrderTest,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("side", c.side)
	req.SetOptionalParam("positionSide", c.positionSide)
	req.SetParam("type", c._type)
	req.SetOptionalParam("reduceOnly", c.reduceOnly)
	req.SetOptionalParam("quantity", c.quantity)
	req.SetOptionalParam("price", c.price)
	req.SetOptionalParam("newClientOrderId", c.newClientOrderId)
	req.SetOptionalParam("stopPrice", c.stopPrice)
	req.SetOptionalParam("closePosition", c.closePosition)
	req.SetOptionalParam("activationPrice", c.activationPrice)
	req.SetOptionalParam("callbackRate", c.callbackRate)
	req.SetOptionalParam("timeInForce", c.timeInForce)
	req.SetOptionalParam("workingType", c.workingType)
	req.SetOptionalParam("priceProtect", c.priceProtect)
	req.SetOptionalParam("newOrderRespType", c.newOrderRespType)
	req.SetOptionalParam("priceMatch", c.priceMatch)
	req.SetOptionalParam("selfTradePreventionMode", c.selfTradePreventionMode)
	req.SetOptionalParam("goodTillDate", c.goodTillDate)

	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("createOrderTestRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*createOrderResponse](resp)
}

// ****************************** Websocket Api *******************************
