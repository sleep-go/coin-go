package consts

// 上述列表的最后4个接口 (api1-api4) 会提供更好的性能，但其稳定性略为逊色。因此，请务必使用最适合的URL。
// 所有接口的响应都是 JSON 格式。
// 响应中如有数组，数组元素以时间升序排列，越早的数据越提前。
// 所有时间、时间戳均为UNIX时间，单位为毫秒。
// 对于仅发送公开市场数据的 API，您可以使用接口的 base URL https://data-api.binance.vision 。请参考 Market Data Only_CN 页面。
const (
	NAME          = "coin-go"
	REST_API      = "https://api.binance.com"
	REST_API_GCP  = "https://api-gcp.binance.com"
	REST_API1     = "https://api1.binance.com"
	REST_API2     = "https://api2.binance.com"
	REST_API3     = "https://api3.binance.com"
	REST_API4     = "https://api4.binance.com"
	REST_API_TEST = "https://testnet.binance.vision"

	WS_STREAM  = "wss://stream.binance.com:443"
	WS_STREAM2 = "wss://stream.binance.com:9443"
	// WS_STREAM_TEST 测试网 Stream base URL
	WS_STREAM_TEST = "wss://testnet.binance.vision"

	WS_API  = "wss://ws-api.binance.com:443/ws-api/v3"
	WS_API2 = "wss://ws-api.binance.com:9443/ws-api/v3"
	// WS_API_TEST 现货测试网的 base URL
	WS_API_TEST = "wss://testnet.binance.vision/ws-api/v3"
)

const (
	// REST_FAPI 期货 rest api
	REST_FAPI = "https://fapi.binance.com"
	// REST_FAPI_TEST 期货测试 rest api
	REST_FAPI_TEST = "https://testnet.binancefuture.com"
	// WS_FUTURE_TEST 期货 Websocket stream 行情推送
	WS_FSTREAM_TEST = "wss://fstream.binancefuture.com"

	// WS_FAPI 期货websocket api
	WS_FAPI = "wss://ws-fapi.binance.com/ws-fapi/v1"
	// WS_FAPI_TEST 期货测试websocket api
	WS_FAPI_TEST = "wss://testnet.binancefuture.com/ws-fapi/v1"
)
