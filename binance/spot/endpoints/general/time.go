package general

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
)

type Time interface {
	Call(ctx context.Context) (body *timeResponse, err error)
}

type timeRequest struct {
	*binance.Client
}

func NewTime(client *binance.Client) Time {
	return &timeRequest{Client: client}
}

func (t *timeRequest) Call(ctx context.Context) (body *timeResponse, err error) {
	r := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTime,
	}
	res, err := t.Client.Do(ctx, r)
	if err != nil {
		t.Debugf("pingRequest response err: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type timeResponse struct {
	ServerTime int64 `json:"serverTime"`
}
