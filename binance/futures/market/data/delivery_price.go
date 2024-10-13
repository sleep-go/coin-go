package data

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type DeliveryPrice interface {
	Call(ctx context.Context) (body []*deliveryPriceResponse, err error)
}

type deliveryPriceRequest struct {
	*binance.Client
	pair  string
	limit enums.LimitType
}
type deliveryPriceResponse struct {
	DeliveryTime  int64   `json:"deliveryTime"`
	DeliveryPrice float64 `json:"deliveryPrice"`
}

// NewDeliveryPrice 季度合约历史结算价
func NewDeliveryPrice(c *binance.Client, pair string) DeliveryPrice {
	return &deliveryPriceRequest{
		Client: c,
		pair:   pair,
	}
}

func (d *deliveryPriceRequest) Call(ctx context.Context) (body []*deliveryPriceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiDataDeliveryPrice,
	}
	req.SetParam("pair", d.pair)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*deliveryPriceResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
