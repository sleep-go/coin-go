package general

import (
	"context"
	"fmt"
	"testing"

	"github.com/sleep-go/exchange-go/binance/consts"

	"github.com/sleep-go/exchange-go/binance"
)

func TestExchangeInfo_Do(t *testing.T) {
	client := binance.NewClient(
		"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
		consts.TESTNET,
	)
	ex := ExchangeInfo{
		Client: client,
	}
	do, err := ex.Do(context.Background(), []string{"BTCUSDT"})
	if err != nil {
		return
	}
	for _, limit := range do.RateLimits {
		fmt.Println(limit.RateLimitType, limit.IntervalNum, limit.Limit)
	}
}
