package exchange_go

import (
	"context"
	"fmt"
	"testing"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
	"github.com/sleep-go/exchange-go/binance/spot/endpoints/general"
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
	response, err := general.NewExchangeInfo(client).Call(context.Background(), []string{"ETHUSDT"}, nil)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(response)
}
