package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// OTO 发送一个新的 OTO 订单。
//
// 一个 OTO 订单（One-Triggers-the-Other）是一个包含了两个订单的订单列表.
// 第一个订单被称为生效订单，必须为 LIMIT 或 LIMIT_MAKER 类型的订单。最初，订单簿上只有生效订单。
// 第二个订单被称为待处理订单。它可以是任何订单类型，但不包括使用参数 quoteOrderQty 的 MARKET 订单。只有当生效订单完全成交时，待处理订单才会被自动下单。
// 如果生效订单或者待处理订单中的任意一个被单独取消，订单列表中剩余的那个订单也会被随之取消或过期。
// 如果生效订单在下订单列表后立即完全成交，则可能会得到订单响应。其中，生效订单的状态为 FILLED ，但待处理订单的状态为 PENDING_NEW。针对这类情况，如果需要检查当前状态，您可以查询相关的待处理订单。
// OTO 订单将2 个订单添加到未成交订单计数，EXCHANGE_MAX_NUM_ORDERS 过滤器和 MAX_NUM_ORDERS 过滤器中。
type OTO interface {
	SetListClientOrderId(listClientOrderId string) OTO
	SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) OTO
	SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) OTO
	SetWorkingType(workingType enums.OrderType) OTO
	SetWorkingSide(workingSide enums.SideType) OTO
	SetWorkingClientOrderId(workingClientOrderId string) OTO
	SetWorkingPrice(workingPrice string) OTO
	SetWorkingQuantity(workingQuantity string) OTO
	SetWorkingIcebergQty(workingIcebergQty string) OTO
	SetWorkingTimeInForce(workingTimeInForce enums.TimeInForceType) OTO
	SetWorkingStrategyId(workingStrategyId int64) OTO
	SetWorkingStrategyType(workingStrategyType int64) OTO
	SetPendingType(pendingType enums.OrderType) OTO
	SetPendingSide(pendingSide enums.SideType) OTO
	SetPendingClientOrderId(pendingClientOrderId string) OTO
	SetPendingPrice(pendingPrice string) OTO
	SetPendingStopPrice(pendingStopPrice string) OTO
	SetPendingTrailingDelta(pendingTrailingDelta string) OTO
	SetPendingQuantity(pendingQuantity string) OTO
	SetPendingIcebergQty(pendingIcebergQty string) OTO
	SetPendingTimeInForce(pendingTimeInForce enums.TimeInForceType) OTO
	SetPendingStrategyId(pendingStrategyId int64) OTO
	SetPendingStrategyType(pendingStrategyType int64) OTO
	SetRecvWindow(recvWindow int64) OTO
	SetTimestamp(timestamp int64) OTO
	Call(ctx context.Context) (body *otoResponse, err error)
}

