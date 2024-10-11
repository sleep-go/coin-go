package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type CreateOrder interface {
	SetSymbol(symbol string) *createOrderRequest
	SetSide(side enums.SideType) *createOrderRequest
	SetType(orderType enums.OrderType) *createOrderRequest
	SetTimeInForce(timeInForce enums.TimeInForceType) *createOrderRequest
	SetQuantity(quantity string) *createOrderRequest
	SetQuoteOrderQty(quoteOrderQty string) *createOrderRequest
	SetPrice(price string) *createOrderRequest
	SetNewClientOrderId(newClientOrderId string) *createOrderRequest
	SetStrategyId(strategyId int64) *createOrderRequest
	SetStrategyType(strategyType int64) *createOrderRequest
	SetStopPrice(stopPrice string) *createOrderRequest
	SetTrailingDelta(trailingDelta int64) *createOrderRequest
	SetIcebergQty(icebergQty string) *createOrderRequest
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *createOrderRequest
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *createOrderRequest
	Call(ctx context.Context) (body *createOrderResponse, err error)
	CallTest(ctx context.Context, computeCommissionRates bool) (body *createOrderTestResponse, err error)
}

type createOrderRequest struct {
	*binance.Client
	symbol                  string
	side                    enums.SideType        //订单方向
	_type                   enums.OrderType       //订单类型
	timeInForce             enums.TimeInForceType //生效时间
	quantity                *string
	quoteOrderQty           *string
	price                   *string
	newClientOrderId        *string                //用户自定义的orderId，如空缺系统会自动赋值。
	strategyId              *int64                 //策略ID
	strategyType            *int64                 //策略类型，不能低于 1000000
	stopPrice               *string                //仅 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT 需要此参数。
	trailingDelta           *int64                 //用于 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, 和 TAKE_PROFIT_LIMIT 类型的订单。
	icebergQty              *string                //仅有限价单(包括条件限价单与限价做事单)可以使用该参数，含义为创建冰山订单并指定冰山订单的数量。
	newOrderRespType        enums.NewOrderRespType //指定响应类型 ACK, RESULT, or FULL; MARKET 与 LIMIT 订单默认为FULL, 其他默认为ACK。
	selfTradePreventionMode enums.StpModeType      //允许的 ENUM 取决于交易对的配置。
	recvWindow              int64
	timestamp               int64
}

