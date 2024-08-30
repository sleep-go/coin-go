package consts

import (
	"net/http"
	"testing"
)

func TestOrder(t *testing.T) {
	g := &General{
		BaseURL:    TESTNET,
		HTTPClient: http.DefaultClient,
		Debug:      true,
	}
	g.Ping()
}
