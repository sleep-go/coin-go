package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type QueryOrder interface {
	SetSymbol(symbol string) QueryOrder
	SetOrderId(orderId int64) QueryOrder
	SetOrigClientOrderId(origClientOrderId string) QueryOrder
	Call(ctx context.Context) (body *queryOrderResponse, err error)
}

// 至少需要发送 orderId 与 origClientOrderId中的一个
// orderId在symbol维度是自增的
type queryOrderRequest struct {
	*binance.Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
}
type queryOrderResponse struct {
	AvgPrice                string                 `json:"avgPrice"`
	ClientOrderId           string                 `json:"clientOrderId"`
	CumQuote                string                 `json:"cumQuote"`
	ExecutedQty             string                 `json:"executedQty"`
	OrderId                 int                    `json:"orderId"`
	OrigQty                 string                 `json:"origQty"`
	OrigType                enums.OrderType        `json:"origType"`
	Price                   string                 `json:"price"`
	ReduceOnly              bool                   `json:"reduceOnly"`
	Side                    enums.SideType         `json:"side"`
	PositionSide            enums.PositionSideType `json:"positionSide"`
	Status                  enums.StatusType       `json:"status"`
	StopPrice               string                 `json:"stopPrice"`
	ClosePosition           bool                   `json:"closePosition"`
	Symbol                  string                 `json:"symbol"`
	Time                    int64                  `json:"time"`
	TimeInForce             enums.TimeInForceType  `json:"timeInForce"`
	Type                    enums.OrderType        `json:"type"`
	ActivatePrice           string                 `json:"activatePrice"`
	PriceRate               string                 `json:"priceRate"`
	UpdateTime              int64                  `json:"updateTime"`
	WorkingType             enums.WorkingType      `json:"workingType"`
	PriceProtect            bool                   `json:"priceProtect"`
	PriceMatch              enums.PriceMatchType   `json:"priceMatch"`
	SelfTradePreventionMode enums.StpModeType      `json:"selfTradePreventionMode"`
	GoodTillDate            int                    `json:"goodTillDate"`
}

func NewQueryOrder(client *binance.Client, symbol string) QueryOrder {
	return &queryOrderRequest{Client: client, symbol: symbol}
}

func (d *queryOrderRequest) SetSymbol(symbol string) QueryOrder {
	d.symbol = symbol
	return d
}
func (d *queryOrderRequest) SetOrderId(orderId int64) QueryOrder {
	d.orderId = &orderId
	return d
}

func (d *queryOrderRequest) SetOrigClientOrderId(origClientOrderId string) QueryOrder {
	d.origClientOrderId = &origClientOrderId
	return d
}

func (d *queryOrderRequest) Call(ctx context.Context) (body *queryOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiOrder,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	req.SetOptionalParam("orderId", d.orderId)
	req.SetOptionalParam("origClientOrderId", d.origClientOrderId)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("queryOrderRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*queryOrderResponse](resp)
}
