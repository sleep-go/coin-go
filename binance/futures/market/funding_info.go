package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type FundingInfo interface {
	Call(ctx context.Context) (body []*fundingInfoResponse, err error)
}

type fundingInfoRequest struct {
	*binance.Client
}

// NewFundingInfo 查询资金费率信息
func NewFundingInfo(client *binance.Client) FundingInfo {
	return &fundingInfoRequest{
		Client: client,
	}
}

type fundingInfoResponse struct {
	Symbol                   string `json:"symbol"`
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"`
	FundingIntervalHours     int    `json:"fundingIntervalHours"`
	Disclaimer               bool   `json:"disclaimer"`
}

// Call 这个只有正式环境返回正确数据
func (t *fundingInfoRequest) Call(ctx context.Context) (body []*fundingInfoResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketFundingInfo,
	}
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*fundingInfoResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
