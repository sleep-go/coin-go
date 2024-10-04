package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// OTOCO 发送一个新的 OTOCO 订单。
//
// 一个 OTOCO 订单（One-Triggers-One-Cancels-the-Other）是一个包含了三个订单的订单列表。
// 第一个订单被称为生效订单，必须为 LIMIT 或 LIMIT_MAKER 类型的订单。最初，订单簿上只有生效订单。
// 生效订单的行为与此一致 OTO
// 一个OTOCO订单有两个待处理订单（pending above 和 pending below），它们构成了一个 OCO 订单列表。只有当生效订单完全成交时，待处理订单们才会被自动下单。
// 待处理上方(pending above)订单和待处理下方(pending below)订单都遵循与 OCO 订单列表相同的规则 Order List OCO。
// OTOCO 在未成交订单计数，EXCHANGE_MAX_NUM_ORDERS 过滤器和 MAX_NUM_ORDERS 过滤器的基础上添加3个订单。
type OTOCO interface {
	SetListClientOrderId(listClientOrderId string) OTOCO
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) OTOCO
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) OTOCO
	SetWorkingType(workingType enums.OrderType) OTOCO
	SetWorkingSide(workingSide enums.SideType) OTOCO
	SetWorkingClientOrderId(workingClientOrderId string) OTOCO
	SetWorkingPrice(workingPrice string) OTOCO
	SetWorkingQuantity(workingQuantity string) OTOCO
	SetWorkingIcebergQty(workingIcebergQty string) OTOCO
	SetWorkingTimeInForce(workingTimeInForce enums.TimeInForceType) OTOCO
	SetWorkingStrategyId(workingStrategyId int64) OTOCO
	SetWorkingStrategyType(workingStrategyType int64) OTOCO
	SetPendingSide(pendingSide enums.SideType) OTOCO
	SetPendingQuantity(pendingQuantity string) OTOCO
	SetPendingAboveType(pendingAboveType enums.OrderType) OTOCO
	SetPendingAboveClientOrderId(pendingAboveClientOrderId string) OTOCO
	SetPendingAbovePrice(pendingAbovePrice string) OTOCO
	SetPendingAboveStopPrice(pendingAboveStopPrice string) OTOCO
	SetPendingAboveTrailingDelta(pendingAboveTrailingDelta string) OTOCO
	SetPendingAboveIcebergQty(pendingAboveIcebergQty string) OTOCO
	SetPendingAboveTimeInForce(pendingAboveTimeInForce enums.TimeInForceType) OTOCO
	SetPendingAboveStrategyId(pendingAboveStrategyId int64) OTOCO
	SetPendingAboveStrategyType(pendingAboveStrategyType int64) OTOCO
	SetPendingBelowType(pendingBelowType enums.OrderType) OTOCO
	SetPendingBelowClientOrderId(pendingBelowClientOrderId string) OTOCO
	SetPendingBelowPrice(pendingBelowPrice string) OTOCO
	SetPendingBelowStopPrice(pendingBelowStopPrice string) OTOCO
	SetPendingBelowTrailingDelta(pendingBelowTrailingDelta string) OTOCO
	SetPendingBelowIcebergQty(pendingBelowIcebergQty string) OTOCO
	SetPendingBelowTimeInForce(pendingBelowTimeInForce enums.TimeInForceType) OTOCO
	SetPendingBelowStrategyId(pendingBelowStrategyId int64) OTOCO
	SetPendingBelowStrategyType(pendingBelowStrategyType int64) OTOCO
	SetRecvWindow(recvWindow int64) OTOCO
	SetTimestamp(timestamp int64) OTOCO
	Call(ctx context.Context) (body *otocoResponse, err error)
}

