package consts

const (
	// FApiOrder 下单 (TRADE)
	FApiOrder = "/fapi/v1/order"

	// FApiBatchOrders 批量下单(TRADE)
	FApiBatchOrders = "/fapi/v1/batchOrders"

	FApiTradingOrderTest = "/fapi/v1/order/test"
)

const (
	// ApiOrder
	//POST 下单 (TRADE)
	//GET 查询订单 (USER_DATA)
	//DELETE 撤销订单 (TRADE)
	ApiOrder = "/api/v3/order"

	// ApiTradingOrderTest 测试下单接口 (TRADE),用于测试订单请求，但不会提交到撮合引擎
	ApiTradingOrderTest = "/api/v3/order/test"

	// ApiOpenOrders
	//查看账户当前挂单 (USER_DATA)
	//撤销单一交易对的所有挂单 (TRADE)
	ApiOpenOrders = "/api/v3/openOrders"

	// ApiTradingCancelReplace 撤消挂单再下单 (TRADE)
	//撤消挂单并在同个交易对上重新下单。
	//在撤消订单和下单前会判断: 1) 过滤器参数, 以及 2) 目前下单数量。
	//即使请求中没有尝试发送新订单，比如(newOrderResult: NOT_ATTEMPTED)，下单的数量仍然会加1。
	ApiTradingCancelReplace = "/api/v3/order/cancelReplace"

	// ApiTradingAllOrders 查询所有订单（包括历史订单） (USER_DATA)
	ApiTradingAllOrders = "/api/v3/allOrders"

	// ApiOrderList
	//查询订单列表 (USER_DATA)
	//根据提供的可选参数检索特定的订单列表。
	ApiOrderList = "/api/v3/orderList"

	// ApiTradingOrderListOCO 发送新 one-cancels-the-other (OCO) 订单，激活其中一个订单会立即取消另一个订单。
	//OCO 有 2 legs，称为 上方 leg 和 下方 leg。
	//其中一条 leg 必须是 LIMIT_MAKER 订单，另一条 leg 必须是 STOP_LOSS 或 STOP_LOSS_LIMIT 订单。
	//针对价格限制：
	//如果 OCO 订单方向是 SELL：LIMIT_MAKER price > 最后交易价格 > stopPrice
	//如果 OCO 订单方向是 BUY：LIMIT_MAKER price < 最后交易价格 < stopPrice
	//OCO 将2 个订单添加到未成交订单计数，EXCHANGE_MAX_ORDERS 过滤器和 MAX_NUM_ORDERS 过滤器中。
	ApiTradingOrderListOCO = "/api/v3/orderList/oco"

	// ApiTradingOrderListOTO 发送一个新的 OTO 订单。
	//一个 OTO 订单（One-Triggers-the-Other）是一个包含了两个订单的订单列表.
	//第一个订单被称为生效订单，必须为 LIMIT 或 LIMIT_MAKER 类型的订单。最初，订单簿上只有生效订单。
	//第二个订单被称为待处理订单。它可以是任何订单类型，但不包括使用参数 quoteOrderQty 的 MARKET 订单。只有当生效订单完全成交时，待处理订单才会被自动下单。
	//如果生效订单或者待处理订单中的任意一个被单独取消，订单列表中剩余的那个订单也会被随之取消或过期。
	//如果生效订单在下订单列表后立即完全成交，则可能会得到订单响应。其中，生效订单的状态为 FILLED ，但待处理订单的状态为 PENDING_NEW。针对这类情况，如果需要检查当前状态，您可以查询相关的待处理订单。
	//OTO 订单将2 个订单添加到未成交订单计数，EXCHANGE_MAX_NUM_ORDERS 过滤器和 MAX_NUM_ORDERS 过滤器中。
	ApiTradingOrderListOTO = "/api/v3/orderList/oto"

	// ApiTradingOrderListOTOCO 发送一个新的 OTOCO 订单。
	//一个 OTOCO 订单（One-Triggers-One-Cancels-the-Other）是一个包含了三个订单的订单列表。
	//第一个订单被称为生效订单，必须为 LIMIT 或 LIMIT_MAKER 类型的订单。最初，订单簿上只有生效订单。
	//生效订单的行为与此一致 OTO
	//一个OTOCO订单有两个待处理订单（pending above 和 pending below），它们构成了一个 OCO 订单列表。只有当生效订单完全成交时，待处理订单们才会被自动下单。
	//待处理上方(pending above)订单和待处理下方(pending below)订单都遵循与 OCO 订单列表相同的规则 Order List OCO。
	//OTOCO 在未成交订单计数，EXCHANGE_MAX_NUM_ORDERS 过滤器和 MAX_NUM_ORDERS 过滤器的基础上添加3个订单。
	ApiTradingOrderListOTOCO = "/api/v3/orderList/otoco"

	// ApiTradingSorOrder 下 SOR 订单 (TRADE)
	ApiTradingSorOrder = "/api/v3/sor/order"

	// ApiTradingSorOrderTest 测试 SOR 下单接口 (TRADE)
	//用于测试使用智能订单路由 (SOR) 的订单请求，但不会提交到撮合引擎
	ApiTradingSorOrderTest = "/api/v3/sor/order/test"
)
