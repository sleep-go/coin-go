package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type OrderList interface {
	SetOrderListId(orderListId int64) *orderListRequest
	SetOrigClientOrderId(origClientOrderId string) *orderListRequest
	SetNewClientOrderId(newClientOrderId string) *orderListRequest
	Call(ctx context.Context) (body *orderListResponse, err error)
	CallDelete(ctx context.Context) (body *deleteOrderListResponse, err error)
}

type orderListRequest struct {
	*binance.Client
	orderListId       *int64  //orderListId 或 origClientOrderId 必须提供一个。
	origClientOrderId *string //orderListId 或 origClientOrderId 必须提供一个。
	newClientOrderId  *string //用户自定义的本次撤销操作的ID(注意不是被撤销的订单的自定义ID)。如无指定会自动赋值。
}

type orderListResponse struct {
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
}
type deleteOrderListResponse struct {
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
		Symbol              string                `json:"symbol"`
		OrigClientOrderId   string                `json:"origClientOrderId"`
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
		StopPrice           string                `json:"stopPrice,omitempty"`
	} `json:"orderReports"`
}

func (o *orderListRequest) SetNewClientOrderId(newClientOrderId string) *orderListRequest {
	o.newClientOrderId = &newClientOrderId
	return o
}
func (o *orderListRequest) SetOrigClientOrderId(origClientOrderId string) *orderListRequest {
	o.origClientOrderId = &origClientOrderId
	return o
}

func (o *orderListRequest) SetOrderListId(orderListId int64) *orderListRequest {
	o.orderListId = &orderListId
	return o
}

func NewOrderList(client *binance.Client) OrderList {
	return &orderListRequest{Client: client}
}

func (o *orderListRequest) Call(ctx context.Context) (body *orderListResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiOrderList,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderListId", o.orderListId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("orderListRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*orderListResponse](resp)
}

// CallDelete 取消订单列表 (TRADE)
// 其他注意点:
//
// 取消订单列表中的单个订单将取消整个订单列表.
// 如果 orderListId 和 listClientOrderId 一起发送, orderListId 优先被考虑.
func (o *orderListRequest) CallDelete(ctx context.Context) (body *deleteOrderListResponse, err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.ApiOrderList,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderListId", o.orderListId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	req.SetOptionalParam("newClientOrderId", o.newClientOrderId)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("orderListRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*deleteOrderListResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiOrderList interface {
	binance.WsApi[*WsApiOrderListResponse]
	OrderList
	SendDelete(ctx context.Context) (*WsApiDeleteOrderListResponse, error)
}
type WsApiOrderListResponse struct {
	binance.WsApiResponse
	Result *orderListResponse `json:"result"`
}
type WsApiDeleteOrderListResponse struct {
	binance.WsApiResponse
	Result *deleteOrderListResponse `json:"result"`
}

// NewWsApiOrderList 查询订单列表 (USER_DATA)
// 检查订单列表的执行状态。
//
// 对于单个订单的执行状态，使用 order.status。
// 备注：
//
// origClientOrderId 指的是订单列表本身的 listClientOrderId。
//
// 如果同时指定了 origClientOrderId 和 orderListId 参数，仅使用 origClientOrderId 并忽略 orderListId。
func NewWsApiOrderList(c *binance.Client) WsApiOrderList {
	return &orderListRequest{Client: c}
}

// Send 查询订单列表 (USER_DATA)
func (o *orderListRequest) Send(ctx context.Context) (*WsApiOrderListResponse, error) {
	req := &binance.Request{Path: "orderList.status"}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderListId", o.orderListId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	return binance.WsApiHandler[*WsApiOrderListResponse](ctx, o.Client, req)
}

// SendDelete 撤销订单列表订单(TRADE)
// 取消整个订单列表。
func (o *orderListRequest) SendDelete(ctx context.Context) (*WsApiDeleteOrderListResponse, error) {
	req := &binance.Request{Path: "orderList.cancel"}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderListId", o.orderListId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	req.SetOptionalParam("newClientOrderId", o.newClientOrderId)
	return binance.WsApiHandler[*WsApiDeleteOrderListResponse](ctx, o.Client, req)
}
