package consts

import (
	"fmt"
	"net/http"
)

// 上述列表的最后4个接口 (api1-api4) 会提供更好的性能，但其稳定性略为逊色。因此，请务必使用最适合的URL。
// 所有接口的响应都是 JSON 格式。
// 响应中如有数组，数组元素以时间升序排列，越早的数据越提前。
// 所有时间、时间戳均为UNIX时间，单位为毫秒。
// 对于仅发送公开市场数据的 API，您可以使用接口的 base URL https://data-api.binance.vision 。请参考 Market Data Only_CN 页面。
const (
	NAME          = "exchange-go"
	REST_API      = "https://api.binance.com"
	REST_API_GCP  = "https://api-gcp.binance.com"
	REST_API1     = "https://api1.binance.com"
	REST_API2     = "https://api2.binance.com"
	REST_API3     = "https://api3.binance.com"
	REST_API4     = "https://api4.binance.com"
	REST_DATA_API = "https://data-api.binance.vision"

	TESTNET = "https://testnet.binance.vision"
)

type CommonService interface {
	Ping()
	Time()
}

type General struct {
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
}

func (g *General) Ping() {
	response, err := g.HTTPClient.Get(TESTNET + "/api/v3/ping")
	if err != nil {
		return
	}
	fmt.Println(response.Status)
}

func (g *General) Time() {
	//TODO implement me
	panic("implement me")
}

func Order() {
	//symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559
	api := fmt.Sprintf(REST_API + "/api/v3/depth?symbol=BTCUSDT&limit=5")
	fmt.Println(api)
}
