package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

// OrderAmendment 查询订单修改历史
type OrderAmendment interface {
	SetOrderId(orderId int64) *orderAmendmentRequest
	SetOrigClientOrderId(origClientOrderId string) *orderAmendmentRequest
	SetSymbol(symbol string) *orderAmendmentRequest
	SetLimit(limit enums.LimitType) *orderAmendmentRequest
	SetStartTime(startTime int64) *orderAmendmentRequest
	SetEndTime(endTime int64) *orderAmendmentRequest
}

type orderAmendmentRequest struct {
	*binance.Client
	symbol            string
	orderId           *int64          //系统订单号
	origClientOrderId *string         //用户自定义的订单号
	startTime         *int64          //起始时间
	endTime           *int64          //结束时间
	limit             enums.LimitType //返回的结果集数量 默认值:50 最大值:100
}
type orderAmendmentResponse struct {
	AmendmentId   int    `json:"amendmentId"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	Time          int64  `json:"time"`
	Amendment     struct {
		Price struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"price"`
		OrigQty struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"origQty"`
		Count int `json:"count"` // 修改记数，代表该修改记录是这笔订单第几次修改
	} `json:"amendment"`
}

func NewOrderAmendment(client *binance.Client, limit enums.LimitType) OrderAmendment {
	return &orderAmendmentRequest{Client: client, limit: limit}
}
func (o *orderAmendmentRequest) SetOrderId(orderId int64) *orderAmendmentRequest {
	o.orderId = &orderId
	return o
}
func (o *orderAmendmentRequest) SetOrigClientOrderId(origClientOrderId string) *orderAmendmentRequest {
	o.origClientOrderId = &origClientOrderId
	return o
}
func (o *orderAmendmentRequest) SetSymbol(symbol string) *orderAmendmentRequest {
	o.symbol = symbol
	return o
}
func (o *orderAmendmentRequest) SetLimit(limit enums.LimitType) *orderAmendmentRequest {
	o.limit = limit
	return o
}

func (o *orderAmendmentRequest) SetStartTime(startTime int64) *orderAmendmentRequest {
	o.startTime = &startTime
	return o
}

func (o *orderAmendmentRequest) SetEndTime(endTime int64) *orderAmendmentRequest {
	o.endTime = &endTime
	return o
}

// Call 查询订单修改历史 (USER_DATA)
func (o *orderAmendmentRequest) Call(ctx context.Context) (body []*orderAmendmentResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiAccountOrderAmendment,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	req.SetOptionalParam("orderId", o.orderId)
	req.SetOptionalParam("origClientOrderId", o.origClientOrderId)
	req.SetOptionalParam("startTime", o.startTime)
	req.SetOptionalParam("endTime", o.endTime)
	req.SetOptionalParam("limit", o.limit)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("orderAmendmentRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*orderAmendmentResponse](resp)
}

// ****************************** Websocket Api *******************************
