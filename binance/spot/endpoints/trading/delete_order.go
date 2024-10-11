package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type DeleteOrder interface {
	Call(ctx context.Context) (body *deleteOrderResponse, err error)
	SetSymbol(symbol string) *deleteOrderRequest
	SetOrderId(orderId int64) *deleteOrderRequest
	SetOrigClientOrderId(origClientOrderId string) *deleteOrderRequest
	SetNewClientOrderId(newClientOrderId string) *deleteOrderRequest
	SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) *deleteOrderRequest
}

// deleteOrderRequest orderId 与 origClientOrderId 必须至少发送一个.
// 如果两个参数一起发送, orderId优先被考虑.
type deleteOrderRequest struct {
	*binance.Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
	newClientOrderId  *string //用户自定义的本次撤销操作的ID(注意不是被撤销的订单的自定义ID)。如无指定会自动赋值
	//支持的值:
	//ONLY_NEW - 如果订单状态为 NEW，撤销将成功。
	//ONLY_PARTIALLY_FILLED - 如果订单状态为 PARTIALLY_FILLED，撤销将成功。
	cancelRestrictions enums.CancelRestrictionsType
}

type deleteOrderResponse struct {
	Symbol                  string                `json:"symbol"`
	OrderId                 int                   `json:"orderId"` //// 除非此单是订单列表的一部分, 否则此值为 -1
	OrderListId             int                   `json:"orderListId"`
	OrigClientOrderId       string                `json:"origClientOrderId"`
	ClientOrderId           string                `json:"clientOrderId"`
	TransactTime            int64                 `json:"transactTime"`
	Price                   string                `json:"price"`
	OrigQty                 string                `json:"origQty"`
	ExecutedQty             string                `json:"executedQty"`
	CummulativeQuoteQty     string                `json:"cummulativeQuoteQty"`
	Status                  enums.OrderStatusType `json:"status"`
	TimeInForce             enums.TimeInForceType `json:"timeInForce"`
	Type                    enums.OrderType       `json:"type"`
	Side                    enums.SideType        `json:"side"`
	SelfTradePreventionMode enums.StpModeType     `json:"selfTradePreventionMode"`
	Orders                  []struct {
		Symbol        string `json:"symbol"`
		OrderId       int    `json:"orderId"`
		ClientOrderId string `json:"clientOrderId"`
	} `json:"orders,omitempty"`
	OrderReports []struct {
		Symbol                  string                `json:"symbol"`
		OrigClientOrderId       string                `json:"origClientOrderId"`
		OrderId                 int                   `json:"orderId"`
		OrderListId             int                   `json:"orderListId"`
		ClientOrderId           string                `json:"clientOrderId"`
		TransactTime            int64                 `json:"transactTime"`
		Price                   string                `json:"price"`
		OrigQty                 string                `json:"origQty"`
		ExecutedQty             string                `json:"executedQty"`
		CummulativeQuoteQty     string                `json:"cummulativeQuoteQty"`
		Status                  enums.OrderStatusType `json:"status"`
		TimeInForce             enums.TimeInForceType `json:"timeInForce"`
		Type                    enums.OrderType       `json:"type"`
		Side                    enums.SideType        `json:"side"`
		StopPrice               string                `json:"stopPrice,omitempty"`
		IcebergQty              string                `json:"icebergQty"`
		SelfTradePreventionMode enums.StpModeType     `json:"selfTradePreventionMode"`
	} `json:"orderReports,omitempty"`
}

func NewDeleteOrder(client *binance.Client, symbol string) DeleteOrder {
	return &deleteOrderRequest{Client: client, symbol: symbol}
}

func (d *deleteOrderRequest) SetSymbol(symbol string) *deleteOrderRequest {
	d.symbol = symbol
	return d
}
func (d *deleteOrderRequest) SetOrderId(orderId int64) *deleteOrderRequest {
	d.orderId = &orderId
	return d
}

func (d *deleteOrderRequest) SetOrigClientOrderId(origClientOrderId string) *deleteOrderRequest {
	d.origClientOrderId = &origClientOrderId
	return d
}

func (d *deleteOrderRequest) SetNewClientOrderId(newClientOrderId string) *deleteOrderRequest {
	d.newClientOrderId = &newClientOrderId
	return d
}

func (d *deleteOrderRequest) SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) *deleteOrderRequest {
	d.cancelRestrictions = cancelRestrictions
	return d
}

func (d *deleteOrderRequest) Call(ctx context.Context) (body *deleteOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.ApiOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	req.SetOptionalParam("orderId", d.orderId)
	req.SetOptionalParam("origClientOrderId", d.origClientOrderId)
	req.SetOptionalParam("newClientOrderId", d.newClientOrderId)
	req.SetOptionalParam("cancelRestrictions", d.cancelRestrictions)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("deleteOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*deleteOrderResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiDeleteOrder interface {
	binance.WsApi[*WsApiDeleteOrderResponse]
	DeleteOrder
}
type WsApiDeleteOrderResponse struct {
	binance.WsApiResponse
	Result *deleteOrderResponse `json:"result"`
}

// NewWsApiDeleteOrder 撤销订单 (TRADE)
// 备注：
//
// 如果同时指定了 orderId 和 origClientOrderId 参数，仅使用 orderId 并忽略 origClientOrderId。
//
// newClientOrderId 将替换已取消订单的 clientOrderId，为新订单腾出空间。
//
// 如果您取消属于订单列表的订单，则整个订单列表将被取消。
func NewWsApiDeleteOrder(c *binance.Client) WsApiDeleteOrder {
	return &deleteOrderRequest{Client: c}
}

func (d *deleteOrderRequest) Send(ctx context.Context) (*WsApiDeleteOrderResponse, error) {
	req := &binance.Request{Path: "order.cancel"}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	req.SetOptionalParam("orderId", d.orderId)
	req.SetOptionalParam("origClientOrderId", d.origClientOrderId)
	req.SetOptionalParam("newClientOrderId", d.newClientOrderId)
	req.SetOptionalParam("cancelRestrictions", d.cancelRestrictions)
	return binance.WsApiHandler[*WsApiDeleteOrderResponse](ctx, d.Client, req)
}
