package trading

import (
	"context"
	"errors"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
)

type GetOrder interface {
	Call(ctx context.Context) (body *GetOrderResponse, err error)
	SetOrderId(orderId int64) GetOrder
	SetRecvWindow(recvWindow int64) GetOrder
	SetOrigClientOrderId(origClientOrderId string) GetOrder
	SetTimestamp(timestamp int64) GetOrder
}
type GetOrderRequest struct {
	*binance.Client
	symbol            string
	orderId           int64
	origClientOrderId string
	recvWindow        int64
	timestamp         int64
}

func (o *GetOrderRequest) Call(ctx context.Context) (body *GetOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTradingOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", o.symbol)
	if o.orderId > 0 {
		req.SetParam("orderId", o.orderId)
	}
	if o.origClientOrderId != "" {
		req.SetParam("origClientOrderId", o.origClientOrderId)
	}
	if o.recvWindow > 0 {
		req.SetParam("recvWindow", o.recvWindow)
	}
	req.SetParam("timestamp", o.timestamp)
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("GetOrderRequest response err:%v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		o.Debugf("StatusCode no ok:%v", resp.Body)
		return nil, errors.New(resp.Status)
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		o.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}

// GetOrderResponse 查询订单 (USER_DATA)
// 至少需要发送 orderId 与 origClientOrderId中的一个
// 某些订单中cummulativeQuoteQty<0，是由于这些订单是cummulativeQuoteQty功能上线之前的订单。
type GetOrderResponse struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int    `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	IcebergQty              string `json:"icebergQty"`
	Time                    int64  `json:"time"`
	UpdateTime              int64  `json:"updateTime"`
	IsWorking               bool   `json:"isWorking"`
	WorkingTime             int64  `json:"workingTime"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

func NewGetOrder(client *binance.Client, symbol string) *GetOrderRequest {
	return &GetOrderRequest{Client: client, symbol: symbol}
}

func (o *GetOrderRequest) SetOrderId(orderId int64) GetOrder {
	o.orderId = orderId
	return o
}

func (o *GetOrderRequest) SetOrigClientOrderId(origClientOrderId string) GetOrder {
	o.origClientOrderId = origClientOrderId
	return o
}

func (o *GetOrderRequest) SetRecvWindow(recvWindow int64) GetOrder {
	o.recvWindow = recvWindow
	return o
}

func (o *GetOrderRequest) SetTimestamp(timestamp int64) GetOrder {
	o.timestamp = timestamp
	return o
}
