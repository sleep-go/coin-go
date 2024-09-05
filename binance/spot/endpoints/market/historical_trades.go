package market

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
)

type historyTradesRequest struct {
	*binance.Client
	symbol string
	limit  TradesLimitType //Default 500; max 1000.
	fromId *uint64         //从哪一条成交id开始返回. 缺省返回最近的成交记录
}

type HistoryTrades interface {
	Call(ctx context.Context) (body []*tradesResponse, err error)
	SetFromId(fromId uint64) HistoryTrades
}

func NewHistoryTrades(client *binance.Client, symbol string, limit TradesLimitType) HistoryTrades {
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
	res, err := t.Client.Do(ctx, req)
	if err != nil {
		t.Client.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Client.Debugf("ReadAll err:%v", err)
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
