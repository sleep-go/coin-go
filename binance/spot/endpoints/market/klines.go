package market

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
)

// KlineIntervalType 支持的K线间隔 （区分大小写）
type KlineIntervalType string

const (
	//seconds -> 秒	1s

	KlineIntervals1 KlineIntervalType = "1s"

	//分钟级别 minutes -> 分钟	1m， 3m， 5m， 15m， 30m

	KlineInterval1m  KlineIntervalType = "1m"
	KlineInterval3m  KlineIntervalType = "3m"
	KlineInterval5m  KlineIntervalType = "5m"
	KlineInterval15m KlineIntervalType = "15m"
	KlineInterval30m KlineIntervalType = "30m"

	//  小时级别 hours -> 小时	1h， 2h， 4h， 6h， 8h， 12h

	KlineInterval1h  KlineIntervalType = "1h"
	KlineInterval2h  KlineIntervalType = "2h"
	KlineInterval4h  KlineIntervalType = "4h"
	KlineInterval6h  KlineIntervalType = "6h"
	KlineInterval8h  KlineIntervalType = "8h"
	KlineInterval12h KlineIntervalType = "12h"

	// 天级别 days -> 天	1d， 3d

	KlineInterval1d KlineIntervalType = "1d"
	KlineInterval3d KlineIntervalType = "3d"

	//周级别 weeks -> 周	1w

	KlineInterval1w KlineIntervalType = "1w"

	// 月级别 months -> 月	1M

	KlineInterval1M KlineIntervalType = "1M"
)

type Klines interface {
	Call(ctx context.Context) (body []*klinesResponse, err error)
	SetInterval(interval KlineIntervalType) Klines
	SetStartTime(startTime int64) Klines
	SetEndTime(endTime int64) Klines
	SetTimeZone(timeZone string) Klines
}
type klinesRequest struct {
	*binance.Client
	symbol    string
	limit     TradesLimitType   //Default 500; max 1000.
	interval  KlineIntervalType //	请参考 K线间隔
	startTime *int64
	endTime   *int64
	timeZone  string
}

// [
//
//	[
//	  1499040000000,      // 开盘时间
//	  "0.01634790",       // 开盘价
//	  "0.80000000",       // 最高价
//	  "0.01575800",       // 最低价
//	  "0.01577100",       // 收盘价(当前K线未结束的即为最新价)
//	  "148976.11427815",  // 成交量
//	  1499644799999,      // 收盘时间
//	  "2434.19055334",    // 成交额
//	  308,                // 成交笔数
//	  "1756.87402397",    // 主动买入成交量
//	  "28.46694368",      // 主动买入成交额
//	  "17928899.62484339" // 请忽略该参数
//	]
//
// ]
type klinesResponse [12]any

func NewKlines(client *binance.Client, symbol string, limit TradesLimitType) Klines {
	return &klinesRequest{Client: client, symbol: symbol, limit: limit}
}

// SetInterval k线间隔 必传
func (k *klinesRequest) SetInterval(interval KlineIntervalType) Klines {
	k.interval = interval
	return k
}

func (k *klinesRequest) SetStartTime(startTime int64) Klines {
	k.startTime = &startTime
	return k
}

func (k *klinesRequest) SetEndTime(endTime int64) Klines {
	k.endTime = &endTime
	return k
}

func (k *klinesRequest) SetTimeZone(timeZone string) Klines {
	k.timeZone = timeZone
	return k
}

func (k *klinesRequest) Call(ctx context.Context) (body []*klinesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketKLines,
	}
	req.SetParam("symbol", k.symbol)
	req.SetParam("limit", k.limit)
	if k.interval != "" {
		req.SetParam("interval", string(k.interval))
	}
	if k.startTime != nil {
		req.SetParam("startTime", *k.startTime)
	}
	if k.endTime != nil {
		req.SetParam("endTime", *k.endTime)
	}
	req.SetParam("timeZone", "0")
	if k.timeZone != "" {
		req.SetParam("timeZone", k.timeZone)
	}
	res, err := k.Client.Do(ctx, req)
	if err != nil {
		k.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		k.Debugf("ReadAll err:%v", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", bytes)
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
