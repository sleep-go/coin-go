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
	SetLimit(limit enums.LimitType) *allOrderListRequest
	SetFormId(fromId int64) *allOrderListRequest
	SetStartTime(startTime int64) *allOrderListRequest
	Call(ctx context.Context) (body []*allOrderListResponse, err error)
}

type allOrderListRequest struct {
	*binance.Client
	symbol    string
	fromId    *int64 //提供该项后, startTime 和 endTime 都不可提供
	startTime *int64
	endTime   *int64
	limit     enums.LimitType
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
		OrderId       int    `json:"fromId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

func NewAllOrderList(client *binance.Client, limit enums.LimitType) AllOrderList {
	return &allOrderListRequest{Client: client, limit: limit}
}

func (o *allOrderListRequest) SetLimit(limit enums.LimitType) *allOrderListRequest {
	o.limit = limit
	return o
}
func (o *allOrderListRequest) SetFormId(orderId int64) *allOrderListRequest {
	o.fromId = &orderId
	return o
}

func (o *allOrderListRequest) SetStartTime(startTime int64) *allOrderListRequest {
	o.startTime = &startTime
	return o
}

func (o *allOrderListRequest) SetEndTime(endTime int64) *allOrderListRequest {
	o.endTime = &endTime
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
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("allOrderListRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*allOrderListResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiAllOrderList interface {
	binance.WsApi[*WsApiAllOrderListResponse]
	AllOrderList
}
type WsApiAllOrderListResponse struct {
	binance.WsApiResponse
	Result []*allOrderListResponse `json:"result"`
}

// NewWsApiAllOrderList 账户订单列表历史 (USER_DATA)
// 查询所有订单列表的信息，按时间范围过滤。
func NewWsApiAllOrderList(c *binance.Client) WsApiAllOrderList {
	return &allOrderListRequest{Client: c}
}

func (o *allOrderListRequest) Send(ctx context.Context) (*WsApiAllOrderListResponse, error) {
	req := &binance.Request{Path: "allOrderLists"}
	req.SetNeedSign(true)
	req.SetOptionalParam("fromId", o.fromId)
	req.SetOptionalParam("startTime", o.startTime)
	req.SetOptionalParam("endTime", o.endTime)
	req.SetOptionalParam("limit", o.limit)
	return binance.WsApiHandler[*WsApiAllOrderListResponse](ctx, o.Client, req)
}
