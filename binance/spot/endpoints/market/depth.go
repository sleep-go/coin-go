package market

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
)

type Depth interface {
	Call(ctx context.Context) (body *depthResponse, err error)
}

// 名称	类型	是否必须	描述
// Symbol	STRING	YES
// limit	INT	NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
// 如果 limit > 5000, 最多返回5000条数据.
type depthRequest struct {
	*binance.Client
	Symbol string
	limit  enums.LimitType
}
type depthResponse struct {
	consts.ErrorResponse
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// NewDepth 深度信息
func NewDepth(c *binance.Client, symbol string, limit enums.LimitType) Depth {
	return &depthRequest{
		Client: c,
		Symbol: symbol,
		limit:  limit,
	}
}

// Call 深度信息
// 默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
// 如果 limit > 5000, 最多返回5000条数据.
// 注意: limit=0 返回全部orderbook，但数据量会非常非常非常非常大！
func (d *depthRequest) Call(ctx context.Context) (body *depthResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiMarketDepth,
	}
	req.SetParam("symbol", d.Symbol)
	req.SetParam("limit", d.limit)
	res, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("response err:%v", err)
		return nil, err
	}
	err = netutil.ParseHttpResponse(res, &body)
	if err != nil {
		d.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
