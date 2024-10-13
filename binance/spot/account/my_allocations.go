package account

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/spot/enums"
	"github.com/sleep-go/coin-go/pkg/utils"
)

type MyAllocations interface {
	SetSymbol(symbol string) *myAllocationsRequest
	SetLimit(limit enums.LimitType) *myAllocationsRequest
	SetStartTime(startTime uint64) *myAllocationsRequest
	SetEndTime(endTime uint64) *myAllocationsRequest
	SetFromAllocationId(fromAllocationId int64) *myAllocationsRequest
	SetOrderId(orderId int64) *myAllocationsRequest
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
}

func NewMyAllocations(client *binance.Client, symbol string, limit enums.LimitType) MyAllocations {
	return &myAllocationsRequest{Client: client, symbol: symbol, limit: limit}
}

type myAllocationsResponse struct {
	Symbol          string `json:"symbol"`
	AllocationId    int    `json:"allocationId"`
	AllocationType  string `json:"allocationType"`
	OrderId         int    `json:"fromId"`
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

func (m *myAllocationsRequest) SetSymbol(symbol string) *myAllocationsRequest {
	m.symbol = symbol
	return m
}

func (m *myAllocationsRequest) SetLimit(limit enums.LimitType) *myAllocationsRequest {
	m.limit = limit
	return m
}

func (m *myAllocationsRequest) SetStartTime(startTime uint64) *myAllocationsRequest {
	m.startTime = &startTime
	return m
}

func (m *myAllocationsRequest) SetEndTime(endTime uint64) *myAllocationsRequest {
	m.endTime = &endTime
	return m
}

func (m *myAllocationsRequest) SetFromAllocationId(fromAllocationId int64) *myAllocationsRequest {
	m.fromAllocationId = &fromAllocationId
	return m
}

func (m *myAllocationsRequest) SetOrderId(orderId int64) *myAllocationsRequest {
	m.orderId = &orderId
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
// symbol + fromId	按从最旧到最新排序并和特定订单关联的分配
// symbol + fromId + fromAllocationId	从指定 AllocationId 开始并和特定订单关联的分配
// 注意: startTime 和 endTime 之间的时间不能超过 24 小时。
func (m *myAllocationsRequest) Call(ctx context.Context) (body []*myAllocationsResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccountMyAllocations,
	}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetOptionalParam("limit", m.limit)
	req.SetOptionalParam("startTime", m.startTime)
	req.SetOptionalParam("endTime", m.endTime)
	req.SetOptionalParam("fromAllocationId", m.fromAllocationId)
	req.SetOptionalParam("orderId", m.orderId)
	resp, err := m.Do(ctx, req)
	if err != nil {
		m.Debugf("myAllocationsRequest response err:%v", err)
		return nil, err
	}
	return utils.ParseHttpResponse[[]*myAllocationsResponse](resp)
}

// ****************************** Websocket Api *******************************

type WsApiMyAllocations interface {
	binance.WsApi[*WsApiMyAllocationsResponse]
	MyAllocations
}
type WsApiMyAllocationsResponse struct {
	binance.WsApiResponse
	Result []*myAllocationsResponse `json:"result"`
}

// NewWsApiMyAllocations 查询分配结果 (USER_DATA)
// 检索由 SOR 订单生成引起的分配结果。
// 注意: startTime 和 endTime 之间的时间不能超过 24 小时。
func NewWsApiMyAllocations(c *binance.Client) WsApiMyAllocations {
	return &myAllocationsRequest{Client: c}
}

func (m *myAllocationsRequest) Send(ctx context.Context) (*WsApiMyAllocationsResponse, error) {
	req := &binance.Request{Path: "myAllocations"}
	req.SetNeedSign(true)
	req.SetParam("symbol", m.symbol)
	req.SetOptionalParam("orderId", m.orderId)
	req.SetOptionalParam("startTime", m.startTime)
	req.SetOptionalParam("endTime", m.endTime)
	req.SetOptionalParam("limit", m.limit)
	return binance.WsApiHandler[*WsApiMyAllocationsResponse](ctx, m.Client, req)
}
