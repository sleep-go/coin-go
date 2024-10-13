package data

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type TopLongShortPositionRatio interface {
	Call(ctx context.Context, symbol string) (body []*topLongShortPositionRatioResponse, err error)
	SetSymbol(symbol string) *topLongShortPositionRatioRequest
	SetLimit(limit enums.LimitType) *topLongShortPositionRatioRequest
	SetPeriod(period enums.KlineIntervalType) *topLongShortPositionRatioRequest
	SetStartTime(startTime int64) *topLongShortPositionRatioRequest
	SetEndTime(endTime int64) *topLongShortPositionRatioRequest
}

type topLongShortPositionRatioRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType
	period    enums.KlineIntervalType
	startTime *int64
	endTime   *int64
}

func (t *topLongShortPositionRatioRequest) SetSymbol(symbol string) *topLongShortPositionRatioRequest {
	t.symbol = symbol
	return t
}

func (t *topLongShortPositionRatioRequest) SetLimit(limit enums.LimitType) *topLongShortPositionRatioRequest {
	t.limit = limit
	return t
}

func (t *topLongShortPositionRatioRequest) SetPeriod(period enums.KlineIntervalType) *topLongShortPositionRatioRequest {
	t.period = period
	return t
}

func (t *topLongShortPositionRatioRequest) SetStartTime(startTime int64) *topLongShortPositionRatioRequest {
	t.startTime = &startTime
	return t
}

func (t *topLongShortPositionRatioRequest) SetEndTime(endTime int64) *topLongShortPositionRatioRequest {
	t.endTime = &endTime
	return t
}

// NewTopLongShortPositionRatio 大户持仓量多空比
// 大户的多头和空头总持仓量占比，大户指保证金余额排名前20%的用户。
// 多仓持仓量比例 = 大户多仓持仓量 / 大户总持仓量 空仓持仓量比例 = 大户空仓持仓量 / 大户总持仓量 多空持仓量比值 = 多仓持仓量比例 / 空仓持仓量比例
func NewTopLongShortPositionRatio(client *binance.Client) TopLongShortPositionRatio {
	return &topLongShortPositionRatioRequest{
		Client: client,
	}
}

type topLongShortPositionRatioResponse struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int    `json:"timestamp"`
}

func (t *topLongShortPositionRatioRequest) Call(ctx context.Context, symbol string) (body []*topLongShortPositionRatioResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiDataTopLongShortPositionRatio,
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
	return utils.ParseHttpResponse[[]*topLongShortPositionRatioResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
