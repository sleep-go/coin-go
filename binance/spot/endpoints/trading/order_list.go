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
	SetOrderListId(orderListId int64) OrderList
	SetOrigClientOrderId(origClientOrderId string) OrderList
	SetRecvWindow(recvWindow int64) OrderList
	SetTimestamp(timestamp int64) OrderList
	Call(ctx context.Context) (body *orderListResponse, err error)
	CallDelete(ctx context.Context) (body *deleteOrderListResponse, err error)
}

type orderListRequest struct {
	*binance.Client
	orderListId       *int64  //orderListId 或 origClientOrderId 必须提供一个。
	origClientOrderId *string //orderListId 或 origClientOrderId 必须提供一个。
	newClientOrderId  *string //用户自定义的本次撤销操作的ID(注意不是被撤销的订单的自定义ID)。如无指定会自动赋值。
	recvWindow        int64
	timestamp         int64
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

func (o *orderListRequest) SetTimestamp(timestamp int64) OrderList {
	o.timestamp = timestamp
	return o
}

func (o *orderListRequest) SetRecvWindow(recvWindow int64) OrderList {
	o.recvWindow = recvWindow
	return o
}

func (o *orderListRequest) SetOrigClientOrderId(origClientOrderId string) OrderList {
	o.origClientOrderId = &origClientOrderId
	return o
}

func (o *orderListRequest) SetOrderListId(orderListId int64) OrderList {
	o.orderListId = &orderListId
	return o
}

func NewOrderList(client *binance.Client) OrderList {
	return &orderListRequest{Client: client}
}

func (o *orderListRequest) Call(ctx context.Context) (body *orderListResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTradingOrderList,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderListId", o.orderListId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	req.SetOptionalParam("recvWindow", o.recvWindow)
	req.SetParam("timestamp", o.timestamp)
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
		Path:   consts.ApiTradingOrderList,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderListId", o.orderListId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	req.SetOptionalParam("newClientOrderId", o.newClientOrderId)
	req.SetOptionalParam("recvWindow", o.recvWindow)
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("orderListRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*deleteOrderListResponse](resp)
}
