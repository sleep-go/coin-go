package events

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type GetEventsByBlockNumberResp struct {
	base.Msg
	Data []struct {
		BlockNumber           int    `json:"block_number"`
		BlockTimestamp        int64  `json:"block_timestamp"`
		CallerContractAddress string `json:"caller_contract_address"`
		ContractAddress       string `json:"contract_address"`
		EventIndex            int    `json:"event_index"`
		EventName             string `json:"event_name"`
		Result                struct {
			Field1 string `json:"0"`
			Field2 string `json:"1"`
			Field3 string `json:"2"`
			From   string `json:"from"`
			To     string `json:"to"`
			Value  string `json:"value"`
		} `json:"result"`
		ResultType struct {
			From  string `json:"from"`
			To    string `json:"to"`
			Value string `json:"value"`
		} `json:"result_type"`
		Event         string `json:"event"`
		TransactionId string `json:"transaction_id"`
	} `json:"data"`
	Meta struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}
type GetEventsByBlockNumberReq struct {
	BlockNumber   int32
	OnlyConfirmed bool
	Limit         int32
	Fingerprint   string
}

func (e *Events) GetEventsByBlockNumber(req *GetEventsByBlockNumberReq) (*GetEventsByBlockNumberResp, error) {
	values := url.Values{}
	values.Set("only_confirmed", fmt.Sprintf("%v", req.OnlyConfirmed))
	if req.Limit != 0 {
		values.Set("limit", fmt.Sprintf("%d", req.Limit))
	}
	values.Set("fingerprint", req.Fingerprint)
	path := fmt.Sprintf("/v1/blocks/%d/events", req.BlockNumber)
	response, err := e.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetEventsByBlockNumberResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