type otocoRequest struct {
	*binance.Client
	symbol string
	//整个订单列表的唯一ID。 如果未发送则自动生成。
	//仅当前一个订单列表已填满或完全过期时，才会接受含有相同 listClientOrderId 值的新订单列表。
	//listClientOrderId 与 workingClientOrderId 和 pendingClientOrderId 不同。
	listClientOrderId         *string
	newOrderRespType          enums.NewOrderRespType
	selfTradePreventionMode   enums.StpModeType
	workingType               enums.OrderType //支持的数值： LIMIT， LIMIT_MAKER
	workingSide               enums.SideType
	workingClientOrderId      *string //用于标识生效订单的唯一ID。 如果未发送则自动生成。
	workingPrice              *string
	workingQuantity           *string //用于设置生效订单的数量。
	workingIcebergQty         *string //只有当 workingTimeInForce 为 GTC 时才能使用。
	workingTimeInForce        enums.TimeInForceType
	workingStrategyId         *int64 //订单策略中用于标识生效订单的 ID。
	workingStrategyType       *int64 //用于标识生效订单策略的任意数值。 小于 1000000 的值被保留，无法使用。
	pendingSide               enums.SideType
	pendingQuantity           *string         //用于设置待处理订单的数量。
	pendingAboveType          enums.OrderType //支持的数值： LIMIT_MAKER，STOP_LOSS 和 STOP_LOSS_LIMIT
	pendingAboveClientOrderId *string         //用于标识待处理上方订单的唯一ID。如果未发送则自动生成。
	pendingAbovePrice         *string
	pendingAboveStopPrice     *string
	pendingAboveTrailingDelta *string
	pendingAboveIcebergQty    *string //只有当 pendingTimeInForce 为 GTC 或者当 pendingType 为 LIMIT_MAKER 时才能使用。
	pendingAboveTimeInForce   enums.TimeInForceType
	pendingAboveStrategyId    *int64          //订单策略中用于标识待处理订单的 ID。
	pendingAboveStrategyType  *int64          //用于标识待处理订单策略的任意数值。小于 1000000 的值被保留，无法使用。
	pendingBelowType          enums.OrderType //支持的数值： LIMIT_MAKER，STOP_LOSS 和 STOP_LOSS_LIMIT
	pendingBelowClientOrderId *string         //用于标识待处理下方订单的唯一ID。如果未发送则自动生成。
	pendingBelowPrice         *string
	pendingBelowStopPrice     *string
	pendingBelowTrailingDelta *string
	pendingBelowIcebergQty    *string
	pendingBelowTimeInForce   enums.TimeInForceType
	pendingBelowStrategyId    *int64
	pendingBelowStrategyType  *int64
	recvWindow                int64
	timestamp                 int64
}
type otocoResponse struct {
	OrderListId       int                       `json:"orderListId"`
	ContingencyType   enums.ContingencyType     `json:"contingencyType"`
	ListStatusType    enums.ListStatusType      `json:"listStatusType"`
	ListOrderStatus   enums.ListOrderStatusType `json:"listOrderStatus"`
	ListClientOrderId string                    `json:"listClientOrderId"`
	TransactionTime   int64                     `json:"transactionTime"`
	Symbol            string                    `json:"symbol"`
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
		TimeInForce             enums.TimeInForceType `json:"timeInForce"`
		Type                    enums.OrderType       `json:"type"`
		Side                    enums.SideType        `json:"side"`
		WorkingTime             int64                 `json:"workingTime"`
		SelfTradePreventionMode enums.StpModeType     `json:"selfTradePreventionMode"`
		StopPrice               string                `json:"stopPrice,omitempty"`
	} `json:"orderReports"`
}

func NewOtoco(client *binance.Client, symbol string) OTOCO {
	return &otocoRequest{Client: client, symbol: symbol}
}

func (o *otocoRequest) SetListClientOrderId(listClientOrderId string) OTOCO {
	o.listClientOrderId = &listClientOrderId
	return o
}

func (o *otocoRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) OTOCO {
	o.newOrderRespType = newOrderRespType
	return o
}

func (o *otocoRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) OTOCO {
	o.selfTradePreventionMode = selfTradePreventionMode
	return o
}

func (o *otocoRequest) SetWorkingType(workingType enums.OrderType) OTOCO {
	o.workingType = workingType
	return o
}

