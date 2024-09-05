package ticker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
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
	res, err := t.Client.Do(ctx, req)
	if err != nil {
		t.Client.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Client.Debugf("ReadAll err:%v", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", bytes)
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
