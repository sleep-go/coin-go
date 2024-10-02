package coin_go

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sleep-go/coin-go/binance/spot/endpoints/stream"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/account"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/general"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market/ticker"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/trading"
	"github.com/spf13/cast"
)

var client *binance.Client

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
	client = binance.NewED25519Client(API_KEY, PRIVATE_KEY_PATH, consts.TESTNET)
	client.Debug = true
}
func TestPing(t *testing.T) {
	res, err := general.NewPing(client).Call(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res.Status, res.Code)
}
func TestNewExchangeInfo(t *testing.T) {
	response, err := general.NewExchangeInfo(client, []string{"ETHUSDT", "BTCUSDT"}, nil).Call(context.Background())
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
	res, err := market.NewTrades(client, "BTCUSDT", enums.Limit20).Call(context.Background())
	if err != nil {
		fmt.Println(res)
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestHistoryTrades(t *testing.T) {
	res, err := market.NewHistoryTrades(client, "BTCUSDT", 1).
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
	res, err := market.NewAggTrades(client, "BTCUSDT", 1).
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
	k := market.NewKlines(client, "BTCUSDT", enums.Limit100).
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
	res, err := market.NewAvgPrice(client, "BTCUSDT").Call(context.Background())
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
	res, err := ticker.NewTradingDay(client, []string{"ETHUSDT", "BNBBTC"}, "8", enums.TickerTypeFull).Call(context.Background())
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
	res, err = ticker.NewTicker(client, []string{"ETHUSDT", "BTCUSDT"}, enums.TickerTypeFull).SetDay(1).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestQueryOrder(t *testing.T) {
	res, err := trading.NewQueryOrder(client, "BTCUSDT").
		//SetOrderId，SetOrigClientOrderId 二选一
		SetOrderId(30102167318).
		//SetOrderId，SetOrigClientOrderId 二选一
		//SetOrigClientOrderId("ios_e5556c10ddda4b4e8520c300cbab4c73").
		SetTimestamp(time.Now().UnixMilli()).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", res)
}
func TestNewOrder(t *testing.T) {
	res, err := trading.NewOrder(client, "BTCUSDT").
		SetQuantity("1").
		SetType(enums.OrderTypeMarket).
		SetSide(enums.SideTypeBuy).
		SetTimestamp(time.Now().UnixMilli()).
		CallTest(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestDeleteOrder(t *testing.T) {
	response, err := trading.NewDeleteOrder(client, "BTCUSDT").
		SetOrderId(394763750).
		SetTimestamp(time.Now().UnixMilli()).
		SetCancelRestrictions(enums.CancelRestrictionsTypeOnlyNew).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", response)
}
func TestGetAccount(t *testing.T) {
	response, err := account.NewGetAccount(client).
		SetOmitZeroBalances(true).
		SetTimestamp(time.Now().UnixMilli()).
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
	res, err := account.NewMyTrades(client, "BTCUSDT", 500).
		SetTimestamp(time.Now().UnixMilli()).
		//SetOrderId(11750571916).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestRateLimitOrder(t *testing.T) {
	res, err := account.NewRateLimitOrder(client).SetTimestamp(time.Now().UnixMilli()).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestMyPreventedMatches(t *testing.T) {
	res, err := account.NewMyPreventedMatches(client, "BTCUSDT", enums.Limit20).
		SetOrderId(11750571916).
		SetTimestamp(time.Now().UnixMilli()).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestMyAllocations(t *testing.T) {
	res, err := account.NewMyAllocations(client, "BTCUSDT", enums.Limit20).
		SetTimestamp(time.Now().UnixMilli()).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestCommission(t *testing.T) {
	res, err := account.NewCommission(client, "BTCUSDT").
		SetTimestamp(time.Now().UnixMilli()).
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
