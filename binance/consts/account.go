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
)
