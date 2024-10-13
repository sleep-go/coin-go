package enums

type (
	ContractType       string //合约类型
	ContractStatusType string //合约状态
	StatusType         string //订单状态
	OrderType          string //订单种类
	SideType           string //订单方向 (side)
	PositionSideType   string //持仓方向
	TimeInForceType    string //有效方式 (timeInForce)
	WorkingType        string //条件价格触发类型 (WorkingType)
	NewOrderRespType   string //响应类型 (NewOrderRespType)
	KlineIntervalType  string //K线间隔
	StpModeType        string //防止自成交模式
	PriceMatchType     string //盘口价下单模式
	RateLimitType      string //限制种类 (RateLimitType)
	LimitType          int
)

// 合约类型 (contractType):
const (
	ContractTypePerpetual           ContractType = "PERPETUAL"            //永续合约
	ContractTypeCurrentMonth        ContractType = "CURRENT_MONTH"        //当月交割合约
	ContractTypeNextMonth           ContractType = "NEXT_MONTH"           //次月交割合约
	ContractTypeCurrentQuarter      ContractType = "CURRENT_QUARTER"      //当季交割合约
	ContractTypeNextQuarter         ContractType = "NEXT_QUARTER"         //次季交割合约
	ContractTypePerpetualDelivering ContractType = "PERPETUAL_DELIVERING" //交割结算中合约
)

// 合约状态 (contractStatus, status):
const (
	ContractStatusTypePendingTrading ContractStatusType = "PENDING_TRADING" //待上市
	ContractStatusTypeTrading        ContractStatusType = "TRADING"         // 交易中
	ContractStatusTypePreDelivering  ContractStatusType = "PRE_DELIVERING"  //预交割
	ContractStatusTypeDelivering     ContractStatusType = "DELIVERING"      //交割中
	ContractStatusTypeDelivered      ContractStatusType = "DELIVERED"       //已交割
	ContractStatusTypePreSettle      ContractStatusType = "PRE_SETTLE"      //预结算
	ContractStatusTypeSettings       ContractStatusType = "SETTLING"        // 结算中
	ContractStatusTypeClose          ContractStatusType = "CLOSE"           //  已下架
)

// 订单状态 (status):
const (
	StatusTypeNew             StatusType = "NEW"              //新建订单
	StatusTypePartiallyFilled StatusType = "PARTIALLY_FILLED" //部分成交
	StatusTypeFilled          StatusType = "FILLED"           //全部成交
	StatusTypeCancelled       StatusType = "CANCELLED"        //已撤销
	StatusTypeRejected        StatusType = "REJECTED"         //订单被拒绝
	StatusTypeExpired         StatusType = "EXPIRED"          //订单过期(根据timeInForce参数规则)
	StatusTypeExpiredInMatch  StatusType = "EXPIRED_IN_MATCH" //订单被STP过期
)

// 订单种类 (orderTypes, type):
const (
	OrderTypeLimit              OrderType = "LIMIT"                //限价单
	OrderTypeMarket             OrderType = "MARKET"               //市价单
	OrderTypeStop               OrderType = "STOP"                 //止损限价单
	OrderTypeStopMarket         OrderType = "STOP_MARKET"          //止损市价单
	OrderTypeTakeProfit         OrderType = "TAKE_PROFIT"          //止盈限价单
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"   //止盈市价单
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET" //跟踪止损单
)

// 订单方向 (side):
const (
	SideTypeBuy  SideType = "BUY"  //买入
	SideTypeSell SideType = "SELL" //卖出
)

// 持仓方向
const (
	PositionSideTypeBoth  = "BOTH"  //单一持仓方向
	PositionSideTypeLong  = "LONG"  //多头(双向持仓下)
	PositionSideTypeShort = "SHORT" //空头(双向持仓下)
)

// 有效方式 (timeInForce):
const (
	timeInForceTypeGTC TimeInForceType = "GTC" //Good Till Cancel 成交为止（下单后仅有1年有效期，1年后自动取消）
	TimeInForceTypeIOC TimeInForceType = "IOC" //Immediate or Cancel 无法立即成交(吃单)的部分就撤销
	TimeInForceTypeFOK TimeInForceType = "FOK" //Fill or Kill 无法全部立即成交就撤销
	TimeInForceTypeGTX TimeInForceType = "GTX" //GTX - Good Till Crossing 无法成为挂单方就撤销
	TimeInForceTypeGTD TimeInForceType = "GTD" //GTD - Good Till Date 在特定时间之前有效，到期自动撤销
)

// 响应类型 (newOrderRespType)
const (
	NewOrderRespTypeAck    NewOrderRespType = "ACK"
	NewOrderRespTypeResult NewOrderRespType = "RESULT"
)

// K线间隔:
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

// STP MODE 防止自成交模式
const (
	StpModeTypeNONE        StpModeType = "NONE"         //此模式使订单免于自我交易预防。
	StpModeTypeExpireMaker StpModeType = "EXPIRE_MAKER" //此模式通过立即使潜在挂单者(maker)的剩余数量过期来预防交易。
	StpModeTypeExpireTaker StpModeType = "EXPIRE_TAKER" // 此模式通过立即使吃单者(taker)的剩余数量过期来预防交易。
	StpModeTypeExpireBoth  StpModeType = "EXPIRE_BOTH"  //此模式通过立即同时使吃单和挂单者的剩余数量过期来预防交易。
)

// 盘口价下单模式
const (
	PriceMatchTypeOpponent   = "OPPONENT"    //盘口对手价
	PriceMatchTypeOpponent5  = "OPPONENT_5"  //盘口对手5档价
	PriceMatchTypeOpponent10 = "OPPONENT_10" //盘口对手10档价
	PriceMatchTypeOpponent20 = "OPPONENT_20" //盘口对手20档价
	PriceMatchTypeQueue      = "QUEUE"       //盘口同向价
	PriceMatchTypeQueue5     = "QUEUE_5"     //盘口同向排队5档价
	PriceMatchTypeQueue10    = "QUEUE_10"    //盘口同向排队10档价
	PriceMatchTypeQueue20    = "QUEUE_20"    //盘口同向排队20档价
)
const (
	RateLimitTypeRequestWeight RateLimitType = "REQUEST_WEIGHT" //单位时间请求权重之和上限
	RateLimitTypeOrders        RateLimitType = "ORDERS"         //单位时间下单(撤单)次数上限
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
