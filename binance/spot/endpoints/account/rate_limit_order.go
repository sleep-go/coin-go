package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type RateLimitOrder interface {
	Call(ctx context.Context) (body []*rateLimitOrderResponse, err error)
}

type rateLimitOrderRequest struct {
	*binance.Client
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

// Call 查询未成交的订单计数 (USER_DATA)
// 显示用户在所有时间间隔内的未成交订单计数。
func (r *rateLimitOrderRequest) Call(ctx context.Context) (body []*rateLimitOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountRateLimitOrder,
	}
	req.SetNeedSign(true)
	resp, err := r.Do(ctx, req)
	if err != nil {
		r.Debugf("rateLimitOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*rateLimitOrderResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiRateLimitOrder interface {
	binance.WsApi[*WsApiRateLimitOrderResponse]
}
type WsApiRateLimitOrderResponse struct {
	binance.WsApiResponse
	Result []*rateLimitOrderResponse `json:"result"`
}

// NewWsApiWsApiRateLimitOrder 查询未成交的订单计数 (USER_DATA)
// 显示用户在所有时间间隔内的未成交订单计数。
func NewWsApiWsApiRateLimitOrder(c *binance.Client) WsApiRateLimitOrder {
	return &rateLimitOrderRequest{Client: c}
}

func (r *rateLimitOrderRequest) Send(ctx context.Context) (*WsApiRateLimitOrderResponse, error) {
	req := &binance.Request{Path: "account.rateLimits.orders"}
	req.SetNeedSign(true)
	return binance.WsApiHandler[*WsApiRateLimitOrderResponse](ctx, r.Client, req)
}
