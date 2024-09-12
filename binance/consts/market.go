package consts

const (
	// ApiMarketDepth 深度信息
	ApiMarketDepth = "/api/v3/depth"

	// ApiMarketTrades 近期成交
	ApiMarketTrades = "/api/v3/trades"

	// ApiMarketHistoricalTrades 查询历史成交
	ApiMarketHistoricalTrades = "/api/v3/historicalTrades"

	// ApiMarketAggTrades 近期成交(归集)
	//与trades的区别是，同一个taker在同一时间同一价格与多个maker的成交会被合并为一条记录
	ApiMarketAggTrades = "/api/v3/aggTrades"

	// ApiMarketKLines K线数据
	ApiMarketKLines = "/api/v3/klines"

	// ApiMarketUIKLines UIK线数据
	ApiMarketUIKLines = "/api/v3/uiKlines"

	// ApiMarketAvgPrice 当前平均价格
	ApiMarketAvgPrice = "/api/v3/avgPrice"

	// ApiMarketTicker24Hr 24hr价格变动情况
	ApiMarketTicker24Hr = "/api/v3/ticker/24hr"

	// ApiMarketTickerTradingDay 交易日行情(Ticker)
	//交易日价格变动统计。
	ApiMarketTickerTradingDay = "/api/v3/ticker/tradingDay"

	// ApiMarketTickerPrice 最新价格接口
	ApiMarketTickerPrice = "/api/v3/ticker/price"

	// ApiMarketTickerBookTicker 最优挂单接口
	ApiMarketTickerBookTicker = "/api/v3/ticker/bookTicker"

	// ApiMarketTicker 滚动窗口价格变动统计
	ApiMarketTicker = "/api/v3/ticker"
)
