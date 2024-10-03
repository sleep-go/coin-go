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
	SetSide(side enums.SideType) CancelReplace
	SetType(_type enums.OrderType) CancelReplace
	SetTimeInForce(timeInForce enums.TimeInForceType) CancelReplace
	SetQuantity(quantity string) CancelReplace
	SetQuoteOrderQty(quoteOrderQty string) CancelReplace
	SetPrice(price string) CancelReplace
	SetNewClientOrderId(newClientOrderId string) CancelReplace
	SetStrategyId(strategyId int64) CancelReplace
	SetStrategyType(strategyType int64) CancelReplace
	SetStopPrice(stopPrice string) CancelReplace
	SetTrailingDelta(trailingDelta int64) CancelReplace
	SetIcebergQty(icebergQty string) CancelReplace
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) CancelReplace
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) CancelReplace
	SetRecvWindow(recvWindow int64) CancelReplace
	SetTimestamp(timestamp int64) CancelReplace

	SetCancelReplaceMode(cancelReplaceMode enums.CancelReplaceModeType) CancelReplace
	SetCancelNewClientOrderId(cancelNewClientOrderId string) CancelReplace
	SetCancelOrigClientOrderId(cancelOrigClientOrderId string) CancelReplace
	SetCancelOrderId(cancelOrderId int64) CancelReplace
	SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) CancelReplace
	SetOrderRateLimitExceededMode(orderRateLimitExceededMode enums.OrderRateLimitExceededModeType) CancelReplace
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

type Data struct {
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
}
type cancelReplaceResponse struct {
	Code  int64  `json:"code,omitempty"`
	Msg   string `json:"msg,omitempty"`
	*Data `json:"data,omitempty"`
}

func NewCancelReplace(client *binance.Client, symbol string) CancelReplace {
	return &cancelReplaceRequest{
		Client:             client,
		createOrderRequest: createOrderRequest{symbol: symbol},
	}
}

func (c *cancelReplaceRequest) SetCancelReplaceMode(cancelReplaceMode enums.CancelReplaceModeType) CancelReplace {
	c.cancelReplaceMode = cancelReplaceMode
	return c
}

func (c *cancelReplaceRequest) SetCancelNewClientOrderId(cancelNewClientOrderId string) CancelReplace {
	c.cancelNewClientOrderId = &cancelNewClientOrderId
	return c
}

func (c *cancelReplaceRequest) SetCancelOrigClientOrderId(cancelOrigClientOrderId string) CancelReplace {
	c.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return c
}

func (c *cancelReplaceRequest) SetCancelOrderId(cancelOrderId int64) CancelReplace {
	c.cancelOrderId = &cancelOrderId
	return c
}

func (c *cancelReplaceRequest) SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) CancelReplace {
	c.cancelRestrictions = cancelRestrictions
	return c
}

func (c *cancelReplaceRequest) SetOrderRateLimitExceededMode(orderRateLimitExceededMode enums.OrderRateLimitExceededModeType) CancelReplace {
	c.orderRateLimitExceededMode = orderRateLimitExceededMode
	return c
}

func (c *cancelReplaceRequest) SetSide(side enums.SideType) CancelReplace {
	c.side = side
	return c
}

func (c *cancelReplaceRequest) SetType(_type enums.OrderType) CancelReplace {
	c._type = _type
	//强制要求的参数
	switch c._type {
	case enums.OrderTypeLimit:
		if c.timeInForce == "" {
			panic("timeInForce not set")
		}
		if c.price == nil {
			panic("price not set")
		}
		if c.quantity == nil {
			panic("quantity not set")
		}
	case enums.OrderTypeMarket:
		//市价买卖单可用quoteOrderQty参数来设置quote asset数量. 正确的quantity取决于市场的流动性与quoteOrderQty
		//例如: 市价 BUY BTCUSDT，单子会基于quoteOrderQty- USDT 的数量，购买 BTC.
		//市价 SELL BTCUSDT，单子会卖出 BTC 来满足quoteOrderQty- USDT 的数量.
		if c.quantity == nil {
			panic("quantity not set")
		}
	case enums.OrderTypeStopLoss:
		if c.quantity == nil {
			panic("quantity not set")
		}
		if c.stopPrice == nil {
			panic("stopPrice not set")
		}
		if c.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeStopLossLimit:
		if c.timeInForce == "" {
			panic("timeInForce not set")
		}
		if c.quantity == nil {
			panic("quantity not set")
		}
		if c.price == nil {
			panic("price not set")
		}
		if c.stopPrice == nil {
			panic("stopPrice not set")
		}
		if c.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeTakeProfit:
		if c.quantity == nil {
			panic("quantity not set")
		}
		if c.stopPrice == nil {
			panic("stopPrice not set")
		}
		if c.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeTakeProfitLimit:
		if c.timeInForce == "" {
			panic("timeInForce not set")
		}
		if c.quantity == nil {
			panic("quantity not set")
		}
		if c.price == nil {
			panic("price not set")
		}
		if c.stopPrice == nil {
			panic("stopPrice not set")
		}
		if c.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeLimitMaker:
		if c.quantity == nil {
			panic("quantity not set")
		}
		if c.price == nil {
			panic("price not set")
		}
	}
	return c
}

