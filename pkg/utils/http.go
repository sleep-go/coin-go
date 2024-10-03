package utils

import (
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/pkg/errors"
)

func ParseHttpResponse[T any](resp *http.Response) (body T, err error) {
	if resp.StatusCode != http.StatusOK {
		var e *errors.Error
		err = netutil.ParseHttpResponse(resp, &e)
		return body, e
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		return body, err
	}
	return body, nil
}
