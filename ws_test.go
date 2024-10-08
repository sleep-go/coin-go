package coin_go

import (
	"fmt"
	"testing"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market/ticker"
)

var err error
var wsClient *binance.WsClient

func init() {
	wsClient = binance.NewWsClient(false, true, consts.WS_TEST_STREAM)
}

func TestDepthWs(t *testing.T) {
	// {"e":"depthUpdate","E":1728284450004,"s":"BTCUSDT","U":3999521,"u":3999550,"b":[["63698.69000000","0.00691000"],["63698.27000000","0.00409000"],["63698.26000000","0.00589000"],["63689.07000000","0.00000000"],["63688.87000000","0.00000000"],["63686.36000000","0.00000000"],["63686.06000000","0.00000000"]],"a":[["63699.99000000","0.00785000"],["63701.67000000","0.00000000"],["63701.79000000","0.00000000"],["63701.99000000","0.00479000"],["63702.17000000","0.00000000"],["63702.32000000","0.00000000"],["63702.82000000","0.00409000"],["63702.83000000","0.00000000"],["63703.48000000","0.00000000"],["63703.51000000","0.00000000"],["63705.30000000","0.00424000"],["63705.37000000","0.00542000"],["63706.07000000","0.00668000"],["63707.17000000","0.00699000"],["63708.00000000","0.00000000"],["63708.02000000","0.00605000"],["63709.39000000","0.00636000"],["63709.40000000","0.00691000"],["63709.94000000","0.00746000"],["63711.34000000","0.00424000"],["63731.16000000","0.00000000"],["63733.34000000","0.00000000"],["63733.67000000","0.00000000"]]}
	if wsClient.IsCombined {
		err := market.NewStreamDepth(wsClient, []string{BTCUSDT, "ETHUSDT"}, func(event *market.StreamDepthEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	} else {
		err := market.NewWsDepth(wsClient, []string{BTCUSDT, "ETHUSDT"}, func(event *market.WsDepthEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}
func TestWsDepthLevel(t *testing.T) {
	if wsClient.IsCombined {
		err := market.NewStreamDepthLevels(wsClient, map[string]enums.LimitType{
			BTCUSDT: enums.Limit5,
		}, func(event market.StreamDepthLevelsEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	} else {
		err := market.NewWsDepthLevels(wsClient, map[string]enums.LimitType{
			BTCUSDT: enums.Limit5,
		}, func(event market.WsDepthLevelsEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}
func TestWsAggTrade(t *testing.T) {
	//{"e":"aggTrade","E":1728319526169,"s":"BTCUSDT","a":549858,"p":"63778.14000000","q":"0.00024000","f":566460,"l":566460,"T":1728319526169,"m":false,"M":true}
	if wsClient.IsCombined {
		err = market.NewStreamAggTrade(wsClient, []string{BTCUSDT}, func(event market.StreamAggTradeEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})

	} else {
		err = market.NewWsAggTrade(wsClient, []string{BTCUSDT}, func(event market.WsAggTradeEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})

	}
	if err != nil {
		t.Fatal(err)
	}
}
func TestWsTrade(t *testing.T) {
	if wsClient.IsCombined {
		err = market.NewStreamTrade(wsClient, []string{BTCUSDT}, func(event market.StreamTradeEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {})
	} else {
		err = market.NewWsTrade(wsClient, []string{BTCUSDT}, func(event market.WsTradeEvent) {
			fmt.Printf("%+v\n", event)
		}, func(messageType int, err error) {})
	}
	if err != nil {
		t.Fatal(err)
	}
}
func TestWsKline(t *testing.T) {
	//设置带有时区偏移量的K线
	wsClient.Timezone = "+08:00"
	if wsClient.IsCombined {
		err = market.NewStreamKline(wsClient, map[string]enums.KlineIntervalType{
			BTCUSDT: enums.KlineIntervalType1d,
		}, func(event market.StreamKlineEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	} else {
		err = market.NewWsKline(wsClient, map[string]enums.KlineIntervalType{
			BTCUSDT: enums.KlineIntervalType1d,
		}, func(event market.WsKlineEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	}
	if err != nil {
		t.Fatal(err)
	}
}
func TestWsMiniTicker(t *testing.T) {
	if wsClient.IsCombined {
		err = ticker.NewStreamMiniTicker(wsClient, []string{BTCUSDT, ETHUSDT}, func(event ticker.StreamMiniTickerEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	} else {
		err = ticker.NewWsMiniTicker(wsClient, []string{BTCUSDT, ETHUSDT}, func(event ticker.WsMiniTickerEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	}
}
func TestWsAllMiniTicker(t *testing.T) {
	if wsClient.IsCombined {
		err = ticker.NewStreamAllMiniTicker(wsClient, func(event ticker.StreamAllMiniTickerEvent) {
			for _, e := range event.Data {
				if e.Symbol == BTCUSDT {
					fmt.Println(event.Stream, e)
				}
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	} else {
		err = ticker.NewWsAllMiniTicker(wsClient, func(event []ticker.WsMiniTickerEvent) {
			for _, e := range event {
				fmt.Println(e)
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	}
}
func TestAllTicker(t *testing.T) {
	if wsClient.IsCombined {
		err = ticker.NewStreamAllTicker(wsClient, func(event ticker.StreamAllTickerEvent) {
			for _, e := range event.Data {
				fmt.Println(e)
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	} else {
		err = ticker.NewWsTicker(wsClient, []string{BTCUSDT, ETHUSDT}, func(event ticker.WsTickerEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	}
	if err != nil {
		t.Fatal(err)
	}
}
func TestWsBookTicker(t *testing.T) {
	if wsClient.IsCombined {
		err = ticker.NewStreamBookTicker(wsClient, []string{BTCUSDT, ETHUSDT}, func(event ticker.StreamBookTickerEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	} else {
		err = ticker.NewWsBookTicker(wsClient, []string{BTCUSDT, ETHUSDT}, func(event ticker.WsBookTickerEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
	}
	if err != nil {
		t.Fatal(err)
	}
}
