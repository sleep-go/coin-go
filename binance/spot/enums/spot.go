package enums

type (
	// OrderType 订单类型（orderTypes, type）
	OrderType string

	// NewOrderRespType 订单返回类型 （newOrderRespType）:
	NewOrderRespType string

	// SideType 订单方向 (side):
	SideType string

	// TimeInForceType 生效时间 （timeInForce）:
	TimeInForceType string

	// TickerType 可接受值: FULL or MINI.  //默认值: FULL
	TickerType string

	// KlineIntervalType 支持的K线间隔 （区分大小写）
	KlineIntervalType string

	// StpModeType selfTradePreventionMode 自我交易模式
	StpModeType string

	// OrderStatusType 订单状态
	OrderStatusType string

	// ListStatusType 订单组（order list）状态 （状态类型集 ListStatusType）:
	ListStatusType string

	// ListOrderStatusType 订单组（order list）中的订单状态 （订单状态集 listOrderStatus）:
	ListOrderStatusType string

	CancelRestrictionsType string

	// ContingencyType 订单组的类型
	ContingencyType string

	// LimitType 是一个表示可选 limit 的类型
	LimitType int

	// OrderRateLimitExceededModeType “DO_NOTHING”（默认值）- 仅在账户未超过未成交订单频率限制时，会尝试取消订单。
	//
	//“CANCEL_ONLY” - 将始终取消订单。
	OrderRateLimitExceededModeType string

	// CancelReplaceModeType 指定类型：STOP_ON_FAILURE - 如果撤消订单失败将不会继续重新下单。
	//ALLOW_FAILURE - 不管撤消订单是否成功都会继续重新下单。
	CancelReplaceModeType string

	// AccountDataEventType 账户信息推送事件
	AccountDataEventType string
)

func (f TimeInForceType) String() string {
	return string(f)
}

const (
	OrderTypeLimit           OrderType = "LIMIT"             //限价单
	OrderTypeMarket          OrderType = "MARKET"            // 市价单
	OrderTypeStopLoss        OrderType = "STOP_LOSS"         //止损单
	OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"   //限价止损单
	OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"       //止盈单
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT" //限价止盈单
	OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"       //限价做市单
)

const (
	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeResult NewOrderRespType = "RESULT"
	NewOrderRespTypeFull   NewOrderRespType = "FULL"
)

const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"
)

const (
	// TimeInForceTypeGTC Good Til Canceled
	//除非订单被取消，否则订单将记录在册上。
	TimeInForceTypeGTC TimeInForceType = "GTC"
	// TimeInForceTypeIOC mmediate Or Cancel
	//订单将尝试在订单到期前尽可能多地填写订单。
	TimeInForceTypeIOC TimeInForceType = "IOC"
	//TimeInForceTypeFOK	Fill or Kill
	//如果执行时无法履行全部订单，则订单将过期。
	TimeInForceTypeFOK TimeInForceType = "FOK"
)

func (o OrderType) String() string {
	return string(o)
}
func (o SideType) String() string {
	return string(o)
}

const (
	TickerTypeFull TickerType = "FULL"
	TickerTypeMINI TickerType = "MINI"
)

func (o TickerType) String() string {
	return string(o)
}

const (
	//seconds -> 秒	1s

	KlineIntervalType1s KlineIntervalType = "1s"

	//分钟级别 minutes -> 分钟	1m， 3m， 5m， 15m， 30m

	KlineIntervalType1m  KlineIntervalType = "1m"
	KlineIntervalType3m  KlineIntervalType = "3m"
	KlineIntervalType5m  KlineIntervalType = "5m"
	KlineIntervalType15m KlineIntervalType = "15m"
	KlineIntervalType30m KlineIntervalType = "30m"

	//  小时级别 hours -> 小时	1h， 2h， 4h， 6h， 8h， 12h

	KlineIntervalType1h  KlineIntervalType = "1h"
	KlineIntervalType2h  KlineIntervalType = "2h"
	KlineIntervalType4h  KlineIntervalType = "4h"
	KlineIntervalType6h  KlineIntervalType = "6h"
	KlineIntervalType8h  KlineIntervalType = "8h"
	KlineIntervalType12h KlineIntervalType = "12h"

	// 天级别 days -> 天	1d， 3d

	KlineIntervalType1d KlineIntervalType = "1d"
	KlineIntervalType3d KlineIntervalType = "3d"

	//周级别 weeks -> 周	1w

	KlineIntervalType1w KlineIntervalType = "1w"

	// 月级别 months -> 月	1M

	KlineIntervalType1M KlineIntervalType = "1M"
)

