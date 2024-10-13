package ticker

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Price interface {
	CallV1(ctx context.Context, symbol string) (body *priceResponse, err error)
	CallAllV1(ctx context.Context) (body []*priceResponse, err error)
	CallV2(ctx context.Context, symbol string) (body *priceResponse, err error)
	CallAllV2(ctx context.Context) (body []*priceResponse, err error)
}

type priceRequest struct {
	*binance.Client
}

type priceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}

// NewPrice 最新价格接口
func NewPrice(client *binance.Client) Price {
	return &priceRequest{Client: client}
}

// CallV1 最新价格接口
// 返回最近价格
func (t *priceRequest) CallV1(ctx context.Context, symbol string) (body *priceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTickerPrice,
	}
	req.SetParam("symbol", symbol)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*priceResponse](resp)
}
func (t *priceRequest) CallAllV1(ctx context.Context) (body []*priceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTickerPrice,
	}
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*priceResponse](resp)
}

// CallV2 最新价格接口
// 返回最近价格
func (t *priceRequest) CallV2(ctx context.Context, symbol string) (body *priceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTickerPriceV2,
	}
	req.SetParam("symbol", symbol)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*priceResponse](resp)
}
func (t *priceRequest) CallAllV2(ctx context.Context) (body []*priceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketTickerPriceV2,
	}
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*priceResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiTickerPrice interface {
	binance.WsApi[*WsApiTickerPriceResponse]
}
type WsApiTickerPriceResponse struct {
	binance.WsApiResponse
	Result []*priceResponse `json:"result"`
}

// NewWsApiTickerPrice 最新价格
// 获取交易对最新价格
//
// 如果需要访问实时价格更新，请考虑使用 WebSocket Streams:
//
// <symbol>@aggTrade
// <symbol>@trade
func NewWsApiTickerPrice(c *binance.Client) WsApiTickerPrice {
	return &priceRequest{Client: c}
}

func (t *priceRequest) Send(ctx context.Context) (*WsApiTickerPriceResponse, error) {
	req := &binance.Request{Path: "ticker.price"}
	return binance.WsApiHandler[*WsApiTickerPriceResponse](ctx, t.Client, req)
}
