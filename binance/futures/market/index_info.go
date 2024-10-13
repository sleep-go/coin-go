package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type IndexInfo interface {
	Call(ctx context.Context) (body *indexInfoResponse, err error)
}

type indexInfoRequest struct {
	*binance.Client
	symbol string
}

// NewIndexInfo 综合指数交易对信息
// 获取交易对为综合指数的基础成分信息。
func NewIndexInfo(client *binance.Client, symbol string) IndexInfo {
	return &indexInfoRequest{
		Client: client,
		symbol: symbol,
	}
}

type indexInfoResponse struct {
	Symbol        string `json:"symbol"`
	Time          int64  `json:"time"`      // 请求时间
	Component     string `json:"component"` //成分资产
	BaseAssetList []struct {
		BaseAsset          string `json:"baseAsset"`          // 基础资产
		QuoteAsset         string `json:"quoteAsset"`         // 报价资产
		WeightInQuantity   string `json:"weightInQuantity"`   //权重(数量)
		WeightInPercentage string `json:"weightInPercentage"` //权重(比例)
	} `json:"baseAssetList"`
}

func (t *indexInfoRequest) Call(ctx context.Context) (body *indexInfoResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketIndexInfo,
	}
	req.SetOptionalParam("symbol", t.symbol)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*indexInfoResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
