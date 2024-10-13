package ticker

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Price interface {
	Call(ctx context.Context) (body []*priceResponse, err error)
}

type priceRequest struct {
	*binance.Client
	symbols []string
}

type priceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// NewPrice 最新价格接口
func NewPrice(client *binance.Client, symbols []string) Price {
	return &priceRequest{Client: client, symbols: symbols}
}

// Call 最新价格接口
func (t *priceRequest) Call(ctx context.Context) (body []*priceResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketTickerPrice,
	}
	if len(t.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(t.symbols, `","`))
		req.SetParam("symbol", result)
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
	SetSymbols(symbols []string) WsApiTickerPrice
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

func (t *priceRequest) SetSymbols(symbols []string) WsApiTickerPrice {
	t.symbols = symbols
	return t
}

func (t *priceRequest) Send(ctx context.Context) (*WsApiTickerPriceResponse, error) {
	req := &binance.Request{Path: "ticker.price"}
	if len(t.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(t.symbols, `","`))
		req.SetParam("symbol", result)
	}
	return binance.WsApiHandler[*WsApiTickerPriceResponse](ctx, t.Client, req)
}
