package coin_go

import (
	"context"
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
	res, err := market.NewWsApiDepth(wsApiClient).
		SetSymbol(ETHUSDT).
		SetLimit(enums.Limit5).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiTrades(t *testing.T) {
	res, err := market.NewWsApiTrades(wsApiClient).
		SetSymbol(BTCUSDT).
		SetLimit(enums.Limit5).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiHistory(t *testing.T) {
	res, err := market.NewWsApiHistoryTrades(wsApiClient).
		SetSymbol(BTCUSDT).
		SetLimit(enums.Limit5).
		SetFromId(1).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiAggTrades(t *testing.T) {
	res, err := market.NewWsApiAggTrades(wsApiClient).
		SetSymbol(BTCUSDT).
		SetLimit(enums.Limit5).
		//SetFromId(1).
		SetStartTime(time.Now().UnixMilli() - 60*60*60).
		SetEndTime(time.Now().UnixMilli()).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiKline(t *testing.T) {
	res, err := market.NewWsApiKlines(wsApiClient).
		SetSymbol(BTCUSDT).
		SetLimit(enums.Limit5).
		SetInterval(enums.KlineIntervalType1d).
		SetTimeZone("+08:00").
		SetEndTime(time.Now().UnixMilli()).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiAvgPrice(t *testing.T) {
	res, err := market.NewWsApiAvgPrice(wsApiClient).
		SetSymbol(BTCUSDT).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiHr24(t *testing.T) {
	res, err := ticker.NewWsApiHr24(wsApiClient).
		SetSymbols([]string{BTCUSDT, ETHUSDT}).
		SetType(enums.TickerTypeFull).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		for _, v := range res.Result {
			fmt.Printf("%+v\n", v)
		}
	}
}
func TestWsApiTradingDay(t *testing.T) {
	res, err := ticker.NewWsApiTradingDay(wsApiClient).SetSymbols([]string{BTCUSDT, ETHUSDT}).
		SetTimeZone("+08:00").
		SetType(enums.TickerTypeFull).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiTicker(t *testing.T) {
	res, err := ticker.NewWsApiTicker(wsApiClient).SetSymbols([]string{BTCUSDT, ETHUSDT}).
		SetDay(3).
		SetType(enums.TickerTypeFull).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiTickerPrice(t *testing.T) {
	res, err := ticker.NewWsApiTickerPrice(wsApiClient).
		SetSymbols([]string{BTCUSDT, ETHUSDT}).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiBookTicker(t *testing.T) {
	res, err := ticker.NewWsApiBookTicker(wsApiClient).
		SetSymbols([]string{BTCUSDT, ETHUSDT}).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiCreateOrder(t *testing.T) {
	//{"id":"08fbb7b7-741d-4d14-b1aa-d5f8800d369a","status":200,"result":{"symbol":"ETHUSDT","orderId":4241828,"orderListId":-1,"clientOrderId":"fvh0x6K2e2s0Gk4oqw7seI","transactTime":1728615836724,"price":"2000.00000000","origQty":"0.01000000","executedQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1728615836724,"fills":[],"selfTradePreventionMode":"EXPIRE_MAKER"},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":1},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":5},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":3}]}
	res, err := trading.NewWsApiCreateOrder(wsApiClient).
		SetSymbol(ETHUSDT).
		SetSide(enums.SideTypeBuy).
		SetQuantity("0.01").
		SetTimeInForce(enums.TimeInForceTypeGTC).
		SetPrice("2000").
		SetType(enums.OrderTypeLimit).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiCreateOrderTest(t *testing.T) {
	//{"id":"08fbb7b7-741d-4d14-b1aa-d5f8800d369a","status":200,"result":{"symbol":"ETHUSDT","orderId":4241828,"orderListId":-1,"clientOrderId":"fvh0x6K2e2s0Gk4oqw7seI","transactTime":1728615836724,"price":"2000.00000000","origQty":"0.01000000","executedQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1728615836724,"fills":[],"selfTradePreventionMode":"EXPIRE_MAKER"},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":1},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":5},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":3}]}
	res, err := trading.NewWsApiCreateOrder(wsApiClient).SetSymbol(ETHUSDT).
		SetSide(enums.SideTypeBuy).
		SetQuantity("0.01").
		SetTimeInForce(enums.TimeInForceTypeGTC).
		SetPrice("2000").
		SetType(enums.OrderTypeLimit).
		SendTest(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiQueryOrder(t *testing.T) {
	res, err := trading.NewWsApiQueryOrder(wsApiClient).
		SetOrderId(736954).
		SetSymbol(BTCUSDT).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiOpenOrders(t *testing.T) {
	res, err := trading.NewWsApiQueryOrder(wsApiClient).
		SetOrderId(736954).
		SetSymbol(BTCUSDT).SendOpenOrders(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiDeleteOrder(t *testing.T) {
	res, err := trading.NewWsApiDeleteOrder(wsApiClient).
		SetOrderId(736954).SetSymbol(BTCUSDT).
		SetCancelRestrictions(enums.CancelRestrictionsTypeOnlyNew).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
func TestWsApiCancelReplace(t *testing.T) {
	res, err := trading.NewWsApiCancelReplace(wsApiClient).
		SetSymbol(BTCUSDT).
		SetCancelReplaceMode(enums.CancelReplaceModeTypeAllowFailure).
		SetSide(enums.SideTypeBuy).
		SetType(enums.OrderTypeLimit).
		SetTimeInForce(enums.TimeInForceTypeGTC).
		SetCancelOrderId(736954).
		SetQuantity("0.01").
		SetPrice("23416").Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Result)
	}
}
