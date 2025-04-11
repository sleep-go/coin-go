package contracts

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type Contracts struct {
	Client *base.Client
}

type GetTransactionInfoByContractAddressResp struct {
	base.Msg
	Data []struct {
		Ret []struct {
			ContractRet string `json:"contractRet"`
		} `json:"ret"`
		Signature        []string `json:"signature"`
		TxID             string   `json:"txID"`
		NetUsage         int      `json:"net_usage"`
		RawDataHex       string   `json:"raw_data_hex"`
		NetFee           int      `json:"net_fee"`
		EnergyUsage      int      `json:"energy_usage"`
		BlockTimestamp   string   `json:"block_timestamp"`
		BlockNumber      string   `json:"blockNumber"`
		EnergyFee        int      `json:"energy_fee"`
		EnergyUsageTotal int      `json:"energy_usage_total"`
		RawData          struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Data            string `json:"data,omitempty"`
						OwnerAddress    string `json:"owner_address"`
						ContractAddress string `json:"contract_address,omitempty"`
						NewContract     struct {
							Bytecode                   string `json:"bytecode"`
							ConsumeUserResourcePercent int    `json:"consume_user_resource_percent"`
							Name                       string `json:"name"`
							OriginAddress              string `json:"origin_address"`
							Abi                        struct {
								Entrys []struct {
									Inputs []struct {
										Name    string `json:"name"`
										Type    string `json:"type"`
										Indexed bool   `json:"indexed,omitempty"`
									} `json:"inputs,omitempty"`
									StateMutability string `json:"stateMutability,omitempty"`
									Type            string `json:"type"`
									Name            string `json:"name,omitempty"`
									Outputs         []struct {
										Type string `json:"type"`
									} `json:"outputs,omitempty"`
								} `json:"entrys"`
							} `json:"abi"`
							OriginEnergyLimit int `json:"origin_energy_limit"`
						} `json:"new_contract,omitempty"`
					} `json:"value"`
					TypeUrl string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			FeeLimit      int    `json:"fee_limit"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		InternalTransactions []interface{} `json:"internal_transactions"`
	} `json:"data"`
	Meta struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}
type GetTransactionInfoByContractAddressReq struct {
	ContractAddress   string
	OnlyConfirmed     bool
	OnlyUnconfirmed   bool
	MinBlockTimestamp *time.Time //Minimal block timestamp
	MaxBlockTimestamp *time.Time //Maximal block timestamp
	OrderBy           string
	Fingerprint       string
	Limit             int32
	SearchInternal    bool
}

func (c *Contracts) GetTransactionInfoByContractAddress(req *GetTransactionInfoByContractAddressReq) (*GetTransactionInfoByContractAddressResp, error) {
	values := url.Values{}
	values.Set("only_confirmed", fmt.Sprint(req.OnlyConfirmed))
	values.Set("only_unconfirmed", fmt.Sprint(req.OnlyUnconfirmed))
	if req.MinBlockTimestamp != nil {
		values.Set("min_block_timestamp", fmt.Sprintf("%d", req.MinBlockTimestamp.UnixMilli()))
	}
	if req.MaxBlockTimestamp != nil {
		values.Set("max_block_timestamp", fmt.Sprintf("%d", req.MaxBlockTimestamp.UnixMilli()))
	}
	if req.OrderBy != "" {
		values.Set("order_by", req.OrderBy)
	}
	if req.Fingerprint != "" {
		values.Set("fingerprint", req.Fingerprint)
	}
	if req.Limit != 0 {
		values.Set("limit", strconv.Itoa(int(req.Limit)))
	}
	values.Set("search_internal", fmt.Sprint(req.SearchInternal))
	path := fmt.Sprintf("/v1/contracts/%s/transactions", req.ContractAddress)
	response, err := c.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetTransactionInfoByContractAddressResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
