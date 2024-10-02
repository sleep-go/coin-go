package trading

import (
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts/enums"
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
	aboveStopPrice     *string         //如果 aboveType 是 STOP_LOSS 或 STOP_LOSS_LIMIT 才能使用。必须指定 aboveStopPrice 或 aboveTrailingDelta 或两者。
	aboveTrailingDelta *int64          //请看 追踪止盈止损(Trailing Stop)订单常见问题。
	aboveTimeInForce   *string         //如果 aboveType 是 STOP_LOSS_LIMIT，则为必填项。
	aboveStrategyId    *int64          //订单策略中上方订单的 ID。
	aboveStrategyType  *int64          //上方订单策略的任意数值。小于 1000000 的值被保留，无法使用。
	belowType          enums.OrderType //支持值：STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER。
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
	recvWindow              int64
	timestamp               int64
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
