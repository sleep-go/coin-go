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
		req.SetParam("symbols", result)
	}
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*priceResponse](resp)
}
