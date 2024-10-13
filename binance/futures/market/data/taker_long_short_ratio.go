package data

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type TakerLongShortRatio interface {
	Call(ctx context.Context, symbol string) (body []*takerLongShortRatioResponse, err error)
	SetSymbol(symbol string) *takerLongShortRatioRequest
	SetLimit(limit enums.LimitType) *takerLongShortRatioRequest
	SetPeriod(period enums.KlineIntervalType) *takerLongShortRatioRequest
	SetStartTime(startTime int64) *takerLongShortRatioRequest
	SetEndTime(endTime int64) *takerLongShortRatioRequest
}

type takerLongShortRatioRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType
	period    enums.KlineIntervalType
	startTime *int64
	endTime   *int64
}

func (t *takerLongShortRatioRequest) SetSymbol(symbol string) *takerLongShortRatioRequest {
	t.symbol = symbol
	return t
}

func (t *takerLongShortRatioRequest) SetLimit(limit enums.LimitType) *takerLongShortRatioRequest {
	t.limit = limit
	return t
}

func (t *takerLongShortRatioRequest) SetPeriod(period enums.KlineIntervalType) *takerLongShortRatioRequest {
	t.period = period
	return t
}

func (t *takerLongShortRatioRequest) SetStartTime(startTime int64) *takerLongShortRatioRequest {
	t.startTime = &startTime
	return t
}

func (t *takerLongShortRatioRequest) SetEndTime(endTime int64) *takerLongShortRatioRequest {
	t.endTime = &endTime
	return t
}

// NewTakerLongShortRatio 合约主动买卖量
func NewTakerLongShortRatio(client *binance.Client) TakerLongShortRatio {
	return &takerLongShortRatioRequest{
		Client: client,
	}
}

type takerLongShortRatioResponse struct {
	BuySellRatio string `json:"buySellRatio"`
	BuyVol       string `json:"buyVol"`  // 主动买入量
	SellVol      string `json:"sellVol"` // 主动卖出量
	Timestamp    int    `json:"timestamp"`
}

func (t *takerLongShortRatioRequest) Call(ctx context.Context, symbol string) (body []*takerLongShortRatioResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiDataTakerLongShortRatio,
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
	return utils.ParseHttpResponse[[]*takerLongShortRatioResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
