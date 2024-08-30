package binance

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/sleep-go/exchange-go/binance/consts"
)

func TestNewClient(t *testing.T) {
	client := NewClient(
		"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
		consts.TESTNET,
	)
	client.TimeOffset = 1000
	r := &Request{
		Method: http.MethodGet,
		Path:   consts.ExchangeInfoApi,
	}
	symbols := []string{"ETHUSDT"}
	result := fmt.Sprintf(`["%s"]`, strings.Join(symbols, `","`))
	r.SetParam("symbols", result)
	request, err := client.request(context.Background(), r)
	if err != nil {
		return
	}
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()
	res, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(res))
}

func TestNewRequest(t *testing.T) {
	fmt.Println("1725013324892")
	fmt.Println(time.Now().UnixMilli())
}
