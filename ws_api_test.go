package coin_go

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market/ticker"
)

var wsApiClient *binance.Client
var done = make(chan struct{}, 1)

func init() {
	// 设置身份验证
	file, err := os.ReadFile("./.test.env")
	if err != nil {
		panic(err)
		return
	}
	API_KEY := strings.TrimSpace(string(file))
	wsApiClient = binance.NewWsApiClient(API_KEY, consts.WS_API_TEST)
	wsApiClient.Debug = true
}

func TestWsApiDepth(t *testing.T) {
	defer close(done)
	wsApiDepth := market.NewWsApiDepth(wsApiClient)
	go func() {
		err := wsApiDepth.Receive(func(event market.WsApiDepthResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				fmt.Println(event.Result)
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
			<-done
		})
		if err != nil {
			panic(err)
			return
		}
	}()
	time.Sleep(2 * time.Second)
	_ = wsApiDepth.SetSymbol(ETHUSDT).SetLimit(enums.Limit5).Send()
	time.Sleep(2 * time.Second)
	_ = wsApiDepth.SetSymbol(BTCUSDT).SetLimit(enums.Limit5).Send()
	time.Sleep(2 * time.Second)
}
func TestWsApiTrades(t *testing.T) {
	defer close(done)
	trades := market.NewWsApiTrades(wsApiClient)
	go func() {
		err := trades.Receive(func(event market.WsApiTradesResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				for _, res := range event.Result {
					fmt.Println(res)
				}
				for _, limit := range event.RateLimits {
					fmt.Println(limit)
				}
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
			<-done
		})
		if err != nil {
			return
		}
	}()
	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
		err := trades.SetSymbol(BTCUSDT).SetLimit(enums.Limit5).Send()
		time.Sleep(1 * time.Second)
		err = trades.SetSymbol(ETHUSDT).SetLimit(enums.Limit5).Send()
		if err != nil {
			return
		}
	}
}
func TestWsApiHistory(t *testing.T) {
	defer close(done)
	historyTrades := market.NewWsApiHistoryTrades(wsApiClient)
	go func() {
		err := historyTrades.Receive(func(event market.WsApiHistoryTradesResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				for _, res := range event.Result {
					fmt.Println(res)
				}
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
		}
	}()
	time.Sleep(2 * time.Second)
	err := historyTrades.
		SetSymbol(BTCUSDT).
		SetLimit(enums.Limit5).
		SetFromId(1).
		Send()
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)
}
func TestWsApiAggTrades(t *testing.T) {
	defer close(done)
	aggTrades := market.NewWsApiAggTrades(wsApiClient)
	go func() {
		err := aggTrades.Receive(func(event market.WsApiAggTradesResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				for _, res := range event.Result {
					fmt.Println(res)
				}
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
		}
	}()
	time.Sleep(2 * time.Second)
	err := aggTrades.
		SetSymbol(BTCUSDT).
		SetLimit(enums.Limit5).
		//SetFromId(1).
		SetStartTime(time.Now().UnixMilli() - 60*60*60).
		SetEndTime(time.Now().UnixMilli()).
		Send()
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)
}
func TestWsApiKline(t *testing.T) {
	defer close(done)
	klines := market.NewWsApiKlines(wsApiClient)
	go func() {
		err := klines.Receive(func(event market.WsApiKlinesResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				for _, res := range event.Result {
					fmt.Println(res)
				}
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
		}
	}()
	for {
		time.Sleep(2 * time.Second)
		err := klines.
			SetSymbol(BTCUSDT).
			SetLimit(enums.Limit5).
			SetInterval(enums.KlineIntervalType1d).
			SetTimeZone("+08:00").
			SetEndTime(time.Now().UnixMilli()).
			Send()
		if err != nil {
			continue
		}
	}
}
func TestWsApiAvgPrice(t *testing.T) {
	defer close(done)
	avgPrice := market.NewWsApiAvgPrice(wsApiClient)
	go func() {
		err := avgPrice.Receive(func(event market.WsApiAvgPriceResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				fmt.Println(event.Result, event.RateLimits)
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
		}
	}()
	for {
		time.Sleep(2 * time.Second)
		err := avgPrice.SetSymbol(BTCUSDT).Send()
		if err != nil {
			continue
		}
	}
}
func TestWsApiHr24(t *testing.T) {
	defer close(done)
	hr24 := ticker.NewWsApiHr24(wsApiClient)
	go func() {
		err := hr24.Receive(func(event ticker.WsApiHr24Response) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				fmt.Println(event.Result, event.RateLimits)
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			return
		}
	}()
	for {
		time.Sleep(2 * time.Second)
		err := hr24.SetSymbols([]string{BTCUSDT, ETHUSDT}).
			SetType(enums.TickerTypeFull).
			Send()
		if err != nil {
			continue
		}
	}
}
func TestWsApiTradingDay(t *testing.T) {
	defer close(done)
	tradingDay := ticker.NewWsApiTradingDay(wsApiClient)
	go func() {
		err := tradingDay.Receive(func(event ticker.WsApiTradingDayResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				for _, res := range event.Result {
					fmt.Println(res)
				}
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			return
		}
	}()
	for {
		time.Sleep(2 * time.Second)
		err := tradingDay.SetSymbols([]string{BTCUSDT, ETHUSDT}).
			SetTimeZone("+08:00").
			SetType(enums.TickerTypeFull).
			Send()
		if err != nil {
			continue
		}
	}
}
func TestWsApiTicker(t *testing.T) {
	defer close(done)
	tk := ticker.NewWsApiTicker(wsApiClient)
	go func() {
		err := tk.Receive(func(event ticker.WsApiTickerResponse) {
			if event.Error != nil {
				fmt.Println(event.Error)
			} else {
				for _, res := range event.Result {
					fmt.Println(res)
				}
			}
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			return
		}
	}()
	for {
		time.Sleep(2 * time.Second)
		err := tk.SetSymbols([]string{BTCUSDT, ETHUSDT}).
			SetDay(3).
			SetType(enums.TickerTypeFull).
			Send()
		if err != nil {
			continue
		}
	}
}
