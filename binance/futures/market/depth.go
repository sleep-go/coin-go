package market

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type Depth interface {
	Call(ctx context.Context) (body *depthResponse, err error)
}

// 名称	类型	是否必须	描述
// symbol	STRING	YES
// limit	INT	NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
// 如果 limit > 5000, 最多返回5000条数据.
type depthRequest struct {
	*binance.Client
	symbol string
	limit  enums.LimitType
}

//	{
//	 "lastUpdateId": 1027024,
//	 "E": 1589436922972,   // 消息时间
//	 "T": 1589436922959,   // 撮合引擎时间
//	 "bids": [				// 买单
//	   [
//	     "4.00000000",     // 价格
//	     "431.00000000"    // 数量
//	   ]
//	 ],
//	 "asks": [				// 卖单
//	   [
//	     "4.00000200",		// 价格
//	     "12.00000000"		// 数量
//	   ]
//	 ]
//	}

type depthResponse struct {
	LastUpdateId int        `json:"lastUpdateId"`
	E            int64      `json:"E"`    // 消息时间
	T            int64      `json:"T"`    // 撮合引擎时间
	Bids         [][]string `json:"bids"` // 买单
	Asks         [][]string `json:"asks"` // 卖单
}

// NewDepth 深度信息
func NewDepth(c *binance.Client, symbol string, limit enums.LimitType) Depth {
	return &depthRequest{
		Client: c,
		symbol: symbol,
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
		Path:   consts.FApiMarketDepth,
	}
	req.SetParam("symbol", d.symbol)
	req.SetParam("limit", d.limit)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*depthResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

type StreamDepthEvent struct {
	Stream string        `json:"stream"`
	Data   *WsDepthEvent `json:"data"`
}
type WsDepthEvent struct {
	Event         string     `json:"e"`
	Time          int64      `json:"E"`
	Symbol        string     `json:"s"`
	FirstUpdateID int        `json:"U"`
	LastUpdateID  int        `json:"u"`
	Bids          [][]string `json:"b"`
	Asks          [][]string `json:"a"`
}

// NewWsDepth 增量深度信息
// 每秒推送orderbook的变化部分（如果有）
func NewWsDepth(c *binance.Client, symbols []string, handler binance.Handler[*WsDepthEvent], exception binance.ErrorHandler) error {
	return wsDepth[*WsDepthEvent](c, symbols, handler, exception)
}

// NewStreamDepth 增量深度信息
// // 每秒推送orderbook的变化部分（如果有）
func NewStreamDepth(c *binance.Client, symbols []string, handler binance.Handler[*StreamDepthEvent], exception binance.ErrorHandler) error {
	return wsDepth[*StreamDepthEvent](c, symbols, handler, exception)
}
func wsDepth[T *StreamDepthEvent | *WsDepthEvent](c *binance.Client, symbols []string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.BaseURL
	for _, s := range symbols {
		if c.IsFast {
			endpoint += fmt.Sprintf("%s@depth@100ms", strings.ToLower(s)) + "/"
		} else {
			endpoint += fmt.Sprintf("%s@depth", strings.ToLower(s)) + "/"
		}
	}
	endpoint = endpoint[:len(endpoint)-1]
	return binance.WsHandler(c, endpoint, handler, exception)
}

type StreamDepthLevelsEvent struct {
	Stream string             `json:"stream"`
	Data   WsDepthLevelsEvent `json:"data"`
}
type WsDepthLevelsEvent struct {
	depthResponse
}

// NewWsDepthLevels 有限档深度信息
// 每秒推送有限档深度信息。levels 表示几档买卖单信息, 可选 5/10/20档
func NewWsDepthLevels(c *binance.Client, symbolLevels map[string]enums.LimitType, handler binance.Handler[WsDepthLevelsEvent], exception binance.ErrorHandler) error {
	return wsDepthLevels[WsDepthLevelsEvent](c, symbolLevels, handler, exception)
}

// NewStreamDepthLevels 有限档深度信息
// 每秒推送有限档深度信息。levels 表示几档买卖单信息, 可选 5/10/20档
func NewStreamDepthLevels(c *binance.Client, symbolLevels map[string]enums.LimitType, handler binance.Handler[StreamDepthLevelsEvent], exception binance.ErrorHandler) error {
	return wsDepthLevels[StreamDepthLevelsEvent](c, symbolLevels, handler, exception)
}
func wsDepthLevels[T WsDepthLevelsEvent | StreamDepthLevelsEvent](c *binance.Client, symbolLevels map[string]enums.LimitType, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.BaseURL
	for s, l := range symbolLevels {
		if c.IsFast {
			endpoint += fmt.Sprintf("%s@depth%d@100ms", strings.ToLower(s), l) + "/"
		} else {
			endpoint += fmt.Sprintf("%s@depth%d", strings.ToLower(s), l) + "/"
		}
	}
	endpoint = endpoint[:len(endpoint)-1]
	return binance.WsHandler(c, endpoint, handler, exception)
}

// ****************************** Websocket Api *******************************

type WsApiDepth interface {
	binance.WsApi[*WsApiDepthResponse]
	SetSymbol(symbol string) WsApiDepth
	SetLimit(limit enums.LimitType) WsApiDepth
}
type WsApiDepthResponse struct {
	binance.WsApiResponse
	Result *depthResponse `json:"result"`
}

// NewWsApiDepth 获取当前深度信息。
//
// 请注意，此请求返回有限的市场深度。
//
// 如果需要持续监控深度信息更新，请考虑使用 WebSocket Streams：
//
// <symbol>@depth<levels>
// <symbol>@depth
// 如果需要维护本地orderbook，您可以将 depth 请求与 <symbol>@depth streams 一起使用。
func NewWsApiDepth(c *binance.Client) WsApiDepth {
	return &depthRequest{
		Client: c,
	}
}
func (d *depthRequest) SetSymbol(symbol string) WsApiDepth {
	d.symbol = symbol
	return d
}

func (d *depthRequest) SetLimit(limit enums.LimitType) WsApiDepth {
	d.limit = limit
	return d
}
func (d *depthRequest) Send(ctx context.Context) (*WsApiDepthResponse, error) {
	req := &binance.Request{Path: "depth"}
	req.SetOptionalParam("symbol", d.symbol)
	req.SetParam("limit", d.limit)
	return binance.WsApiHandler[*WsApiDepthResponse](ctx, d.Client, req)
}
