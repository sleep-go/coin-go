package market

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

type historyTradesRequest struct {
	*binance.Client
	symbol string
	limit  enums.LimitType //Default 500; max 1000.
	fromId *uint64         //从哪一条成交id开始返回. 缺省返回最近的成交记录
}

type HistoryTrades interface {
	Call(ctx context.Context) (body []*tradesResponse, err error)
	SetFromId(fromId uint64) HistoryTrades
}

func NewHistoryTrades(client *binance.Client, symbol string, limit enums.LimitType) HistoryTrades {
	return &historyTradesRequest{Client: client, symbol: symbol, limit: limit}
}

func (t *historyTradesRequest) SetFromId(fromId uint64) HistoryTrades {
	t.fromId = &fromId
	return t
}

// Call 查询历史成交
// 权重: 25
// 名称	类型	是否必需	描述
// symbol	STRING	YES
// limit	INT	NO	Default 500; max 1000.
// fromId	LONG	NO	从哪一条成交id开始返回. 缺省返回最近的成交记录
func (t *historyTradesRequest) Call(ctx context.Context) (body []*tradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketHistoricalTrades,
	}
	req.SetParam("symbol", t.symbol)
	req.SetParam("limit", t.limit)
	if t.fromId != nil {
		req.SetParam("fromId", *t.fromId)
	}
	res, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		var e *consts.ErrorResponse
		err = netutil.ParseHttpResponse(res, &e)
		return nil, consts.Error(e.Code, e.Msg)
	}
	err = netutil.ParseHttpResponse(res, &body)
	if err != nil {
		t.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
