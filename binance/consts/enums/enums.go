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

	// OrderListStatusType 订单组（order list）状态 （状态类型集 listStatusType）
	OrderListStatusType string
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

const (
	StpModeTypeNONE        StpModeType = "NONE"
	StpModeTypeExpireMaker StpModeType = "EXPIRE_MAKER"
	StpModeTypeExpireTaker StpModeType = "EXPIRE_TAKER"
	StpModeTypeExpireBoth  StpModeType = "EXPIRE_BOTH"
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
