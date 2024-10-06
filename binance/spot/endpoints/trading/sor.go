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
	SetSide(side enums.SideType) SOR
	SetType(_type enums.OrderType) SOR
	SetTimeInForce(timeInForce enums.TimeInForceType) SOR
	SetQuantity(quantity string) SOR
	SetPrice(price string) SOR
	SetNewClientOrderId(newClientOrderId string) SOR
	SetStrategyId(strategyId int64) SOR
	SetStrategyType(strategyType int64) SOR
	SetIcebergQty(icebergQty string) SOR
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) SOR
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) SOR
	SetComputeCommissionRates(computeCommissionRates bool) SOR
	SetRecvWindow(recvWindow int64) SOR
	SetTimestamp(timestamp int64) SOR
	Call(ctx context.Context) (body *sorResponse, err error)
	CallTest(ctx context.Context) (body *sorTestResponse, err error)
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
	computeCommissionRates  *bool
	recvWindow              int64
	timestamp               int64
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

func (s *sorRequest) SetComputeCommissionRates(computeCommissionRates bool) SOR {
	s.computeCommissionRates = &computeCommissionRates
	return s
}

func (s *sorRequest) SetSide(side enums.SideType) SOR {
	s.side = side
	return s
}

func (s *sorRequest) SetType(_type enums.OrderType) SOR {
	s._type = _type
	return s
}

func (s *sorRequest) SetTimeInForce(timeInForce enums.TimeInForceType) SOR {
	s.timeInForce = timeInForce
	return s
}

func (s *sorRequest) SetQuantity(quantity string) SOR {
	s.quantity = &quantity
	return s
}

func (s *sorRequest) SetPrice(price string) SOR {
	s.price = &price
	return s
}

func (s *sorRequest) SetNewClientOrderId(newClientOrderId string) SOR {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *sorRequest) SetStrategyId(strategyId int64) SOR {
	s.strategyId = &strategyId
	return s
}

func (s *sorRequest) SetStrategyType(strategyType int64) SOR {
	s.strategyType = &strategyType
	return s
}

func (s *sorRequest) SetIcebergQty(icebergQty string) SOR {
	s.icebergQty = &icebergQty
	return s
}

func (s *sorRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) SOR {
	s.newOrderRespType = newOrderRespType
	return s
}

func (s *sorRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) SOR {
	s.selfTradePreventionMode = selfTradePreventionMode
	return s
}

func (s *sorRequest) SetRecvWindow(recvWindow int64) SOR {
	s.recvWindow = recvWindow
	return s
}

func (s *sorRequest) SetTimestamp(timestamp int64) SOR {
	s.timestamp = timestamp
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
	req.SetOptionalParam("recvWindow", s.recvWindow)
	req.SetParam("timestamp", s.timestamp)
	resp, err := s.Do(ctx, req)
	if err != nil {
		s.Debugf("sorRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*sorResponse](resp)
}

// CallTest 测试 SOR 下单接口 (TRADE)
// 用于测试使用智能订单路由 (SOR) 的订单请求，但不会提交到撮合引擎
func (s *sorRequest) CallTest(ctx context.Context) (body *sorTestResponse, err error) {
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
	req.SetOptionalParam("computeCommissionRates", s.computeCommissionRates)
	req.SetOptionalParam("selfTradePreventionMode", s.selfTradePreventionMode)
	req.SetOptionalParam("recvWindow", s.recvWindow)
	req.SetParam("timestamp", s.timestamp)
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
