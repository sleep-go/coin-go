package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Constituents interface {
	Call(ctx context.Context, symbol string) (body *constituentsResponse, err error)
}

type constituentsRequest struct {
	*binance.Client
}

// NewConstituents 查询指数价格成分
func NewConstituents(client *binance.Client) Constituents {
	return &constituentsRequest{
		Client: client,
	}
}

type constituentsResponse struct {
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
	Constituents []struct {
		Exchange string `json:"exchange"`
		Symbol   string `json:"symbol"`
	} `json:"constituents"`
}

func (t *constituentsRequest) Call(ctx context.Context, symbol string) (body *constituentsResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketConstituents,
	}
	req.SetParam("symbol", symbol)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*constituentsResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
