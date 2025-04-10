package accounts

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type Accounts struct {
	Client *base.Client
}
type GetAccountInfoByAddressResp struct {
	base.Msg
	Data []struct {
		LatestOprationTime int64 `json:"latest_opration_time"`
		OwnerPermission    struct {
			Keys []struct {
				Address string `json:"address"`
				Weight  int    `json:"weight"`
			} `json:"keys"`
			Threshold      int    `json:"threshold"`
			PermissionName string `json:"permission_name"`
		} `json:"owner_permission"`
		AccountResource struct {
			EnergyWindowOptimized      bool  `json:"energy_window_optimized"`
			LatestConsumeTimeForEnergy int64 `json:"latest_consume_time_for_energy"`
			EnergyWindowSize           int   `json:"energy_window_size"`
		} `json:"account_resource"`
		ActivePermission []struct {
			Operations string `json:"operations"`
			Keys       []struct {
				Address string `json:"address"`
				Weight  int    `json:"weight"`
			} `json:"keys"`
			Threshold      int    `json:"threshold"`
			Id             int    `json:"id"`
			Type           string `json:"type"`
			PermissionName string `json:"permission_name"`
		} `json:"active_permission"`
		FrozenV2 []struct {
			Type string `json:"type,omitempty"`
		} `json:"frozenV2"`
		Address    string `json:"address"`
		Balance    int    `json:"balance"`
		CreateTime int64  `json:"create_time"`
		Trc20      []struct {
			TDLVXu6Mvt34KRRmHJtDc26BR99D7Eu7No string `json:"TDLVXu6mvt34kRRmHJtDc26bR99d7eu7No"`
		} `json:"trc20"`
		LatestConsumeFreeTime int64 `json:"latest_consume_free_time"`
		NetWindowSize         int   `json:"net_window_size"`
		NetWindowOptimized    bool  `json:"net_window_optimized"`
	} `json:"data"`
	Meta struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}
type GetAccountInfoByAddressReq struct {
	Address         string //owner address in base58 or hex
	OnlyConfirmed   bool   //true (If no param is specified, then only confirmed)
	OnlyUnconfirmed bool   //true (If no param is specified, then only confirmed)
}

func (a *Accounts) GetAccountInfoByAddress(req *GetAccountInfoByAddressReq) (*GetAccountInfoByAddressResp, error) {
	values := url.Values{}
	values.Add("only_confirmed", fmt.Sprintf("%v", req.OnlyConfirmed))
	values.Add("only_unconfirmed", fmt.Sprintf("%v", req.OnlyUnconfirmed))
	path := fmt.Sprintf("/v1/accounts/%s", req.Address)
	response, err := a.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetAccountInfoByAddressResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
