package data

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type TopLongShortAccountRatio interface {
	Call(ctx context.Context, symbol string) (body []*topLongShortAccountRatioResponse, err error)
	SetSymbol(symbol string) *topLongShortAccountRatioRequest
	SetLimit(limit enums.LimitType) *topLongShortAccountRatioRequest
	SetPeriod(period enums.KlineIntervalType) *topLongShortAccountRatioRequest
	SetStartTime(startTime int64) *topLongShortAccountRatioRequest
	SetEndTime(endTime int64) *topLongShortAccountRatioRequest
}

type topLongShortAccountRatioRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType
	period    enums.KlineIntervalType
	startTime *int64
	endTime   *int64
}

func (t *topLongShortAccountRatioRequest) SetSymbol(symbol string) *topLongShortAccountRatioRequest {
	t.symbol = symbol
	return t
}

func (t *topLongShortAccountRatioRequest) SetLimit(limit enums.LimitType) *topLongShortAccountRatioRequest {
	t.limit = limit
	return t
}

func (t *topLongShortAccountRatioRequest) SetPeriod(period enums.KlineIntervalType) *topLongShortAccountRatioRequest {
	t.period = period
	return t
}

func (t *topLongShortAccountRatioRequest) SetStartTime(startTime int64) *topLongShortAccountRatioRequest {
	t.startTime = &startTime
	return t
}

func (t *topLongShortAccountRatioRequest) SetEndTime(endTime int64) *topLongShortAccountRatioRequest {
	t.endTime = &endTime
	return t
}

// NewTopLongShortAccountRatio 大户账户数多空比
// 持仓大户的净持仓多头和空头账户数占比，大户指保证金余额排名前20%的用户。一个账户记一次。
// 多仓账户数比例 = 持多仓大户数 / 总持仓大户数 空仓账户数比例 = 持空仓大户数 / 总持仓大户数 多空账户数比值 = 多仓账户数比例 / 空仓账户数比例
func NewTopLongShortAccountRatio(client *binance.Client) TopLongShortAccountRatio {
	return &topLongShortAccountRatioRequest{
		Client: client,
	}
}

type topLongShortAccountRatioResponse struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int    `json:"timestamp"`
}

func (t *topLongShortAccountRatioRequest) Call(ctx context.Context, symbol string) (body []*topLongShortAccountRatioResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiDataTopLongShortAccountRatio,
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
	return utils.ParseHttpResponse[[]*topLongShortAccountRatioResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
