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

type BookTicker interface {
	Call(ctx context.Context) (body []*bookTickerResponse, err error)
}

type bookTickerRequest struct {
	*binance.Client
	symbols []string
}
type bookTickerResponse struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

func NewBookTicker(client *binance.Client, symbols []string) BookTicker {
	return &bookTickerRequest{Client: client, symbols: symbols}
}
func (b *bookTickerRequest) Call(ctx context.Context) (body []*bookTickerResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketTickerBookTicker,
	}
	if len(b.symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(b.symbols, `","`))
		req.SetParam("symbols", result)
	}
	res, err := b.Client.Do(ctx, req)
	if err != nil {
		b.Client.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		b.Client.Debugf("ReadAll err:%v", err)
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
