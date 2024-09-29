package binance

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sleep-go/coin-go/binance/consts"
)

var client *Client

func init() {
	// 设置身份验证
	file, err := os.ReadFile("../.env")
	if err != nil {
		fmt.Println(err)
		return
	}
	API_KEY := strings.TrimSpace(string(file))
	PRIVATE_KEY_PATH := "../private.pem"
	fmt.Println(API_KEY)
	client = NewRsaClient(API_KEY, PRIVATE_KEY_PATH)
	client.Debug = false
}

// todo 待封装order相关接口
func TestOrder(t *testing.T) {
	req := Request{
		Method:   http.MethodGet,
		Path:     consts.ApiTradingOpenOrders,
		needSign: true,
	}
	req.SetParam("symbol", "BTCUSDT")
	req.SetParam("timestamp", time.Now().UnixMilli())
	resp, err := client.Do(context.Background(), &req)
	fmt.Println(req.query.Encode())
	if err != nil {
		return
	}
	//signature=SpDQS1R24mQGZrYRAZe9hkBuRzaGGEbg7Z5Z1cI0YGGUmSF0v39pAQbz5z1P72RqFezaPYtyn4AqwjB4RZKGsjvKgFb4b5W04zZYBni%2FbnamYuWZuAwc5%2BCdo33kZDZaN%2B6hwGmLZq8g2c8EdXpzgkQ8Ik1sy91wkrhVw%2BOwpLY%3D&symbol=BTCUSDT&timestamp=1727624123207
	//// 设置请求参数
	//params := url.Values{
	//	"symbol": {"BNBUSDT"},
	//	//"side":        {enums.BUY.String()},
	//	//"type":        {enums.LIMIT.String()},
	//	//"quantity":    {"0.1"},
	//	//"timeInForce": {enums.GTC.String()},
	//	//"price":       {"400"},
	//	//"orderId": {"5747170632"},
	//}
	//
	//// 参数中加时间戳
	//timestamp := fmt.Sprintf("%d", time.Now().UnixMilli()) // 以毫秒为单位的 UNIX 时间戳
	//params.Set("timestamp", timestamp)
	//
	//// 参数中加签名
	//payload := params.Encode()
	//signature := signPayload(payload, client.PrivateKey)
	//params.Set("signature", signature)
	//
	//// 发送请求
	//httpClient := &http.Client{}
	////req, err := http.NewRequest(http.MethodPost, "https://api.binance.com/api/v3/order", bytes.NewBufferString(params.Encode()))
	//req, err := http.NewRequest(http.MethodGet, "https://api.binance.com/api/v3/openOrders?"+params.Encode(), nil)
	////req, err := http.NewRequest(http.MethodGet, "https://api.binance.com/api/v3/order?"+params.Encode(), nil)
	////req, err := http.NewRequest(http.MethodDelete, "https://api.binance.com/api/v3/order", bytes.NewBufferString(params.Encode()))
	//if err != nil {
	//	panic(err)
	//}
	//req.Header.Set("X-MBX-APIKEY", client.APIKey)
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//resp, err := httpClient.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