// STP 的发生取决于 Taker 订单 的 STP 模式。
// 因此，订单薄上的订单的 STP 模式不再有效果，并且将在所有未来的订单处理中被忽略。
const (
	StpModeTypeNONE        StpModeType = "NONE"         //此模式使订单免于自我交易预防。
	StpModeTypeExpireMaker StpModeType = "EXPIRE_MAKER" //此模式通过立即使潜在挂单者(maker)的剩余数量过期来预防交易。
	StpModeTypeExpireTaker StpModeType = "EXPIRE_TAKER" // 此模式通过立即使吃单者(taker)的剩余数量过期来预防交易。
	StpModeTypeExpireBoth  StpModeType = "EXPIRE_BOTH"  //此模式通过立即同时使吃单和挂单者的剩余数量过期来预防交易。
)

const (
	OrderStatusTypeNew OrderStatusType = "NEW" //该订单被交易引擎接受。
	// OrderStatusTypePendingNew 该订单处于待处理 (PENDING) 阶段，直到其所属订单组（order list） 中的 working order 完全成交。
	OrderStatusTypePendingNew OrderStatusType = "PENDING_NEW"
	// OrderStatusTypePartiallyFilled 部分订单已被成交。
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	// OrderStatusTypeFilled 订单已完全成交。
	OrderStatusTypeFilled OrderStatusType = "FILLED"
	// OrderStatusTypeCanceled 用户撤销了订单。
	OrderStatusTypeCanceled OrderStatusType = "CANCELED"
	// OrderStatusTypePendingCancel 撤销中(目前并未使用)。
	OrderStatusTypePendingCancel OrderStatusType = "PENDING_CANCEL"
	// OrderStatusTypeRejected 订单没有被交易引擎接受，也没被处理。
	OrderStatusTypeRejected OrderStatusType = "REJECTED"
	// OrderStatusTypeExpiredCanceled 该订单根据订单类型的规则被取消（例如，没有成交的 LIMIT FOK 订单, LIMIT IOC 或部分成交的 MARKET 订单）
	//或者被交易引擎取消（例如，在强平期间被取消的订单，在交易所维护期间被取消的订单）
	OrderStatusTypeExpiredCanceled OrderStatusType = "EXPIRED"
	// OrderStatusTypeExpiredInMatch 表示订单由于 STP 而过期。（例如，带有 EXPIRE_TAKER 的订单与账簿上同属相同帐户或相同 tradeGroupId 的现有订单匹配）
	OrderStatusTypeExpiredInMatch OrderStatusType = "EXPIRED_IN_MATCH"
)
const (
	// ListStatusTypeResponse 在 ListStatus 用于响应失败的操作时会被使用。（例如，下订单组或取消订单组）
	ListStatusTypeResponse = "RESPONSE"
	// ListStatusTypeExecStarted 订单组已被下达或订单组状态有更新。
	ListStatusTypeExecStarted = "EXEC_STARTED"
	// ListStatusTypeAllDone 订单组执行结束，因此不再处于活动状态。
	ListStatusTypeAllDone = "ALL_DONE"
)

