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

// TradesLimitType 是一个表示可选 limit 的类型
type TradesLimitType int

// 定义可选的 limit 值的枚举
const (
	TradesLimit500  TradesLimitType = 500
	TradesLimit1000 TradesLimitType = 1000
)

type Trades interface {
	Call(ctx context.Context) (body []*tradesResponse, err error)
}

// tradesRequest 近期成交
// 名称	类型	是否必须	描述
// symbol	STRING	YES
// limit	INT	NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
// 如果 limit > 5000, 最多返回5000条数据.
type tradesRequest struct {
	*binance.Client
	symbol string
	limit  TradesLimitType //默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
}

func NewTrades(client *binance.Client, symbol string, limit TradesLimitType) Trades {
	return &tradesRequest{
		Client: client,
		symbol: symbol,
		limit:  limit,
	}
}

type tradesResponse struct {
	Id           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

// Call 获取近期成交
// 名称	类型	是否必需	描述
// symbol	STRING	YES
// limit	INT	NO	Default 500; max 1000.
func (t *tradesRequest) Call(ctx context.Context) (body []*tradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketTrades,
	}
	req.SetParam("symbol", t.symbol)
	req.SetParam("limit", t.limit)
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
