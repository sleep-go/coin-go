package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type OpenInterest interface {
	Call(ctx context.Context, symbol string) (body *openInterestResponse, err error)
}

type openInterestRequest struct {
	*binance.Client
}

// NewOpenInterest 获取未平仓合约数
func NewOpenInterest(client *binance.Client) OpenInterest {
	return &openInterestRequest{
		Client: client,
	}
}

type openInterestResponse struct {
	OpenInterest string `json:"openInterest"` // 未平仓合约数量
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
}

func (t *openInterestRequest) Call(ctx context.Context, symbol string) (body *openInterestResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketOpenInterest,
	}
	req.SetParam("symbol", symbol)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*openInterestResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
