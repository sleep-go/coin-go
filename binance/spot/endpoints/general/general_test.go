package general

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
)

var client *binance.Client

func init() {
	binance.LogLevel = os.Stdout
	client = binance.NewClient(
		"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
		consts.TESTNET,
	)
}
func TestGetIP(t *testing.T) {
	curl := "https://www.ip.cn/api/index?ip&type=0"
	get, err := http.Get(curl)
	if err != nil {
		return
	}
	defer get.Body.Close()
	all, err := io.ReadAll(get.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))
}
func TestExchangeInfo_Do(t *testing.T) {
	ex := exchangeInfoRequest{
		Client:      client,
		symbols:     []string{},
		permissions: nil,
	}
	do, err := ex.Call(context.Background())
	if err != nil {
		return
	}
	for _, limit := range do.RateLimits {
		client.Debugf(limit.RateLimitType, limit.IntervalNum, limit.Limit)
	}
	for _, s := range do.Symbols {
		client.Logger.Println(s.Symbol, s.Filters)
	}
}
func TestPing_Do(t *testing.T) {
	ping := NewPing(client)
	do, err := ping.Call(context.Background())
	if err != nil {
		return
	}
	binance.LogLevel = os.Stdout
	t.Log(do.Status, do.Code)
}
func TestTime_Do(t *testing.T) {
	tr := NewTime(client)
	for i := 0; i < 3; i++ {
		res, err := tr.Call(context.Background())
		if err != nil {
			return
		}
		t.Log(time.UnixMilli(res.ServerTime))
		time.Sleep(1 * time.Second)
	}
}
