package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// OCO 发送新 one-cancels-the-other (OCO) 订单，激活其中一个订单会立即取消另一个订单。
//
// OCO 包含了两个订单，分别被称为 上方订单 和 下方订单。
// 其中一个订单必须是 LIMIT_MAKER 订单，另一个订单必须是 STOP_LOSS 或 STOP_LOSS_LIMIT 订单。
// 针对价格限制：
// 如果 OCO 订单方向是 SELL：LIMIT_MAKER price > 最后交易价格 > stopPrice
// 如果 OCO 订单方向是 BUY：LIMIT_MAKER price < 最后交易价格 < stopPrice
// OCO 将2 个订单添加到未成交订单计数，EXCHANGE_MAX_ORDERS 过滤器和 MAX_NUM_ORDERS 过滤器中。
type OCO interface {
	SetSymbol(symbol string) *ocoRequest
	SetListClientOrderId(listClientOrderId string) *ocoRequest
	SetSide(side enums.SideType) *ocoRequest
	SetQuantity(quantity string) *ocoRequest
	SetAboveType(aboveType enums.OrderType) *ocoRequest
	SetAboveClientOrderId(aboveClientOrderId string) *ocoRequest
	SetAboveIcebergQty(aboveIcebergQty int64) *ocoRequest
	SetAbovePrice(abovePrice string) *ocoRequest
	SetAboveStopPrice(aboveStopPrice string) *ocoRequest
	SetAboveTrailingDelta(aboveTrailingDelta int64) *ocoRequest
	SetAboveTimeInForce(aboveTimeInForce enums.TimeInForceType) *ocoRequest
	SetAboveStrategyId(aboveStrategyId int64) *ocoRequest
	SetAboveStrategyType(aboveStrategyType int64) *ocoRequest
	SetBelowType(belowType enums.OrderType) *ocoRequest
	SetBelowClientOrderId(belowClientOrderId string) *ocoRequest
	SetBelowIcebergQty(belowIcebergQty int64) *ocoRequest
	SetBelowPrice(belowPrice string) *ocoRequest
	SetBelowStopPrice(belowStopPrice string) *ocoRequest
	SetBelowTrailingDelta(belowTrailingDelta int64) *ocoRequest
	SetBelowTimeInForce(belowTimeInForce enums.TimeInForceType) *ocoRequest
	SetBelowStrategyId(belowStrategyId int64) *ocoRequest
	SetBelowStrategyType(belowStrategyType int64) *ocoRequest
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *ocoRequest
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *ocoRequest
	Call(ctx context.Context) (body *ocoResponse, err error)
}

type ocoRequest struct {
	*binance.Client
	symbol string
	//整个 OCO order list 的唯一ID。 如果未发送则自动生成。
	//仅当前一个订单已填满或完全过期时，才会接受具有相同的listClientOrderId。
	//listClientOrderId 与 aboveClientOrderId 和 belowCLientOrderId 不同。
	listClientOrderId  string
	side               enums.SideType
	quantity           *string         //两个订单的数量。
	aboveType          enums.OrderType //支持值：STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER。
	aboveClientOrderId *string         //上方订单的唯一ID。 如果未发送则自动生成。
	aboveIcebergQty    *int64          //请注意，只有当 aboveTimeInForce 为 GTC 时才能使用。
	abovePrice         *string
	//如果 aboveType 是 STOP_LOSS 或 STOP_LOSS_LIMIT 才能使用。
	//必须指定 aboveStopPrice 或 aboveTrailingDelta 或两者。
	aboveStopPrice     *string                //如果 aboveType 是 STOP_LOSS 或 STOP_LOSS_LIMIT 才能使用。必须指定 aboveStopPrice 或 aboveTrailingDelta 或两者。
	aboveTrailingDelta *int64                 //请看 追踪止盈止损(Trailing Stop)订单常见问题。
	aboveTimeInForce   *enums.TimeInForceType //如果 aboveType 是 STOP_LOSS_LIMIT，则为必填项。
	aboveStrategyId    *int64                 //订单策略中上方订单的 ID。
	aboveStrategyType  *int64                 //上方订单策略的任意数值。小于 1000000 的值被保留，无法使用。
	belowType          enums.OrderType        //支持值：STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER。
	belowClientOrderId *string
	belowIcebergQty    *int64 //请注意，只有当 belowTimeInForce 为 GTC 时才能使用。
	belowPrice         *string
	//如果 belowType 是 STOP_LOSS 或 STOP_LOSS_LIMIT 才能使用。
	//必须指定 belowStopPrice 或 belowTrailingDelta 或两者。
	belowStopPrice          *string
	belowTrailingDelta      *int64                 //请看 追踪止盈止损(Trailing Stop)订单常见问题。
	belowTimeInForce        enums.TimeInForceType  //如果belowType 是 STOP_LOSS_LIMIT，则为必须配合提交的值。
	belowStrategyId         *int64                 //订单策略中下方订单的 ID。
	belowStrategyType       *int64                 //下方订单策略的任意数值。 小于 1000000 的值被保留，无法使用。
	newOrderRespType        enums.NewOrderRespType //响应格式可选值: ACK, RESULT, FULL。
	selfTradePreventionMode enums.StpModeType      //允许的 ENUM 取决于交易对上的配置。 支持值：STP 模式。
}

