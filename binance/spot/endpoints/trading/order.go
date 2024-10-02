package trading

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

type CreateOrder interface {
	SetSide(side enums.SideType) CreateOrder
	SetType(orderType enums.OrderType) CreateOrder
	SetTimeInForce(timeInForce enums.TimeInForceType) CreateOrder
	SetQuantity(quantity string) CreateOrder
	SetQuoteOrderQty(quoteOrderQty string) CreateOrder
	SetPrice(price string) CreateOrder
	SetNewClientOrderId(newClientOrderId string) CreateOrder
	SetStrategyId(strategyId int64) CreateOrder
	SetStrategyType(strategyType int64) CreateOrder
	SetStopPrice(stopPrice string) CreateOrder
	SetTrailingDelta(trailingDelta int64) CreateOrder
	SetIcebergQty(icebergQty string) CreateOrder
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) CreateOrder
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) CreateOrder
	SetRecvWindow(recvWindow int64) CreateOrder
	SetTimestamp(timestamp int64) CreateOrder
	Call(ctx context.Context) (body *createOrderResponse, err error)
	CallTest(ctx context.Context, computeCommissionRates bool) (body *CreateOrderTestResponse, err error)
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
	consts.ErrorResponse
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

func (o *createOrderRequest) SetSide(side enums.SideType) CreateOrder {
	o.side = side
	return o
}

func (o *createOrderRequest) SetType(_type enums.OrderType) CreateOrder {
	o._type = _type
	//强制要求的参数
	switch o._type {
	case enums.OrderTypeLimit:
		if o.timeInForce == "" {
			panic("timeInForce not set")
		}
		if o.price == nil {
			panic("price not set")
		}
		if o.quantity == nil {
			panic("quantity not set")
		}
	case enums.OrderTypeMarket:
		//市价买卖单可用quoteOrderQty参数来设置quote asset数量. 正确的quantity取决于市场的流动性与quoteOrderQty
		//例如: 市价 BUY BTCUSDT，单子会基于quoteOrderQty- USDT 的数量，购买 BTC.
		//市价 SELL BTCUSDT，单子会卖出 BTC 来满足quoteOrderQty- USDT 的数量.
		if o.quantity == nil {
			panic("quantity not set")
		}
	case enums.OrderTypeStopLoss:
		if o.quantity == nil {
			panic("quantity not set")
		}
		if o.stopPrice == nil {
			panic("stopPrice not set")
		}
		if o.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeStopLossLimit:
		if o.timeInForce == "" {
			panic("timeInForce not set")
		}
		if o.quantity == nil {
			panic("quantity not set")
		}
		if o.price == nil {
			panic("price not set")
		}
		if o.stopPrice == nil {
			panic("stopPrice not set")
		}
		if o.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeTakeProfit:
		if o.quantity == nil {
			panic("quantity not set")
		}
		if o.stopPrice == nil {
			panic("stopPrice not set")
		}
		if o.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeTakeProfitLimit:
		if o.timeInForce == "" {
			panic("timeInForce not set")
		}
		if o.quantity == nil {
			panic("quantity not set")
		}
		if o.price == nil {
			panic("price not set")
		}
		if o.stopPrice == nil {
			panic("stopPrice not set")
		}
		if o.trailingDelta == nil {
			panic("trailingDelta not set")
		}
	case enums.OrderTypeLimitMaker:
		if o.quantity == nil {
			panic("quantity not set")
		}
		if o.price == nil {
			panic("price not set")
		}
	}
	return o
}

func (o *createOrderRequest) SetTimeInForce(timeInForce enums.TimeInForceType) CreateOrder {
	o.timeInForce = timeInForce
	return o
}

func (o *createOrderRequest) SetQuantity(quantity string) CreateOrder {
	o.quantity = &quantity
	return o
}

func (o *createOrderRequest) SetQuoteOrderQty(quoteOrderQty string) CreateOrder {
	o.quoteOrderQty = &quoteOrderQty
	return o
}

func (o *createOrderRequest) SetPrice(price string) CreateOrder {
	o.price = &price
	return o
}

func (o *createOrderRequest) SetNewClientOrderId(newClientOrderId string) CreateOrder {
	o.newClientOrderId = &newClientOrderId
	return o
}

func (o *createOrderRequest) SetStrategyId(strategyId int64) CreateOrder {
	o.strategyId = &strategyId
	return o
}

