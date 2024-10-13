package general

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
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
	return utils.ParseHttpResponse[*timeResponse](res)
}

type timeResponse struct {
	ServerTime int64 `json:"serverTime"`
}
