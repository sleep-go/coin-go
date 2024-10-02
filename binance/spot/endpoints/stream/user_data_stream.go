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
	SetListenKey(listenKey string) UserDataStream
	CallCreate(ctx context.Context) (body *userDataStreamResponse, err error)
	CallUpdate(ctx context.Context) (err error)
	CallDelete(ctx context.Context) (err error)
}
type userDataStreamRequest struct {
	*binance.Client
	listenKey string
}

type userDataStreamResponse struct {
	consts.ErrorResponse
	ListenKey string `json:"listenKey"` //用于订阅的数据流名
}

func NewUserDataStream(client *binance.Client) UserDataStream {
	return &userDataStreamRequest{Client: client}
}

func (o *userDataStreamRequest) SetListenKey(listenKey string) UserDataStream {
	o.listenKey = listenKey
	return o
}

// CallCreate 新建用户数据流 (USER_STREAM)
// 从创建起60分钟有效
func (o *userDataStreamRequest) CallCreate(ctx context.Context) (body *userDataStreamResponse, err error) {
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

// CallUpdate 延长用户数据流有效期到60分钟之后。 建议每30分钟调用一次
func (o *userDataStreamRequest) CallUpdate(ctx context.Context) (err error) {
	req := &binance.Request{
		Method: http.MethodPut,
		Path:   consts.ApiStreamUserDataStream,
	}
	req.SetParam("listenKey", o.listenKey)
	_, err = o.Do(ctx, req)
	if err != nil {
		o.Debugf("userDataStreamRequest response err:%v", err)
		return err
	}
	return nil
}

// CallDelete 关闭用户数据流 (USER_STREAM)
func (o *userDataStreamRequest) CallDelete(ctx context.Context) (err error) {
	req := &binance.Request{
		Method: http.MethodDelete,
		Path:   consts.ApiStreamUserDataStream,
	}
	req.SetParam("listenKey", o.listenKey)
	_, err = o.Do(ctx, req)
	if err != nil {
		o.Debugf("userDataStreamRequest response err:%v", err)
		return err
	}
	return nil
}
