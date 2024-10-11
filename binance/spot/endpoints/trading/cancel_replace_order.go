package trading

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

type CancelReplace interface {
	SetSymbol(symbol string) *cancelReplaceRequest
	SetSide(side enums.SideType) *cancelReplaceRequest
	SetType(_type enums.OrderType) *cancelReplaceRequest
	SetTimeInForce(timeInForce enums.TimeInForceType) *cancelReplaceRequest
	SetQuantity(quantity string) *cancelReplaceRequest
	SetQuoteOrderQty(quoteOrderQty string) *cancelReplaceRequest
	SetPrice(price string) *cancelReplaceRequest
	SetNewClientOrderId(newClientOrderId string) *cancelReplaceRequest
	SetStrategyId(strategyId int64) *cancelReplaceRequest
	SetStrategyType(strategyType int64) *cancelReplaceRequest
	SetStopPrice(stopPrice string) *cancelReplaceRequest
	SetTrailingDelta(trailingDelta int64) *cancelReplaceRequest
	SetIcebergQty(icebergQty string) *cancelReplaceRequest
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *cancelReplaceRequest
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *cancelReplaceRequest

	SetCancelReplaceMode(cancelReplaceMode enums.CancelReplaceModeType) *cancelReplaceRequest
	SetCancelNewClientOrderId(cancelNewClientOrderId string) *cancelReplaceRequest
	SetCancelOrigClientOrderId(cancelOrigClientOrderId string) *cancelReplaceRequest
	SetCancelOrderId(cancelOrderId int64) *cancelReplaceRequest
	SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) *cancelReplaceRequest
	SetOrderRateLimitExceededMode(orderRateLimitExceededMode enums.OrderRateLimitExceededModeType) *cancelReplaceRequest
	Call(ctx context.Context) (body *cancelReplaceResponse, err error)
}

type cancelReplaceRequest struct {
	*binance.Client
	createOrderRequest
	//指定类型：STOP_ON_FAILURE - 如果撤消订单失败将不会继续重新下单。
	//ALLOW_FAILURE - 不管撤消订单是否成功都会继续重新下单。
	cancelReplaceMode enums.CancelReplaceModeType
	//用户自定义的id，如空缺系统会自动赋值
	cancelNewClientOrderId *string
	//必须提供cancelOrigClientOrderId 或者 cancelOrderId。 如果两个参数都提供, cancelOrderId 会占优先。
	cancelOrigClientOrderId *string
	//必须提供cancelOrigClientOrderId 或者 cancelOrderId。 如果两个参数都提供，cancelOrderId 会占优先。
	cancelOrderId      *int64
	cancelRestrictions enums.CancelRestrictionsType
	//“DO_NOTHING”（默认值）- 仅在账户未超过未成交订单频率限制时，会尝试取消订单。
	//“CANCEL_ONLY” - 将始终取消订单。
	orderRateLimitExceededMode enums.OrderRateLimitExceededModeType
}