func (o *createOrderRequest) SetStrategyType(strategyType int64) CreateOrder {
	o.strategyType = &strategyType
	return o
}

func (o *createOrderRequest) SetStopPrice(stopPrice string) CreateOrder {
	o.stopPrice = &stopPrice
	return o
}

func (o *createOrderRequest) SetTrailingDelta(trailingDelta int64) CreateOrder {
	o.trailingDelta = &trailingDelta
	return o
}

func (o *createOrderRequest) SetIcebergQty(icebergQty string) CreateOrder {
	o.icebergQty = &icebergQty
	return o
}

func (o *createOrderRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) CreateOrder {
	o.newOrderRespType = newOrderRespType
	return o
}

func (o *createOrderRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) CreateOrder {
	o.selfTradePreventionMode = selfTradePreventionMode
	return o
}

func (o *createOrderRequest) SetRecvWindow(recvWindow int64) CreateOrder {
	o.recvWindow = recvWindow
	return o
}

func (o *createOrderRequest) SetTimestamp(timestamp int64) CreateOrder {
	o.timestamp = timestamp
	return o
}
func (o *createOrderRequest) Call(ctx context.Context) (body *createOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetParam("side", o.side)
	req.SetParam("type", o._type)
	if o.timeInForce != "" {
		req.SetForm("timeInForce", o.timeInForce)
	}
	req.SetParam("quantity", o.quantity)
	if o.quoteOrderQty != nil {
		req.SetParam("quoteOrderQty", o.quoteOrderQty)
	}
	if o.price != nil {
		req.SetParam("price", o.price)
	}
	if o.newClientOrderId != nil {
		req.SetParam("newClientOrderId", o.newClientOrderId)
	}
	if o.strategyId != nil {
		req.SetParam("strategyId", o.strategyId)
	}
	if o.stopPrice != nil {
		req.SetParam("stopPrice", o.stopPrice)
	}
	if o.trailingDelta != nil {
		req.SetParam("trailingDelta", o.trailingDelta)
	}
	if o.icebergQty != nil {
		req.SetParam("icebergQty", o.icebergQty)
	}
	if o.newOrderRespType != "" {
		req.SetParam("newOrderRespType", o.newOrderRespType)
	}
	if o.selfTradePreventionMode != "" {
		req.SetParam("selfTradePreventionMode", o.selfTradePreventionMode)
	}
	if o.recvWindow > 0 {
		req.SetParam("recvWindow", o.recvWindow)
	}
	if o.timestamp > 0 {
		req.SetParam("timestamp", o.timestamp)
	}
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("createOrderRequest response err:%v", err)
		return nil, err
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		o.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}

type CreateOrderTestResponse struct {
	consts.ErrorResponse
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

func (o *createOrderRequest) CallTest(ctx context.Context, computeCommissionRates bool) (body *CreateOrderTestResponse, err error) {
	// 没有 computeCommissionRates返回空
	if computeCommissionRates == false {
		return nil, nil
	}
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingOrderTest,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetParam("side", o.side)
	req.SetParam("type", o._type)
	req.SetParam("computeCommissionRates", computeCommissionRates)
	if o.timeInForce != "" {
		req.SetForm("timeInForce", o.timeInForce)
	}
	req.SetParam("quantity", o.quantity)
	if o.quoteOrderQty != nil {
		req.SetParam("quoteOrderQty", o.quoteOrderQty)
	}
	if o.price != nil {
		req.SetParam("price", o.price)
	}
	if o.newClientOrderId != nil {
		req.SetParam("newClientOrderId", o.newClientOrderId)
	}
	if o.strategyId != nil {
		req.SetParam("strategyId", o.strategyId)
	}
	if o.stopPrice != nil {
		req.SetParam("stopPrice", o.stopPrice)
	}
	if o.trailingDelta != nil {
		req.SetParam("trailingDelta", o.trailingDelta)
	}
	if o.icebergQty != nil {
		req.SetParam("icebergQty", o.icebergQty)
	}
	if o.newOrderRespType != "" {
		req.SetParam("newOrderRespType", o.newOrderRespType)
	}
	if o.selfTradePreventionMode != "" {
		req.SetParam("selfTradePreventionMode", o.selfTradePreventionMode)
	}
	if o.recvWindow > 0 {
		req.SetParam("recvWindow", o.recvWindow)
	}
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("createOrderRequest response err:%v", err)
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
