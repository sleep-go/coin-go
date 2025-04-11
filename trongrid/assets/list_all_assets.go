package assets

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type Assets struct {
	Client *base.Client
}
type ListAllAssetsResp struct {
	base.Msg
	Data []struct {
		Id           int    `json:"id"`
		Abbr         string `json:"abbr"`
		Description  string `json:"description"`
		Name         string `json:"name"`
		Num          int    `json:"num"`
		Precision    int    `json:"precision"`
		Url          string `json:"url"`
		TotalSupply  int64  `json:"total_supply"`
		TrxNum       int    `json:"trx_num"`
		VoteScore    int    `json:"vote_score"`
		OwnerAddress string `json:"owner_address"`
		StartTime    int64  `json:"start_time"`
		EndTime      int64  `json:"end_time"`
	} `json:"data"`
	Meta struct {
		At          int64  `json:"at"`
		Fingerprint string `json:"fingerprint"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
		PageSize int `json:"page_size"`
	} `json:"meta"`
	Success bool `json:"success"`
}
type ListAllAssetsReq struct {
	OrderBy     string //order_by = total_supply,asc | total_supply,desc | start_time,asc | start_time,desc | end_time,asc | end_time,desc | id,asc | id,desc (default)
	Limit       int32  //number of assets per page, default 20, max 200
	Fingerprint string //fingerprint of the last asset returned by the previous page; when using it, the other parameters and filters should remain the same
}

func (a *Assets) ListAllAssets(req *ListAllAssetsReq) (*ListAllAssetsResp, error) {
	values := url.Values{}
	if req.OrderBy != "" {
		values.Set("order_by", req.OrderBy)
	}
	if req.Limit != 0 {
		values.Set("limit", strconv.Itoa(int(req.Limit)))
	}
	if req.Fingerprint != "" {
		values.Set("fingerprint", req.Fingerprint)
	}
	path := fmt.Sprintf("/v1/assets")
	response, err := a.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(ListAllAssetsResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
