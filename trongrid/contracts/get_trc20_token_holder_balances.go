package contracts

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type GetTrc20TokenHolderBalancesResp struct {
	base.Msg
	Data []struct {
		TGNBhSEXcaxYcsFavVyZEbuWPqtz7MNACF string `json:"TGNBhSEXcaxYcsFavVyZEbuWPqtz7mNACF,omitempty"`
		TY4YmYKXRzA1SkRYCVe1QyvyV97UDtWuBB string `json:"TY4ymYKXRzA1SkRYCVe1qyvyV97UDtWuBB,omitempty"`
	} `json:"data"`
	Meta struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}
type GetTrc20TokenHolderBalancesReq struct {
	ContractAddress string
	OnlyConfirmed   bool
	OnlyUnconfirmed bool
	OrderBy         string
	Fingerprint     string
	Limit           int32
}

func (c *Contracts) GetTrc20TokenHolderBalances(req *GetTrc20TokenHolderBalancesReq) (*GetTrc20TokenHolderBalancesResp, error) {
	values := url.Values{}
	values.Set("only_confirmed", fmt.Sprint(req.OnlyConfirmed))
	values.Set("only_unconfirmed", fmt.Sprint(req.OnlyUnconfirmed))
	if req.OrderBy != "" {
		values.Add("order_by", req.OrderBy)
	}
	if req.Fingerprint != "" {
		values.Add("fingerprint", req.Fingerprint)
	}
	if req.Limit != 0 {
		values.Add("limit", strconv.Itoa(int(req.Limit)))
	}
	path := fmt.Sprintf("/v1/contracts/%s/tokens", req.ContractAddress)
	response, err := c.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetTrc20TokenHolderBalancesResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