type createOrderResponse struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int    `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactTime            int64  `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	WorkingTime             int64  `json:"workingTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	Fills                   []struct {
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int    `json:"tradeId"`
	} `json:"fills"`
}

func NewOrder(client *binance.Client, symbol string) CreateOrder {
	return &createOrderRequest{Client: client, symbol: symbol}
}
func (c *createOrderRequest) SetSymbol(symbol string) *createOrderRequest {
	c.symbol = symbol
	return c
}
func (c *createOrderRequest) SetSide(side enums.SideType) *createOrderRequest {
	c.side = side
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
func (c *createOrderRequest) SetQuoteOrderQty(quoteOrderQty string) *createOrderRequest {
	c.quoteOrderQty = &quoteOrderQty
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
func (c *createOrderRequest) SetStrategyId(strategyId int64) *createOrderRequest {
	c.strategyId = &strategyId
	return c
}
func (c *createOrderRequest) SetStrategyType(strategyType int64) *createOrderRequest {
	c.strategyType = &strategyType
	return c
}
func (c *createOrderRequest) SetStopPrice(stopPrice string) *createOrderRequest {
	c.stopPrice = &stopPrice
	return c
}
func (c *createOrderRequest) SetTrailingDelta(trailingDelta int64) *createOrderRequest {
	c.trailingDelta = &trailingDelta
	return c
}
func (c *createOrderRequest) SetIcebergQty(icebergQty string) *createOrderRequest {
	c.icebergQty = &icebergQty
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
func (c *createOrderRequest) Call(ctx context.Context) (body *createOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("side", c.side)
	req.SetParam("type", c._type)
	req.SetParam("quantity", c.quantity)
	req.SetOptionalParam("timeInForce", c.timeInForce)
	req.SetOptionalParam("quoteOrderQty", c.quoteOrderQty)
	req.SetOptionalParam("price", c.price)
	req.SetOptionalParam("newClientOrderId", c.newClientOrderId)
	req.SetOptionalParam("strategyId", c.strategyId)
	req.SetOptionalParam("stopPrice", c.stopPrice)
	req.SetOptionalParam("trailingDelta", c.trailingDelta)
	req.SetOptionalParam("icebergQty", c.icebergQty)
	req.SetOptionalParam("newOrderRespType", c.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", c.selfTradePreventionMode)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("createOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*createOrderResponse](resp)
}

type createOrderTestResponse struct {
	// 订单交易的标准佣金率
	StandardCommissionForOrder struct {
		Maker string `json:"maker"`
		Taker string `json:"taker"`
	} `json:"standardCommissionForOrder"`
	// 订单交易的税率
	TaxCommissionForOrder struct {
		Maker string `json:"maker"`
		Taker string `json:"taker"`
	} `json:"taxCommissionForOrder"`
	// 以BNB支付时的标准佣金折扣。
	Discount struct {
		EnabledForAccount bool   `json:"enabledForAccount"`
		EnabledForSymbol  bool   `json:"enabledForSymbol"`
		DiscountAsset     string `json:"discountAsset"`
		Discount          string `json:"discount"` // 当用BNB支付佣金时，在标准佣金上按此比率打折
	} `json:"discount"`
}

func (c *createOrderRequest) CallTest(ctx context.Context, computeCommissionRates bool) (body *createOrderTestResponse, err error) {
	// 没有 computeCommissionRates返回空
	if computeCommissionRates == false {
		return nil, nil
	}
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingOrderTest,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("side", c.side)
	req.SetParam("type", c._type)
	req.SetParam("computeCommissionRates", computeCommissionRates)
	req.SetParam("quantity", c.quantity)
	req.SetOptionalParam("timeInForce", c.timeInForce)
	req.SetOptionalParam("quoteOrderQty", c.quoteOrderQty)
	req.SetOptionalParam("price", c.price)
	req.SetOptionalParam("newClientOrderId", c.newClientOrderId)
	req.SetOptionalParam("strategyId", c.strategyId)
	req.SetOptionalParam("stopPrice", c.stopPrice)
	req.SetOptionalParam("trailingDelta", c.trailingDelta)
	req.SetOptionalParam("icebergQty", c.icebergQty)
	req.SetOptionalParam("newOrderRespType", c.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", c.selfTradePreventionMode)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("createOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*createOrderTestResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiCreateOrder interface {
	binance.WsApi[WsApiCreateOrderResponse]
	CreateOrder
}
type WsApiCreateOrderResponse struct {
	binance.WsApiResponse
	Result *createOrderResponse `json:"result"`
}

func NewWsApiCreateOrder(c *binance.Client) WsApiCreateOrder {
	return &createOrderRequest{Client: c}
}

func (c *createOrderRequest) Receive(handler binance.Handler[WsApiCreateOrderResponse], exception binance.ErrorHandler) error {
	return binance.WsHandler(c.Client, c.BaseURL, handler, exception)
}

func (c *createOrderRequest) Send() error {
	req := &binance.Request{Path: "order.place"}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("side", c.side)
	req.SetParam("type", c._type)
	req.SetOptionalParam("quantity", c.quantity)
	req.SetOptionalParam("timeInForce", c.timeInForce)
	req.SetOptionalParam("quoteOrderQty", c.quoteOrderQty)
	req.SetOptionalParam("price", c.price)
	req.SetOptionalParam("newClientOrderId", c.newClientOrderId)
	req.SetOptionalParam("strategyId", c.strategyId)
	req.SetOptionalParam("stopPrice", c.stopPrice)
	req.SetOptionalParam("trailingDelta", c.trailingDelta)
	req.SetOptionalParam("icebergQty", c.icebergQty)
	req.SetOptionalParam("newOrderRespType", c.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", c.selfTradePreventionMode)
	return c.SendMessage(req)
}
