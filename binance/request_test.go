package binance

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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
		Path:   consts.ApiExchangeInfo,
	}
	symbols := []string{"BTCUSDT"}
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

// todo 待封装order相关接口
func TestOrder(t *testing.T) {
	// 设置身份验证
	file, err := os.ReadFile("../.env")
	if err != nil {
		fmt.Println(err)
		return
	}
	API_KEY := strings.TrimSpace(string(file))
	PRIVATE_KEY_PATH := "../private.pem"

	// 加载 private key
	privateKey, err := loadRsaPrivateKey(PRIVATE_KEY_PATH)
	if err != nil {
		panic(err)
	}

	// 设置请求参数
	params := url.Values{
		"symbol": {"BNBUSDT"},
		//"side":        {enums.BUY.String()},
		//"type":        {enums.LIMIT.String()},
		//"quantity":    {"0.1"},
		//"timeInForce": {enums.GTC.String()},
		//"price":       {"400"},
		//"orderId": {"5747170632"},
	}

	// 参数中加时间戳
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli()) // 以毫秒为单位的 UNIX 时间戳
	params.Set("timestamp", timestamp)

	// 参数中加签名
	payload := params.Encode()
	signature, err := signRsaPayload(payload, privateKey)
	if err != nil {
		panic(err)
	}
	params.Set("signature", signature)

	// 发送请求
	client := &http.Client{}
	//req, err := http.NewRequest(http.MethodPost, "https://api.binance.com/api/v3/order", bytes.NewBufferString(params.Encode()))
	req, err := http.NewRequest(http.MethodGet, "https://api.binance.com/api/v3/openOrders?"+params.Encode(), nil)
	//req, err := http.NewRequest(http.MethodGet, "https://api.binance.com/api/v3/order?"+params.Encode(), nil)
	//req, err := http.NewRequest(http.MethodDelete, "https://api.binance.com/api/v3/order", bytes.NewBufferString(params.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-MBX-APIKEY", API_KEY)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// loadPrivateKey 加载并解析 PEM 编码的 ECDSA 私钥
func loadPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	privKeyPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privKeyPEM)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
func loadRsaPrivateKey(path string) (*rsa.PrivateKey, error) {
	privKeyPEM, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// signPayload 用 ECDSA 私钥签名数据
func signPayload(payload string, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := sha256.Sum256([]byte(payload))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	return base64.StdEncoding.EncodeToString(signature), nil
}
func signRsaPayload(payload string, privateKey *rsa.PrivateKey) (string, error) {
	hash := sha256.Sum256([]byte(payload))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
