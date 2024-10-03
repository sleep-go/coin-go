package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type RateLimitOrder interface {
	SetRecvWindow(recvWindow int64) RateLimitOrder
	SetTimestamp(timestamp int64) RateLimitOrder
	Call(ctx context.Context) (body []*rateLimitOrderResponse, err error)
}

type rateLimitOrderRequest struct {
	*binance.Client
	recvWindow int64
	timestamp  int64
}

type rateLimitOrderResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

func NewRateLimitOrder(client *binance.Client) RateLimitOrder {
	return &rateLimitOrderRequest{Client: client}
}

func (r *rateLimitOrderRequest) SetRecvWindow(recvWindow int64) RateLimitOrder {
	r.recvWindow = recvWindow
	return r
}

func (r *rateLimitOrderRequest) SetTimestamp(timestamp int64) RateLimitOrder {
	r.timestamp = timestamp
	return r
}

// Call 查询未成交的订单计数 (USER_DATA)
// 显示用户在所有时间间隔内的未成交订单计数。
func (r *rateLimitOrderRequest) Call(ctx context.Context) (body []*rateLimitOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountRateLimitOrder,
	}
	req.SetNeedSign(true)
	if r.recvWindow > 0 {
		req.SetParam("recvWindow", r.recvWindow)
	}
	req.SetParam("timestamp", r.timestamp)
	resp, err := r.Do(ctx, req)
	if err != nil {
		r.Debugf("rateLimitOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*rateLimitOrderResponse](resp)
}
