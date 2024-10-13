package stream

import (
	"context"
	"net/http"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/pkg/utils"
)

//	新建用户数据流 (USER_STREAM)
//
// 从创建起60分钟有效
type UserDataStream interface {
	SetListenKey(listenKey string) *userDataStreamRequest
	CallCreate(ctx context.Context) (body *userDataStreamResponse, err error)
	CallUpdate(ctx context.Context) (err error)
	CallDelete(ctx context.Context) (err error)
}
type userDataStreamRequest struct {
	*binance.Client
	listenKey string
}

type userDataStreamResponse struct {
	ListenKey string `json:"listenKey"` //用于订阅的数据流名
}

func NewUserDataStream(client *binance.Client) UserDataStream {
	return &userDataStreamRequest{Client: client}
}

func (o *userDataStreamRequest) SetListenKey(listenKey string) *userDataStreamRequest {
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
	return utils.ParseHttpResponse[*userDataStreamResponse](resp)
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

// ****************************** Websocket Api *******************************

type WsApiUserDataStream interface {
	UserDataStream
	SendStart(ctx context.Context) (*WsApiUserDataStreamResponse, error)
	SendPing(ctx context.Context) (*WsApiUserDataStreamResponse, error)
	SendStop(ctx context.Context) (*WsApiUserDataStreamResponse, error)
}
type WsApiUserDataStreamResponse struct {
	binance.WsApiResponse
	Result *userDataStreamResponse `json:"result"`
}

func NewWsApiUserDataStream(c *binance.Client) WsApiUserDataStream {
	return &userDataStreamRequest{Client: c}
}

func (o *userDataStreamRequest) SendStart(ctx context.Context) (*WsApiUserDataStreamResponse, error) {
	req := &binance.Request{Path: "userDataStream.start"}
	return binance.WsApiHandler[*WsApiUserDataStreamResponse](ctx, o.Client, req)
}
func (o *userDataStreamRequest) SendPing(ctx context.Context) (*WsApiUserDataStreamResponse, error) {
	req := &binance.Request{Path: "userDataStream.ping"}
	req.SetParam("listenKey", o.listenKey)
	return binance.WsApiHandler[*WsApiUserDataStreamResponse](ctx, o.Client, req)
}
func (o *userDataStreamRequest) SendStop(ctx context.Context) (*WsApiUserDataStreamResponse, error) {
	req := &binance.Request{Path: "userDataStream.stop"}
	req.SetParam("listenKey", o.listenKey)
	return binance.WsApiHandler[*WsApiUserDataStreamResponse](ctx, o.Client, req)
}