type ocoResponse struct {
	OrderListId       int    `json:"orderListId"`
	ContingencyType   string `json:"contingencyType"`
	ListStatusType    string `json:"listStatusType"`
	ListOrderStatus   string `json:"listOrderStatus"`
	ListClientOrderId string `json:"listClientOrderId"`
	TransactionTime   int64  `json:"transactionTime"`
	Symbol            string `json:"symbol"`
	Orders            []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
	OrderReports []struct {
		Symbol                  string                `json:"symbol"`
		OrderId                 int                   `json:"orderId"`
		OrderListId             int                   `json:"orderListId"`
		ClientOrderId           string                `json:"clientOrderId"`
		TransactTime            int64                 `json:"transactTime"`
		Price                   string                `json:"price"`
		OrigQty                 string                `json:"origQty"`
		ExecutedQty             string                `json:"executedQty"`
		CummulativeQuoteQty     string                `json:"cummulativeQuoteQty"`
		Status                  enums.OrderStatusType `json:"status"`
		TimeInForce             string                `json:"timeInForce"`
		Type                    enums.OrderType       `json:"type"`
		Side                    enums.SideType        `json:"side"`
		StopPrice               string                `json:"stopPrice,omitempty"`
		WorkingTime             int64                 `json:"workingTime"`
		IcebergQty              string                `json:"icebergQty,omitempty"`
		SelfTradePreventionMode enums.StpModeType     `json:"selfTradePreventionMode"`
	} `json:"orderReports"`
}

func (o *ocoRequest) SetSymbol(symbol string) *ocoRequest {
	o.symbol = symbol
	return o
}
func (o *ocoRequest) SetListClientOrderId(listClientOrderId string) *ocoRequest {
	o.listClientOrderId = listClientOrderId
	return o
}
func (o *ocoRequest) SetSide(side enums.SideType) *ocoRequest {
	o.side = side
	return o
}
func (o *ocoRequest) SetQuantity(quantity string) *ocoRequest {
	o.quantity = &quantity
	return o
}
func (o *ocoRequest) SetAboveType(aboveType enums.OrderType) *ocoRequest {
	o.aboveType = aboveType
	return o
}
func (o *ocoRequest) SetAboveClientOrderId(aboveClientOrderId string) *ocoRequest {
	o.aboveClientOrderId = &aboveClientOrderId
	return o
}
func (o *ocoRequest) SetAboveIcebergQty(aboveIcebergQty int64) *ocoRequest {
	o.aboveIcebergQty = &aboveIcebergQty
	return o
}
func (o *ocoRequest) SetAbovePrice(abovePrice string) *ocoRequest {
	o.abovePrice = &abovePrice
	return o
}
func (o *ocoRequest) SetAboveStopPrice(aboveStopPrice string) *ocoRequest {
	o.aboveStopPrice = &aboveStopPrice
	return o
}
func (o *ocoRequest) SetAboveTrailingDelta(aboveTrailingDelta int64) *ocoRequest {
	o.aboveTrailingDelta = &aboveTrailingDelta
	return o
}
func (o *ocoRequest) SetAboveTimeInForce(aboveTimeInForce enums.TimeInForceType) *ocoRequest {
	o.aboveTimeInForce = &aboveTimeInForce
	return o
}
func (o *ocoRequest) SetAboveStrategyId(aboveStrategyId int64) *ocoRequest {
	o.aboveStrategyId = &aboveStrategyId
	return o
}
func (o *ocoRequest) SetAboveStrategyType(aboveStrategyType int64) *ocoRequest {
	o.aboveStrategyType = &aboveStrategyType
	return o
}

