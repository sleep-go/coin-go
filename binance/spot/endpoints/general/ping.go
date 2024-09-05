package general

import (
	"context"
	"fmt"
	"io"
	"net/http"

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

func (p *pingRequest) Call(ctx context.Context) (*pingResponse, error) {
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
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", bytes)
	}
	return &pingResponse{
		Status: res.Status,
		Code:   res.StatusCode,
	}, nil
}
