package trongrid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sleep-go/coin-go/trongrid/accounts"
	"github.com/sleep-go/coin-go/trongrid/assets"
	"github.com/sleep-go/coin-go/trongrid/base"
	"github.com/sleep-go/coin-go/trongrid/contracts"
)

var a accounts.Accounts
var a1 assets.Assets
var c contracts.Contracts

func init() {
	client := base.NewClient(http.DefaultClient, true)
	a = accounts.Accounts{Client: client}
	a1 = assets.Assets{Client: client}
	c = contracts.Contracts{Client: client}
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
