package binance

import (
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/spf13/cast"
)

type Request struct {
	Method   string //请求方法
	Path     string //请求路径
	fullURL  string
	query    url.Values
	form     url.Values
	header   http.Header
	body     io.Reader
	needSign bool
}

func (r *Request) SetNeedSign(needSign bool) *Request {
	r.needSign = needSign
	return r
}

// SetParam set param with key/value to query string
func (r *Request) SetParam(key string, value any) *Request {
	if r.query == nil {
		r.query = url.Values{}
	}
	switch value.(type) {
	case *int, *int8, *int16, *int32, *int64:
		r.query.Set(key, fmt.Sprintf("%v", cast.ToString(value)))
	case *uint, *uint8, *uint16, *uint32, *uint64:
		r.query.Set(key, fmt.Sprintf("%v", cast.ToUint(value)))
	case *string:
		r.query.Set(key, fmt.Sprintf("%v", cast.ToString(value)))
	default:
		r.query.Set(key, fmt.Sprintf("%v", value))
	}
	return r
}
func (r *Request) SetOptionalParam(key string, value any) *Request {
	// 如果 value 为 nil，直接返回
	if value == nil {
		return r
	}
	// 使用反射获取实际的值
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Ptr:
		// 如果是指针并且指向 nil，则不设置参数
		if v.IsNil() {
			return r
		}
		// 获取指针指向的实际值
		value = v.Elem().Interface()
		if value == 0 {
			return r
		}
	case reflect.String:
		if v.String() == "" {
			return r
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.Int() == 0 {
			return r
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0 {
			return r
		}
	default:
		// 设置查询参数
		r.query.Set(key, fmt.Sprintf("%v", value))
	}
	return r
}

func (r *Request) SetOptionalParam2(key string, value any) *Request {
	// 如果 value 为 nil，直接返回
	if value == nil {
		return r
	}
	// 使用反射获取实际的值
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return r
		}
		// 获取指针指向的实际值
		value = v.Elem().Interface()
	case reflect.String:
		if v.String() == "" {
			return r
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.Int() == 0 {
			return r
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0 {
			return r
		}
	case reflect.Bool:
		if !v.Bool() {
			return r
		}
	default:
	}
	// 设置查询参数
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// SetForm set form with key/value to body string
func (r *Request) SetForm(key string, value any) *Request {
	if r.form == nil {
		r.form = url.Values{}
	}
	// 使用反射获取实际的值，避免重复处理不同的指针类型
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		// 如果是指针并且指向nil，则不设置参数
		if v.IsNil() {
			return r
		}
		// 获取指针指向的实际值
		value = v.Elem().Interface()
	}
	// 设置查询参数
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}
func (r *Request) SetOptionalForm(key string, value any) *Request {
	// 如果 value 为 nil，直接返回
	if value == nil {
		return r
	}
	// 使用反射获取实际的值
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		// 如果是指针并且指向 nil，则不设置参数
		if v.IsNil() {
			return r
		}
		// 获取指针指向的实际值
		value = v.Elem().Interface()
	}
	// 设置查询参数
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

func signPayload(payload string, privateKey any) string {
	switch privateKey.(type) {
	case string:
		h := hmac.New(sha256.New, []byte(privateKey.(string)))
		_, err := h.Write([]byte(payload))
		if err != nil {
			return ""
		}
		// 计算 HMAC 并将结果转换为十六进制字符串
		return hex.EncodeToString(h.Sum(nil))
	case *rsa.PrivateKey:
		hash := sha256.Sum256([]byte(payload))
		signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash[:])
		if err != nil {
			return ""
		}
		return base64.StdEncoding.EncodeToString(signature)
	case ed25519.PrivateKey:
		signature := ed25519.Sign(privateKey.(ed25519.PrivateKey), []byte(payload))
		return base64.StdEncoding.EncodeToString(signature)
	default:
		log.Println(fmt.Errorf("unsupported signing algorithm"))
		return ""
	}
}
