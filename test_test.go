package exchange_go

import (
	"context"
	"fmt"
	"testing"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/spot/endpoints/general"
)

func TestName(t *testing.T) {
	c := binance.NewClient("", "")
	response, err := general.NewExchangeInfo(c).Call(context.Background(), []string{"ETHUSDT"}, nil)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(response)
}
