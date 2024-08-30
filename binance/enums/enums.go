package enums

type (
	// OrderTypes 订单类型（orderTypes, type）
	OrderTypes string
	// NewOrderRespType 订单返回类型 （newOrderRespType）:
	NewOrderRespType string
	// Side 订单方向 (side):
	Side string
	// TimeInForce 生效时间 （timeInForce）:
	TimeInForce string
)

func (f TimeInForce) String() string {
	return string(f)
}

const (
	LIMIT             OrderTypes = "LIMIT"             //限价单
	MARKET            OrderTypes = "MARKET"            // 市价单
	STOP_LOSS         OrderTypes = "STOP_LOSS"         //止损单
	STOP_LOSS_LIMIT   OrderTypes = "STOP_LOSS_LIMIT"   //限价止损单
	TAKE_PROFIT       OrderTypes = "TAKE_PROFIT"       //止盈单
	TAKE_PROFIT_LIMIT OrderTypes = "TAKE_PROFIT_LIMIT" //限价止盈单
	LIMIT_MAKER       OrderTypes = "LIMIT_MAKER"       //限价做市单
)

const (
	ACK    NewOrderRespType = "ACK"
	RESULT NewOrderRespType = "RESULT"
	FULL   NewOrderRespType = "FULL"
)

const (
	BUY  Side = "BUY"
	SELL Side = "SELL"
)

const (
	// GTC Good Til Canceled
	//除非订单被取消，否则订单将记录在册上。
	GTC TimeInForce = "GTC"
	// IOC mmediate Or Cancel
	//订单将尝试在订单到期前尽可能多地填写订单。
	IOC TimeInForce = "IOC"
	//FOK	Fill or Kill
	//如果执行时无法履行全部订单，则订单将过期。
	FOK TimeInForce = "FOK"
)

func (o OrderTypes) String() string {
	return string(o)
}
func (o Side) String() string {
	return string(o)
}
