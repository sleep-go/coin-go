package spot_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/futures/general"
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
