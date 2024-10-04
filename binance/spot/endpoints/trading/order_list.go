package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance/consts/enums"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type OrderList interface {
	SetOrderListId(orderListId int64) OrderList
	SetOrigClientOrderId(origClientOrderId string) OrderList
	SetRecvWindow(recvWindow int64) OrderList
	SetTimestamp(timestamp int64) OrderList
	Call(ctx context.Context) (body []*orderListResponse, err error)
}

type orderListRequest struct {
	*binance.Client
	orderListId       *int64  //orderListId 或 origClientOrderId 必须提供一个。
	origClientOrderId *string //orderListId 或 origClientOrderId 必须提供一个。
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

func (o *orderListRequest) Call(ctx context.Context) (body []*orderListResponse, err error) {
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
	return utils.ParseHttpResponse[[]*orderListResponse](resp)
}
