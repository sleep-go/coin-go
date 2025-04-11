package assets

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type GetAssetByIdOrIssuerResp struct {
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
type GetAssetByIdOrIssuerReq struct {
	Identifier    string //id of the asset or the owner address of the asset in base58 or hex
	OnlyConfirmed bool   //true | false. If false (default), it returns both confirmed and unconfirmed transactions.
}

func (a *Assets) GetAssetByIdOrIssuer(req *GetAssetByIdOrIssuerReq) (*GetAssetByIdOrIssuerResp, error) {
	values := url.Values{}
	values.Set("only_confirmed", fmt.Sprintf("%v", req.OnlyConfirmed))
	path := fmt.Sprintf("/v1/assets/%s", req.Identifier)
	response, err := a.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetAssetByIdOrIssuerResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
