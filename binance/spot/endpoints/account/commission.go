package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/pkg/utils"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
)

type Commission interface {
	SetTimestamp(timestamp int64) Commission
	Call(ctx context.Context) (body *commissionResponse, err error)
}

// commissionRequest 查询佣金费率
// 获取当前账户的佣金费率。(USER_DATA)
type commissionRequest struct {
	*binance.Client
	symbol    string
	timestamp int64
}

type commissionResponse struct {
	Symbol             string `json:"symbol"`
	StandardCommission struct {
		Maker  string `json:"maker"`
		Taker  string `json:"taker"`
		Buyer  string `json:"buyer"`
		Seller string `json:"seller"`
	} `json:"standardCommission"`
	TaxCommission struct {
		Maker  string `json:"maker"`
		Taker  string `json:"taker"`
		Buyer  string `json:"buyer"`
		Seller string `json:"seller"`
	} `json:"taxCommission"`
	Discount struct {
		EnabledForAccount bool   `json:"enabledForAccount"`
		EnabledForSymbol  bool   `json:"enabledForSymbol"`
		DiscountAsset     string `json:"discountAsset"`
		Discount          string `json:"discount"`
	} `json:"discount"`
}

func NewCommission(client *binance.Client, symbol string) Commission {
	return &commissionRequest{
		Client: client,
		symbol: symbol,
	}
}
func (c *commissionRequest) SetTimestamp(timestamp int64) Commission {
	c.timestamp = timestamp
	return c
}

func (c *commissionRequest) Call(ctx context.Context) (body *commissionResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountCommission,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	req.SetParam("timestamp", c.timestamp)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("commissionRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[commissionResponse](resp)
}
