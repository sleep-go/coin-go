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
	"github.com/sleep-go/coin-go/binance/spot/endpoints/account"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/general"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market/ticker"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/stream"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/trading"
	"github.com/spf13/cast"
)

var client *binance.Client

const (
	BTCUSDT = "BTCUSDT"
	ETHUSDT = "ETHUSDT"
)

func init() {
	// 设置身份验证
	file, err := os.ReadFile("./.test.env")
	if err != nil {
		panic(err)
		return
	}
	API_KEY := strings.TrimSpace(string(file))
	PRIVATE_KEY_PATH := "./test-prv-key.pem"
	fmt.Println(API_KEY)
	client = binance.NewED25519Client(API_KEY, PRIVATE_KEY_PATH, consts.REST_API_TEST)
	client.Debug = true
}
func TestPing(t *testing.T) {
	res, err := general.NewPing(client).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
func TestTime(t *testing.T) {
	res, err := general.NewTime(client).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
func TestNewExchangeInfo(t *testing.T) {
	response, err := general.NewExchangeInfo(client, []string{"ETHUSDT", BTCUSDT}, nil).Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(response)
}
func TestDepth(t *testing.T) {
	response, err := market.NewDepth(client, "ETCUSDT", enums.Limit20).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(response.Asks))
	fmt.Println(len(response.Bids))
	fmt.Println(response.LastUpdateId)
}
func TestTrades(t *testing.T) {
	res, err := market.NewTrades(client, BTCUSDT, enums.Limit20).Call(context.Background())
	if err != nil {
		fmt.Println(res)
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestHistoryTrades(t *testing.T) {
	res, err := market.NewHistoryTrades(client, BTCUSDT, 1).
		SetFromId(3049539).
		Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, r := range res {
		fmt.Println(r)
	}
}
func TestAggTrades(t *testing.T) {
	res, err := market.NewAggTrades(client, BTCUSDT, 1).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		//SetFromId(3031206).
		Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, r := range res {
		fmt.Println(r)
	}
}
func TestKlines(t *testing.T) {
	k := market.NewKlines(client, BTCUSDT, enums.Limit100).
		SetInterval(enums.KlineIntervalType1M).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		SetTimeZone("0")
	res, err := k.Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	res1, err := k.CallUI(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for i, r := range res {
		fmt.Print(time.UnixMilli(cast.ToInt64(r[0])).Format(time.DateTime), "开盘时间 ") // 开盘时间
		fmt.Print(r[1], " ", res1[i][1], "开盘价 ")                                     // 开盘价
		fmt.Print(r[2], " ", res1[i][2], "最高价 ")                                     // 最高价
		fmt.Print(r[3], " ", res1[i][3], "最低价 ")                                     // 最低价
		fmt.Print(r[4], " ", res1[i][4], "收盘价 ")                                     // 收盘价(当前K线未结束的即为最新价)
		fmt.Print(r[5], "成交量 ")                                                      // 成交量
		fmt.Print(time.UnixMilli(cast.ToInt64(r[6])), "收盘时间 ")                       // 收盘时间
		fmt.Print(r[7], "成交额 ")                                                      // 成交额
		fmt.Print(r[8], "成交笔数 ")                                                     // 成交笔数
		fmt.Print(r[9], "主动买入成交量 ")                                                  // 主动买入成交量
		fmt.Print(r[10], "主动买入成交额 ")                                                 // 主动买入成交额
		fmt.Println(r[11])                                                           // 请忽略该参数
	}
}
func TestAvgPrice(t *testing.T) {
	res, err := market.NewAvgPrice(client, BTCUSDT).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	fmt.Println(res)
}
func TestHr24(t *testing.T) {
	res, err := ticker.NewHr24(client, []string{"ETHUSDT", "BNBBTC"}, enums.TickerTypeFull).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
func TestTradingDay(t *testing.T) {
	res, err := ticker.NewTradingDay(client, []string{BTCUSDT}, "+8:00", enums.TickerTypeFull).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
func TestNewPrice(t *testing.T) {
	res, err := ticker.NewPrice(client, []string{"ETHUSDT", "BNBBTC"}).Call(context.Background())
	if err != nil {
		return
	}
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
func TestBookTicker(t *testing.T) {
	res, err := ticker.NewBookTicker(client, []string{"ETHUSDT", "BNBBTC"}).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
func TestTicker(t *testing.T) {
	var u uint8 = 255
	fmt.Println(u)
	res, err := ticker.NewTicker(client, []string{"ETHUSDT"}, enums.TickerTypeFull).SetMinute(1).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}
	res, err = ticker.NewTicker(client, []string{"ETHUSDT", BTCUSDT}, enums.TickerTypeFull).SetDay(1).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestQueryOrder(t *testing.T) {
	res, err := trading.NewQueryOrder(client, BTCUSDT).
		//SetFormId，SetOrigClientOrderId 二选一
		SetOrderId(30102167318).
		//SetFormId，SetOrigClientOrderId 二选一
		//SetOrigClientOrderId("ios_e5556c10ddda4b4e8520c300cbab4c73").
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", res)
}
func TestOpenOrders(t *testing.T) {
	res, err := trading.NewQueryOrder(client, BTCUSDT).
		//SetTimestamp(time.Now().UnixMilli()).
		CallOpenOrders(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", res)
}
func TestCreateOrder(t *testing.T) {
	res, err := trading.NewOrder(client, BTCUSDT).
		SetQuantity("1").
		SetType(enums.OrderTypeMarket).
		SetSide(enums.SideTypeBuy).
		CallTest(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestAllOrders(t *testing.T) {
	res, err := account.NewAllOrders(client, BTCUSDT, enums.Limit20).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestCancelReplace(t *testing.T) {
	res, err := trading.NewCancelReplace(client, BTCUSDT).
		SetSide(enums.SideTypeBuy).
		SetQuantity("0.0001").
		SetCancelReplaceMode(enums.CancelReplaceModeTypeStopOnFailure).
		SetCancelOrderId(123).
		SetType(enums.OrderTypeMarket).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestDeleteOrder(t *testing.T) {
	response, err := trading.NewDeleteOrder(client, BTCUSDT).
		SetOrderId(394763750).
		SetCancelRestrictions(enums.CancelRestrictionsTypeOnlyNew).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", response)
}
func TestDeleteOpenOrders(t *testing.T) {
	response, err := trading.NewDeleteOpenOrders(client, BTCUSDT).Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", response)
}
func TestGetAccount(t *testing.T) {
	response, err := account.NewGetAccount(client).
		SetOmitZeroBalances(true).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, balance := range response.Balances {
		fmt.Printf("%+v\n", balance)
	}
	fmt.Printf("%+v\n", response.Permissions)
}
func TestMyTrades(t *testing.T) {
	res, err := account.NewMyTrades(client, BTCUSDT, 500).
		//SetFormId(11750571916).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestRateLimitOrder(t *testing.T) {
	res, err := account.NewRateLimitOrder(client).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestMyPreventedMatches(t *testing.T) {
	res, err := account.NewMyPreventedMatches(client, BTCUSDT, enums.Limit20).
		SetOrderId(11750571916).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestMyAllocations(t *testing.T) {
	res, err := account.NewMyAllocations(client, BTCUSDT, enums.Limit20).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestCommission(t *testing.T) {
	res, err := account.NewCommission(client, BTCUSDT).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestUserDataStream(t *testing.T) {
	res, err := stream.NewUserDataStream(client).CallCreate(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
	err = stream.NewUserDataStream(client).
		SetListenKey(res.ListenKey).
		CallUpdate(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	err = stream.NewUserDataStream(client).
		SetListenKey(res.ListenKey).
		CallDelete(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
func TestOCO(t *testing.T) {
	res, err := trading.NewOco(client, BTCUSDT).
		SetSide(enums.SideTypeSell).
		SetQuantity("1").
		SetBelowType(enums.OrderTypeStopLossLimit).
		SetBelowPrice("1").
		SetBelowStopPrice("1").
		SetBelowTrailingDelta(1).
		SetBelowTimeInForce(enums.TimeInForceTypeGTC).
		SetAboveType(enums.OrderTypeStopLossLimit).
		SetAbovePrice("1").
		SetAboveTrailingDelta(1).
		SetAboveTimeInForce(enums.TimeInForceTypeIOC).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestOTO(t *testing.T) {
	res, err := trading.NewOTO(client, BTCUSDT).
		SetWorkingType(enums.OrderTypeLimit).
		SetWorkingSide(enums.SideTypeSell).
		SetWorkingPrice("1").
		SetWorkingQuantity("1").
		SetPendingType(enums.OrderTypeStopLossLimit).
		SetPendingSide(enums.SideTypeSell).
		SetPendingQuantity("1").
		SetPendingPrice("1").
		SetPendingTrailingDelta("1").
		SetWorkingTimeInForce(enums.TimeInForceTypeGTC).
		SetPendingTimeInForce(enums.TimeInForceTypeGTC).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestOTOCO(t *testing.T) {
	res, err := trading.NewOtoco(client, BTCUSDT).
		SetWorkingType(enums.OrderTypeMarket).
		SetWorkingSide(enums.SideTypeSell).
		SetWorkingPrice("1").
		SetWorkingQuantity("1").
		SetPendingSide(enums.SideTypeSell).
		SetPendingQuantity("1").
		SetPendingAboveType(enums.OrderTypeStopLossLimit).
		SetPendingBelowType(enums.OrderTypeStopLossLimit).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestOrderList(t *testing.T) {
	res, err := trading.NewOrderList(client).
		SetOrderListId(123456).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestOpenOrderList(t *testing.T) {
	res, err := account.NewOpenOrderList(client).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestSor(t *testing.T) {
	res, err := trading.NewSor(client, BTCUSDT).
		SetSide(enums.SideTypeBuy).
		SetType(enums.OrderTypeMarket).
		SetQuantity("0.0001").
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestSorTest(t *testing.T) {
	res, err := trading.NewSor(client, BTCUSDT).
		SetSide(enums.SideTypeBuy).
		SetType(enums.OrderTypeMarket).
		SetQuantity("0.0001").
		CallTest(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
