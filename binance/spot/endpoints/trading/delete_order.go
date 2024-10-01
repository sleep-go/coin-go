package trading

import (
	"context"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"net/http"
)

type DeleteOrder interface {
	Call(ctx context.Context) (body *deleteOrderResponse, err error)
	SetOrderId(orderId int64) DeleteOrder
	SetOrigClientOrderId(origClientOrderId string) DeleteOrder
	SetNewClientOrderId(newClientOrderId string) DeleteOrder
	SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) DeleteOrder
	SetRecvWindow(recvWindow int64) DeleteOrder
	SetTimestamp(timestamp int64) DeleteOrder
}

// deleteOrderRequest orderId 与 origClientOrderId 必须至少发送一个.
// 如果两个参数一起发送, orderId优先被考虑.
type deleteOrderRequest struct {
	*binance.Client
	symbol            string
	orderId           int64
	origClientOrderId string
	newClientOrderId  string //用户自定义的本次撤销操作的ID(注意不是被撤销的订单的自定义ID)。如无指定会自动赋值
	//支持的值:
	//ONLY_NEW - 如果订单状态为 NEW，撤销将成功。
	//ONLY_PARTIALLY_FILLED - 如果订单状态为 PARTIALLY_FILLED，撤销将成功。
	cancelRestrictions enums.CancelRestrictionsType
	recvWindow         int64
	timestamp          int64
}

type deleteOrderResponse struct {
	consts.ErrorResponse
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
}

func NewDeleteOrder(client *binance.Client, symbol string) DeleteOrder {
	return &deleteOrderRequest{Client: client, symbol: symbol}
}

func (d *deleteOrderRequest) SetOrderId(orderId int64) DeleteOrder {
	d.orderId = orderId
	return d
}

func (d *deleteOrderRequest) SetOrigClientOrderId(origClientOrderId string) DeleteOrder {
	d.origClientOrderId = origClientOrderId
	return d
}

func (d *deleteOrderRequest) SetNewClientOrderId(newClientOrderId string) DeleteOrder {
	d.newClientOrderId = newClientOrderId
	return d
}

func (d *deleteOrderRequest) SetCancelRestrictions(cancelRestrictions enums.CancelRestrictionsType) DeleteOrder {
	d.cancelRestrictions = cancelRestrictions
	return d
}

func (d *deleteOrderRequest) SetRecvWindow(recvWindow int64) DeleteOrder {
	d.recvWindow = recvWindow
	return d
}

func (d *deleteOrderRequest) SetTimestamp(timestamp int64) DeleteOrder {
	d.timestamp = timestamp
	return d
}
func (d *deleteOrderRequest) Call(ctx context.Context) (body *deleteOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.ApiTradingOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	if d.orderId > 0 {
		req.SetParam("orderId", d.orderId)
	}
	if d.origClientOrderId != "" {
		req.SetParam("origClientOrderId", d.origClientOrderId)
	}
	if d.recvWindow > 0 {
		req.SetParam("recvWindow", d.recvWindow)
	}
	req.SetParam("timestamp", d.timestamp)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("deleteOrderRequest response err:%v", err)
		return nil, err
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		d.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
