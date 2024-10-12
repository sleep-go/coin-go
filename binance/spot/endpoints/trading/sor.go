package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// SOR 下 SOR 订单 (TRADE)
// 发送使用智能订单路由 (SOR) 的新订单。
// 请注意: POST /api/v3/sor/order 只支持 限价 和 市场 单， 并不支持 quoteOrderQty。
type SOR interface {
	SetSymbol(symbol string) *sorRequest
	SetSide(side enums.SideType) *sorRequest
	SetType(_type enums.OrderType) *sorRequest
	SetTimeInForce(timeInForce enums.TimeInForceType) *sorRequest
	SetQuantity(quantity string) *sorRequest
	SetPrice(price string) *sorRequest
	SetNewClientOrderId(newClientOrderId string) *sorRequest
	SetStrategyId(strategyId int64) *sorRequest
	SetStrategyType(strategyType int64) *sorRequest
	SetIcebergQty(icebergQty string) *sorRequest
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *sorRequest
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *sorRequest
	Call(ctx context.Context) (body *sorResponse, err error)
	CallTest(ctx context.Context, computeCommissionRates bool) (body *sorTestResponse, err error)
}
type sorRequest struct {
	*binance.Client
	symbol      string
	side        enums.SideType
	_type       enums.OrderType
	timeInForce enums.TimeInForceType
	quantity    *string
	price       *string
	//用户自定义的orderid，如空缺系统会自动赋值。如果几个订单具有相同的 newClientOrderID 赋值，
	//那么只有在前一个订单成交后才可以接受下一个订单，否则该订单将被拒绝。
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int64
	icebergQty              *string
	newOrderRespType        enums.NewOrderRespType //指定响应类型 ACK, RESULT 或 FULL; 默认为 FULL。
	selfTradePreventionMode enums.StpModeType
}

type sorResponse struct {
	Symbol              string                `json:"symbol"`
	OrderId             int                   `json:"orderId"`
	OrderListId         int                   `json:"orderListId"`
	ClientOrderId       string                `json:"clientOrderId"`
	TransactTime        int64                 `json:"transactTime"`
	Price               string                `json:"price"`
	OrigQty             string                `json:"origQty"`
	ExecutedQty         string                `json:"executedQty"`
	CummulativeQuoteQty string                `json:"cummulativeQuoteQty"`
	Status              enums.OrderStatusType `json:"status"`
	TimeInForce         enums.TimeInForceType `json:"timeInForce"`
	Type                enums.OrderType       `json:"type"`
	Side                enums.SideType        `json:"side"`
	WorkingTime         int64                 `json:"workingTime"`
	Fills               []struct {
		MatchType       string `json:"matchType"`
		Price           string `json:"price"`
		Qty             string `json:"qty"`
		Commission      string `json:"commission"`
		CommissionAsset string `json:"commissionAsset"`
		TradeId         int    `json:"tradeId"`
		AllocId         int    `json:"allocId"`
	} `json:"fills"`
	WorkingFloor            string            `json:"workingFloor"`
	SelfTradePreventionMode enums.StpModeType `json:"selfTradePreventionMode"`
	UsedSor                 bool              `json:"usedSor"`
}

func (s *sorRequest) SetSymbol(symbol string) *sorRequest {
	s.symbol = symbol
	return s
}

func (s *sorRequest) SetSide(side enums.SideType) *sorRequest {
	s.side = side
	return s
}

func (s *sorRequest) SetType(_type enums.OrderType) *sorRequest {
	s._type = _type
	return s
}

func (s *sorRequest) SetTimeInForce(timeInForce enums.TimeInForceType) *sorRequest {
	s.timeInForce = timeInForce
	return s
}

func (s *sorRequest) SetQuantity(quantity string) *sorRequest {
	s.quantity = &quantity
	return s
}

func (s *sorRequest) SetPrice(price string) *sorRequest {
	s.price = &price
	return s
}

func (s *sorRequest) SetNewClientOrderId(newClientOrderId string) *sorRequest {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *sorRequest) SetStrategyId(strategyId int64) *sorRequest {
	s.strategyId = &strategyId
	return s
}

func (s *sorRequest) SetStrategyType(strategyType int64) *sorRequest {
	s.strategyType = &strategyType
	return s
}

func (s *sorRequest) SetIcebergQty(icebergQty string) *sorRequest {
	s.icebergQty = &icebergQty
	return s
}