type cancelReplaceResponse struct {
	Code           int64  `json:"code,omitempty"`
	Msg            string `json:"msg,omitempty"`
	CancelResult   string `json:"cancelResult,omitempty"`
	NewOrderResult string `json:"newOrderResult,omitempty"`
	CancelResponse *struct {
		Code                    int    `json:"code,omitempty"`
		Msg                     string `json:"msg,omitempty"`
		Symbol                  string `json:"symbol,omitempty"`
		OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
		OrderId                 int64  `json:"orderId,omitempty"`
		OrderListId             int64  `json:"orderListId,omitempty"`
		ClientOrderId           string `json:"clientOrderId,omitempty"`
		Price                   string `json:"price,omitempty"`
		OrigQty                 string `json:"origQty,omitempty"`
		ExecutedQty             string `json:"executedQty,omitempty"`
		CumulativeQuoteQty      string `json:"cumulativeQuoteQty,omitempty"`
		Status                  string `json:"status,omitempty"`
		TimeInForce             string `json:"timeInForce,omitempty"`
		Type                    string `json:"type,omitempty"`
		Side                    string `json:"side,omitempty"`
		SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
	} `json:"cancelResponse,omitempty"`
	NewOrderResponse *struct {
		Code                    int64    `json:"code,omitempty"`
		Msg                     string   `json:"msg,omitempty"`
		Symbol                  string   `json:"symbol,omitempty"`
		OrderId                 int64    `json:"orderId,omitempty"`
		OrderListId             int64    `json:"orderListId,omitempty"`
		ClientOrderId           string   `json:"clientOrderId,omitempty"`
		TransactTime            uint64   `json:"transactTime,omitempty"`
		Price                   string   `json:"price,omitempty"`
		OrigQty                 string   `json:"origQty,omitempty"`
		ExecutedQty             string   `json:"executedQty,omitempty"`
		CumulativeQuoteQty      string   `json:"cumulativeQuoteQty,omitempty"`
		Status                  string   `json:"status,omitempty"`
		TimeInForce             string   `json:"timeInForce,omitempty"`
		Type                    string   `json:"type,omitempty"`
		Side                    string   `json:"side,omitempty"`
		Fills                   []string `json:"fills,omitempty"`
		SelfTradePreventionMode string   `json:"selfTradePreventionMode,omitempty"`
	} `json:"newOrderResponse,omitempty"`
	Data *struct {
		CancelResult   string `json:"cancelResult,omitempty"`
		NewOrderResult string `json:"newOrderResult,omitempty"`
		CancelResponse *struct {
			Code                    int64  `json:"code,omitempty"`
			Msg                     string `json:"msg,omitempty"`
			Symbol                  string `json:"symbol,omitempty"`
			OrigClientOrderId       string `json:"origClientOrderId,omitempty"`
			OrderId                 int64  `json:"orderId,omitempty"`
			OrderListId             int64  `json:"orderListId,omitempty"`
			ClientOrderId           string `json:"clientOrderId,omitempty"`
			Price                   string `json:"price,omitempty"`
			OrigQty                 string `json:"origQty,omitempty"`
			ExecutedQty             string `json:"executedQty,omitempty"`
			CumulativeQuoteQty      string `json:"cumulativeQuoteQty,omitempty"`
			Status                  string `json:"status,omitempty"`
			TimeInForce             string `json:"timeInForce,omitempty"`
			Type                    string `json:"type,omitempty"`
			Side                    string `json:"side,omitempty"`
			SelfTradePreventionMode string `json:"selfTradePreventionMode,omitempty"`
		} `json:"cancelResponse,omitempty"`
		NewOrderResponse struct {
			Code                    int64    `json:"code,omitempty"`
			Msg                     string   `json:"msg,omitempty"`
			Symbol                  string   `json:"symbol,omitempty"`
			OrderId                 int64    `json:"orderId,omitempty"`
			OrderListId             int64    `json:"orderListId,omitempty"`
			ClientOrderId           string   `json:"clientOrderId,omitempty"`
			TransactTime            uint64   `json:"transactTime,omitempty"`
			Price                   string   `json:"price,omitempty"`
			OrigQty                 string   `json:"origQty,omitempty"`
			ExecutedQty             string   `json:"executedQty,omitempty"`
			CumulativeQuoteQty      string   `json:"cumulativeQuoteQty,omitempty"`
			Status                  string   `json:"status,omitempty"`
			TimeInForce             string   `json:"timeInForce,omitempty"`
			Type                    string   `json:"type,omitempty"`
			Side                    string   `json:"side,omitempty"`
			Fills                   []string `json:"fills,omitempty"`
			SelfTradePreventionMode string   `json:"selfTradePreventionMode,omitempty"`
		} `json:"newOrderResponse"`
	} `json:"data,omitempty"`
}

func NewCancelReplace(client *binance.Client, symbol string) CancelReplace {
	return &cancelReplaceRequest{
		Client:             client,
		createOrderRequest: createOrderRequest{symbol: symbol},
	}
}

func (c *cancelReplaceRequest) SetCancelReplaceMode(cancelReplaceMode enums.CancelReplaceModeType) *cancelReplaceRequest {
	c.cancelReplaceMode = cancelReplaceMode
	return c
}

func (c *cancelReplaceRequest) SetCancelNewClientOrderId(cancelNewClientOrderId string) *cancelReplaceRequest {
	c.cancelNewClientOrderId = &cancelNewClientOrderId
	return c
}

func (c *cancelReplaceRequest) SetCancelOrigClientOrderId(cancelOrigClientOrderId string) *cancelReplaceRequest {
	c.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return c
}

func (c *cancelReplaceRequest) SetCancelOrderId(cancelOrderId int64) *cancelReplaceRequest {
	c.cancelOrderId = &cancelOrderId
	return c
}

