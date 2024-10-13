package data

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type GlobalLongShortAccountRatio interface {
	Call(ctx context.Context, symbol string) (body []*globalLongShortAccountRatioResponse, err error)
	SetSymbol(symbol string) *globalLongShortAccountRatioRequest
	SetLimit(limit enums.LimitType) *globalLongShortAccountRatioRequest
	SetPeriod(period enums.KlineIntervalType) *globalLongShortAccountRatioRequest
	SetStartTime(startTime int64) *globalLongShortAccountRatioRequest
	SetEndTime(endTime int64) *globalLongShortAccountRatioRequest
}

type globalLongShortAccountRatioRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType
	period    enums.KlineIntervalType
	startTime *int64
	endTime   *int64
}

func (t *globalLongShortAccountRatioRequest) SetSymbol(symbol string) *globalLongShortAccountRatioRequest {
	t.symbol = symbol
	return t
}

func (t *globalLongShortAccountRatioRequest) SetLimit(limit enums.LimitType) *globalLongShortAccountRatioRequest {
	t.limit = limit
	return t
}

func (t *globalLongShortAccountRatioRequest) SetPeriod(period enums.KlineIntervalType) *globalLongShortAccountRatioRequest {
	t.period = period
	return t
}

func (t *globalLongShortAccountRatioRequest) SetStartTime(startTime int64) *globalLongShortAccountRatioRequest {
	t.startTime = &startTime
	return t
}

func (t *globalLongShortAccountRatioRequest) SetEndTime(endTime int64) *globalLongShortAccountRatioRequest {
	t.endTime = &endTime
	return t
}

// NewGlobalLongShortAccountRatio 多空持仓人数比
func NewGlobalLongShortAccountRatio(client *binance.Client) GlobalLongShortAccountRatio {
	return &globalLongShortAccountRatioRequest{
		Client: client,
	}
}

type globalLongShortAccountRatioResponse struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int    `json:"timestamp"`
}

func (t *globalLongShortAccountRatioRequest) Call(ctx context.Context, symbol string) (body []*globalLongShortAccountRatioResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiDataGlobalLongShortAccountRatio,
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
	return utils.ParseHttpResponse[[]*globalLongShortAccountRatioResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
