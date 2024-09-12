package general

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
)

type Ping interface {
	Call(ctx context.Context) (*pingResponse, error)
}
type pingRequest struct {
	*binance.Client
}

func NewPing(c *binance.Client) Ping {
	return &pingRequest{Client: c}
}

type pingResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func (p *pingRequest) Call(ctx context.Context) (body *pingResponse, err error) {
	r := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiPing,
	}
	res, err := p.Client.Do(ctx, r)
	if err != nil {
		p.Debugf("pingRequest response err: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v", res)
	}
	err = netutil.ParseHttpResponse(res, &body)
	if err != nil {
		return nil, errors.New(res.Status)
	}
	return &pingResponse{
		Status: res.Status,
		Code:   res.StatusCode,
	}, nil
}
