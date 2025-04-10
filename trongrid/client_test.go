package trongrid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sleep-go/coin-go/trongrid/accounts"
	"github.com/sleep-go/coin-go/trongrid/assets"
	"github.com/sleep-go/coin-go/trongrid/base"
	"github.com/sleep-go/coin-go/trongrid/contracts"
	"github.com/sleep-go/coin-go/trongrid/events"
)

var a accounts.Accounts
var a1 assets.Assets
var c contracts.Contracts
var e events.Events

func init() {
	client := base.NewClient(http.DefaultClient, true)
	a = accounts.Accounts{Client: client}
	a1 = assets.Assets{Client: client}
	c = contracts.Contracts{Client: client}
	e = events.Events{Client: client}
}

func TestGetAccountInfoByAddress(t *testing.T) {
	resp, err := a.GetAccountInfoByAddress(&accounts.GetAccountInfoByAddressReq{
		Address:         "TGNBhSEXcaxYcsFavVyZEbuWPqtz7mNACF",
		OnlyConfirmed:   false,
		OnlyUnconfirmed: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestGetTransactionInfoByAccountAddress(t *testing.T) {
	res, err := a.GetTransactionInfoByAccountAddress(&accounts.GetTransactionInfoByAccountAddressReq{
		Address: "TGNBhSEXcaxYcsFavVyZEbuWPqtz7mNACF",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetContractTransactionInfoByAccountAddress(t *testing.T) {
	res, err := a.GetContractTransactionInfoByAccountAddress(&accounts.GetContractTransactionInfoByAccountAddressReq{
		Address:         "TGNBhSEXcaxYcsFavVyZEbuWPqtz7mNACF",
		ContractAddress: "TDLVXu6mvt34kRRmHJtDc26bR99d7eu7No",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestListAllAssets(t *testing.T) {
	res, err := a1.ListAllAssets(&assets.ListAllAssetsReq{
		OrderBy:     "",
		Limit:       1,
		Fingerprint: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetAssetsByName(t *testing.T) {
	res, err := a1.GetAssetsByName(&assets.GetAssetsByNameReq{
		Name:          "name",
		Limit:         0,
		Fingerprint:   "",
		OrderBy:       "",
		OnlyConfirmed: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
func TestGetAssetByIdOrIssuer(t *testing.T) {
	res, err := a1.GetAssetByIdOrIssuer(&assets.GetAssetByIdOrIssuerReq{
		Identifier:    "41c0343ebf132a80a15a5b368af6938f30f5572fda",
		OnlyConfirmed: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetTransactionInfoByContractAddress(t *testing.T) {
	res, err := c.GetTransactionInfoByContractAddress(&contracts.GetTransactionInfoByContractAddressReq{
		ContractAddress: "TDLVXu6mvt34kRRmHJtDc26bR99d7eu7No",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetTrc20TokenHolderBalances(t *testing.T) {
	res, err := c.GetTrc20TokenHolderBalances(&contracts.GetTrc20TokenHolderBalancesReq{
		ContractAddress: "TDLVXu6mvt34kRRmHJtDc26bR99d7eu7No",
		OnlyConfirmed:   false,
		OnlyUnconfirmed: false,
		OrderBy:         "",
		Fingerprint:     "",
		Limit:           0,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetEventsByTransactionId(t *testing.T) {
	res, err := e.GetEventsByTransactionId(&events.GetEventsByTransactionIdReq{
		TransactionID:   "8c37004e25da2ded852f0f0afa57cb22f49dd363c416c351730f0dff3ff489a8",
		OnlyUnconfirmed: false,
		OnlyConfirmed:   false,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetEventsByContractAddress(t *testing.T) {
	res, err := e.GetEventsByContractAddress(&events.GetEventsByContractAddressReq{
		Address:           "TDLVXu6mvt34kRRmHJtDc26bR99d7eu7No",
		EventName:         "",
		BlockNumber:       0,
		OnlyUnconfirmed:   false,
		OnlyConfirmed:     false,
		MinBlockTimestamp: nil,
		MaxBlockTimestamp: nil,
		OrderBy:           "",
		Fingerprint:       "",
		Limit:             0,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetEventsByBlockNumber(t *testing.T) {
	res, err := e.GetEventsByBlockNumber(&events.GetEventsByBlockNumberReq{
		BlockNumber:   29384990,
		OnlyConfirmed: false,
		Limit:         0,
		Fingerprint:   "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestGetEventsOfLatestBlock(t *testing.T) {
	res, err := e.GetEventsOfLatestBlock(&events.GetEventsOfLatestBlockReq{
		OnlyConfirmed: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
