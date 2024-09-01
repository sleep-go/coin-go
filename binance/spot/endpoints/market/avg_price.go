package market

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
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
	res, err := k.Client.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		k.Debugf("ReadAll err:%v", err)
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