func (s *sorRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *sorRequest {
	s.newOrderRespType = newOrderRespType
	return s
}

func (s *sorRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *sorRequest {
	s.selfTradePreventionMode = selfTradePreventionMode
	return s
}

func NewSor(client *binance.Client, symbol string) SOR {
	return &sorRequest{Client: client, symbol: symbol}
}
func (s *sorRequest) Call(ctx context.Context) (body *sorResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingSorOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", s.symbol)
	req.SetParam("side", s.side)
	req.SetParam("type", s._type)
	req.SetOptionalParam("timeInForce", s.timeInForce)
	req.SetParam("quantity", s.quantity)
	req.SetOptionalParam("price", s.price)
	req.SetOptionalParam("newClientOrderId", s.newClientOrderId)
	req.SetOptionalParam("strategyId", s.strategyId)
	req.SetOptionalParam("strategyType", s.strategyType)
	req.SetOptionalParam("icebergQty", s.icebergQty)
	req.SetOptionalParam("newOrderRespType", s.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", s.selfTradePreventionMode)
	resp, err := s.Do(ctx, req)
	if err != nil {
		s.Debugf("sorRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*sorResponse](resp)
}

// CallTest 测试 SOR 下单接口 (TRADE)
// 用于测试使用智能订单路由 (SOR) 的订单请求，但不会提交到撮合引擎
func (s *sorRequest) CallTest(ctx context.Context, computeCommissionRates bool) (body *sorTestResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingSorOrderTest,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", s.symbol)
	req.SetParam("side", s.side)
	req.SetParam("type", s._type)
	req.SetOptionalParam("timeInForce", s.timeInForce)
	req.SetParam("quantity", s.quantity)
	req.SetOptionalParam("price", s.price)
	req.SetOptionalParam("newClientOrderId", s.newClientOrderId)
	req.SetOptionalParam("strategyId", s.strategyId)
	req.SetOptionalParam("strategyType", s.strategyType)
	req.SetOptionalParam("icebergQty", s.icebergQty)
	req.SetOptionalParam("newOrderRespType", s.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", s.selfTradePreventionMode)
	req.SetOptionalParam("computeCommissionRates", computeCommissionRates)
	resp, err := s.Do(ctx, req)
	if err != nil {
		s.Debugf("sorRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*sorTestResponse](resp)
}

type sorTestResponse struct {
	StandardCommissionForOrder struct {
		Maker string `json:"maker"`
		Taker string `json:"taker"`
	} `json:"standardCommissionForOrder"`
	TaxCommissionForOrder struct {
		Maker string `json:"maker"`
		Taker string `json:"taker"`
	} `json:"taxCommissionForOrder"`
	Discount struct {
		EnabledForAccount bool   `json:"enabledForAccount"`
		EnabledForSymbol  bool   `json:"enabledForSymbol"`
		DiscountAsset     string `json:"discountAsset"`
		Discount          string `json:"discount"`
	} `json:"discount"`
}

// ****************************** Websocket Api *******************************

type WsApiSOR interface {
	binance.WsApi[*WsApiSORResponse]
	SOR
	SendTest(ctx context.Context, computeCommissionRates bool) (*WsApiSORTestResponse, error)
}
type WsApiSORResponse struct {
	binance.WsApiResponse
	Result *sorResponse `json:"result"`
}
type WsApiSORTestResponse struct {
	binance.WsApiResponse
	Result *sorTestResponse `json:"result"`
}

func NewWsApiSOR(c *binance.Client) WsApiSOR {
	return &sorRequest{Client: c}
}

// Send 下 SOR 订单 (TRADE)
// 下使用智能订单路由 (SOR) 的新订单。
// 注意: sor.order.place 只支持 限价 和 市场 单， 并不支持 quoteOrderQty。
func (s *sorRequest) Send(ctx context.Context) (*WsApiSORResponse, error) {
	req := &binance.Request{Path: "sor.order.place"}
	req.SetNeedSign(true)
	req.SetParam("symbol", s.symbol)
	req.SetParam("side", s.side)
	req.SetParam("type", s._type)
	req.SetOptionalParam("timeInForce", s.timeInForce)
	req.SetParam("quantity", s.quantity)
	req.SetOptionalParam("price", s.price)
	req.SetOptionalParam("newClientOrderId", s.newClientOrderId)
	req.SetOptionalParam("strategyId", s.strategyId)
	req.SetOptionalParam("strategyType", s.strategyType)
	req.SetOptionalParam("icebergQty", s.icebergQty)
	req.SetOptionalParam("newOrderRespType", s.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", s.selfTradePreventionMode)
	return binance.WsApiHandler[*WsApiSORResponse](ctx, s.Client, req)
}

// SendTest 测试 SOR 下单接口 (TRADE)
func (s *sorRequest) SendTest(ctx context.Context, computeCommissionRates bool) (*WsApiSORTestResponse, error) {
	req := &binance.Request{Path: "sor.order.test"}
	req.SetNeedSign(true)
	req.SetParam("symbol", s.symbol)
	req.SetParam("side", s.side)
	req.SetParam("type", s._type)
	req.SetOptionalParam("timeInForce", s.timeInForce)
	req.SetParam("quantity", s.quantity)
	req.SetOptionalParam("price", s.price)
	req.SetOptionalParam("newClientOrderId", s.newClientOrderId)
	req.SetOptionalParam("strategyId", s.strategyId)
	req.SetOptionalParam("strategyType", s.strategyType)
	req.SetOptionalParam("icebergQty", s.icebergQty)
	req.SetOptionalParam("newOrderRespType", s.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", s.selfTradePreventionMode)
	req.SetOptionalParam("computeCommissionRates", computeCommissionRates)
	return binance.WsApiHandler[*WsApiSORTestResponse](ctx, s.Client, req)
}
