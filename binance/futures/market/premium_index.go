package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type PremiumIndex interface {
	Call(ctx context.Context, symbol string) (body *premiumIndexResponse, err error)
	CallAll(ctx context.Context) (body []*premiumIndexResponse, err error)
}

type premiumIndexRequest struct {
	*binance.Client
}

// NewPremiumIndex Call 最新标记价格和资金费率
// 采集各大交易所数据加权平均
func NewPremiumIndex(client *binance.Client) PremiumIndex {
	return &premiumIndexRequest{
		Client: client,
	}
}

type premiumIndexResponse struct {
	Symbol               string `json:"symbol"`               // 交易对
	MarkPrice            string `json:"markPrice"`            // 标记价格
	IndexPrice           string `json:"indexPrice"`           // 指数价格
	EstimatedSettlePrice string `json:"estimatedSettlePrice"` // 预估结算价,仅在交割开始前最后一小时有意义
	LastFundingRate      string `json:"lastFundingRate"`      // 最近更新的资金费率
	NextFundingTime      int64  `json:"nextFundingTime"`      // 下次资金费时间
	InterestRate         string `json:"interestRate"`         // 标的资产基础利率
	Time                 int64  `json:"time"`                 // 更新时间
}

func (t *premiumIndexRequest) Call(ctx context.Context, symbol string) (body *premiumIndexResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketPremiumIndex,
	}
	req.SetParam("symbol", symbol)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*premiumIndexResponse](resp)
}
func (t *premiumIndexRequest) CallAll(ctx context.Context) (body []*premiumIndexResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketPremiumIndex,
	}
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*premiumIndexResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
