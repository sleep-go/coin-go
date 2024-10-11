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
	"github.com/sleep-go/coin-go/binance/spot/endpoints/trading"
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
	PRIVATE_KEY_PATH := "./test-prv-key.pem"
	wsApiClient = binance.NewWsApiED25519Client(API_KEY, PRIVATE_KEY_PATH, consts.WS_API_TEST)
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
func TestWsApiTickerPrice(t *testing.T) {
	defer close(done)
	tk := ticker.NewWsApiTickerPrice(wsApiClient)
	go func() {
		err := tk.Receive(func(event ticker.WsApiTickerPriceResponse) {
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
		err := tk.SetSymbols([]string{BTCUSDT, ETHUSDT}).Send()
		if err != nil {
			continue
		}
	}
}
func TestWsApiBookTicker(t *testing.T) {
	defer close(done)
	tk := ticker.NewWsApiBookTicker(wsApiClient)
	go func() {
		err := tk.Receive(func(event ticker.WsApiBookTickerResponse) {
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
		err := tk.SetSymbols([]string{BTCUSDT, ETHUSDT}).Send()
		if err != nil {
			continue
		}
	}
}

func TestWsApiCreateOrder(t *testing.T) {
	defer close(done)
	tk := trading.NewWsApiCreateOrder(wsApiClient)
	go func() {
		err := tk.Receive(func(event trading.WsApiCreateOrderResponse) {
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
	//time.Sleep(2 * time.Second)
	//{"id":"08fbb7b7-741d-4d14-b1aa-d5f8800d369a","status":200,"result":{"symbol":"ETHUSDT","orderId":4241828,"orderListId":-1,"clientOrderId":"fvh0x6K2e2s0Gk4oqw7seI","transactTime":1728615836724,"price":"2000.00000000","origQty":"0.01000000","executedQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1728615836724,"fills":[],"selfTradePreventionMode":"EXPIRE_MAKER"},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":1},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":5},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":3}]}
	//err = tk.SetSymbol(ETHUSDT).
	//	SetSide(enums.SideTypeBuy).
	//	SetQuantity("0.01").
	//	SetTimeInForce(enums.TimeInForceTypeGTC).
	//	SetPrice("2000").
	//	SetType(enums.OrderTypeLimit).
	//	Send()
	time.Sleep(2 * time.Second)
	//{"id":"b54b5319-5787-4f3b-89a1-b338755ce24c","status":200,"result":{"symbol":"BTCUSDT","orderId":4336196,"orderListId":-1,"clientOrderId":"fDBM0cAaaaGpcffPQ0Wuud","transactTime":1728614800782,"price":"0.00000000","origQty":"0.01000000","executedQty":"0.01000000","cummulativeQuoteQty":"605.00045610","status":"FILLED","timeInForce":"GTC","type":"MARKET","side":"BUY","workingTime":1728614800782,"fills":[{"price":"60500.04000000","qty":"0.00439000","commission":"0.00000000","commissionAsset":"BTC","tradeId":1043813},{"price":"60500.05000000","qty":"0.00561000","commission":"0.00000000","commissionAsset":"BTC","tradeId":1043814}],"selfTradePreventionMode":"EXPIRE_MAKER"},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":1},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":2},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":3}]}

	err = tk.SetSymbol(BTCUSDT).
		SetSide(enums.SideTypeSell).
		SetType(enums.OrderTypeMarket).
		SetQuantity("0.01").Send()
	time.Sleep(2 * time.Second)
}
