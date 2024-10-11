package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// DeleteOpenOrders 撤销单一交易对下所有挂单。这也包括了来自订单列表的挂单。
type DeleteOpenOrders interface {
	SetSymbol(symbol string) *deleteOpenOrdersRequest
	Call(ctx context.Context) (body []*deleteOpenOrdersResponse, err error)
}

type deleteOpenOrdersRequest struct {
	*binance.Client
	symbol string
}

type deleteOpenOrdersResponse struct {
	Symbol                  string                    `json:"symbol"`
	OrigClientOrderId       string                    `json:"origClientOrderId,omitempty"`
	OrderId                 int                       `json:"orderId,omitempty"`
	OrderListId             int                       `json:"orderListId"`
	ClientOrderId           string                    `json:"clientOrderId,omitempty"`
	TransactTime            int64                     `json:"transactTime,omitempty"`
	Price                   string                    `json:"price,omitempty"`
	OrigQty                 string                    `json:"origQty,omitempty"`
	ExecutedQty             string                    `json:"executedQty,omitempty"`
	CummulativeQuoteQty     string                    `json:"cummulativeQuoteQty,omitempty"`
	Status                  string                    `json:"status,omitempty"`
	TimeInForce             string                    `json:"timeInForce,omitempty"`
	Type                    enums.OrderType           `json:"type,omitempty"`
	Side                    enums.SideType            `json:"side,omitempty"`
	SelfTradePreventionMode enums.StpModeType         `json:"selfTradePreventionMode,omitempty"`
	ContingencyType         enums.ContingencyType     `json:"contingencyType,omitempty"`
	ListStatusType          enums.ListOrderStatusType `json:"listStatusType,omitempty"`
	ListOrderStatus         enums.ListOrderStatusType `json:"listOrderStatus,omitempty"`
	ListClientOrderId       string                    `json:"listClientOrderId,omitempty"`
	TransactionTime         int64                     `json:"transactionTime,omitempty"`
	Orders                  []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders,omitempty"`
	OrderReports []struct {
		Symbol                  string                `json:"symbol"`
		OrigClientOrderId       string                `json:"origClientOrderId"`
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
		StopPrice               string                `json:"stopPrice,omitempty"`
		IcebergQty              string                `json:"icebergQty"`
		SelfTradePreventionMode enums.StpModeType     `json:"selfTradePreventionMode"`
	} `json:"orderReports,omitempty"`
}

func NewDeleteOpenOrders(client *binance.Client, symbol string) DeleteOpenOrders {
	return &deleteOpenOrdersRequest{Client: client, symbol: symbol}
}
func (d *deleteOpenOrdersRequest) SetSymbol(symbol string) *deleteOpenOrdersRequest {
	d.symbol = symbol
	return d
}

func (d *deleteOpenOrdersRequest) Call(ctx context.Context) (body []*deleteOpenOrdersResponse, err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.ApiOpenOrders,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("deleteOpenOrdersRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*deleteOpenOrdersResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiDeleteOpenOrders interface {
	binance.WsApi[*WsApiDeleteOpenOrdersResponse]
	DeleteOpenOrders
}

// WsApiDeleteOpenOrdersResponse 订单和订单列表的撤销报告的格式与 order.cancel 中的格式相同。
type WsApiDeleteOpenOrdersResponse struct {
	binance.WsApiResponse
	Result []*deleteOpenOrdersResponse `json:"result"`
}

// NewWsApiDeleteOpenOrders 撤销单一交易对的所有挂单 (TRADE)
// 撤销单一交易对的所有挂单,包括交易组。
func NewWsApiDeleteOpenOrders(c *binance.Client) WsApiDeleteOpenOrders {
	return &deleteOpenOrdersRequest{Client: c}
}

// Send 撤销单一交易对的所有挂单 (TRADE)
func (d *deleteOpenOrdersRequest) Send(ctx context.Context) (*WsApiDeleteOpenOrdersResponse, error) {
	req := &binance.Request{Path: "openOrders.cancelAll"}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	return binance.WsApiHandler[*WsApiDeleteOpenOrdersResponse](ctx, d.Client, req)
}
