package exchange_go

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sleep-go/exchange-go/binance/enums"

	"github.com/sleep-go/exchange-go/binance/spot/endpoints/market/ticker"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
	"github.com/sleep-go/exchange-go/binance/spot/endpoints/general"
	"github.com/sleep-go/exchange-go/binance/spot/endpoints/market"
	"github.com/spf13/cast"
)

var client *binance.Client

func init() {
	client = binance.NewClient(
		"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
		consts.REST_API,
	)
	client.Debug = true
}
func TestNewExchangeInfo(t *testing.T) {
	response, err := general.NewExchangeInfo(client, []string{"ETHUSDT"}, nil).Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(response)
}

func TestDepth(t *testing.T) {
	response, err := market.NewDepth(client, "ETCUSDT", market.DepthLimit20).Call(context.Background())
	if err != nil {
		return
	}
	fmt.Println(len(response.Asks))
	fmt.Println(len(response.Bids))
	fmt.Println(response.LastUpdateId)
}
func TestTrades(t *testing.T) {
	res, err := market.NewTrades(client, "BTCUSDT", market.TradesLimit500).Call(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
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
	k := market.NewKlines(client, "BTCUSDT", market.TradesLimit500).
		SetInterval(market.KlineInterval1m).
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