type otoRequest struct {
	*binance.Client
	symbol string
	//整个订单列表的唯一ID。 如果未发送则自动生成。
	//仅当前一个订单列表已填满或完全过期时，才会接受含有相同 listClientOrderId 值的新订单列表。
	//listClientOrderId 与 workingClientOrderId 和 pendingClientOrderId 不同。
	listClientOrderId       *string
	newOrderRespType        enums.NewOrderRespType
	selfTradePreventionMode enums.StpModeType
	workingType             enums.OrderType //支持的数值： LIMIT， LIMIT_MAKER
	workingSide             enums.SideType
	workingClientOrderId    *string //用于标识生效订单的唯一ID。 如果未发送则自动生成。
	workingPrice            *string
	workingQuantity         *string //用于设置生效订单的数量。
	workingIcebergQty       *string //只有当 workingTimeInForce 为 GTC 时才能使用。
	workingTimeInForce      enums.TimeInForceType
	workingStrategyId       *int64 //订单策略中用于标识生效订单的 ID。
	workingStrategyType     *int64 //用于标识生效订单策略的任意数值。 小于 1000000 的值被保留，无法使用。
	pendingType             enums.OrderType
	pendingSide             enums.SideType
	pendingClientOrderId    *string
	pendingPrice            *string
	pendingStopPrice        *string
	pendingTrailingDelta    *string
	pendingQuantity         *string //用于设置待处理订单的数量。
	pendingIcebergQty       *string //只有当 pendingTimeInForce 为 GTC 或者当 pendingType 为 LIMIT_MAKER 时才能使用。
	pendingTimeInForce      enums.TimeInForceType
	pendingStrategyId       *int64 //订单策略中用于标识待处理订单的 ID。
	pendingStrategyType     *int64 //用于标识待处理订单策略的任意数值。小于 1000000 的值被保留，无法使用。
	recvWindow              int64
	timestamp               int64
}
type otoResponse struct {
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

func (o *otoRequest) SetListClientOrderId(listClientOrderId string) OTO {
	o.listClientOrderId = &listClientOrderId
	return o
}

func (o *otoRequest) SetNewOrderRespType(newOrderRespType enums.NewOrderRespType) OTO {
	o.newOrderRespType = newOrderRespType
	return o
}

func (o *otoRequest) SetSelfTradePreventionMode(selfTradePreventionMode enums.StpModeType) OTO {
	o.selfTradePreventionMode = selfTradePreventionMode
	return o
}

func (o *otoRequest) SetWorkingType(workingType enums.OrderType) OTO {
	o.workingType = workingType
	return o
}

func (o *otoRequest) SetWorkingSide(workingSide enums.SideType) OTO {
	o.workingSide = workingSide
	return o
}

func (o *otoRequest) SetWorkingClientOrderId(workingClientOrderId string) OTO {
	o.workingClientOrderId = &workingClientOrderId
	return o
}

func (o *otoRequest) SetWorkingPrice(workingPrice string) OTO {
	o.workingPrice = &workingPrice
	return o
}

func (o *otoRequest) SetWorkingQuantity(workingQuantity string) OTO {
	o.workingQuantity = &workingQuantity
	return o
}

func (o *otoRequest) SetWorkingIcebergQty(workingIcebergQty string) OTO {
	o.workingIcebergQty = &workingIcebergQty
	return o
}

func (o *otoRequest) SetWorkingTimeInForce(workingTimeInForce enums.TimeInForceType) OTO {
	o.workingTimeInForce = workingTimeInForce
	return o
}

func (o *otoRequest) SetWorkingStrategyId(workingStrategyId int64) OTO {
	o.workingStrategyId = &workingStrategyId
	return o
}

func (o *otoRequest) SetWorkingStrategyType(workingStrategyType int64) OTO {
	o.workingStrategyType = &workingStrategyType
	return o
}

func (o *otoRequest) SetPendingType(pendingType enums.OrderType) OTO {
	o.pendingType = pendingType
	return o
}

func (o *otoRequest) SetPendingSide(pendingSide enums.SideType) OTO {
	o.pendingSide = pendingSide
	return o
}

func (o *otoRequest) SetPendingClientOrderId(pendingClientOrderId string) OTO {
	o.pendingClientOrderId = &pendingClientOrderId
	return o
}

func (o *otoRequest) SetPendingPrice(pendingPrice string) OTO {
	o.pendingPrice = &pendingPrice
	return o
}

func (o *otoRequest) SetPendingStopPrice(pendingStopPrice string) OTO {
	o.pendingStopPrice = &pendingStopPrice
	return o
}

func (o *otoRequest) SetPendingTrailingDelta(pendingTrailingDelta string) OTO {
	o.pendingTrailingDelta = &pendingTrailingDelta
	return o
}

func (o *otoRequest) SetPendingQuantity(pendingQuantity string) OTO {
	o.pendingQuantity = &pendingQuantity
	return o
}

func (o *otoRequest) SetPendingIcebergQty(pendingIcebergQty string) OTO {
	o.pendingIcebergQty = &pendingIcebergQty
	return o
}

func (o *otoRequest) SetPendingTimeInForce(pendingTimeInForce enums.TimeInForceType) OTO {
	o.pendingTimeInForce = pendingTimeInForce
	return o
}

func (o *otoRequest) SetPendingStrategyId(pendingStrategyId int64) OTO {
	o.pendingStrategyId = &pendingStrategyId
	return o
}

func (o *otoRequest) SetPendingStrategyType(pendingStrategyType int64) OTO {
	o.pendingStrategyType = &pendingStrategyType
	return o
}

func (o *otoRequest) SetRecvWindow(recvWindow int64) OTO {
	o.recvWindow = recvWindow
	return o
}

func (o *otoRequest) SetTimestamp(timestamp int64) OTO {
	o.timestamp = timestamp
	return o
}

func NewOto(client *binance.Client, symbol string) OTO {
	return &otoRequest{Client: client, symbol: symbol}
}

func (o *otoRequest) Call(ctx context.Context) (body *otoResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiTradingOrderListOTO,
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
	req.SetParam("pendingType", o.pendingType)
	req.SetParam("pendingSide", o.pendingSide)
	req.SetOptionalParam("pendingClientOrderId", o.pendingClientOrderId)
	req.SetOptionalParam("pendingPrice", o.pendingPrice)
	req.SetOptionalParam("pendingStopPrice", o.pendingStopPrice)
	req.SetOptionalParam("pendingTrailingDelta", o.pendingTrailingDelta)
	req.SetParam("pendingQuantity", o.pendingQuantity)
	req.SetOptionalParam("pendingIcebergQty", o.pendingIcebergQty)
	req.SetOptionalParam("pendingTimeInForce", o.pendingTimeInForce)
	req.SetOptionalParam("pendingStrategyId", o.pendingStrategyId)
	req.SetOptionalParam("pendingStrategyType", o.pendingStrategyType)
	req.SetOptionalParam("recvWindow", o.recvWindow)
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("queryOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*otoResponse](resp)
}
