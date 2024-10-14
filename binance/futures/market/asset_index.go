package market

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type AssetIndex interface {
	Call(ctx context.Context, symbol string) (body *assetIndexResponse, err error)
	CallAll(ctx context.Context) (body []*assetIndexResponse, err error)
}

type assetIndexRequest struct {
	*binance.Client
}

// NewAssetIndex Call 最新标记价格和资金费率
// 采集各大交易所数据加权平均
func NewAssetIndex(client *binance.Client) AssetIndex {
	return &assetIndexRequest{
		Client: client,
	}
}

type assetIndexResponse struct {
	Symbol                string `json:"symbol"`
	Time                  int64  `json:"time"`
	Index                 string `json:"index"`
	BidBuffer             string `json:"bidBuffer"`
	AskBuffer             string `json:"askBuffer"`
	BidRate               string `json:"bidRate"`
	AskRate               string `json:"askRate"`
	AutoExchangeBidBuffer string `json:"autoExchangeBidBuffer"`
	AutoExchangeAskBuffer string `json:"autoExchangeAskBuffer"`
	AutoExchangeBidRate   string `json:"autoExchangeBidRate"`
	AutoExchangeAskRate   string `json:"autoExchangeAskRate"`
}

func (t *assetIndexRequest) Call(ctx context.Context, symbol string) (body *assetIndexResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketAssetIndex,
	}
	req.SetParam("symbol", symbol)
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*assetIndexResponse](resp)
}
func (t *assetIndexRequest) CallAll(ctx context.Context) (body []*assetIndexResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiMarketAssetIndex,
	}
	resp, err := t.Do(ctx, req)
	if err != nil {
		t.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*assetIndexResponse](resp)
}

// ****************************** Websocket 行情推送 *******************************

// ****************************** Websocket Api *******************************
