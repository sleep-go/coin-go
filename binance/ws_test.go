package binance

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/sleep-go/coin-go/binance/consts"
)

var wsClient *WsClient

func init() {
	wsClient = NewWsClient(true, consts.WS_TEST_STREAM)
	wsClient.Endpoint += "btcusdt@avgPrice"
	wsClient.Connection()
	fmt.Println(wsClient.Endpoint)

}

// {"e":"depthUpdate","E":1728284450004,"s":"BTCUSDT","U":3999521,"u":3999550,"b":[["63698.69000000","0.00691000"],["63698.27000000","0.00409000"],["63698.26000000","0.00589000"],["63689.07000000","0.00000000"],["63688.87000000","0.00000000"],["63686.36000000","0.00000000"],["63686.06000000","0.00000000"]],"a":[["63699.99000000","0.00785000"],["63701.67000000","0.00000000"],["63701.79000000","0.00000000"],["63701.99000000","0.00479000"],["63702.17000000","0.00000000"],["63702.32000000","0.00000000"],["63702.82000000","0.00409000"],["63702.83000000","0.00000000"],["63703.48000000","0.00000000"],["63703.51000000","0.00000000"],["63705.30000000","0.00424000"],["63705.37000000","0.00542000"],["63706.07000000","0.00668000"],["63707.17000000","0.00699000"],["63708.00000000","0.00000000"],["63708.02000000","0.00605000"],["63709.39000000","0.00636000"],["63709.40000000","0.00691000"],["63709.94000000","0.00746000"],["63711.34000000","0.00424000"],["63731.16000000","0.00000000"],["63733.34000000","0.00000000"],["63733.67000000","0.00000000"]]}
type streamDepthEvent struct {
	Stream string `json:"stream"`
	Data   struct {
		Event         string     `json:"e"`
		Time          int64      `json:"E"`
		Symbol        string     `json:"s"`
		FirstUpdateID int        `json:"U"`
		LastUpdateID  int        `json:"u"`
		Bids          [][]string `json:"b"`
		Asks          [][]string `json:"a"`
	} `json:"data"`
}
type wsDepthEvent struct {
	Event string     `json:"e"`
	E1    int64      `json:"E"`
	S     string     `json:"s"`
	U     int        `json:"U"`
	U1    int        `json:"u"`
	B     [][]string `json:"b"`
	A     [][]string `json:"a"`
}

var ws = MessageHandler(func(mt int, msg []byte) {})
var errHandler = ErrorHandler(func(mt int, err error) {})

func TestDepthWs(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer wsClient.conn.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			mt, message, err := wsClient.conn.ReadMessage()
			if err != nil {
				errHandler(mt, err)
				log.Println("read:", err)
				return
			}
			ws(mt, message)
			log.Printf("recv: %s", message)
		}
	}()
	for {
		select {
		case <-done:
			log.Println("done")
			return
		case <-interrupt:
			log.Println("interrupt")
			err := wsClient.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
		}
	}
}
