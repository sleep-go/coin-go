package enums

type (
	ContractType       string //合约类型
	ContractStatusType string //合约状态
	StatusType         string //订单状态
	orderType          string //订单种类
	SideType           string //订单方向 (side)
	PositionSideType   string //持仓方向
	TimeInForceType    string //有效方式 (timeInForce)
	WorkingType        string //条件价格触发类型 (WorkingType)
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
	OrderTypeLimit              orderType = "LIMIT"                //限价单
	OrderTypeMarket             orderType = "MARKET"               //市价单
	OrderTypeStop               orderType = "STOP"                 //止损限价单
	OrderTypeStopMarket         orderType = "STOP_MARKET"          //止损市价单
	OrderTypeTakeProfit         orderType = "TAKE_PROFIT"          //止盈限价单
	OrderTypeTakeProfitMarket   orderType = "TAKE_PROFIT_MARKET"   //止盈市价单
	OrderTypeTrailingStopMarket orderType = "TRAILING_STOP_MARKET" //跟踪止损单
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
