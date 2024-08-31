package binance

import (
	"bytes"
	"context"
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
}

// NewClient 创建客户端函数来初始化客户端
func NewClient(apiKey string, secretKey string, baseURL ...string) *Client {
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

func (c *Client) request(ctx context.Context, r *Request) (*http.Request, error) {
	r.header = http.Header{}
	r.header.Set("X-MBX-APIKEY", c.APIKey)
	r.header.Set("Content-Type", "application/x-www-form-urlencoded")
	//获取 query url
	queryString := r.query.Encode()
	//获取body
	bodyString := r.form.Encode()
	if bodyString != "" {
		r.body = &bytes.Buffer{}
		r.body = bytes.NewBufferString(bodyString)
		//设置签名参数
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		r.SetParam("timestamp", time.Now().UnixMilli()-c.TimeOffset)
		r.SetParam("signature", sign(raw, c.SecretKey))
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