func (c *cancelReplaceRequest) SetTimeInForce(timeInForce enums.TimeInForceType) CancelReplace {
	c.timeInForce = timeInForce
	return c
}

func (c *cancelReplaceRequest) SetQuantity(quantity string) CancelReplace {
	c.quantity = &quantity
	return c
}

func (c *cancelReplaceRequest) SetQuoteOrderQty(quoteOrderQty string) CancelReplace {
	c.quoteOrderQty = &quoteOrderQty
	return c
}

func (c *cancelReplaceRequest) SetPrice(price string) CancelReplace {
	c.price = &price
	return c
}

func (c *cancelReplaceRequest) SetNewClientOrderId(newClientOrderId string) CancelReplace {
	c.newClientOrderId = &newClientOrderId
	return c
}

func (c *cancelReplaceRequest) SetStrategyId(strategyId int64) CancelReplace {
	c.strategyId = &strategyId
	return c
}

func (c *cancelReplaceRequest) SetStrategyType(strategyType int64) CancelReplace {
	c.strategyType = &strategyType
	return c
}

func (c *cancelReplaceRequest) SetStopPrice(stopPrice string) CancelReplace {
	c.stopPrice = &stopPrice
	return c
}

func (c *cancelReplaceRequest) SetTrailingDelta(trailingDelta int64) CancelReplace {
	c.trailingDelta = &trailingDelta
	return c
}

func (c *cancelReplaceRequest) SetIcebergQty(icebergQty string) CancelReplace {
	c.icebergQty = &icebergQty
	return c
}

func (c *cancelReplaceRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) CancelReplace {
	c.newOrderRespType = newOrderRespType
	return c
}

func (c *cancelReplaceRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) CancelReplace {
	c.selfTradePreventionMode = selfTradePreventionMode
	return c
}

func (c *cancelReplaceRequest) SetRecvWindow(recvWindow int64) CancelReplace {
	c.recvWindow = recvWindow
	return c
}

func (c *cancelReplaceRequest) SetTimestamp(timestamp int64) CancelReplace {
	c.timestamp = timestamp
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
		Path:   consts.ApiTradingOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("side", c.side)
	req.SetParam("type", c._type)
	if c.timeInForce != "" {
		req.SetForm("timeInForce", c.timeInForce)
	}
	req.SetParam("quantity", c.quantity)
	if c.quoteOrderQty != nil {
		req.SetParam("quoteOrderQty", c.quoteOrderQty)
	}
	if c.price != nil {
		req.SetParam("price", c.price)
	}
	if c.newClientOrderId != nil {
		req.SetParam("newClientOrderId", c.newClientOrderId)
	}
	if c.strategyId != nil {
		req.SetParam("strategyId", c.strategyId)
	}
	if c.stopPrice != nil {
		req.SetParam("stopPrice", c.stopPrice)
	}
	if c.trailingDelta != nil {
		req.SetParam("trailingDelta", c.trailingDelta)
	}
	if c.icebergQty != nil {
		req.SetParam("icebergQty", c.icebergQty)
	}
	if c.newOrderRespType != "" {
		req.SetParam("newOrderRespType", c.newOrderRespType)
	}
	if c.selfTradePreventionMode != "" {
		req.SetParam("selfTradePreventionMode", c.selfTradePreventionMode)
	}
	if c.recvWindow > 0 {
		req.SetParam("recvWindow", c.recvWindow)
	}
	req.SetParam("timestamp", c.timestamp)
	if c.cancelNewClientOrderId != nil {
		req.SetParam("cancelNewClientOrderId", c.cancelNewClientOrderId)
	}
	if c.cancelOrigClientOrderId != nil {
		req.SetParam("cancelOrigClientOrderId", c.cancelOrigClientOrderId)
	}
	if c.cancelOrderId != nil {
		req.SetParam("cancelOrderId", c.cancelOrderId)
	}
	if c.cancelRestrictions != "" {
		req.SetParam("cancelRestrictions", c.cancelRestrictions)
	}
	if c.orderRateLimitExceededMode != "" {
		req.SetParam("orderRateLimitExceededMode", c.orderRateLimitExceededMode)
	}
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("cancelReplaceRequest response err:%v", err)
		return nil, err
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		c.Debugf("ParseHttpResponse err:%v", err.Error())
		return nil, err
	}
	return body, nil
}
