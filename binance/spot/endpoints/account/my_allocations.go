package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type MyAllocations interface {
	SetStartTime(startTime uint64) MyAllocations
	SetEndTime(endTime uint64) MyAllocations
	SetFromAllocationId(fromAllocationId int64) MyAllocations
	SetOrderId(orderId int64) MyAllocations
	SetRecvWindow(recvWindow int64) MyAllocations
	SetTimestamp(timestamp int64) MyAllocations
	Call(ctx context.Context) (body []*myAllocationsResponse, err error)
}

type myAllocationsRequest struct {
	*binance.Client
	symbol           string
	startTime        *uint64
	endTime          *uint64
	fromAllocationId *int64
	limit            enums.LimitType
	orderId          *int64
	recvWindow       int64
	timestamp        int64
}

func NewMyAllocations(client *binance.Client, symbol string, limit enums.LimitType) MyAllocations {
	return &myAllocationsRequest{Client: client, symbol: symbol, limit: limit}
}

type myAllocationsResponse struct {
	Symbol          string `json:"symbol"`
	AllocationId    int    `json:"allocationId"`
	AllocationType  string `json:"allocationType"`
	OrderId         int    `json:"orderId"`
	OrderListId     int    `json:"orderListId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsAllocator     bool   `json:"isAllocator"`
}

func (m *myAllocationsRequest) SetStartTime(startTime uint64) MyAllocations {
	m.startTime = &startTime
	return m
}

func (m *myAllocationsRequest) SetEndTime(endTime uint64) MyAllocations {
	m.endTime = &endTime
	return m
}

func (m *myAllocationsRequest) SetFromAllocationId(fromAllocationId int64) MyAllocations {
	m.fromAllocationId = &fromAllocationId
	return m
}

func (m *myAllocationsRequest) SetOrderId(orderId int64) MyAllocations {
	m.orderId = &orderId
	return m
}

func (m *myAllocationsRequest) SetRecvWindow(recvWindow int64) MyAllocations {
	m.recvWindow = recvWindow
	return m
}

func (m *myAllocationsRequest) SetTimestamp(timestamp int64) MyAllocations {
	m.timestamp = timestamp
	return m
}

// Call 检索由 SOR 订单生成引起的分配结果。
// 支持的参数组合:
//
// 参数	响应
// symbol	按从最旧到最新排序的分配
// symbol + startTime	从 startTime 开始的最旧的分配
// symbol + endTime	到 endTime 为止的最新的分配
// symbol + startTime + endTime	在指定时间范围内的分配
// symbol + fromAllocationId	从指定 AllocationId 开始的分配
// symbol + orderId	按从最旧到最新排序并和特定订单关联的分配
// symbol + orderId + fromAllocationId	从指定 AllocationId 开始并和特定订单关联的分配
// 注意: startTime 和 endTime 之间的时间不能超过 24 小时。
func (m *myAllocationsRequest) Call(ctx context.Context) (body []*myAllocationsResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountMyAllocations,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetParam("limit", m.limit)
	if m.startTime != nil {
		req.SetParam("startTime", m.startTime)
	}
	if m.endTime != nil {
		req.SetParam("endTime", m.endTime)
	}
	if m.fromAllocationId != nil {
		req.SetParam("fromAllocationId", m.fromAllocationId)
	}
	if m.orderId != nil {
		req.SetParam("orderId", m.orderId)
	}
	if m.recvWindow > 0 {
		req.SetParam("recvWindow", m.recvWindow)
	}
	req.SetParam("timestamp", m.timestamp)
	resp, err := m.Do(ctx, req)
	if err != nil {
		m.Debugf("myAllocationsRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*myAllocationsResponse](resp)
}