func (o *ocoRequest) SetBelowType(belowType enums.OrderType) *ocoRequest {
	o.belowType = belowType
	return o
}
func (o *ocoRequest) SetBelowClientOrderId(belowClientOrderId string) *ocoRequest {
	o.belowClientOrderId = &belowClientOrderId
	return o
}
func (o *ocoRequest) SetBelowIcebergQty(belowIcebergQty int64) *ocoRequest {
	o.belowIcebergQty = &belowIcebergQty
	return o
}
func (o *ocoRequest) SetBelowPrice(belowPrice string) *ocoRequest {
	o.belowPrice = &belowPrice
	return o
}
func (o *ocoRequest) SetBelowStopPrice(belowStopPrice string) *ocoRequest {
	o.belowStopPrice = &belowStopPrice
	return o
}
func (o *ocoRequest) SetBelowTrailingDelta(belowTrailingDelta int64) *ocoRequest {
	o.belowTrailingDelta = &belowTrailingDelta
	return o
}
func (o *ocoRequest) SetBelowTimeInForce(belowTimeInForce enums.TimeInForceType) *ocoRequest {
	o.belowTimeInForce = belowTimeInForce
	return o
}
func (o *ocoRequest) SetBelowStrategyId(belowStrategyId int64) *ocoRequest {
	o.belowStrategyId = &belowStrategyId
	return o
}
func (o *ocoRequest) SetBelowStrategyType(belowStrategyType int64) *ocoRequest {
	o.belowStrategyType = &belowStrategyType
	return o
}
func (o *ocoRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) *ocoRequest {
	o.newOrderRespType = newOrderRespType
	return o
}
func (o *ocoRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) *ocoRequest {
	o.selfTradePreventionMode = selfTradePreventionMode
	return o
}
func NewOco(client *binance.Client, symbol string) OCO {
	return &ocoRequest{Client: client, symbol: symbol}
}
func (o *ocoRequest) Call(ctx context.Context) (body *ocoResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingOrderListOCO,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("listClientOrderId", o.listClientOrderId)
	req.SetParam("side", o.side)
	req.SetParam("quantity", o.quantity)
	req.SetParam("aboveType", o.aboveType)
	req.SetOptionalParam("aboveClientOrderId", o.aboveClientOrderId)
	req.SetOptionalParam("aboveIcebergQty", o.aboveIcebergQty)
	req.SetOptionalParam("abovePrice", o.abovePrice)
	req.SetOptionalParam("aboveStopPrice", o.aboveStopPrice)
	req.SetOptionalParam("aboveTrailingDelta", o.aboveTrailingDelta)
	req.SetOptionalParam("aboveTimeInForce", o.aboveTimeInForce)
	req.SetOptionalParam("aboveStrategyId", o.aboveStrategyId)
	req.SetOptionalParam("aboveStrategyType", o.aboveStrategyType)
	req.SetParam("belowType", o.belowType)
	req.SetOptionalParam("belowClientOrderId", o.belowClientOrderId)
	req.SetOptionalParam("belowIcebergQty", o.belowIcebergQty)
	req.SetOptionalParam("belowPrice", o.belowPrice)
	req.SetOptionalParam("belowStopPrice", o.belowStopPrice)
	req.SetOptionalParam("belowTrailingDelta", o.belowTrailingDelta)
	req.SetOptionalParam("belowTimeInForce", o.belowTimeInForce)
	req.SetOptionalParam("belowStrategyId", o.belowStrategyId)
	req.SetOptionalParam("belowStrategyType", o.belowStrategyType)
	req.SetOptionalParam("newOrderRespType", o.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", o.selfTradePreventionMode)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("ocoRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*ocoResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiOCO interface {
	binance.WsApi[*WsApiOCOResponse]
	OCO
}
type WsApiOCOResponse struct {
	binance.WsApiResponse
	Result *createOrderResponse `json:"result"`
}

func NewWsApiOCO(c *binance.Client) WsApiOCO {
	return &ocoRequest{Client: c}
}

// Send 下新的订单 (TRADE)
func (o *ocoRequest) Send(ctx context.Context) (*WsApiOCOResponse, error) {
	req := &binance.Request{Path: "orderList.place.oco"}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("listClientOrderId", o.listClientOrderId)
	req.SetParam("side", o.side)
	req.SetParam("quantity", o.quantity)
	req.SetParam("aboveType", o.aboveType)
	req.SetOptionalParam("aboveClientOrderId", o.aboveClientOrderId)
	req.SetOptionalParam("aboveIcebergQty", o.aboveIcebergQty)
	req.SetOptionalParam("abovePrice", o.abovePrice)
	req.SetOptionalParam("aboveStopPrice", o.aboveStopPrice)
	req.SetOptionalParam("aboveTrailingDelta", o.aboveTrailingDelta)
	req.SetOptionalParam("aboveTimeInForce", o.aboveTimeInForce)
	req.SetOptionalParam("aboveStrategyId", o.aboveStrategyId)
	req.SetOptionalParam("aboveStrategyType", o.aboveStrategyType)
	req.SetParam("belowType", o.belowType)
	req.SetOptionalParam("belowClientOrderId", o.belowClientOrderId)
	req.SetOptionalParam("belowIcebergQty", o.belowIcebergQty)
	req.SetOptionalParam("belowPrice", o.belowPrice)
	req.SetOptionalParam("belowStopPrice", o.belowStopPrice)
	req.SetOptionalParam("belowTrailingDelta", o.belowTrailingDelta)
	req.SetOptionalParam("belowTimeInForce", o.belowTimeInForce)
	req.SetOptionalParam("belowStrategyId", o.belowStrategyId)
	req.SetOptionalParam("belowStrategyType", o.belowStrategyType)
	req.SetOptionalParam("newOrderRespType", o.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", o.selfTradePreventionMode)
	return binance.WsApiHandler[*WsApiOCOResponse](ctx, o.Client, req)
}
