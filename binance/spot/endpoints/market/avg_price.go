package market

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
)

type AvgPrice interface {
	Call(ctx context.Context) (body *AvgPriceResponse, err error)
}
type AvgPriceRequest struct {
	*binance.Client
	symbol string
}

// NewAvgPrice 当前平均价格
func NewAvgPrice(client *binance.Client, symbol string) AvgPrice {
	return &AvgPriceRequest{Client: client, symbol: symbol}
}

type AvgPriceResponse struct {
	consts.ErrorResponse
	Mins      int    `json:"mins"`
	Price     string `json:"price"`
	CloseTime int64  `json:"closeTime"`
}

// Call 当前平均价格
func (k *AvgPriceRequest) Call(ctx context.Context) (body *AvgPriceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketAvgPrice,
	}
	req.SetParam("symbol", k.symbol)
	res, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	err = netutil.ParseHttpResponse(res, &body)
	if err != nil {
		k.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
