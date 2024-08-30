package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/sleep-go/exchange-go/binance/consts"
)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
}

type Request struct {
	Method     string //请求方法
	Path       string //请求路径
	fullURL    string
	RecvWindow int64
	query      url.Values
	form       url.Values
	header     http.Header
	body       io.Reader
}

func (r *Request) addParam(key string, value interface{}) *Request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// SetParam set param with key/value to query string
func (r *Request) SetParam(key string, value interface{}) *Request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}
func (r *Request) SetForm(key string, value interface{}) *Request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// NewClient 创建客户端函数来初始化客户端
func NewClient(apiKey string, secretKey string, baseURL ...string) *Client {
	api := consts.REST_API
	if len(baseURL) > 0 {
		api = baseURL[0]
	}
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    api,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (c *Client) request(ctx context.Context, r *Request) (*http.Request, error) {
	r.header = http.Header{}
	queryString := r.query.Encode()
	r.body = &bytes.Buffer{}
	bodyString := r.form.Encode()
	r.header.Set("User-Agent", fmt.Sprintf("%s/%s", consts.NAME, "1.0"))
	if c.APIKey != "" {
		r.header.Set("X-MBX-APIKEY", c.APIKey)
	}
	r.fullURL = fmt.Sprintf("%s%s", c.BaseURL, r.Path)
	if bodyString != "" {
		r.header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.body = bytes.NewBufferString(bodyString)
		//设置签名参数
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		r.SetParam("timestamp", time.Now().UnixMilli()-c.TimeOffset)
		r.SetParam("signature", sign(raw, c.SecretKey))
	}
	if queryString != "" {
		r.fullURL = fmt.Sprintf("%s?%s", r.fullURL, r.query.Encode())
	}
	request, err := http.NewRequest(r.Method, r.fullURL, r.body)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	request.Header = r.header
	c.Logger.Println("r.fullURL", r.fullURL)
	c.Logger.Println("query", r.query)
	c.Logger.Println("form", r.form)
	return request, nil
}

func (c *Client) Do(ctx context.Context, r *Request) (*http.Response, error) {
	request, err := c.request(ctx, r)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(request)
}

func sign(raw string, secret string) string {
	// 创建 HMAC-SHA256 哈希
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(raw))
	if err != nil {
		return ""
	}
	// 计算 HMAC 并将结果转换为十六进制字符串
	return hex.EncodeToString(h.Sum(nil))
}
