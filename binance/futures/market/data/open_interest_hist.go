package data

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type OpenInterestHist interface {
	Call(ctx context.Context, symbol string) (body []*openInterestHistResponse, err error)
	SetSymbol(symbol string) *openInterestHistRequest
	SetLimit(limit enums.LimitType) *openInterestHistRequest
	SetPeriod(period enums.KlineIntervalType) *openInterestHistRequest
	SetStartTime(startTime int64) *openInterestHistRequest
	SetEndTime(endTime int64) *openInterestHistRequest
}

type openInterestHistRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType
	period    enums.KlineIntervalType
	startTime *int64
	endTime   *int64
}

func (t *openInterestHistRequest) SetSymbol(symbol string) *openInterestHistRequest {
	t.symbol = symbol
	return t
}

func (t *openInterestHistRequest) SetLimit(limit enums.LimitType) *openInterestHistRequest {
	t.limit = limit
	return t
}

func (t *openInterestHistRequest) SetPeriod(period enums.KlineIntervalType) *openInterestHistRequest {
	t.period = period
	return t
}

func (t *openInterestHistRequest) SetStartTime(startTime int64) *openInterestHistRequest {
	t.startTime = &startTime
	return t
}

func (t *openInterestHistRequest) SetEndTime(endTime int64) *openInterestHistRequest {
	t.endTime = &endTime
	return t
}

// NewOpenInterestHist 合约持仓量历史
func NewOpenInterestHist(client *binance.Client) OpenInterestHist {
	return &openInterestHistRequest{
		Client: client,
	}
}

type openInterestHistResponse struct {
	Symbol               string `json:"symbol"`
	SumOpenInterest      string `json:"sumOpenInterest"`
	SumOpenInterestValue string `json:"sumOpenInterestValue"`
	Timestamp            int    `json:"timestamp"`
}

func (t *openInterestHistRequest) Call(ctx context.Context, symbol string) (body []*openInterestHistResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiDataOpenInterestHist,
	}
	req.SetParam("symbol", symbol)
	req.SetParam("period", t.period)
	req.SetOptionalParam("limit", t.limit)
	req.SetOptionalParam("startTime", t.startTime)
	req.SetOptionalParam("endTime", t.endTime)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*openInterestHistResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
