package coin_go

import (
	"fmt"
	"testing"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market"
)

var wsClient *binance.WsClient

func init() {
	wsClient = binance.NewWsClient(false, true, consts.WS_TEST_STREAM)
}

// {"e":"depthUpdate","E":1728284450004,"s":"BTCUSDT","U":3999521,"u":3999550,"b":[["63698.69000000","0.00691000"],["63698.27000000","0.00409000"],["63698.26000000","0.00589000"],["63689.07000000","0.00000000"],["63688.87000000","0.00000000"],["63686.36000000","0.00000000"],["63686.06000000","0.00000000"]],"a":[["63699.99000000","0.00785000"],["63701.67000000","0.00000000"],["63701.79000000","0.00000000"],["63701.99000000","0.00479000"],["63702.17000000","0.00000000"],["63702.32000000","0.00000000"],["63702.82000000","0.00409000"],["63702.83000000","0.00000000"],["63703.48000000","0.00000000"],["63703.51000000","0.00000000"],["63705.30000000","0.00424000"],["63705.37000000","0.00542000"],["63706.07000000","0.00668000"],["63707.17000000","0.00699000"],["63708.00000000","0.00000000"],["63708.02000000","0.00605000"],["63709.39000000","0.00636000"],["63709.40000000","0.00691000"],["63709.94000000","0.00746000"],["63711.34000000","0.00424000"],["63731.16000000","0.00000000"],["63733.34000000","0.00000000"],["63733.67000000","0.00000000"]]}
func TestDepthWs(t *testing.T) {
	if wsClient.IsCombined {
		err := market.NewStreamDepth(wsClient, []string{"BTCUSDT", "ETHUSDT"}, func(event *market.StreamDepthEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	} else {
		err := market.NewWsDepth(wsClient, []string{"BTCUSDT", "ETHUSDT"}, func(event *market.WsDepthEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestWsDepthLevel(t *testing.T) {
	if wsClient.IsCombined {
		err := market.NewStreamDepthLevels(wsClient, map[string]enums.LimitType{
			"BTCUSDT": enums.Limit5,
		}, func(event *market.StreamDepthLevelsEvent) {
			fmt.Println(event.Stream, event.Data)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	} else {
		err := market.NewWsDepthLevels(wsClient, map[string]enums.LimitType{
			"BTCUSDT": enums.Limit5,
		}, func(event *market.WsDepthLevelsEvent) {
			fmt.Println(event)
		}, func(messageType int, err error) {
			fmt.Println(messageType, err)
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}
