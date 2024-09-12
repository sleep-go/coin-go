package general

import (
	"context"
	"errors"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
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
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	err = netutil.ParseHttpResponse(res, &body)
	if err != nil {
		t.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}

type timeResponse struct {
	ServerTime int64 `json:"serverTime"`
}