func (o *otocoRequest) SetWorkingSide(workingSide enums.SideType) OTOCO {
	o.workingSide = workingSide
	return o
}

func (o *otocoRequest) SetWorkingClientOrderId(workingClientOrderId string) OTOCO {
	o.workingClientOrderId = &workingClientOrderId
	return o
}

func (o *otocoRequest) SetWorkingPrice(workingPrice string) OTOCO {
	o.workingPrice = &workingPrice
	return o
}

func (o *otocoRequest) SetWorkingQuantity(workingQuantity string) OTOCO {
	o.workingQuantity = &workingQuantity
	return o
}

func (o *otocoRequest) SetWorkingIcebergQty(workingIcebergQty string) OTOCO {
	o.workingIcebergQty = &workingIcebergQty
	return o
}

func (o *otocoRequest) SetWorkingTimeInForce(workingTimeInForce enums.TimeInForceType) OTOCO {
	o.workingTimeInForce = workingTimeInForce
	return o
}

func (o *otocoRequest) SetWorkingStrategyId(workingStrategyId int64) OTOCO {
	o.workingStrategyId = &workingStrategyId
	return o
}

func (o *otocoRequest) SetWorkingStrategyType(workingStrategyType int64) OTOCO {
	o.workingStrategyType = &workingStrategyType
	return o
}

func (o *otocoRequest) SetPendingSide(pendingSide enums.SideType) OTOCO {
	o.pendingSide = pendingSide
	return o
}

func (o *otocoRequest) SetPendingQuantity(pendingQuantity string) OTOCO {
	o.pendingQuantity = &pendingQuantity
	return o
}

func (o *otocoRequest) SetPendingAboveType(pendingAboveType enums.OrderType) OTOCO {
	o.pendingAboveType = pendingAboveType
	return o
}

func (o *otocoRequest) SetPendingAboveClientOrderId(pendingAboveClientOrderId string) OTOCO {
	o.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return o
}

func (o *otocoRequest) SetPendingAbovePrice(pendingAbovePrice string) OTOCO {
	o.pendingAbovePrice = &pendingAbovePrice
	return o
}

func (o *otocoRequest) SetPendingAboveStopPrice(pendingAboveStopPrice string) OTOCO {
	o.pendingAboveStopPrice = &pendingAboveStopPrice
	return o
}

func (o *otocoRequest) SetPendingAboveTrailingDelta(pendingAboveTrailingDelta string) OTOCO {
	o.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return o
}

func (o *otocoRequest) SetPendingAboveIcebergQty(pendingAboveIcebergQty string) OTOCO {
	o.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return o
}

func (o *otocoRequest) SetPendingAboveTimeInForce(pendingAboveTimeInForce enums.TimeInForceType) OTOCO {
	o.pendingAboveTimeInForce = pendingAboveTimeInForce
	return o
}

func (o *otocoRequest) SetPendingAboveStrategyId(pendingAboveStrategyId int64) OTOCO {
	o.pendingAboveStrategyId = &pendingAboveStrategyId
	return o
}

func (o *otocoRequest) SetPendingAboveStrategyType(pendingAboveStrategyType int64) OTOCO {
	o.pendingAboveStrategyType = &pendingAboveStrategyType
	return o
}

func (o *otocoRequest) SetPendingBelowType(pendingBelowType enums.OrderType) OTOCO {
	o.pendingBelowType = pendingBelowType
	return o
}

func (o *otocoRequest) SetPendingBelowClientOrderId(pendingBelowClientOrderId string) OTOCO {
	o.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return o
}

func (o *otocoRequest) SetPendingBelowPrice(pendingBelowPrice string) OTOCO {
	o.pendingBelowPrice = &pendingBelowPrice
	return o
}

func (o *otocoRequest) SetPendingBelowStopPrice(pendingBelowStopPrice string) OTOCO {
	o.pendingBelowStopPrice = &pendingBelowStopPrice
	return o
}

func (o *otocoRequest) SetPendingBelowTrailingDelta(pendingBelowTrailingDelta string) OTOCO {
	o.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return o
}

