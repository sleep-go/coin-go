package trading

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type UpdateOrder interface {
	SetOrderId(orderId int) *UpdateOrderRequest
	SetOrigClientOrderId(origClientOrderId string) *UpdateOrderRequest
	SetSymbol(symbol string) *UpdateOrderRequest
	SetSide(side enums.SideType) *UpdateOrderRequest
	SetPriceMatch(priceMatch enums.PriceMatchType) *UpdateOrderRequest
	SetQuantity(quantity string) *UpdateOrderRequest
	SetPrice(price string) *UpdateOrderRequest
	Call(ctx context.Context) (body *updateOrderResponse, err error)
	CallBatch(ctx context.Context, data []*UpdateOrderRequest) (body []*updateOrderResponse, err error)
}

// orderId 与 origClientOrderId 必须至少发送一个，同时发送则以 order id为准
// quantity 与 price 均必须发送，这点和 dapi 修改订单不同
// 当新订单的quantity 或 price不满足PRICE_FILTER / PERCENT_FILTER / LOT_SIZE限制，修改会被拒绝，原订单依旧被保留
// 订单会在下列情况下被取消：
// 原订单被部分执行且新订单quantity <= executedQty
// 原订单是GTX，新订单的价格会导致订单立刻执行
// 同一订单修改次数最多10000次
// 改单会将selfTradePreventionMode重置为NONE
type UpdateOrderRequest struct {
	*binance.Client
	OrderId           *int                 `json:"orderId,omitempty"`           //系统订单号
	OrigClientOrderId *string              `json:"origClientOrderId,omitempty"` //用户自定义的订单号
	Symbol            string               `json:"symbol,omitempty"`            //交易对
	Side              enums.SideType       `json:"side,omitempty"`              //订单方向
	Quantity          *string              `json:"quantity"`                    //下单数量,使用closePosition不支持此参数。
	Price             *string              `json:"price,omitempty"`             //委托价格
	PriceMatch        enums.PriceMatchType `json:"priceMatch,omitempty"`        //OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}
type updateOrderResponse struct {
	OrderId                 int64                  `json:"orderId"`
	Symbol                  string                 `json:"symbol"`
	Pair                    string                 `json:"pair"`
	Status                  enums.StatusType       `json:"status"`
	ClientOrderId           string                 `json:"clientOrderId"`
	Price                   string                 `json:"price"`
	AvgPrice                string                 `json:"avgPrice"`
	OrigQty                 string                 `json:"origQty"`
	ExecutedQty             string                 `json:"executedQty"`
	CumQty                  string                 `json:"cumQty"`
	CumBase                 string                 `json:"cumBase"`
	TimeInForce             enums.TimeInForceType  `json:"timeInForce"`
	Type                    enums.OrderType        `json:"type"`
	ReduceOnly              bool                   `json:"reduceOnly"`
	ClosePosition           bool                   `json:"closePosition"`
	Side                    enums.SideType         `json:"side"`
	PositionSide            enums.PositionSideType `json:"positionSide"`
	StopPrice               string                 `json:"stopPrice"`
	WorkingType             enums.WorkingType      `json:"workingType"`
	PriceProtect            bool                   `json:"priceProtect"`
	OrigType                enums.OrderType        `json:"origType"`
	PriceMatch              enums.PriceMatchType   `json:"priceMatch"`
	SelfTradePreventionMode enums.StpModeType      `json:"selfTradePreventionMode"`
	GoodTillDate            int                    `json:"goodTillDate"`
	UpdateTime              int64                  `json:"updateTime"`
}

func (c *UpdateOrderRequest) SetOrderId(orderId int) *UpdateOrderRequest {
	c.OrderId = &orderId
	return c
}
func (c *UpdateOrderRequest) SetOrigClientOrderId(origClientOrderId string) *UpdateOrderRequest {
	c.OrigClientOrderId = &origClientOrderId
	return c
}
func (c *UpdateOrderRequest) SetSymbol(symbol string) *UpdateOrderRequest {
	c.Symbol = symbol
	return c
}
func (c *UpdateOrderRequest) SetSide(side enums.SideType) *UpdateOrderRequest {
	c.Side = side
	return c
}
func (c *UpdateOrderRequest) SetPriceMatch(priceMatch enums.PriceMatchType) *UpdateOrderRequest {
	c.PriceMatch = priceMatch
	return c
}
func (c *UpdateOrderRequest) SetQuantity(quantity string) *UpdateOrderRequest {
	c.Quantity = &quantity
	return c
}
func (c *UpdateOrderRequest) SetPrice(price string) *UpdateOrderRequest {
	c.Price = &price
	return c
}

func NewUpdateOrder(client *binance.Client, symbol string) UpdateOrder {
	return &UpdateOrderRequest{Client: client, Symbol: symbol}
}
func (c *UpdateOrderRequest) Call(ctx context.Context) (body *updateOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPut,
		Path:   consts.FApiOrder,
	}
	req.SetNeedSign(true)
	req.SetOptionalParam("orderId", c.OrderId)
	req.SetOptionalParam("origClientOrderId", c.OrigClientOrderId)
	req.SetParam("symbol", c.Symbol)
	req.SetParam("side", c.Side)
	req.SetParam("quantity", c.Quantity)
	req.SetParam("price", c.Price)
	req.SetOptionalParam("priceMatch", c.PriceMatch)
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("UpdateOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*updateOrderResponse](resp)
}

// CallBatch 批量下单(TRADE)
// 具体订单条件规则,与普通修改订单一致
// 批量修改订单采取并发处理,不保证订单撮合顺序
// 批量修改订单的返回内容顺序,与订单列表顺序一致
// 同一订单修改次数最多10000次
// 改单不支持设置selfTradePreventionMode并会将selfTradePreventionMode重置为NONE
func (c *UpdateOrderRequest) CallBatch(ctx context.Context, data []*UpdateOrderRequest) (body []*updateOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPut,
		Path:   consts.FApiBatchOrders,
	}
	req.SetNeedSign(true)
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req.SetParam("batchOrders", string(bytes))
	resp, err := c.Do(ctx, req)
	if err != nil {
		c.Debugf("CallBatch response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*updateOrderResponse](resp)
}

// ****************************** Websocket Api *******************************
