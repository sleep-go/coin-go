package assets

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/sleep-go/coin-go/tron/base"
)

type GetAssetsByNameResp struct {
	base.Msg
	Data []struct {
		Id           int    `json:"id"`
		Abbr         string `json:"abbr"`
		Description  string `json:"description"`
		Name         string `json:"name"`
		Num          int    `json:"num"`
		Precision    int    `json:"precision"`
		Url          string `json:"url"`
		TotalSupply  int    `json:"total_supply"`
		TrxNum       int    `json:"trx_num"`
		VoteScore    int    `json:"vote_score"`
		OwnerAddress string `json:"owner_address"`
		StartTime    int64  `json:"start_time"`
		EndTime      int64  `json:"end_time"`
	} `json:"data"`
	Meta struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}
type GetAssetsByNameReq struct {
	Name          string //name of the asset(s)
	Limit         int32  //number of assets per page, default 20, max 200
	Fingerprint   string //fingerprint of the last asset returned by the previous page; when using it, the other parameters and filters should remain the same
	OrderBy       string //order_by = total_supply,asc | total_supply,desc | start_time,asc | start_time,desc | end_time,asc | end_time,desc | id,asc | id,desc (default)
	OnlyConfirmed bool   //true | false. If false, it returns both confirmed and unconfirmed assets.
}

func (a *Assets) GetAssetsByName(req *GetAssetsByNameReq) (*GetAssetsByNameResp, error) {
	values := url.Values{}
	if req.Limit != 0 {
		values.Add("limit", strconv.Itoa(int(req.Limit)))
	}
	if req.Fingerprint != "" {
		values.Add("fingerprint", req.Fingerprint)
	}
	if req.OrderBy != "" {
		values.Add("order_by", req.OrderBy)
	}
	values.Add("only_confirmed", fmt.Sprintf("%v", req.OnlyConfirmed))
	path := fmt.Sprintf("/v1/assets/%s/list", req.Name)
	response, err := a.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetAssetsByNameResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
