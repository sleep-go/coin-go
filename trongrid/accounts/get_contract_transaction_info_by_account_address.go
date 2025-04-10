package accounts

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type GetContractTransactionInfoByAccountAddressResp struct {
	base.Msg
	Data []struct {
		Ret []struct {
			ContractRet string `json:"contractRet"`
			Fee         int    `json:"fee"`
		} `json:"ret"`
		Signature        []string `json:"signature"`
		TxID             string   `json:"txID"`
		NetUsage         int      `json:"net_usage"`
		RawDataHex       string   `json:"raw_data_hex"`
		NetFee           int      `json:"net_fee"`
		EnergyUsage      int      `json:"energy_usage"`
		BlockNumber      int      `json:"blockNumber"`
		BlockTimestamp   int64    `json:"block_timestamp"`
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
						Amount    int    `json:"amount,omitempty"`
						ToAddress string `json:"to_address,omitempty"`
					} `json:"value"`
					TypeUrl string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			FeeLimit      int    `json:"fee_limit,omitempty"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		InternalTransactions []interface{} `json:"internal_transactions"`
	} `json:"data"`
	Meta struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}
type GetContractTransactionInfoByAccountAddressReq struct {
	Address         string     //owner address in base58 or hex
	OnlyConfirmed   bool       //true | false. If false, it returns both confirmed and unconfirmed transactions. If no param is specified, it returns both confirmed and unconfirmed transactions. Cannot be used at the same time with only_unconfirmed param.
	OnlyUnconfirmed bool       //true | false. If false, it returns both confirmed and unconfirmed transactions. If no param is specified, it returns both confirmed and unconfirmed transactions. Cannot be used at the same time with only_confirmed param.
	Limit           int32      //number of transactions per page, default 20, max 200
	Fingerprint     string     //fingerprint of the last transaction returned by the previous page; when using it, the other parameters and filters should remain the same
	OrderBy         string     //block_timestamp,asc | block_timestamp,desc (default)
	MinTimestamp    *time.Time //minimum block_timestamp, default 0
	MaxTimestamp    *time.Time //maximum block_timestamp, default now
	ContractAddress string     //contract address in base58 or hex
	OnlyTo          bool       //true | false. If true, only transactions to this address, default: false
	OnlyFrom        bool       //true | false. If true, only transactions from this address, default: false
}

// GetContractTransactionInfoByAccountAddress The same time window can get up to 1000 pieces of data. If you need to get more data, you can move the time window to get more data.
func (a *Accounts) GetContractTransactionInfoByAccountAddress(req *GetContractTransactionInfoByAccountAddressReq) (*GetContractTransactionInfoByAccountAddressResp, error) {
	values := url.Values{}
	values.Add("only_confirmed", fmt.Sprintf("%v", req.OnlyConfirmed))
	values.Add("only_unconfirmed", fmt.Sprintf("%v", req.OnlyUnconfirmed))
	if req.Limit != 0 {
		values.Add("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Fingerprint != "" {
		values.Add("fingerprint", req.Fingerprint)
	}
	values.Add("order_by", req.OrderBy)
	if req.MinTimestamp != nil {
		values.Add("min_timestamp", fmt.Sprintf("%d", req.MinTimestamp.UnixMilli()))
	}
	if req.MaxTimestamp != nil {
		values.Add("max_timestamp", fmt.Sprintf("%d", req.MaxTimestamp.UnixMilli()))
	}
	values.Add("contract_address", req.ContractAddress)
	values.Add("only_to", fmt.Sprintf("%v", req.OnlyTo))
	values.Add("only_from", fmt.Sprintf("%v", req.OnlyFrom))
	path := fmt.Sprintf("/v1/accounts/%s/transactions/trc20", req.Address)
	response, err := a.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetContractTransactionInfoByAccountAddressResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
