package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type LvKlines interface {
	Call(ctx context.Context) (body []*lvKlinesResponse, err error)
	SetSymbol(symbol string) *LvKlinesRequest
	SetInterval(interval enums.KlineIntervalType) *LvKlinesRequest
	SetStartTime(startTime int64) *LvKlinesRequest
	SetEndTime(endTime int64) *LvKlinesRequest
	SetLimit(limit enums.LimitType) *LvKlinesRequest
}
type LvKlinesRequest struct {
	*binance.Client
	symbol    string
	interval  enums.KlineIntervalType
	startTime *int64
	endTime   *int64
	limit     enums.LimitType //Default 500; max 1000.
}

// [
//
//		[
//			1598371200000,		// 开盘时间
//			"5.88275270",		// 开盘净值
//	     	"6.03142087",		// 最高净值
//	     	"5.85749741",		// 最低净值
//	     	"5.99403551",		// 收盘净值(当前K线未结束的即为最新净值)
//	     	"2.28602984",		// 收盘真实杠杆
//	     	1598374799999,		// 收盘时间
//	     	"0",				// 请忽略
//	     	6209,				// 净值更新笔数
//	     	"14517.64507907",	// 请忽略
//	     	"0",				// 请忽略
//	     	"0"					// 请忽略
//		]
//
// ]

type lvKlinesResponse [12]any

// NewLvKlines 杠杆代币历史净值K线
// 杠杆代币历史净值K线，杠杆代币净值系统基于合约架构，故该接口采用fapi
func NewLvKlines(client *binance.Client, symbol string, limit enums.LimitType) LvKlines {
	return &LvKlinesRequest{Client: client, symbol: symbol, limit: limit}
}

func (k *LvKlinesRequest) SetSymbol(symbol string) *LvKlinesRequest {
	k.symbol = symbol
	return k
}
func (k *LvKlinesRequest) SetLimit(limit enums.LimitType) *LvKlinesRequest {
	k.limit = limit
	return k
}

// SetInterval k线间隔 必传
func (k *LvKlinesRequest) SetInterval(interval enums.KlineIntervalType) *LvKlinesRequest {
	k.interval = interval
	return k
}

func (k *LvKlinesRequest) SetStartTime(startTime int64) *LvKlinesRequest {
	k.startTime = &startTime
	return k
}

func (k *LvKlinesRequest) SetEndTime(endTime int64) *LvKlinesRequest {
	k.endTime = &endTime
	return k
}

func (k *LvKlinesRequest) Call(ctx context.Context) (body []*lvKlinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketLvKlines,
	}
	req.SetParam("symbol", k.symbol)
	req.SetParam("limit", k.limit)
	req.SetOptionalParam("interval", k.interval)
	req.SetOptionalParam("startTime", k.startTime)
	req.SetOptionalParam("endTime", k.endTime)
	resp, err := k.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*lvKlinesResponse](resp)
}
