package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// AllOrderList 查询所有订单列表 (USER_DATA)
// 根据提供的可选参数检索所有的订单列表。
// 请注意，startTime和endTime之间的时间不能超过 24 小时。
type AllOrderList interface {
	SetRecvWindow(recvWindow int64) AllOrderList
	SetTimestamp(timestamp int64) AllOrderList
	Call(ctx context.Context) (body []*allOrderListResponse, err error)
}

type allOrderListRequest struct {
	*binance.Client
	fromId     *int64 //提供该项后, startTime 和 endTime 都不可提供
	startTime  *int64
	endTime    *int64
	limit      enums.LimitType
	recvWindow int64
	timestamp  int64
}
type allOrderListResponse struct {
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

func NewAllOrderList(client *binance.Client, limit enums.LimitType) AllOrderList {
	return &allOrderListRequest{Client: client, limit: limit}
}
func (o *allOrderListRequest) SetFromId(fromId int64) AllOrderList {
	o.fromId = &fromId
	return o
}

func (o *allOrderListRequest) SetStartTime(startTime int64) AllOrderList {
	o.startTime = &startTime
	return o
}

func (o *allOrderListRequest) SetEndTime(endTime int64) AllOrderList {
	o.endTime = &endTime
	return o
}

func (o *allOrderListRequest) SetTimestamp(timestamp int64) AllOrderList {
	o.timestamp = timestamp
	return o
}

func (o *allOrderListRequest) SetRecvWindow(recvWindow int64) AllOrderList {
	o.recvWindow = recvWindow
	return o
}

// Call 查询所有订单列表 (USER_DATA)
func (o *allOrderListRequest) Call(ctx context.Context) (body []*allOrderListResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountAllOrderList,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("fromId", o.fromId)
	req.SetOptionalParam("startTime", o.startTime)
	req.SetOptionalParam("endTime", o.endTime)
	req.SetOptionalParam("limit", o.limit)
	req.SetOptionalParam("recvWindow", o.recvWindow)
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("allOrderListRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*allOrderListResponse](resp)
}
