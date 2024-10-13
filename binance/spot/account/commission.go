package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Commission interface {
	SetSymbol(symbol string) *commissionRequest
	Call(ctx context.Context) (body *commissionResponse, err error)
}

// commissionRequest 查询佣金费率
// 获取当前账户的佣金费率。(USER_DATA)
type commissionRequest struct {
	*binance.Client
	symbol string
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

func (c *commissionRequest) SetSymbol(symbol string) *commissionRequest {
	c.symbol = symbol
	return c
}

func (c *commissionRequest) Call(ctx context.Context) (body *commissionResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountCommission,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("commissionRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*commissionResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiCommission interface {
	binance.WsApi[*WsApiCommissionResponse]
	Commission
}
type WsApiCommissionResponse struct {
	binance.WsApiResponse
	Result *commissionResponse `json:"result"`
}

// NewWsApiCommission 账户佣金费率 (USER_DATA)
// 获取当前账户的佣金费率。
func NewWsApiCommission(c *binance.Client) WsApiCommission {
	return &commissionRequest{Client: c}
}

func (c *commissionRequest) Send(ctx context.Context) (*WsApiCommissionResponse, error) {
	req := &binance.Request{Path: "account.commission"}
	req.SetNeedSign(true)
	req.SetParam("symbol", c.symbol)
	return binance.WsApiHandler[*WsApiCommissionResponse](ctx, c.Client, req)
}
