package utils

import (
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/pkg/errors"
)

func ParseHttpResponse[T any](resp *http.Response) (body *T, err error) {
	if resp.StatusCode != http.StatusOK {
		var e *errors.Error
		err = netutil.ParseHttpResponse(resp, &e)
		return nil, e
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func ParseHttpSliceResponse[T any](resp *http.Response) (body []*T, err error) {
	if resp.StatusCode != http.StatusOK {
		var e *errors.Error
		err = netutil.ParseHttpResponse(resp, &e)
		return nil, e
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