func (c *cancelReplaceRequest) SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) *cancelReplaceRequest {
	c.cancelRestrictions = cancelRestrictions
	return c
}

func (c *cancelReplaceRequest) SetOrderRateLimitExceededMode(orderRateLimitExceededMode enums.OrderRateLimitExceededModeType) *cancelReplaceRequest {
	c.orderRateLimitExceededMode = orderRateLimitExceededMode
	return c
}
func (c *cancelReplaceRequest) SetSymbol(symbol string) *cancelReplaceRequest {
	c.symbol = symbol
	return c
}
func (c *cancelReplaceRequest) SetSide(side enums.SideType) *cancelReplaceRequest {
	c.side = side
	return c
}

func (c *cancelReplaceRequest) SetType(_type enums.OrderType) *cancelReplaceRequest {
	c._type = _type
	return c
}

func (c *cancelReplaceRequest) SetTimeInForce(timeInForce enums.TimeInForceType) *cancelReplaceRequest {
	c.timeInForce = timeInForce
	return c
}

func (c *cancelReplaceRequest) SetQuantity(quantity string) *cancelReplaceRequest {
	c.quantity = &quantity
	return c
}

func (c *cancelReplaceRequest) SetQuoteOrderQty(quoteOrderQty string) *cancelReplaceRequest {
	c.quoteOrderQty = &quoteOrderQty
	return c
}

func (c *cancelReplaceRequest) SetPrice(price string) *cancelReplaceRequest {
	c.price = &price
	return c
}

func (c *cancelReplaceRequest) SetNewClientOrderId(newClientOrderId string) *cancelReplaceRequest {
	c.newClientOrderId = &newClientOrderId
	return c
}

func (c *cancelReplaceRequest) SetStrategyId(strategyId int64) *cancelReplaceRequest {
	c.strategyId = &strategyId
	return c
}

func (c *cancelReplaceRequest) SetStrategyType(strategyType int64) *cancelReplaceRequest {
	c.strategyType = &strategyType
	return c
}

func (c *cancelReplaceRequest) SetStopPrice(stopPrice string) *cancelReplaceRequest {
	c.stopPrice = &stopPrice
	return c
}

func (c *cancelReplaceRequest) SetTrailingDelta(trailingDelta int64) *cancelReplaceRequest {
	c.trailingDelta = &trailingDelta
	return c
}

func (c *cancelReplaceRequest) SetIcebergQty(icebergQty string) *cancelReplaceRequest {
	c.icebergQty = &icebergQty
	return c
}