const (
	// ListOrderStatusTypeExecuting 订单组已被下达或订单组状态有更新。
	ListOrderStatusTypeExecuting ListOrderStatusType = "EXECUTING"
	// ListOrderStatusTypeAllDone 订单组执行结束，因此不再处于活动状态。
	ListOrderStatusTypeAllDone ListOrderStatusType = "ALL_DONE"
	// ListOrderStatusTypeReject 在 ListStatus 用于响应在下单阶段或取消订单组期间的失败操作时会被使用
	ListOrderStatusTypeReject ListOrderStatusType = "REJECT"
)

const (
	// CancelRestrictionsTypeOnlyNew 如果订单状态为 NEW，撤销将成功
	CancelRestrictionsTypeOnlyNew CancelRestrictionsType = "ONLY_NEW"
	// CancelRestrictionsTypeOnlyPartiallyFilled 如果订单状态为 PARTIALLY_FILLED，撤销将成功。
	CancelRestrictionsTypeOnlyPartiallyFilled CancelRestrictionsType = "ONLY_PARTIALLY_FILLED"
)

// 定义可选的 limit 值的枚举
// 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
const (
	Limit5    LimitType = 5
	Limit10   LimitType = 10
	Limit20   LimitType = 20
	Limit50   LimitType = 50
	Limit100  LimitType = 100
	Limit500  LimitType = 500
	Limit1000 LimitType = 1000
	Limit5000 LimitType = 5000
)

const (
	ContingencyTypeOCO ContingencyType = "OCO"
	ContingencyTypeOTO ContingencyType = "OTO"
)
const (
	// CancelReplaceModeTypeStopOnFailure 如果撤消订单失败将不会继续重新下单。
	CancelReplaceModeTypeStopOnFailure CancelReplaceModeType = "STOP_ON_FAILURE"
	// CancelReplaceModeTypeAllowFailure 不管撤消订单是否成功都会继续重新下单。
	CancelReplaceModeTypeAllowFailure CancelReplaceModeType = "ALLOW_FAILURE"
)

const (
	OrderRateLimitExceededModeTypeDoNothing  OrderRateLimitExceededModeType = "DO_NOTHING"
	OrderRateLimitExceededModeTypeCancelOnly OrderRateLimitExceededModeType = "CANCEL_ONLY"
)

const (
	// AccountDataEventTypeOutboundAccountPosition 会在账户余额发生变化时发送，其中包含可能因产生余额变化的事件而发生变化的资产。
	AccountDataEventTypeOutboundAccountPosition AccountDataEventType = "outboundAccountPosition"
	// AccountDataEventTypeBalanceUpdate 余额更新在以下情况下发生：
	// 1.账户存款或取款
	// 2.账户之间的资金转移（如现货转保证金）
	AccountDataEventTypeBalanceUpdate AccountDataEventType = "balanceUpdate"
	// AccountDataEventTypeExecutionReport 订单更新
	AccountDataEventTypeExecutionReport AccountDataEventType = "executionReport"
	// AccountDataEventTypeListStatus 如果是一个订单组，则除了显示executionReport事件外，还将显示一个名为ListStatus的事件。
	//
	//可能的执行类型:
	//
	//NEW - 新订单已被引擎接受。
	//CANCELED - 订单被用户取消。
	//REPLACED - (保留字段，当前未使用)
	//REJECTED - 新订单被拒绝 （这信息只会在撤消挂单再下单中发生，下新订单被拒绝但撤消挂单请求成功）。
	//TRADE - 订单有新成交。
	//EXPIRED - 订单已根据 Time In Force 参数的规则取消（e.g. 没有成交的 LIMIT FOK 订单或部分成交的 LIMIT IOC 订单）或者被交易所取消（e.g. 强平或维护期间取消的订单）。
	//TRADE_PREVENTION - 订单因 STP 触发而过期。
	AccountDataEventTypeListStatus AccountDataEventType = "listStatus"
	// AccountDataEventTypeListenKeyExpired Listen Key 已过期
	//当监听 listen key 过期时会发送此事件。此后不会再发送任何事件，直到创建新的 listenKey。
	//
	//正常关闭流时不会推送该事件。
	AccountDataEventTypeListenKeyExpired AccountDataEventType = "listenKeyExpired"
)
