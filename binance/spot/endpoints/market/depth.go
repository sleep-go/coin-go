package market

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
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
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type StreamDepthEvent struct {
	Stream string `json:"stream"`
	Data   struct {
		Event         string     `json:"e"`
		Time          int64      `json:"E"`
		Symbol        string     `json:"s"`
		FirstUpdateID int        `json:"U"`
		LastUpdateID  int        `json:"u"`
		Bids          [][]string `json:"b"`
		Asks          [][]string `json:"a"`
	} `json:"data"`
}
type WsDepthEvent struct {
	Event string     `json:"e"`
	E1    int64      `json:"E"`
	S     string     `json:"s"`
	U     int        `json:"U"`
	U1    int        `json:"u"`
	B     [][]string `json:"b"`
	A     [][]string `json:"a"`
}

// NewDepth 深度信息
func NewDepth(c *binance.Client, symbol string, limit enums.LimitType) Depth {
	return &depthRequest{
		Client: c,
		Symbol: symbol,
		limit:  limit,
	}
}

func newWsDepth[T WsDepthEvent | StreamDepthEvent](c *binance.WsClient, symbols []string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.Endpoint
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	fmt.Println(endpoint)
	wsHandler := func(mt int, msg []byte) {
		event := new(T)
		err := json.Unmarshal(msg, &event)
		if err != nil {
			exception(mt, err)
			return
		}
		handler(event)
	}
	return c.WsServe(endpoint, wsHandler, exception)
}
func NewWsDepth(c *binance.WsClient, symbols []string, handler binance.Handler[WsDepthEvent], exception binance.ErrorHandler) error {
	return newWsDepth[WsDepthEvent](c, symbols, handler, exception)
}
func NewStreamDepth(c *binance.WsClient, symbols []string, handler binance.Handler[StreamDepthEvent], exception binance.ErrorHandler) error {
	return newWsDepth[StreamDepthEvent](c, symbols, handler, exception)
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
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*depthResponse](resp)
}
