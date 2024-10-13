package consts

const (
	// FApiMarketDepth 交易对深度信息
	FApiMarketDepth = "/fapi/v1/depth"

	// FApiMarketTrades 近期成交
	FApiMarketTrades = "/fapi/v1/trades"

	// FApiMarketHistoricalTrades 查询历史成交
	FApiMarketHistoricalTrades = "/fapi/v1/historicalTrades"

	// FApiMarketAggTrades 近期成交(归集)
	//与trades的区别是，同一个taker在同一时间同一价格与多个maker的成交会被合并为一条记录
	FApiMarketAggTrades = "/fapi/v1/aggTrades"

	// FApiMarketKLines K线数据
	FApiMarketKLines = "/fapi/v1/klines"

	// FApiMarketContinuousKlines 连续合约K线数据
	// 每根K线的开盘时间可视为唯一ID
	FApiMarketContinuousKlines = "/fapi/v1/continuousKlines"

	// FApiMarketIndexPriceKlines 价格指数K线数据
	//价格指数K线数据，每根K线的开盘时间可视为唯一ID
	FApiMarketIndexPriceKlines = "/fapi/v1/indexPriceKlines"

	// FApiMarketMarkPriceKlines 标记价格K线数据,每根K线的开盘时间可视为唯一ID
	FApiMarketMarkPriceKlines = "/fapi/v1/indexPriceKlines"

	// FApiMarketPremiumIndexKlines 溢价指数K线数据
	//合约溢价指数K线。每根K线的开盘时间可视为唯一ID。
	FApiMarketPremiumIndexKlines = "/fapi/v1/premiumIndexKlines"

	// FApiMarketPremiumIndex 最新标记价格和资金费率
	//采集各大交易所数据加权平均
	FApiMarketPremiumIndex = "/fapi/v1/premiumIndex"

	// FApiMarketFundingRate 查询资金费率历史
	FApiMarketFundingRate = "/fapi/v1/fundingRate"

	// FApiMarketFundingInfo 查询资金费率信息
	FApiMarketFundingInfo = "/fapi/v1/fundingInfo"

	// FApiMarketTicker24Hr 24hr价格变动情况
	FApiMarketTicker24Hr = "/fapi/v1/ticker/24hr"

	// FApiMarketTickerPrice 最新价格接口
	FApiMarketTickerPrice   = "/fapi/v1/ticker/price"
	FApiMarketTickerPriceV2 = "/fapi/v2/ticker/price"

	// FApiMarketTickerBookTicker 最优挂单接口
	FApiMarketTickerBookTicker = "/fapi/v1/ticker/bookTicker"

	// FApiMarketOpenInterest 获取未平仓合约数
	FApiMarketOpenInterest = "/fapi/v1/openInterest"
)

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
