package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/spot/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
	"github.com/tidwall/gjson"
)

type Account interface {
	Call(ctx context.Context) (body *getAccountResponse, err error)
	SetOmitZeroBalances(omitZeroBalances bool) *getAccountRequest
}

type getAccountRequest struct {
	*binance.Client
	omitZeroBalances bool //如果true，将隐藏所有零余额。默认值：false
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

func (g *getAccountRequest) SetOmitZeroBalances(omitZeroBalances bool) *getAccountRequest {
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
	resp, err := g.Do(ctx, req)
	if err != nil {
		g.Debugf("getAccountRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*getAccountResponse](resp)
}

// ****************************** Websocket Stream *******************************

// WsOutboundAccountPositionEvent 账户更新
// 每当帐户余额发生更改时，都会发送一个事件outboundAccountPosition，其中包含可能由生成余额变动的事件而变动的资产。
type WsOutboundAccountPositionEvent struct {
	Event      enums.AccountDataEventType `json:"e"` // 事件类型
	Time       int64                      `json:"E"` // 事件时间
	UpdateTime int64                      `json:"u"` // 账户末次更新时间戳
	Balances   []struct {
		Asset  string `json:"a"`
		Free   string `json:"f"`
		Locked string `json:"l"`
	} `json:"B"`
}

// WsBalanceUpdateEvent 余额更新
// 当下列情形发生时更新:
//
// 账户发生充值或提取
// 交易账户之间发生划转(例如 现货向杠杆账户划转)
type WsBalanceUpdateEvent struct {
	Event           enums.AccountDataEventType `json:"e"` // 事件类型
	Time            int64                      `json:"E"` // 事件时间
	Asset           string                     `json:"a"`
	Change          string                     `json:"d"`
	TransactionTime int64                      `json:"T"`
}

// WsExecutionReportEvent 订单更新
// 订单通过executionReport事件进行更新。
// 备注: 通过将Z除以z可以找到平均价格。
type WsExecutionReportEvent struct {
	Event                   enums.AccountDataEventType `json:"e"` // 事件类型
	Time                    int64                      `json:"E"` // 事件时间
	Symbol                  string                     `json:"s"`
	ClientOrderId           string                     `json:"c"`
	Side                    enums.SideType             `json:"S"`
	Type                    enums.OrderType            `json:"o"`
	TimeInForce             enums.TimeInForceType      `json:"f"`
	Volume                  string                     `json:"q"`
	Price                   string                     `json:"p"`
	StopPrice               string                     `json:"P"`
	TrailingDelta           int64                      `json:"d"` // Trailing Delta
	IceBergVolume           string                     `json:"F"`
	OrderListId             int64                      `json:"g"` // for OCO
	OrigCustomOrderId       string                     `json:"C"` // customized order ID for the original order
	ExecutionType           string                     `json:"x"` // execution type for this event NEW/TRADE...
	Status                  enums.OrderStatusType      `json:"X"` // order status
	RejectReason            string                     `json:"r"`
	Id                      int64                      `json:"i"` // order id
	LatestVolume            string                     `json:"l"` // quantity for the latest trade
	FilledVolume            string                     `json:"z"`
	LatestPrice             string                     `json:"L"` // price for the latest trade
	FeeAsset                string                     `json:"N"`
	FeeCost                 string                     `json:"n"`
	TransactionTime         int64                      `json:"T"`
	TradeId                 int64                      `json:"t"`
	IsInOrderBook           bool                       `json:"w"` // is the order in the order book?
	IsMaker                 bool                       `json:"m"` // is this order maker?
	CreateTime              int64                      `json:"O"`
	FilledQuoteVolume       string                     `json:"Z"` // the quote volume that already filled
	LatestQuoteVolume       string                     `json:"Y"` // the quote volume for the latest trade
	QuoteVolume             string                     `json:"Q"`
	TrailingTime            int64                      `json:"D"` // Trailing Time
	StrategyId              int64                      `json:"j"` // Strategy ID
	StrategyType            int64                      `json:"J"` // Strategy Type
	WorkingTime             int64                      `json:"W"` // Working Time
	SelfTradePreventionMode enums.StpModeType          `json:"V"`
}

type WsListStatusEvent struct {
	Event           enums.AccountDataEventType `json:"e"` // 事件类型
	Time            int64                      `json:"E"` // 事件时间
	Symbol          string                     `json:"s"`
	OrderListId     int64                      `json:"g"`
	ContingencyType enums.ContingencyType      `json:"c"`
	ListStatusType  enums.ListStatusType       `json:"l"`
	ListOrderStatus enums.ListOrderStatusType  `json:"L"`
	RejectReason    string                     `json:"r"`
	ClientOrderId   string                     `json:"C"` // List Client Order ID
	Orders          []struct {
		Symbol        string `json:"s"`
		OrderId       int64  `json:"i"`
		ClientOrderId string `json:"c"`
	} `json:"O"`
}

type WsListenKeyExpiredEvent struct {
	ListenKey string `json:"listenKey"`
}

// NewWsUserData 账户更新
// 每当帐户余额发生更改时，都会发送一个事件outboundAccountPosition，其中包含可能由生成余额变动的事件而变动的资产。
func NewWsUserData(
	c *binance.Client,
	listenKey string,
	oap binance.Handler[*WsOutboundAccountPositionEvent],
	bu binance.Handler[*WsBalanceUpdateEvent],
	er binance.Handler[*WsExecutionReportEvent],
	ls binance.Handler[*WsListStatusEvent],
	lke binance.Handler[*WsListenKeyExpiredEvent],
	exception binance.ErrorHandler,
) error {
	h := func(mt int, msg []byte) {
		e := gjson.Get(string(msg), "e").String()
		switch enums.AccountDataEventType(e) {
		case enums.AccountDataEventTypeOutboundAccountPosition:
			event := new(WsOutboundAccountPositionEvent)
			err := json.Unmarshal(msg, &event)
			if err != nil {
				return
			}
			oap(event)
		case enums.AccountDataEventTypeBalanceUpdate:
			event := new(WsBalanceUpdateEvent)
			err := json.Unmarshal(msg, &event)
			if err != nil {
				return
			}
			bu(event)
		case enums.AccountDataEventTypeExecutionReport:
			//{"stream":"re1kcvyiLnbcX8D7xqHK4dKfdWlSzrvLYHvpYCdP9bKH6JPlJkSc36mp8ezY","data":{"e":"listStatus","E":1728660487437,"s":"ETHUSDT","g":2617,"c":"OTO","l":"EXEC_STARTED","L":"EXECUTING","r":"NONE","C":"7780ba25-a448-4d97-ac27-156bab1bea54","T":1728660487437,"O":[{"s":"ETHUSDT","i":4580466,"c":"MQucRQKc3SWeKPFVoP45Me"},{"s":"ETHUSDT","i":4580467,"c":"CefgUNxEhQq2RPhyti21Oi"}]}}
			event := new(WsExecutionReportEvent)
			err := json.Unmarshal(msg, &event)
			if err != nil {
				return
			}
			er(event)
		case enums.AccountDataEventTypeListStatus:
			event := new(WsListStatusEvent)
			err := json.Unmarshal(msg, &event)
			if err != nil {
				exception(mt, err)
				return
			}
			ls(event)
		case enums.AccountDataEventTypeListenKeyExpired:
			event := new(WsListenKeyExpiredEvent)
			err := json.Unmarshal(msg, &event)
			if err != nil {
				exception(mt, err)
				return
			}
			lke(event)
		}
	}
	endpoint := c.BaseURL + listenKey
	return c.Serve(endpoint, h, exception)
}

// NewStreamUserData 账户更新
// 每当帐户余额发生更改时，都会发送一个事件outboundAccountPosition，其中包含可能由生成余额变动的事件而变动的资产。
func NewStreamUserData(
	c *binance.Client,
	listenKey string,
	oap binance.Handler[*WsOutboundAccountPositionEvent],
	bu binance.Handler[*WsBalanceUpdateEvent],
	er binance.Handler[*WsExecutionReportEvent],
	ls binance.Handler[*WsListStatusEvent],
	lke binance.Handler[*WsListenKeyExpiredEvent],
	exception binance.ErrorHandler,
) error {
	h := func(mt int, msg []byte) {
		e := gjson.Get(string(msg), "data.e").String()
		data := gjson.Get(string(msg), "data").String()
		switch enums.AccountDataEventType(e) {
		case enums.AccountDataEventTypeOutboundAccountPosition:
			event := new(WsOutboundAccountPositionEvent)
			err := json.Unmarshal([]byte(data), &event)
			if err != nil {
				return
			}
			oap(event)
		case enums.AccountDataEventTypeBalanceUpdate:
			event := new(WsBalanceUpdateEvent)
			err := json.Unmarshal([]byte(data), &event)
			if err != nil {
				return
			}
			bu(event)
		case enums.AccountDataEventTypeExecutionReport:
			//{"stream":"re1kcvyiLnbcX8D7xqHK4dKfdWlSzrvLYHvpYCdP9bKH6JPlJkSc36mp8ezY","data":{"e":"listStatus","E":1728660487437,"s":"ETHUSDT","g":2617,"c":"OTO","l":"EXEC_STARTED","L":"EXECUTING","r":"NONE","C":"7780ba25-a448-4d97-ac27-156bab1bea54","T":1728660487437,"O":[{"s":"ETHUSDT","i":4580466,"c":"MQucRQKc3SWeKPFVoP45Me"},{"s":"ETHUSDT","i":4580467,"c":"CefgUNxEhQq2RPhyti21Oi"}]}}
			event := new(WsExecutionReportEvent)
			err := json.Unmarshal([]byte(data), &event)
			if err != nil {
				return
			}
			er(event)
		case enums.AccountDataEventTypeListStatus:
			event := new(WsListStatusEvent)
			err := json.Unmarshal([]byte(data), &event)
			if err != nil {
				exception(mt, err)
				return
			}
			ls(event)
		case enums.AccountDataEventTypeListenKeyExpired:
			event := new(WsListenKeyExpiredEvent)
			err := json.Unmarshal([]byte(data), &event)
			if err != nil {
				exception(mt, err)
				return
			}
			lke(event)
		}
	}
	endpoint := c.BaseURL + listenKey
	return c.Serve(endpoint, h, exception)
}

// ****************************** Websocket Api *******************************

type WsApiAccount interface {
	binance.WsApi[*WsApiAccountResponse]
	Account
}
type WsApiAccountResponse struct {
	binance.WsApiResponse
	Result *getAccountResponse `json:"result"`
}

func NewWsApiAccount(c *binance.Client) WsApiAccount {
	return &getAccountRequest{Client: c}
}

// Send 下 SOR 订单 (TRADE)
// 下使用智能订单路由 (SOR) 的新订单。
// 注意: sor.order.place 只支持 限价 和 市场 单， 并不支持 quoteOrderQty。
func (g *getAccountRequest) Send(ctx context.Context) (*WsApiAccountResponse, error) {
	req := &binance.Request{Path: "account.status"}
	req.SetNeedSign(true)
	req.SetParam("omitZeroBalances", g.omitZeroBalances)
	return binance.WsApiHandler[*WsApiAccountResponse](ctx, g.Client, req)
}
