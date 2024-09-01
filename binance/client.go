package binance

import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sleep-go/exchange-go/binance/consts"
)

var LogLevel = os.Stderr

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
	Logger     *log.Logger
	TimeOffset int64
	Debug      bool
	PrivateKey crypto.Signer
}

// NewClient 创建客户端函数来初始化客户端
func NewClient(apiKey, secretKey string, baseURL ...string) *Client {
	api := consts.REST_API
	if len(baseURL) > 0 {
		api = baseURL[0]
	}
	prefix := "[INFO] "
	switch LogLevel {
	case os.Stderr:
		prefix = "[ERROR] "
	case os.Stdout:
		prefix = "[INFO] "
	}
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    api,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(LogLevel, prefix, log.LstdFlags),
	}
}

func NewRsaClient(apiKey, privateKeyPath string, baseURL ...string) *Client {
	api := consts.REST_API
	if len(baseURL) > 0 {
		api = baseURL[0]
	}
	prefix := "[INFO] "
	switch LogLevel {
	case os.Stderr:
		prefix = "[ERROR] "
	case os.Stdout:
		prefix = "[INFO] "
	}
	// 加载 private key
	privateKey, err := loadRsaPrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	return &Client{
		APIKey:     apiKey,
		PrivateKey: privateKey,
		BaseURL:    api,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(LogLevel, prefix, log.LstdFlags),
	}
}

func NewECDAClient(apiKey, privateKeyPath string, baseURL ...string) *Client {
	api := consts.REST_API
	if len(baseURL) > 0 {
		api = baseURL[0]
	}
	prefix := "[INFO] "
	switch LogLevel {
	case os.Stderr:
		prefix = "[ERROR] "
	case os.Stdout:
		prefix = "[INFO] "
	}
	// 加载 private key
	privateKey, err := loadECDAPrivateKey(privateKeyPath)
	if err != nil {
		panic(err)
	}
	return &Client{
		APIKey:     apiKey,
		PrivateKey: privateKey,
		BaseURL:    api,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(LogLevel, prefix, log.LstdFlags),
	}
}

func (c *Client) request(ctx context.Context, r *Request) (*http.Request, error) {
	r.header = http.Header{}
	r.header.Set("X-MBX-APIKEY", c.APIKey)
	r.header.Set("Content-Type", "application/x-www-form-urlencoded")
	//获取 query url
	queryString := r.query.Encode()
	//获取body
	bodyString := r.form.Encode()
	if r.needSign {
		r.body = &bytes.Buffer{}
		r.body = bytes.NewBufferString(bodyString)
		//设置签名参数
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		r.SetParam("timestamp", time.Now().UnixMilli()-c.TimeOffset)
		if c.SecretKey != "" {
			r.SetParam("signature", signPayload(raw, c.SecretKey))
		} else if c.PrivateKey != nil {
			r.SetParam("signature", signPayload(raw, c.PrivateKey))
		} else {
			c.Println("signature is empty")
		}
	}
	//获取请求地址完整路径
	r.fullURL = fmt.Sprintf("%s%s", c.BaseURL, r.Path)
	if queryString != "" {
		r.fullURL = fmt.Sprintf("%s?%s", r.fullURL, r.query.Encode())
	}
	req, err := http.NewRequest(r.Method, r.fullURL, r.body)
	if err != nil {
		c.Logger.Println(err)
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.Debugf("r.fullURL:%s", r.fullURL)
	c.Debugf("query:%s", r.query)
	c.Debugf("form:%v", r.form)
	return req, nil
}

func (c *Client) Do(ctx context.Context, r *Request) (*http.Response, error) {
	request, err := c.request(ctx, r)
	if err != nil {
		c.Debugf("request err:%v", err)
		return nil, err
	}
	return c.HTTPClient.Do(request)
}

func (c *Client) Debugf(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}
func (c *Client) Println(v ...interface{}) {
	c.Logger.Println(v...)
}

// loadRsaPrivateKey 加载并解析 PEM 编码的 RSA 私钥
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

// loadECDAPrivateKey 加载并解析 PEM 编码的 ECDSA 私钥
func loadECDAPrivateKey(path string) (*ecdsa.PrivateKey, error) {
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
