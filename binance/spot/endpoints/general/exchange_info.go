package general

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/sleep-go/exchange-go/binance"
	"github.com/sleep-go/exchange-go/binance/consts"
)

type ExchangeInfoRequest struct {
	*binance.Client
	log *log.Logger
}

// ExchangeInfoResponse 解释响应中的 permissionSets：
// [["A","B"]] - 有权限"A"或权限"B"的账户可以下订单。
// [["A"],["B"]] - 有权限"A"和权限"B"的账户可以下订单。
// [["A"],["B","C"]] - 有权限"A"和权限"B"或权限"C"的账户可以下订单。（此处应用的是包含或，而不是排除或，因此账户可以同时拥有权限"B"和权限"C"。）
// 数据源: 缓存
type ExchangeInfoResponse struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol                     string   `json:"symbol"`
		Status                     string   `json:"status"`
		BaseAsset                  string   `json:"baseAsset"`
		BaseAssetPrecision         int      `json:"baseAssetPrecision"`
		QuoteAsset                 string   `json:"quoteAsset"`
		QuotePrecision             int      `json:"quotePrecision"`
		QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
		BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
		QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
		OrderTypes                 []string `json:"orderTypes"`
		IcebergAllowed             bool     `json:"icebergAllowed"`
		OcoAllowed                 bool     `json:"ocoAllowed"`
		OtoAllowed                 bool     `json:"otoAllowed"`
		QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
		AllowTrailingStop          bool     `json:"allowTrailingStop"`
		CancelReplaceAllowed       bool     `json:"cancelReplaceAllowed"`
		IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
		IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
		Filters                    []struct {
			FilterType            string `json:"filterType"`
			MinPrice              string `json:"minPrice,omitempty"`
			MaxPrice              string `json:"maxPrice,omitempty"`
			TickSize              string `json:"tickSize,omitempty"`
			MinQty                string `json:"minQty,omitempty"`
			MaxQty                string `json:"maxQty,omitempty"`
			StepSize              string `json:"stepSize,omitempty"`
			Limit                 int    `json:"limit,omitempty"`
			MinTrailingAboveDelta int    `json:"minTrailingAboveDelta,omitempty"`
			MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta,omitempty"`
			MinTrailingBelowDelta int    `json:"minTrailingBelowDelta,omitempty"`
			MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta,omitempty"`
			BidMultiplierUp       string `json:"bidMultiplierUp,omitempty"`
			BidMultiplierDown     string `json:"bidMultiplierDown,omitempty"`
			AskMultiplierUp       string `json:"askMultiplierUp,omitempty"`
			AskMultiplierDown     string `json:"askMultiplierDown,omitempty"`
			AvgPriceMins          int    `json:"avgPriceMins,omitempty"`
			MinNotional           string `json:"minNotional,omitempty"`
			ApplyMinToMarket      bool   `json:"applyMinToMarket,omitempty"`
			MaxNotional           string `json:"maxNotional,omitempty"`
			ApplyMaxToMarket      bool   `json:"applyMaxToMarket,omitempty"`
			MaxNumOrders          int    `json:"maxNumOrders,omitempty"`
			MaxNumAlgoOrders      int    `json:"maxNumAlgoOrders,omitempty"`
		} `json:"filters"`
		Permissions                     []interface{} `json:"permissions"`
		PermissionSets                  [][]string    `json:"permissionSets"`
		DefaultSelfTradePreventionMode  string        `json:"defaultSelfTradePreventionMode"`
		AllowedSelfTradePreventionModes []string      `json:"allowedSelfTradePreventionModes"`
	} `json:"symbols"`
}

// Do ExchangeInfoRequest
// 备注:
// 如果参数 symbol 或者 symbols 提供的交易对不存在, 系统会返回错误并提示交易对不正确.
// 所有的参数都是可选的.
// permissions 支持单个或者多个值, 比如 SPOT, ["MARGIN","LEVERAGED"].
// 如果permissions值没有提供, 其默认值为 ["SPOT","MARGIN","LEVERAGED"].
// 如果想显示所有交易权限，需要分别指定(比如，["SPOT","MARGIN",...]). 从 账户与交易对权限 查看交易权限列表.
func (ex *ExchangeInfoRequest) Do(ctx context.Context, symbols, permissions []string) (body *ExchangeInfoResponse, err error) {
	r := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiExchangeInfo,
	}
	if len(symbols) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(symbols, `","`))
		r.SetParam("symbols", result)
	}
	if len(permissions) > 0 {
		result := fmt.Sprintf(`["%s"]`, strings.Join(permissions, `","`))
		r.SetParam("permissions", result)
	}
	res, err := ex.Client.Do(ctx, r)
	if err != nil {
		ex.log.Println("ExchangeInfoRequest response err:", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
