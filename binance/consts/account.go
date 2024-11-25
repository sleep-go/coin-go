package consts

const (
	// ApiAccount 账户信息 (USER_DATA)
	ApiAccount = "/api/v3/account"

	// ApiAccountMyTrades 账户成交历史 (USER_DATA)
	// 获取某交易对的成交历史
	ApiAccountMyTrades = "/api/v3/myTrades"

	// ApiAccountRateLimitOrder 查询未成交的订单计数 (USER_DATA)
	//显示用户在所有时间间隔内的未成交订单计数。
	ApiAccountRateLimitOrder = "/api/v3/rateLimit/order"

	// ApiAccountMyPreventedMatches 获取因 STP 而过期的订单列表。
	ApiAccountMyPreventedMatches = "/api/v3/myPreventedMatches"

	// ApiAccountMyAllocations 查询分配结果 (USER_DATA)
	//检索由 SOR 订单生成引起的分配结果。
	ApiAccountMyAllocations = "/api/v3/myAllocations"

	// ApiAccountCommission 查询佣金费率 (USER_DATA)
	// 获取当前账户的佣金费率。
	ApiAccountCommission = "/api/v3/account/commission"

	// ApiTradingOpenOrderList 查询订单列表挂单 (USER_DATA)
	ApiTradingOpenOrderList = "/api/v3/openOrderList"

	// ApiAccountAllOrderList 查询所有订单列表 (USER_DATA)
	//根据提供的可选参数检索所有的订单列表。
	//请注意，startTime和endTime之间的时间不能超过 24 小时。
	ApiAccountAllOrderList = "/api/v3/allOrderList"
)

const (
	FApiAccountOrderAmendment = "/fapi/v1/orderAmendment"
)
