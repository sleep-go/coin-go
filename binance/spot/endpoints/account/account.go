package account

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/errors"
)

type Account interface {
	Call(ctx context.Context) (body *getAccountResponse, err error)
	SetOmitZeroBalances(omitZeroBalances bool) Account
	SetTimestamp(timestamp int64) Account
	SetRecvWindow(recvWindow int64) Account
}

type getAccountRequest struct {
	*binance.Client
	omitZeroBalances bool //如果true，将隐藏所有零余额。默认值：false
	recvWindow       int64
	timestamp        int64
}

type getAccountResponse struct {
	AccountType string `json:"accountType"`
	Balances    []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
	Brokered        bool `json:"brokered"`
	BuyerCommission int  `json:"buyerCommission"`
	CanDeposit      bool `json:"canDeposit"`
	CanTrade        bool `json:"canTrade"`
	CanWithdraw     bool `json:"canWithdraw"`
	CommissionRates struct {
		Buyer  string `json:"buyer"`
		Maker  string `json:"maker"`
		Seller string `json:"seller"`
		Taker  string `json:"taker"`
	} `json:"commissionRates"`
	MakerCommission            int      `json:"makerCommission"`
	Permissions                []string `json:"permissions"`
	PreventSor                 bool     `json:"preventSor"`
	RequireSelfTradePrevention bool     `json:"requireSelfTradePrevention"`
	SellerCommission           int      `json:"sellerCommission"`
	TakerCommission            int      `json:"takerCommission"`
	Uid                        int      `json:"uid"`
	UpdateTime                 int64    `json:"updateTime"`
}

func NewGetAccount(client *binance.Client) Account {
	return &getAccountRequest{Client: client}
}

func (g *getAccountRequest) SetTimestamp(timestamp int64) Account {
	g.timestamp = timestamp
	return g
}

func (g *getAccountRequest) SetRecvWindow(recvWindow int64) Account {
	g.recvWindow = recvWindow
	return g
}

func (g *getAccountRequest) SetOmitZeroBalances(omitZeroBalances bool) Account {
	g.omitZeroBalances = omitZeroBalances
	return g
}