func (o *otocoRequest) SetPendingBelowIcebergQty(pendingBelowIcebergQty string) OTOCO {
	o.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return o
}

func (o *otocoRequest) SetPendingBelowTimeInForce(pendingBelowTimeInForce enums.TimeInForceType) OTOCO {
	o.pendingBelowTimeInForce = pendingBelowTimeInForce
	return o
}

func (o *otocoRequest) SetPendingBelowStrategyId(pendingBelowStrategyId int64) OTOCO {
	o.pendingBelowStrategyId = &pendingBelowStrategyId
	return o
}

func (o *otocoRequest) SetPendingBelowStrategyType(pendingBelowStrategyType int64) OTOCO {
	o.pendingBelowStrategyType = &pendingBelowStrategyType
	return o
}

func (o *otocoRequest) SetRecvWindow(recvWindow int64) OTOCO {
	o.recvWindow = recvWindow
	return o
}

func (o *otocoRequest) SetTimestamp(timestamp int64) OTOCO {
	o.timestamp = timestamp
	return o
}

func (o *otocoRequest) Call(ctx context.Context) (body *otocoResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingOrderListOTOCO,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("listClientOrderId", o.listClientOrderId)
	req.SetOptionalParam("newOrderRespType", o.newOrderRespType)
	req.SetOptionalParam("selfTradePreventionMode", o.selfTradePreventionMode)
	req.SetParam("workingType", o.workingType)
	req.SetParam("workingSide", o.workingSide)
	req.SetOptionalParam("workingClientOrderId", o.workingClientOrderId)
	req.SetParam("workingPrice", o.workingPrice)
	req.SetParam("workingQuantity", o.workingQuantity)
	req.SetOptionalParam("workingIcebergQty", o.workingIcebergQty)
	req.SetOptionalParam("workingTimeInForce", o.workingTimeInForce)
	req.SetOptionalParam("workingStrategyId", o.workingStrategyId)
	req.SetOptionalParam("workingStrategyType", o.workingStrategyType)
	req.SetOptionalParam("pendingSide", o.pendingSide)
	req.SetOptionalParam("pendingQuantity", o.pendingQuantity)
	req.SetOptionalParam("pendingAboveType", o.pendingAboveType)
	req.SetOptionalParam("pendingAboveClientOrderId", o.pendingAboveClientOrderId)
	req.SetOptionalParam("pendingAbovePrice", o.pendingAbovePrice)
	req.SetOptionalParam("pendingAboveStopPrice", o.pendingAboveStopPrice)
	req.SetOptionalParam("pendingAboveTrailingDelta", o.pendingAboveTrailingDelta)
	req.SetOptionalParam("pendingAboveIcebergQty", o.pendingAboveIcebergQty)
	req.SetOptionalParam("pendingAboveTimeInForce", o.pendingAboveTimeInForce)
	req.SetOptionalParam("pendingAboveStrategyId", o.pendingAboveStrategyId)
	req.SetOptionalParam("pendingAboveStrategyType", o.pendingAboveStrategyType)
	req.SetOptionalParam("pendingBelowType", o.pendingBelowType)
	req.SetOptionalParam("pendingBelowClientOrderId", o.pendingBelowClientOrderId)
	req.SetOptionalParam("pendingBelowPrice", o.pendingBelowPrice)
	req.SetOptionalParam("pendingBelowStopPrice", o.pendingBelowStopPrice)
	req.SetOptionalParam("pendingBelowTrailingDelta", o.pendingBelowTrailingDelta)
	req.SetOptionalParam("pendingBelowIcebergQty", o.pendingBelowIcebergQty)
	req.SetOptionalParam("pendingBelowTimeInForce", o.pendingBelowTimeInForce)
	req.SetOptionalParam("pendingBelowStrategyId", o.pendingBelowStrategyId)
	req.SetOptionalParam("pendingBelowStrategyType", o.pendingBelowStrategyType)
	req.SetOptionalParam("recvWindow", o.recvWindow)
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("queryOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*otocoResponse](resp)
}