func (c *cancelReplaceRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *cancelReplaceRequest {
	c.newOrderRespType = newOrderRespType
	return c
}

func (c *cancelReplaceRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *cancelReplaceRequest {
	c.selfTradePreventionMode = selfTradePreventionMode
	return c
}

// Call 撤消挂单再下单 (TRADE)
// 撤消挂单并在同个交易对上重新下单。
//
// 在撤消订单和下单前会判断: 1) 过滤器参数, 以及 2) 目前下单数量。
//
// 即使请求中没有尝试发送新订单，比如(newOrderResult: NOT_ATTEMPTED)，下单的数量仍然会加1。
func (c *cancelReplaceRequest) Call(ctx context.Context) (body *cancelReplaceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingCancelReplace,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("side", c.side)
	req.SetParam("type", c._type)
	req.SetOptionalParam("timeInForce", c.timeInForce)
	req.SetParam("quantity", c.quantity)
	req.SetOptionalParam("quoteOrderQty", c.quoteOrderQty)
	req.SetOptionalParam("price", c.price)
	req.SetOptionalParam("newClientOrderId", c.newClientOrderId)
	req.SetOptionalParam("strategyId", c.strategyId)
	req.SetOptionalParam("stopPrice", c.stopPrice)
	req.SetOptionalParam("trailingDelta", c.trailingDelta)
	req.SetOptionalParam("icebergQty", c.icebergQty)
	req.SetOptionalParam("newOrderRespType", c.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", c.selfTradePreventionMode)
	req.SetOptionalParam("cancelNewClientOrderId", c.cancelNewClientOrderId)
	req.SetOptionalParam("cancelReplaceMode", c.cancelReplaceMode)
	req.SetOptionalParam("cancelOrigClientOrderId", c.cancelOrigClientOrderId)
	req.SetOptionalParam("cancelOrderId", c.cancelOrderId)
	req.SetOptionalParam("cancelRestrictions", c.cancelRestrictions)
	req.SetOptionalParam("orderRateLimitExceededMode", c.orderRateLimitExceededMode)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("cancelReplaceRequest response err:%v", err)
		return nil, err
	}
	//{"code":-2022,"msg":"Order cancel-replace failed.","data":{"cancelResult":"FAILURE","newOrderResult":"NOT_ATTEMPTED","cancelResponse":{"code":-2011,"msg":"Unknown order sent."},"newOrderResponse":null}}
	//因为这个返回值跟其他普通错误时返回不通结构，所以单独处理
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		c.Debugf("ParseHttpResponse err:%v", err.Error())
		return nil, err
	}
	return body, nil
}

// ****************************** Websocket Api *******************************

type WsApiCancelReplace interface {
	binance.WsApi[*WsApiCancelReplaceResponse]
	CancelReplace
}
type WsApiCancelReplaceResponse struct {
	binance.WsApiResponse
	Result *cancelReplaceResponse `json:"result"`
}

// NewWsApiCancelReplace 撤消挂单再下单 (TRADE)
// 撤消挂单并在同个交易对上重新下单。
// 类似于 order.place 请求，额外的强制参数 (*) 由新订单的 type 确定。
//
// 可用的 cancelReplaceMode 选项：
//
// STOP_ON_FAILURE – 如果撤销订单请求失败，将不会尝试下新订单。
// ALLOW_FAILURE – 即使撤销订单请求失败，也会尝试下新订单。
// 备注：
//
// 如果同时指定了 cancelOrderId 和 cancelOrigClientOrderId 参数，仅使用 cancelOrderId 并忽略 cancelOrigClientOrderId。
//
// cancelNewClientOrderId 将替换已撤销订单的 clientOrderId，为新订单腾出空间。
//
// newClientOrderId 指定下单的 clientOrderId 值。
//
// 仅当前一个订单已成交或过期时，才会接受具有相同 clientOrderId 的新订单。
//
// 新订单可以重用已取消订单的旧 clientOrderId。
//
// 此 cancel-replace 操作不是事务性的。
//
// 如果一个操作成功但另一个操作失败，则仍然执行成功的操作。
//
// 例如，在 STOP_ON_FAILURE 模式下，如果下新订单达失败，旧订单仍然被撤销。
//
// 过滤器和订单次数限制会在撤销和下订单之前评估。
//
// 如果未尝试下新订单，订单次数仍会增加。
//
// 与 order.cancel 一样，如果您撤销订单列表内的某个订单，则整个订单列表将被撤销。
func NewWsApiCancelReplace(c *binance.Client) WsApiCancelReplace {
	return &cancelReplaceRequest{Client: c}
}

func (c *cancelReplaceRequest) Send(ctx context.Context) (*WsApiCancelReplaceResponse, error) {
	req := &binance.Request{Path: "order.cancelReplace"}
	req.SetNeedSign(true)
	req.SetOptionalParam("symbol", c.symbol)
	req.SetOptionalParam("side", c.side)
	req.SetOptionalParam("type", c._type)
	req.SetOptionalParam("timeInForce", c.timeInForce)
	req.SetOptionalParam("quantity", c.quantity)
	req.SetOptionalParam("quoteOrderQty", c.quoteOrderQty)
	req.SetOptionalParam("price", c.price)
	req.SetOptionalParam("newClientOrderId", c.newClientOrderId)
	req.SetOptionalParam("strategyId", c.strategyId)
	req.SetOptionalParam("stopPrice", c.stopPrice)
	req.SetOptionalParam("trailingDelta", c.trailingDelta)
	req.SetOptionalParam("icebergQty", c.icebergQty)
	req.SetOptionalParam("newOrderRespType", c.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", c.selfTradePreventionMode)
	req.SetOptionalParam("cancelNewClientOrderId", c.cancelNewClientOrderId)
	req.SetOptionalParam("cancelReplaceMode", c.cancelReplaceMode)
	req.SetOptionalParam("cancelOrigClientOrderId", c.cancelOrigClientOrderId)
	req.SetOptionalParam("cancelOrderId", c.cancelOrderId)
	req.SetOptionalParam("cancelRestrictions", c.cancelRestrictions)
	req.SetOptionalParam("orderRateLimitExceededMode", c.orderRateLimitExceededMode)
	return binance.WsApiHandler[*WsApiCancelReplaceResponse](ctx, c.Client, req)
}
