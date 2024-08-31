package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method  string //请求方法
	Path    string //请求路径
	fullURL string
	query   url.Values
	form    url.Values
	header  http.Header
	body    io.Reader
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

// SetForm set form with key/value to body string
func (r *Request) SetForm(key string, value interface{}) *Request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
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
