package market

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

type AggTrades interface {
	Call(ctx context.Context) (body []*aggTradesResponse, err error)
	SetFromId(fromId uint64) AggTrades
	SetStartTime(startTime int64) AggTrades
	SetEndTime(endTime int64) AggTrades
}

type aggTradesRequest struct {
	*binance.Client
	symbol    string
	limit     enums.LimitType //默认 500; 最大 1000.
	fromId    *uint64         //从包含fromID的成交开始返回结果
	startTime *int64          //从该时刻之后的成交记录开始返回结果
	endTime   *int64          //返回该时刻为止的成交记录
}
type aggTradesResponse struct {
	A  int    `json:"a"` //归集成交ID
	P  string `json:"p"` // 成交价
	Q  string `json:"q"` // 成交量
	F  int    `json:"f"` // 被归集的首个成交ID
	L  int    `json:"l"` // 被归集的末个成交ID
	T  int64  `json:"T"` // 成交时间
	M  bool   `json:"m"` // 是否为主动卖出单
	M1 bool   `json:"M"` // 是否为最优撮合单(可忽略，目前总为最优撮合)
}

func NewAggTrades(client *binance.Client, symbol string, limit enums.LimitType) AggTrades {
	return &aggTradesRequest{Client: client, symbol: symbol, limit: limit}
}

func (a *aggTradesRequest) SetFromId(fromId uint64) AggTrades {
	a.fromId = &fromId
	return a
}

func (a *aggTradesRequest) SetStartTime(startTime int64) AggTrades {
	a.startTime = &startTime
	return a
}

func (a *aggTradesRequest) SetEndTime(endTime int64) AggTrades {
	a.endTime = &endTime
	return a
}

// Call 近期成交(归集)
// 与trades的区别是，同一个taker在同一时间同一价格与多个maker的成交会被合并为一条记录
func (a *aggTradesRequest) Call(ctx context.Context) (body []*aggTradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketAggTrades,
	}
	req.SetParam("symbol", a.symbol)
	req.SetParam("limit", a.limit)
	if a.fromId != nil {
		req.SetParam("fromId", *a.fromId)
	}
	if a.startTime != nil {
		req.SetParam("startTime", *a.startTime)
	}
	if a.endTime != nil {
		req.SetParam("endTime", *a.endTime)
	}
	res, err := a.Do(ctx, req)
	if err != nil {
		a.Debugf("response err:%v", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		var e *consts.ErrorResponse
		err = netutil.ParseHttpResponse(res, &e)
		return nil, consts.Error(e.Code, e.Msg)
	}
	err = netutil.ParseHttpResponse(res, &body)
	if err != nil {
		a.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
