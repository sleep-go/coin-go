package general

import (
	"context"
	"log"
	"net/http"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
)

type PingRequest struct {
	*binance.Client
	log *log.Logger
}
type PingResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func (p *PingRequest) Do(ctx context.Context) (*PingResponse, error) {
	r := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiPing,
	}
	res, err := p.Client.Do(ctx, r)
	if err != nil {
		p.log.Println("PingRequest response err:", err)
		return nil, err
	}
	defer res.Body.Close()
	return &PingResponse{
		Status: res.Status,
		Code:   res.StatusCode,
	}, nil
}
