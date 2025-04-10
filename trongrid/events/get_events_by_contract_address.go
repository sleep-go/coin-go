package events

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type GetEventsByContractAddressResp struct {
	base.Msg
	Data []struct {
		BlockNumber           int    `json:"block_number"`
		BlockTimestamp        int64  `json:"block_timestamp"`
		CallerContractAddress string `json:"caller_contract_address"`
		ContractAddress       string `json:"contract_address"`
		EventIndex            int    `json:"event_index"`
		EventName             string `json:"event_name"`
		Result                struct {
			Field1        string `json:"0"`
			Field2        string `json:"1"`
			Field3        string `json:"2,omitempty"`
			From          string `json:"from,omitempty"`
			To            string `json:"to,omitempty"`
			Value         string `json:"value,omitempty"`
			PreviousOwner string `json:"previousOwner,omitempty"`
			NewOwner      string `json:"newOwner,omitempty"`
		} `json:"result"`
		ResultType struct {
			From          string `json:"from,omitempty"`
			To            string `json:"to,omitempty"`
			Value         string `json:"value,omitempty"`
			PreviousOwner string `json:"previousOwner,omitempty"`
			NewOwner      string `json:"newOwner,omitempty"`
		} `json:"result_type"`
		Event         string `json:"event"`
		TransactionId string `json:"transaction_id"`
	} `json:"data"`
	Meta struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}
type GetEventsByContractAddressReq struct {
	Address           string
	EventName         string
	BlockNumber       int64
	OnlyUnconfirmed   bool
	OnlyConfirmed     bool
	MinBlockTimestamp *time.Time
	MaxBlockTimestamp *time.Time
	OrderBy           string
	Fingerprint       string
	Limit             int32
}

func (e *Events) GetEventsByContractAddress(req *GetEventsByContractAddressReq) (*GetEventsByContractAddressResp, error) {
	values := url.Values{}
	values.Add("event_name", req.EventName)
	if req.BlockNumber != 0 {
		values.Add("block_number", fmt.Sprintf("%v", req.BlockNumber))
	}
	values.Set("only_confirmed", fmt.Sprint(req.OnlyConfirmed))
	values.Set("only_unconfirmed", fmt.Sprint(req.OnlyUnconfirmed))
	if req.MinBlockTimestamp != nil {
		values.Add("min_block_timestamp", fmt.Sprintf("%v", req.MinBlockTimestamp.UnixMilli()))
	}
	if req.MaxBlockTimestamp != nil {
		values.Add("max_block_timestamp", fmt.Sprintf("%v", req.MaxBlockTimestamp.UnixMilli()))
	}
	values.Set("order_by", req.OrderBy)
	values.Set("fingerprint", req.Fingerprint)
	if req.Limit != 0 {
		values.Add("limit", fmt.Sprintf("%d", req.Limit))
	}
	path := fmt.Sprintf("/v1/contracts/%s/events", req.Address)
	response, err := e.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetEventsByContractAddressResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
