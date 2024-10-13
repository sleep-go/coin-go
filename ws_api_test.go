package coin_go

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/spot/account"
	"github.com/sleep-go/coin-go/binance/spot/enums"
	"github.com/sleep-go/coin-go/binance/spot/market"
	"github.com/sleep-go/coin-go/binance/spot/market/ticker"
	"github.com/sleep-go/coin-go/binance/spot/stream"
	"github.com/sleep-go/coin-go/binance/spot/trading"
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
	res, err := ticker.NewWsApiTickerPrice(wsApiClient).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Println(len(res.Result), res.Result)
		for _, v := range res.Result {
			fmt.Println(v.Symbol, v.Price)
		}
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
func TestWsApiDeleteOpenOrders(t *testing.T) {
	res, err := trading.NewWsApiDeleteOpenOrders(wsApiClient).SetSymbol(BTCUSDT).Send(context.Background())
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
func TestWsApiOCO(t *testing.T) {
	//OrderListId:2604
	//&{OrderListId:2605 ContingencyType:OCO ListStatusType:EXEC_STARTED ListOrderStatus:EXECUTING ListClientOrderId:8d931268-307c-4bb4-97fb-0a978f974f14 TransactionTime:1728658223336 Symbol:ETHUSDT Orders:[{Symbol:ETHUSDT OrderId:4564037 ClientOrderId:qta1Pd9f1C1SHMAJj3zhyR} {Symbol:ETHUSDT OrderId:4564038 ClientOrderId:OTffGJW6HD7Qvweq1ujCth}] OrderReports:[{Symbol:ETHUSDT OrderId:4564037 OrderListId:2605 ClientOrderId:qta1Pd9f1C1SHMAJj3zhyR TransactTime:1728658223336 Price:2428.77000000 OrigQty:0.01000000 ExecutedQty:0.00000000 CummulativeQuoteQty:0.00000000 Status:NEW TimeInForce:GTC Type:STOP_LOSS_LIMIT Side:BUY StopPrice:2428.87000000 WorkingTime:-1 IcebergQty: SelfTradePreventionMode:EXPIRE_MAKER} {Symbol:ETHUSDT OrderId:4564038 OrderListId:2605 ClientOrderId:OTffGJW6HD7Qvweq1ujCth TransactTime:1728658223336 Price:1500.00000000 OrigQty:0.01000000 ExecutedQty:0.00000000 CummulativeQuoteQty:0.00000000 Status:NEW TimeInForce:GTC Type:LIMIT_MAKER Side:BUY StopPrice: WorkingTime:1728658223336 IcebergQty: SelfTradePreventionMode:EXPIRE_MAKER}]}
	res, err := trading.NewWsApiOCO(wsApiClient).
		SetSymbol(ETHUSDT).
		SetSide(enums.SideTypeBuy).
		SetListClientOrderId(uuid.New().String()).
		SetAboveType(enums.OrderTypeStopLossLimit).
		SetAbovePrice("2428.77000000").
		SetAboveStopPrice("2428.87000000").
		SetAboveTimeInForce(enums.TimeInForceTypeGTC).
		SetBelowType(enums.OrderTypeLimitMaker).
		SetBelowPrice("1500").
		SetQuantity("0.01").
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
	}
}
func TestWsApiOTO(t *testing.T) {
	//&{OrderListId:2609 ContingencyType:OTO ListStatusType:EXEC_STARTED ListOrderStatus:EXECUTING ListClientOrderId:267f613e-116a-44ed-b353-e87e4f2fe412 TransactionTime:1728659162507 Symbol:ETHUSDT Orders:[{Symbol:ETHUSDT OrderId:4571011 ClientOrderId:hto0WXQ2Msjr3PwDgSX224} {Symbol:ETHUSDT OrderId:4571012 ClientOrderId:fV1kDH2dTlMVxM9FL4qiJI}] OrderReports:[{Symbol:ETHUSDT OrderId:4571011 OrderListId:2609 ClientOrderId:hto0WXQ2Msjr3PwDgSX224 TransactTime:1728659162507 Price:2428.77000000 OrigQty:0.01000000 ExecutedQty:0.01000000 CummulativeQuoteQty:24.36850000 Status:FILLED TimeInForce:GTC Type:LIMIT Side:SELL StopPrice: WorkingTime:1728659162507 IcebergQty: SelfTradePreventionMode:EXPIRE_MAKER} {Symbol:ETHUSDT OrderId:4571012 OrderListId:2609 ClientOrderId:fV1kDH2dTlMVxM9FL4qiJI TransactTime:1728659162507 Price:0.00000000 OrigQty:0.01000000 ExecutedQty:0.00000000 CummulativeQuoteQty:0.00000000 Status:PENDING_NEW TimeInForce:GTC Type:MARKET Side:BUY StopPrice: WorkingTime:-1 IcebergQty: SelfTradePreventionMode:EXPIRE_MAKER}]}
	res, err := trading.NewWsApiOTO(wsApiClient).
		SetSymbol(ETHUSDT).
		SetPendingSide(enums.SideTypeBuy).
		SetPendingQuantity("0.01").
		SetListClientOrderId(uuid.New().String()).
		SetPendingType(enums.OrderTypeMarket).
		SetWorkingPrice("2428.77000000").
		SetWorkingQuantity("0.01").
		SetWorkingSide(enums.SideTypeSell).
		SetWorkingType(enums.OrderTypeLimit).
		SetWorkingTimeInForce(enums.TimeInForceTypeGTC).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
	}
}
func TestWsApiOTOCO(t *testing.T) {
	res, err := trading.NewWsApiOTOCO(wsApiClient).
		SetSymbol("LTCBNB").
		SetPendingSide(enums.SideTypeSell).
		SetPendingQuantity("5").
		SetPendingBelowPrice("5").
		SetPendingBelowType(enums.OrderTypeLimitMaker).
		SetPendingAboveStopPrice("0.5").
		SetPendingAboveType(enums.OrderTypeStopLoss).
		SetWorkingPrice("1.5").
		SetWorkingQuantity("1").
		SetWorkingSide(enums.SideTypeBuy).
		SetWorkingTimeInForce(enums.TimeInForceTypeGTC).
		SetWorkingType(enums.OrderTypeLimit).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
	}
}
func TestWsOrderListCancel(t *testing.T) {
	res, err := trading.NewWsApiOrderList(wsApiClient).
		SetOrigClientOrderId("fvh0x6K2e2s0Gk4oqw7seI").
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
	}
}
func TestOpenOrderLists(t *testing.T) {
	res, err := account.NewWsApiOpenOrderList(wsApiClient).Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		for _, v := range res.Result {
			fmt.Printf("%+v\n", v)
		}
	}
}
func TestWsApiSor(t *testing.T) {
	res, err := trading.NewWsApiSOR(wsApiClient).
		SetSymbol(ETHUSDT).
		SetSide(enums.SideTypeBuy).
		SetType(enums.OrderTypeMarket).
		SetQuantity("0.0001").
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
	}
}
func TestWsApiSorTest(t *testing.T) {
	res, err := trading.NewWsApiSOR(wsApiClient).
		SetSymbol(ETHUSDT).
		SetSide(enums.SideTypeBuy).
		SetType(enums.OrderTypeMarket).
		SetQuantity("0.0001").
		SendTest(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
	}
}
func TestWsApiAccount(t *testing.T) {
	res, err := account.NewWsApiAccount(wsApiClient).
		SetOmitZeroBalances(false).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
	}
}
func TestWsApiRateLimitOrder(t *testing.T) {
	res, err := account.NewWsApiWsApiRateLimitOrder(wsApiClient).
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
func TestWsApiNewWsApiAllOrderList(t *testing.T) {
	res, err := account.NewWsApiAllOrderList(wsApiClient).
		SetLimit(enums.Limit20).
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
		fmt.Printf("%+v\n", res.RateLimits)
	}
}
func TestWsApiAllOrders(t *testing.T) {
	res, err := account.NewWsApiAllOrders(wsApiClient).
		SetSymbol(ETHUSDT).
		SetLimit(enums.Limit100).
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
		fmt.Printf("%+v\n", res.RateLimits)
	}
}
func TestWsApiMyTrades(t *testing.T) {
	res, err := account.NewWsApiMyTrades(wsApiClient).
		SetSymbol(ETHUSDT).
		SetFromId(834230).
		SetLimit(enums.Limit5).
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
		fmt.Printf("%+v\n", res.RateLimits)
	}
}
func TestNewWsApiMyPreventedMatches(t *testing.T) {
	res, err := account.NewWsApiMyPreventedMatches(wsApiClient).
		SetSymbol(ETHUSDT).
		SetLimit(enums.Limit5).
		SetOrderId(11750571916).
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
		fmt.Printf("%+v\n", res.RateLimits)
	}
}
func TestWsApiMyAllocations(t *testing.T) {
	res, err := account.NewWsApiMyAllocations(wsApiClient).
		SetSymbol(ETHUSDT).
		SetLimit(enums.Limit5).
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
		fmt.Printf("%+v\n", res.RateLimits)
	}
}
func TestNewWsApiWsApiCommission(t *testing.T) {
	res, err := account.NewWsApiCommission(wsApiClient).
		SetSymbol(ETHUSDT).
		Send(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
	} else {
		fmt.Printf("%+v\n", res.Result)
		fmt.Printf("%+v\n", res.RateLimits)
	}
}
func TestNewWsApiUserDataStream(t *testing.T) {
	ds := stream.NewWsApiUserDataStream(wsApiClient)
	res, err := ds.SendStart(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	fmt.Printf("%+v\n", res.Result)
	fmt.Printf("%+v\n", res.RateLimits)
	listenKey := res.Result.ListenKey
	res, err = ds.SetListenKey(listenKey).SendPing(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	fmt.Printf("%+v\n", res.RateLimits)
	res, err = ds.SetListenKey(listenKey).SendStop(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	fmt.Printf("%+v\n", res.RateLimits)
}
