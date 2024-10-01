package account

import (
	"context"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
)

type Account interface {
	Call(ctx context.Context) (body *getAccountResponse, err error)
	SetOmitZeroBalances(omitZeroBalances bool) Account
	SetTimestamp(timestamp int64) Account
	SetRecvWindow(recvWindow int64) Account
}

type getAccountRequest struct {
	*binance.Client
	omitZeroBalances bool //如果true，将隐藏所有零余额。默认值：false
	recvWindow       int64
	timestamp        int64
}

type getAccountResponse struct {
	consts.ErrorResponse
	AccountType string `json:"accountType"`
	Balances    []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
	Brokered        bool `json:"brokered"`
	BuyerCommission int  `json:"buyerCommission"`
	CanDeposit      bool `json:"canDeposit"`
	CanTrade        bool `json:"canTrade"`
	CanWithdraw     bool `json:"canWithdraw"`
	CommissionRates struct {
		Buyer  string `json:"buyer"`
		Maker  string `json:"maker"`
		Seller string `json:"seller"`
		Taker  string `json:"taker"`
	} `json:"commissionRates"`
	MakerCommission            int      `json:"makerCommission"`
	Permissions                []string `json:"permissions"`
	PreventSor                 bool     `json:"preventSor"`
	RequireSelfTradePrevention bool     `json:"requireSelfTradePrevention"`
	SellerCommission           int      `json:"sellerCommission"`
	TakerCommission            int      `json:"takerCommission"`
	Uid                        int      `json:"uid"`
	UpdateTime                 int64    `json:"updateTime"`
}

func NewGetAccount(client *binance.Client) Account {
	return &getAccountRequest{Client: client}
}

func (g *getAccountRequest) SetTimestamp(timestamp int64) Account {
	g.timestamp = timestamp
	return g
}

func (g *getAccountRequest) SetRecvWindow(recvWindow int64) Account {
	g.recvWindow = recvWindow
	return g
}

func (g *getAccountRequest) SetOmitZeroBalances(omitZeroBalances bool) Account {
	g.omitZeroBalances = omitZeroBalances
	return g
}

func (g *getAccountRequest) Call(ctx context.Context) (body *getAccountResponse, err error) {
	req := &binance.Request{
		Method: http.MethodGet,
		Path:   consts.ApiAccount,
	}
	req.SetNeedSign(true)
	req.SetParam("omitZeroBalances", g.omitZeroBalances)
	if g.recvWindow > 0 {
		req.SetParam("recvWindow", g.recvWindow)
	}
	req.SetParam("timestamp", g.timestamp)
	resp, err := g.Do(ctx, req)
	if err != nil {
		g.Debugf("getAccountRequest response err:%v", err)
		return nil, err
	}
	err = netutil.ParseHttpResponse(resp, &body)
	if err != nil {
		g.Debugf("ParseHttpResponse err:%v", err)
		return nil, err
	}
	return body, nil
}
