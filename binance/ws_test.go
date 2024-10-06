package binance

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sleep-go/coin-go/binance/consts"
)

var wsClient *WsClient

func init() {
	wsClient = NewWsClient(false, consts.WS_TEST_STREAM)
}
func TestDepthWs(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	h := http.Header{}
	//h.Add("User-Agent", "coin-go/v0.1")
	conn, _, err := websocket.DefaultDialer.Dial(consts.WS_TEST_STREAM, h)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("test")
	defer conn.Close()
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
