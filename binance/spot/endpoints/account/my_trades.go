package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type MyTrades interface {
	SetOrderId(orderId int64) MyTrades
	SetStartTime(startTime uint64) MyTrades
	SetEndTime(endTime uint64) MyTrades
	SetFromId(fromId int64) MyTrades
	SetRecvWindow(recvWindow int64) MyTrades
	SetTimestamp(timestamp int64) MyTrades
	Call(ctx context.Context) (body []*myTradesResponse, err error)
}

// 备注:
//
// 如果设置了fromId, 会返回ID大于此fromId的交易. 不然则会返回最近的交易.
// startTime和endTime设置的时间间隔不能超过24小时.
// 支持的所有参数组合:
// symbol
// symbol + orderId
// symbol + startTime
// symbol + endTime
// symbol + fromId
// symbol + startTime + endTime
// symbol+ orderId + fromId
type myTradesRequest struct {
	*binance.Client
	symbol     string
	orderId    *int64 //必须要和参数symbol一起使用.
	startTime  *uint64
	endTime    *uint64
	fromId     *int64          //返回该fromId之后的成交，缺省返回最近的成交
	limit      enums.LimitType //Default 500; max 1000.
	recvWindow int64
	timestamp  int64
}

type myTradesResponse struct {
	Symbol          string `json:"symbol"`
	Id              int    `json:"id"`
	OrderId         int    `json:"fromId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

func NewMyTrades(client *binance.Client, symbol string, limit enums.LimitType) MyTrades {
	return &myTradesRequest{Client: client, symbol: symbol, limit: limit}
}

func (m *myTradesRequest) SetOrderId(orderId int64) MyTrades {
	m.orderId = &orderId
	return m
}

func (m *myTradesRequest) SetStartTime(startTime uint64) MyTrades {
	m.startTime = &startTime
	return m
}

func (m *myTradesRequest) SetEndTime(endTime uint64) MyTrades {
	m.endTime = &endTime
	return m
}

func (m *myTradesRequest) SetFromId(fromId int64) MyTrades {
	m.fromId = &fromId
	return m
}

func (m *myTradesRequest) SetRecvWindow(recvWindow int64) MyTrades {
	m.recvWindow = recvWindow
	return m
}

func (m *myTradesRequest) SetTimestamp(timestamp int64) MyTrades {
	m.timestamp = timestamp
	return m
}

func (m *myTradesRequest) Call(ctx context.Context) (body []*myTradesResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountMyTrades,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	if m.orderId != nil {
		req.SetParam("fromId", m.orderId)
	}
	if m.startTime != nil {
		req.SetParam("startTime", m.startTime)
	}
	if m.endTime != nil {
		req.SetParam("endTime", m.endTime)
	}
	if m.fromId != nil {
		req.SetParam("fromId", m.fromId)
	}
	if m.limit > 0 {
		req.SetParam("limit", m.limit)
	}
	if m.recvWindow > 0 {
		req.SetParam("recvWindow", m.recvWindow)
	}
	req.SetParam("timestamp", m.timestamp)
	resp, err := m.Do(ctx, req)
	if err != nil {
		m.Debugf("myTradesRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*myTradesResponse](resp)
}
