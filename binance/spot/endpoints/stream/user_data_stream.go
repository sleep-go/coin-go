package stream

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
)

//	新建用户数据流 (USER_STREAM)
//
// 从创建起60分钟有效
type UserDataStream interface {
	SetRecvWindow(recvWindow int64) UserDataStream
	SetTimestamp(timestamp int64) UserDataStream
	Call(ctx context.Context) (body *userDataStreamResponse, err error)
}
type userDataStreamRequest struct {
	*binance.Client
	recvWindow int64
	timestamp  int64
}
type userDataStreamResponse struct {
	consts.ErrorResponse
	ListenKey string `json:"listenKey"` //用于订阅的数据流名
}

func NewUserDataStream(client *binance.Client) UserDataStream {
	return &userDataStreamRequest{Client: client}
}

func (o *userDataStreamRequest) SetRecvWindow(recvWindow int64) UserDataStream {
	o.recvWindow = recvWindow
	return o
}

func (o *userDataStreamRequest) SetTimestamp(timestamp int64) UserDataStream {
	o.timestamp = timestamp
	return o
}

// Call 新建用户数据流 (USER_STREAM)
// 从创建起60分钟有效
func (o *userDataStreamRequest) Call(ctx context.Context) (body *userDataStreamResponse, err error) {
	req := &binance.Request{
		Method: http.MethodPost,
		Path:   consts.ApiStreamUserDataStream,
	}
	resp, err := o.Do(ctx, req)
	if err != nil {
		o.Debugf("userDataStreamRequest response err:%v", err)
		return nil, err
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		o.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
