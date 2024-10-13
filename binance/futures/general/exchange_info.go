package general

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance/futures/enums"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type ExchangeInfo interface {
	Call(ctx context.Context) (body *exchangeInfoResponse, err error)
}
type exchangeInfoRequest struct {
	*binance.Client
}

func NewExchangeInfo(client *binance.Client) ExchangeInfo {
	return &exchangeInfoRequest{
		Client: client,
	}
}

type exchangeInfoResponse struct {
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	RateLimits      []struct {
		Interval      string              `json:"interval"`
		IntervalNum   int                 `json:"intervalNum"`
		Limit         int                 `json:"limit"`
		RateLimitType enums.RateLimitType `json:"rateLimitType"`
	} `json:"rateLimits"`
	ServerTime int64 `json:"serverTime"`
	Assets     []struct {
		Asset             string  `json:"asset"`
		MarginAvailable   bool    `json:"marginAvailable"`
		AutoAssetExchange *string `json:"autoAssetExchange"`
	} `json:"assets"`
	Symbols []struct {
		Symbol                string             `json:"symbol"`
		Pair                  string             `json:"pair"`
		ContractType          enums.ContractType `json:"contractType"`
		DeliveryDate          int64              `json:"deliveryDate"`
		OnboardDate           int64              `json:"onboardDate"`
		Status                enums.StatusType   `json:"status"`
		MaintMarginPercent    string             `json:"maintMarginPercent"`
		RequiredMarginPercent string             `json:"requiredMarginPercent"`
		BaseAsset             string             `json:"baseAsset"`
		QuoteAsset            string             `json:"quoteAsset"`
		MarginAsset           string             `json:"marginAsset"`
		PricePrecision        int                `json:"pricePrecision"`
		QuantityPrecision     int                `json:"quantityPrecision"`
		BaseAssetPrecision    int                `json:"baseAssetPrecision"`
		QuotePrecision        int                `json:"quotePrecision"`
		UnderlyingType        string             `json:"underlyingType"`
		UnderlyingSubType     []string           `json:"underlyingSubType"`
		SettlePlan            int                `json:"settlePlan"`
		TriggerProtect        string             `json:"triggerProtect"`
		Filters               []struct {
			FilterType        string  `json:"filterType"`
			MaxPrice          string  `json:"maxPrice,omitempty"`
			MinPrice          string  `json:"minPrice,omitempty"`
			TickSize          string  `json:"tickSize,omitempty"`
			MaxQty            string  `json:"maxQty,omitempty"`
			MinQty            string  `json:"minQty,omitempty"`
			StepSize          string  `json:"stepSize,omitempty"`
			Limit             int     `json:"limit,omitempty"`
			Notional          string  `json:"notional,omitempty"`
			MultiplierUp      string  `json:"multiplierUp,omitempty"`
			MultiplierDown    string  `json:"multiplierDown,omitempty"`
			MultiplierDecimal *string `json:"multiplierDecimal,omitempty"`
		} `json:"filters"`
		OrderType       []enums.OrderType       `json:"OrderType"`
		TimeInForce     []enums.TimeInForceType `json:"timeInForce"`
		LiquidationFee  string                  `json:"liquidationFee"`
		MarketTakeBound string                  `json:"marketTakeBound"`
	} `json:"symbols"`
	Timezone string `json:"timezone"`
}

// Call
// 备注:
// 如果参数 symbol 或者 symbols 提供的交易对不存在, 系统会返回错误并提示交易对不正确.
// 所有的参数都是可选的.
// permissions 支持单个或者多个值, 比如 SPOT, ["MARGIN","LEVERAGED"].
// 如果permissions值没有提供, 其默认值为 ["SPOT","MARGIN","LEVERAGED"].
// 如果想显示所有交易权限，需要分别指定(比如，["SPOT","MARGIN",...]). 从 账户与交易对权限 查看交易权限列表.
func (ex *exchangeInfoRequest) Call(ctx context.Context) (body *exchangeInfoResponse, err error) {
	r := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.FApiExchangeInfo,
	}
	resp, err := ex.Do(ctx, r)
	if err != nil {
		ex.Debugf("response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*exchangeInfoResponse](resp)
}
