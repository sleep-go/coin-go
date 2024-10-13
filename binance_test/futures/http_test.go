package spot_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/enums"
	"github.com/sleep-go/coin-go/binance/futures/general"
	"github.com/sleep-go/coin-go/binance/futures/market"
	"github.com/sleep-go/coin-go/binance/futures/market/data"
	"github.com/sleep-go/coin-go/binance/futures/market/ticker"
	"github.com/spf13/cast"
)

var client *binance.Client

const (
	BTCUSDT = "BTCUSDT"
	ETHUSDT = "ETHUSDT"
)

func init() {
	// 设置身份验证
	file, err := os.ReadFile("./.env")
	if err != nil {
		panic(err)
		return
	}
	API_KEY := strings.TrimSpace(string(file))
	PRIVATE_KEY_PATH := "./private.pem"
	fmt.Println(API_KEY)
	client = binance.NewRsaClient(API_KEY, PRIVATE_KEY_PATH, consts.REST_FAPI_TEST)
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
	response, err := general.NewExchangeInfo(client).
		Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(len(response.Symbols))
	fmt.Printf("%+v\n", response)
}
func TestDepth(t *testing.T) {
	res, err := market.NewDepth(client, ETHUSDT, enums.Limit20).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Asks)
	fmt.Println(len(res.Asks))
	fmt.Println(len(res.Bids))
	fmt.Println(res.LastUpdateId)
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
		SetFromId(290414224).
		Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, r := range res {
		fmt.Printf("%+v\n", r)
	}
}
func TestAggTrades(t *testing.T) {
	res, err := market.NewAggTrades(client, BTCUSDT, enums.Limit20).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		//SetFromId(3031206).
		Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, r := range res {
		fmt.Printf("%+v\n", r)
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
	for i, r := range res {
		fmt.Print(time.UnixMilli(cast.ToInt64(r[0])).Format(time.DateTime), "开盘时间 ") // 开盘时间
		fmt.Print(r[1], " ", res[i][1], "开盘价 ")                                      // 开盘价
		fmt.Print(r[2], " ", res[i][2], "最高价 ")                                      // 最高价
		fmt.Print(r[3], " ", res[i][3], "最低价 ")                                      // 最低价
		fmt.Print(r[4], " ", res[i][4], "收盘价 ")                                      // 收盘价(当前K线未结束的即为最新价)
		fmt.Print(r[5], "成交量 ")                                                      // 成交量
		fmt.Print(time.UnixMilli(cast.ToInt64(r[6])), "收盘时间 ")                       // 收盘时间
		fmt.Print(r[7], "成交额 ")                                                      // 成交额
		fmt.Print(r[8], "成交笔数 ")                                                     // 成交笔数
		fmt.Print(r[9], "主动买入成交量 ")                                                  // 主动买入成交量
		fmt.Print(r[10], "主动买入成交额 ")                                                 // 主动买入成交额
		fmt.Println(r[11])                                                           // 请忽略该参数
	}
}
func TestContinuousKlines(t *testing.T) {
	res, err := market.NewKlines(client, BTCUSDT, enums.Limit100).
		SetContractType(enums.ContractTypePerpetual).
		SetInterval(enums.KlineIntervalType1M).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		SetTimeZone("0").
		CallMarkPriceKlines(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestCallIndexPriceKlines(t *testing.T) {
	res, err := market.NewKlines(client, BTCUSDT, enums.Limit100).
		SetContractType(enums.ContractTypePerpetual).
		SetInterval(enums.KlineIntervalType1M).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		SetTimeZone("0").
		CallIndexPriceKlines(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, r := range res {
		fmt.Print("开盘时间:", time.UnixMilli(cast.ToInt64(r[0])).Format(time.DateTime), " ") // 开盘时间
		fmt.Print("开盘价:", r[1], " ")                                                      // 开盘价
		fmt.Print("最高价:", r[2], " ")                                                      // 最高价
		fmt.Print("最低价:", r[3], " ")                                                      // 最低价
		fmt.Print("收盘价:", r[4], " ")                                                      // 收盘价(当前K线未结束的即为最新价)
		fmt.Print("成交量:", r[5], " ")                                                      // 成交量
		fmt.Print("收盘时间:", time.UnixMilli(cast.ToInt64(r[6])), " ")                       // 收盘时间
		fmt.Print("成交额:", r[7], " ")                                                      // 成交额
		fmt.Print("成交笔数:", r[8], " ")                                                     // 成交笔数
		fmt.Print("主动买入成交量:", r[9], " ")                                                  // 主动买入成交量
		fmt.Print("主动买入成交额:", r[10], " ")                                                 // 主动买入成交额
		fmt.Println(r[11])                                                                // 请忽略该参数
	}
}
func TestCallMarkPriceKlines(t *testing.T) {
	res, err := market.NewKlines(client, BTCUSDT, enums.Limit100).
		SetContractType(enums.ContractTypePerpetual).
		SetInterval(enums.KlineIntervalType1M).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		SetTimeZone("0").
		CallIndexPriceKlines(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, r := range res {
		fmt.Print("开盘时间:", time.UnixMilli(cast.ToInt64(r[0])).Format(time.DateTime), " ") // 开盘时间
		fmt.Print("开盘价:", r[1], " ")                                                      // 开盘价
		fmt.Print("最高价:", r[2], " ")                                                      // 最高价
		fmt.Print("最低价:", r[3], " ")                                                      // 最低价
		fmt.Print("收盘价:", r[4], " ")                                                      // 收盘价(当前K线未结束的即为最新价)
		fmt.Print("成交量:", r[5], " ")                                                      // 成交量
		fmt.Print("收盘时间:", time.UnixMilli(cast.ToInt64(r[6])), " ")                       // 收盘时间
		fmt.Print("成交额:", r[7], " ")                                                      // 成交额
		fmt.Print("成交笔数:", r[8], " ")                                                     // 成交笔数
		fmt.Print("主动买入成交量:", r[9], " ")                                                  // 主动买入成交量
		fmt.Print("主动买入成交额:", r[10], " ")                                                 // 主动买入成交额
		fmt.Println(r[11])                                                                // 请忽略该参数
	}
}
func TestCallPremiumIndexKlines(t *testing.T) {
	res, err := market.NewKlines(client, BTCUSDT, enums.Limit100).
		SetContractType(enums.ContractTypePerpetual).
		SetInterval(enums.KlineIntervalType1M).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		SetTimeZone("0").
		CallPremiumIndexKlines(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, r := range res {
		fmt.Print("开盘时间:", time.UnixMilli(cast.ToInt64(r[0])).Format(time.DateTime), " ") // 开盘时间
		fmt.Print("开盘价:", r[1], " ")                                                      // 开盘价
		fmt.Print("最高价:", r[2], " ")                                                      // 最高价
		fmt.Print("最低价:", r[3], " ")                                                      // 最低价
		fmt.Print("收盘价:", r[4], " ")                                                      // 收盘价(当前K线未结束的即为最新价)
		fmt.Print("成交量:", r[5], " ")                                                      // 成交量
		fmt.Print("收盘时间:", time.UnixMilli(cast.ToInt64(r[6])), " ")                       // 收盘时间
		fmt.Print("成交额:", r[7], " ")                                                      // 成交额
		fmt.Print("成交笔数:", r[8], " ")                                                     // 成交笔数
		fmt.Print("主动买入成交量:", r[9], " ")                                                  // 主动买入成交量
		fmt.Print("主动买入成交额:", r[10], " ")                                                 // 主动买入成交额
		fmt.Println(r[11])                                                                // 请忽略该参数
	}
}
func TestCallPremiumIndex(t *testing.T) {
	resp, err := market.NewPremiumIndex(client).
		Call(context.Background(), BTCUSDT)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(resp)
	res, err := market.NewPremiumIndex(client).
		CallAll(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestCallFundingRate(t *testing.T) {
	res, err := market.NewFundingRate(client).
		Call(context.Background(), ETHUSDT)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestCallFundingInfo(t *testing.T) {
	client.BaseURL = consts.REST_FAPI
	res, err := market.NewFundingInfo(client).Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestHr24(t *testing.T) {
	res, err := ticker.NewHr24(client, ETHUSDT).Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	fmt.Println(res)
	ress, err := ticker.NewHr24(client, ETHUSDT).CallAll(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range ress {
		fmt.Printf("%+v\n", v)
	}
}
func TestNewPrice(t *testing.T) {
	res, err := ticker.NewPrice(client).CallAllV2(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range res {
		fmt.Println(v)
	}
	resp, err := ticker.NewPrice(client).CallV2(context.Background(), ETHUSDT)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(resp)
}
func TestBookTicker(t *testing.T) {
	res, err := ticker.NewBookTicker(client).Call(context.Background(), ETHUSDT)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	fmt.Printf("%+v\n", res)
	resp, err := ticker.NewBookTicker(client).CallAll(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, v := range resp {
		fmt.Println(v)
	}
}
func TestDeliveryPrice(t *testing.T) {
	client.BaseURL = consts.REST_FAPI
	res, err := data.NewDeliveryPrice(client, ETHUSDT).Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
func TestOpenInterest(t *testing.T) {
	res, err := market.NewOpenInterest(client).Call(context.Background(), ETHUSDT)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
func TestOpenInterestHist(t *testing.T) {
	client.BaseURL = consts.REST_FAPI
	res, err := data.NewOpenInterestHist(client).
		SetPeriod(enums.KlineIntervalType5m).
		SetLimit(enums.Limit5).
		Call(context.Background(), ETHUSDT)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
