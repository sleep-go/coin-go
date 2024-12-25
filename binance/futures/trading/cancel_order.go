package trading

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type CancelOrder interface {
	SetSymbol(symbol string) CancelOrder
	CallCountdownCancelAll(ctx context.Context, countdownTime uint64) (body *cancelOrderResponse, err error)
}

type cancelOrderRequest struct {
	*binance.Client
	symbol        string
	countdownTime uint64
}
type cancelOrderResponse struct {
	Symbol        string `json:"symbol"`
	CountdownTime string `json:"countdownTime"`
}

func NewCancelOrder(client *binance.Client, symbol string) CancelOrder {
	return &cancelOrderRequest{Client: client, symbol: symbol}
}

func (d *cancelOrderRequest) SetSymbol(symbol string) CancelOrder {
	d.symbol = symbol
	return d
}
func (d *cancelOrderRequest) SetCountdownTime(symbol string) CancelOrder {
	d.symbol = symbol
	return d
}

// CallCountdownCancelAll 倒计时撤销所有订单 (TRADE)
//
// 接口描述
// 该接口可以被用于确保在倒计时结束时撤销指定symbol上的所有挂单。 在使用这个功能时，接口应像心跳一样在倒计时内被反复调用，以便可以取消既有的倒计时并开始新的倒数计时设置。
//
// 用法示例： 以30s的间隔重复此接口，每次倒计时countdownTime设置为120000(120s)。
// 如果在120秒内未再次调用此接口，则您指定symbol上的所有挂单都会被自动撤销。
// 如果在120秒内以将countdownTime设置为0，则倒数计时器将终止，自动撤单功能取消。
//
// 系统会大约每10毫秒检查一次所有倒计时情况，因此请注意，使用此功能时应考虑足够的冗余。
// 我们不建议将倒记时设置得太精确或太小。
// countdownTime	LONG	YES	倒计时。 1000 表示 1 秒； 0 表示取消倒计时撤单功能。
func (d *cancelOrderRequest) CallCountdownCancelAll(ctx context.Context, countdownTime uint64) (body *cancelOrderResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.FApiCountdownCancelAll,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", d.symbol)
	req.SetParam("countdownTime", countdownTime)
	resp, err := d.Do(ctx, req)
	if err != nil {
		d.Debugf("CallCountdownCancelAll response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[*cancelOrderResponse](resp)
}

// ****************************** Websocket Api *******************************
