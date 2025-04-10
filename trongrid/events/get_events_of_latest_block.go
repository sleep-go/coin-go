package events

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/sleep-go/coin-go/trongrid/base"
)

type GetEventsOfLatestBlockResp struct {
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
type GetEventsOfLatestBlockReq struct {
	OnlyConfirmed bool
}

func (e *Events) GetEventsOfLatestBlock(req *GetEventsOfLatestBlockReq) (*GetEventsOfLatestBlockResp, error) {
	values := url.Values{}
	values.Set("only_confirmed", fmt.Sprintf("%v", req.OnlyConfirmed))
	path := fmt.Sprintf("/v1/blocks/latest/events")
	response, err := e.Client.Get(path, values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp = new(GetEventsOfLatestBlockResp)
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