func (g *getAccountRequest) Call(ctx context.Context) (body *getAccountResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccount,
	}
	req.SetNeedSign(true)
	req.SetParam("omitZeroBalances", g.omitZeroBalances)
	if g.recvWindow > 0 {
		req.SetParam("recvWindow", g.recvWindow)
	}
	req.SetParam("timestamp", g.timestamp)
	resp, err := g.Do(ctx, req)
	if err != nil {
		g.Debugf("getAccountRequest response err:%v", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var e *errors.Error
		err = netutil.ParseHttpResponse(resp, &e)
		return nil, e
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		g.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}

// StreamAccountUpdateEvent 账户更新
type StreamAccountUpdateEvent struct {
	Stream string             `json:"stream"`
	Data   WsAccountDataEvent `json:"data"`
}
type WsAccountDataEvent struct {
	Event           enums.AccountDataEventType `json:"e"` // 事件类型
	Time            int64                      `json:"E"` // 事件时间
	UpdateTime      int64                      `json:"u"` // 账户末次更新时间戳
	TransactionTime int64                      `json:"T"`
	Balances        []Balance                  `json:"B"`
	*BalanceUpdateEvent
	*OrderUpdateEvent
	*OCOUpdateEvent
}

// Balance 账户更新
// 每当帐户余额发生更改时，都会发送一个事件outboundAccountPosition，其中包含可能由生成余额变动的事件而变动的资产。
type Balance struct {
	Asset  string `json:"a"`
	Free   string `json:"f"`
	Locked string `json:"l"`
}

// BalanceUpdateEvent 余额更新
// 当下列情形发生时更新:
//
// 账户发生充值或提取
// 交易账户之间发生划转(例如 现货向杠杆账户划转)
type BalanceUpdateEvent struct {
	Asset  string `json:"a"`
	Change string `json:"d"`
}

// OrderUpdateEvent 订单更新
// 订单通过executionReport事件进行更新。
// 备注: 通过将Z除以z可以找到平均价格。
type OrderUpdateEvent struct {
	Symbol                  string                `json:"s"`
	ClientOrderId           string                `json:"c"`
	Side                    enums.SideType        `json:"S"`
	Type                    enums.OrderType       `json:"o"`
	TimeInForce             enums.TimeInForceType `json:"f"`
	Volume                  string                `json:"q"`
	Price                   string                `json:"p"`
	StopPrice               string                `json:"P"`
	TrailingDelta           int64                 `json:"d"` // Trailing Delta
	IceBergVolume           string                `json:"F"`
	OrderListId             int64                 `json:"g"` // for OCO
	OrigCustomOrderId       string                `json:"C"` // customized order ID for the original order
	ExecutionType           string                `json:"x"` // execution type for this event NEW/TRADE...
	Status                  enums.OrderStatusType `json:"X"` // order status
	RejectReason            string                `json:"r"`
	Id                      int64                 `json:"i"` // order id
	LatestVolume            string                `json:"l"` // quantity for the latest trade
	FilledVolume            string                `json:"z"`
	LatestPrice             string                `json:"L"` // price for the latest trade
	FeeAsset                string                `json:"N"`
	FeeCost                 string                `json:"n"`
	TransactionTime         int64                 `json:"T"`
	TradeId                 int64                 `json:"t"`
	IsInOrderBook           bool                  `json:"w"` // is the order in the order book?
	IsMaker                 bool                  `json:"m"` // is this order maker?
	CreateTime              int64                 `json:"O"`
	FilledQuoteVolume       string                `json:"Z"` // the quote volume that already filled
	LatestQuoteVolume       string                `json:"Y"` // the quote volume for the latest trade
	QuoteVolume             string                `json:"Q"`
	TrailingTime            int64                 `json:"D"` // Trailing Time
	StrategyId              int64                 `json:"j"` // Strategy ID
	StrategyType            int64                 `json:"J"` // Strategy Type
	WorkingTime             int64                 `json:"W"` // Working Time
	SelfTradePreventionMode enums.StpModeType     `json:"V"`
}

type OCOUpdateEvent struct {
	Symbol          string                    `json:"s"`
	OrderListId     int64                     `json:"g"`
	ContingencyType enums.ContingencyType     `json:"c"`
	ListStatusType  enums.ListStatusType      `json:"l"`
	ListOrderStatus enums.ListOrderStatusType `json:"L"`
	RejectReason    string                    `json:"r"`
	ClientOrderId   string                    `json:"C"` // List Client Order ID
	Orders          []OCOOrder                `json:"O"`
}
type OCOOrder struct {
	Symbol        string `json:"s"`
	OrderId       int64  `json:"i"`
	ClientOrderId string `json:"c"`
}

// NewWsUserData 账户更新
// 每当帐户余额发生更改时，都会发送一个事件outboundAccountPosition，其中包含可能由生成余额变动的事件而变动的资产。
func NewWsUserData(c *binance.WsClient, listenKey string, handler binance.Handler[WsAccountDataEvent], exception binance.ErrorHandler) error {
	return userData(c, listenKey, handler, exception)
}

// NewStreamUserData 账户更新
// 每当帐户余额发生更改时，都会发送一个事件outboundAccountPosition，其中包含可能由生成余额变动的事件而变动的资产。
func NewStreamUserData(c *binance.WsClient, listenKey string, handler binance.Handler[StreamAccountUpdateEvent], exception binance.ErrorHandler) error {
	return userData(c, listenKey, handler, exception)
}

func userData[T WsAccountDataEvent | StreamAccountUpdateEvent](c *binance.WsClient, listenKey string, handler binance.Handler[T], exception binance.ErrorHandler) error {
	endpoint := c.Endpoint + listenKey
	return binance.WsHandler(c, endpoint, handler, exception)
}
