package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type FundingRate interface {
	SetSymbol(symbol string) *fundingRateRequest
	SetStartTime(startTime int64) *fundingRateRequest
	SetEndTime(endTime int64) *fundingRateRequest
	SetLimit(limit enums.LimitType) *fundingRateRequest
	Call(ctx context.Context, symbol string) (body []*fundingRateResponse, err error)
}

type fundingRateRequest struct {
	*binance.Client
	symbol    string
	startTime *int64
	endTime   *int64
	limit     enums.LimitType
}

func (t *fundingRateRequest) SetSymbol(symbol string) *fundingRateRequest {
	t.symbol = symbol
	return t
}

func (t *fundingRateRequest) SetStartTime(startTime int64) *fundingRateRequest {
	t.startTime = &startTime
	return t
}

func (t *fundingRateRequest) SetEndTime(endTime int64) *fundingRateRequest {
	t.endTime = &endTime
	return t
}

func (t *fundingRateRequest) SetLimit(limit enums.LimitType) *fundingRateRequest {
	t.limit = limit
	return t
}

// NewFundingRate 查询资金费率历史
func NewFundingRate(client *binance.Client) FundingRate {
	return &fundingRateRequest{
		Client: client,
	}
}

type fundingRateResponse struct {
	Symbol      string `json:"symbol"`      // 交易对
	FundingRate string `json:"fundingRate"` // 资金费率
	FundingTime int64  `json:"fundingTime"` // 资金费时间
	MarkPrice   string `json:"markPrice"`   // 资金费对应标记价格
}

func (t *fundingRateRequest) Call(ctx context.Context, symbol string) (body []*fundingRateResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketFundingRate,
	}
	req.SetOptionalParam("symbol", symbol)
	req.SetOptionalParam("startTime", t.startTime)
	req.SetOptionalParam("endTime", t.endTime)
	req.SetOptionalParam("limit", t.limit)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*fundingRateResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
