package tron

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sleep-go/coin-go/tron/accounts"
	"github.com/sleep-go/coin-go/tron/assets"
	"github.com/sleep-go/coin-go/tron/base"
)

var a accounts.Accounts
var a1 assets.Assets

func init() {
	client := base.NewClient(http.DefaultClient, true)
	a = accounts.Accounts{Client: client}
	a1 = assets.Assets{Client: client}
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
