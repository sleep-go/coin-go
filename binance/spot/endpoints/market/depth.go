package market

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
)

// DepthLimitType 是一个表示可选 limit 的类型
type DepthLimitType int

// 定义可选的 limit 值的枚举
const (
	DepthLimit5    DepthLimitType = 5
	DepthLimit10   DepthLimitType = 10
	DepthLimit20   DepthLimitType = 20
	DepthLimit50   DepthLimitType = 50
	DepthLimit100  DepthLimitType = 100
	DepthLimit500  DepthLimitType = 500
	DepthLimit1000 DepthLimitType = 1000
	DepthLimit5000 DepthLimitType = 5000
)

type Depth interface {
	Call(ctx context.Context) (body *depthResponse, err error)
}

// 名称	类型	是否必须	描述
// Symbol	STRING	YES
// Limit	INT	NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
// 如果 Limit > 5000, 最多返回5000条数据.
type depthRequest struct {
	*binance.Client
	Symbol string
	Limit  DepthLimitType
}
type depthResponse struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// NewDepth 深度信息
func NewDepth(c *binance.Client, symbol string, limit DepthLimitType) Depth {
	return &depthRequest{
		Client: c,
		Symbol: symbol,
		Limit:  limit,
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
	req.SetParam("limit", d.Limit)
	res, err := d.Client.Do(ctx, req)
	if err != nil {
		d.Client.Debugf("response err:%v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		d.Client.Debugf("ReadAll err:%v", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
