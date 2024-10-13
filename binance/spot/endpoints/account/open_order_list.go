package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/spot/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type OpenOrderList interface {
	Call(ctx context.Context) (body []*openOrderListResponse, err error)
}

type openOrderListRequest struct {
	*binance.Client
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
		OrderId       int    `json:"fromId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders"`
}

func NewOpenOrderList(client *binance.Client) OpenOrderList {
	return &openOrderListRequest{Client: client}
}

// Call 查询订单列表挂单 (USER_DATA)
func (o *openOrderListRequest) Call(ctx context.Context) (body []*openOrderListResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTradingOpenOrderList,
	}
	req.SetNeedSign(true)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("openOrderListRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*openOrderListResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiOpenOrderList interface {
	binance.WsApi[*WsApiOpenOrderListResponse]
}
type WsApiOpenOrderListResponse struct {
	binance.WsApiResponse
	Result []*openOrderListResponse `json:"result"`
}

// NewWsApiOpenOrderList 查询订单列表挂单 (USER_DATA)
// 查询所有订单列表挂单的执行状态。
//
// 如果您需要持续监控订单状态更新，请考虑使用 WebSocket Streams：
//
// userDataStream.start 请求
// executionReport 更新
func NewWsApiOpenOrderList(c *binance.Client) WsApiOpenOrderList {
	return &openOrderListRequest{Client: c}
}

func (o *openOrderListRequest) Send(ctx context.Context) (*WsApiOpenOrderListResponse, error) {
	req := &binance.Request{Path: "openOrderLists.status"}
	req.SetNeedSign(true)
	return binance.WsApiHandler[*WsApiOpenOrderListResponse](ctx, o.Client, req)
}
