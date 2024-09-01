package exchange_go

import (
	"context"
	"fmt"
	"testing"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
	"github.com/sleep-go/exchange-go/binance/spot/endpoints/general"
	"github.com/sleep-go/exchange-go/binance/spot/endpoints/market"
)

var client *binance.Client

func init() {
	client = binance.NewClient(
		"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
		consts.TESTNET,
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
