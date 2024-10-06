package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type OpenOrderList interface {
	SetOrderListId(orderListId int64) OpenOrderList
	SetRecvWindow(recvWindow int64) OpenOrderList
	SetTimestamp(timestamp int64) OpenOrderList
	Call(ctx context.Context) (body []*openOrderListResponse, err error)
}

type openOrderListRequest struct {
	*binance.Client
	orderListId *int64
	recvWindow  int64
	timestamp   int64
}

type openOrderListResponse struct {
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

func (o *openOrderListRequest) SetOrderListId(orderListId int64) OpenOrderList {
	o.orderListId = &orderListId
	return o
}
func (o *openOrderListRequest) SetTimestamp(timestamp int64) OpenOrderList {
	o.timestamp = timestamp
	return o
}
func (o *openOrderListRequest) SetRecvWindow(recvWindow int64) OpenOrderList {
	o.recvWindow = recvWindow
	return o
}

func NewOpenOrderList(client *binance.Client) OpenOrderList {
	return &openOrderListRequest{Client: client}
}

// Call 查询订单列表挂单 (USER_DATA)
func (o *openOrderListRequest) Call(ctx context.Context) (body []*openOrderListResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTradingOrderList,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderListId", o.orderListId)
	req.SetOptionalParam("recvWindow", o.recvWindow)
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("openOrderListRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*openOrderListResponse](resp)
}
