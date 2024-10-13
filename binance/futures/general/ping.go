package general

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
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
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiPing,
	}
	resp, err := p.Do(ctx, req)
	if err != nil {
		p.Debugf("pingRequest response err: %v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*pingResponse](resp)
}
