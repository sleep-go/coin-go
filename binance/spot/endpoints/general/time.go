package general

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/sleep-go/exchange-go/binance"

	"github.com/sleep-go/exchange-go/binance/consts"
)

type TimeRequest struct {
	*binance.Client
	log *log.Logger
}

func (t *TimeRequest) Do(ctx context.Context) (body *TimeResponse, err error) {
	r := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiTime,
	}
	res, err := t.Client.Do(ctx, r)
	if err != nil {
		t.log.Println("PingRequest response err:", err)
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

type TimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}
