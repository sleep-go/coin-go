package market

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
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
	limit  enums.LimitType //默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
}

func NewTrades(client *binance.Client, symbol string, limit enums.LimitType) Trades {
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
